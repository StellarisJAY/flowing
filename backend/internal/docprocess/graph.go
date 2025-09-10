package docprocess

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
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
	for _, chunk := range chunks {
		startTime := time.Now()
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
		kg.Entities = append(kg.Entities, entities...)

		messages := []*schema.Message{
			{
				Role:    schema.System,
				Content: b.relationPrompt,
			},
			{
				Role:    schema.User,
				Content: fmt.Sprintf("实体: %s\n\n文本: %s", response.Content, chunk.Content),
			},
		}
		// 提取关系
		response, err = b.chatModel.Generate(ctx, messages)
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
		kg.Relations = append(kg.Relations, relations...)
		slog.Info("chunk extract entity and relation done", "chunk_id", chunk.ID, "duration(s)", time.Now().Sub(startTime).Seconds(), "entity_count", len(entities), "relation_count", len(relations))
	}

	// TODO 合并实体，合并关系
	return kg, nil
}
