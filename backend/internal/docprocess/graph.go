package docprocess

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"maps"
	"slices"
	"strings"
	"time"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
	"golang.org/x/sync/errgroup"
)

// Entity 知识图谱实体
type Entity struct {
	Id          string   `json:"-"`
	ChunkIds    []string `json:"-"`           // 实体在哪些chunk中出现
	Name        string   `json:"name"`        // 实体名称
	Type        string   `json:"type"`        // 实体类型, 比如：PERSON, ORGANIZATION, LOCATION, etc.
	Description string   `json:"description"` // 实体描述
	Frequency   int      `json:"-"`           // 实体在文档中出现的次数
}

// Relationship 知识图谱关系
type Relationship struct {
	Id          string   `json:"-"`
	ChunkIds    []string `json:"-"`           // 关系在哪些chunk中出现
	Source      string   `json:"source"`      // 关系的源实体名称
	Target      string   `json:"target"`      // 关系的目标实体名称
	SourceId    string   `json:"-"`           // 关系的源实体ID
	TargetId    string   `json:"-"`           // 关系的目标实体ID
	Description string   `json:"description"` // 关系描述
	Strength    int      `json:"strength"`    // 关系强度
}

// KnowledgeGraph 知识图谱
type KnowledgeGraph struct {
	Entities  []*Entity       `json:"entities"`
	Relations []*Relationship `json:"relations"`
}

type KnowledgeGraphBuilder struct {
	entityPrompt   string
	relationPrompt string
	chatModel      model.BaseChatModel
}

func NewKnowledgeGraphBuilder(entityPrompt, relationPrompt string, chatModel model.BaseChatModel) (*KnowledgeGraphBuilder, error) {
	return &KnowledgeGraphBuilder{
		entityPrompt:   entityPrompt,
		relationPrompt: relationPrompt,
		chatModel:      chatModel,
	}, nil
}

func (b *KnowledgeGraphBuilder) Build(ctx context.Context, chunks []*schema.Document) (*KnowledgeGraph, error) {
	kg := &KnowledgeGraph{}
	entityGroup, gCtx := errgroup.WithContext(ctx)
	entityGroup.SetLimit(8)
	chunkEntities := make([][]*Entity, len(chunks))
	// 提取实体
	for i, chunk := range chunks {
		entityGroup.Go(func() error {
			startTime := time.Now()
			entities, err := b.modelGenEntities(gCtx, chunk)
			if err != nil {
				slog.Debug("extract entity failed", "chunkId", chunk.ID, "err", err)
				return nil
			}
			chunkEntities[i] = entities
			slog.Debug("extract entity done", "chunkId", chunk.ID, "duration(s)", time.Since(startTime).Seconds())
			return nil
		})
	}

	if err := entityGroup.Wait(); err != nil {
		return nil, err
	}
	// 将同类、同名的实体合并
	entities := make(map[string]*Entity)
	for _, ce := range chunkEntities {
		for _, entity := range ce {
			id := fmt.Sprintf("%s:%s", entity.Type, entity.Name)
			if existEntity, ok := entities[id]; ok {
				existEntity.ChunkIds = append(existEntity.ChunkIds, entity.ChunkIds...)
			} else {
				entity.Id = id
				entities[id] = entity
			}
		}
	}

	kg.Entities = slices.Collect(maps.Values(entities))

	// 提取关系
	relGroup, rCtx := errgroup.WithContext(ctx)
	relGroup.SetLimit(8)
	chunkRelations := make([][]*Relationship, len(chunks))
	for i, chunk := range chunks {
		relGroup.Go(func() error {
			startTime := time.Now()
			relations, err := b.modelGenRelations(rCtx, chunk, chunkEntities[i])
			if err != nil {
				slog.Debug("extract relation failed", "chunkId", chunk.ID, "err", err)
				return nil
			}
			chunkRelations[i] = relations
			slog.Debug("extract relation done", "chunkId", chunk.ID, "relationCount", len(relations), "duration(s)", time.Since(startTime).Seconds())
			return nil
		})
	}

	if err := relGroup.Wait(); err != nil {
		return nil, err
	}

	// 合并关系
	relations := make(map[string]*Relationship)
	for _, cr := range chunkRelations {
		for _, rel := range cr {
			id := fmt.Sprintf("%s-%s->%s", rel.Source, rel.Description, rel.Target)
			if existRelation, ok := relations[id]; ok {
				existRelation.ChunkIds = append(existRelation.ChunkIds, rel.ChunkIds...)
			} else {
				rel.Id = id
				relations[id] = rel
			}
		}
	}
	kg.Relations = slices.Collect(maps.Values(relations))
	return kg, nil
}

func (b *KnowledgeGraphBuilder) modelGenEntities(ctx context.Context, chunk *schema.Document) ([]*Entity, error) {
	// 提取实体
	response, err := b.chatModel.Generate(ctx, []*schema.Message{
		{
			Role:    schema.System,
			Content: b.entityPrompt,
		},
		{
			Role:    schema.User,
			Content: chunk.Content,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("extract entity failed: %w", err)
	}
	var entities []*Entity
	err = json.Unmarshal([]byte(strings.TrimSuffix(strings.TrimPrefix(response.Content, "```json"), "```")), &entities)
	if err != nil {
		return nil, fmt.Errorf("unmarshal entity failed: %w", err)
	}
	// 实体设置chunkId
	for _, entity := range entities {
		entity.ChunkIds = append(entity.ChunkIds, chunk.ID)
	}
	return entities, nil
}

func (b *KnowledgeGraphBuilder) modelGenRelations(ctx context.Context, chunk *schema.Document, entities []*Entity) ([]*Relationship, error) {
	ent, _ := json.Marshal(entities)
	messages := []*schema.Message{
		{
			Role:    schema.System,
			Content: b.relationPrompt,
		},
		{
			Role:    schema.User,
			Content: fmt.Sprintf("实体: %s\n\n文本: %s", string(ent), chunk.Content),
		},
	}
	// 提取关系
	response, err := b.chatModel.Generate(ctx, messages)
	if err != nil {
		return nil, fmt.Errorf("extract relation failed: %w", err)
	}
	var relations []*Relationship
	err = json.Unmarshal([]byte(strings.TrimSuffix(strings.TrimPrefix(response.Content, "```json"), "```")), &relations)
	if err != nil {
		return nil, fmt.Errorf("unmarshal relation failed: %w", err)
	}
	// 关系设置chunkId
	for _, relation := range relations {
		relation.ChunkIds = append(relation.ChunkIds, chunk.ID)
	}
	return relations, nil
}
