package kb

import (
	"context"
	"errors"
	"flowing/global"
	"flowing/internal/model/ai"
	"flowing/internal/model/common"
	"flowing/internal/model/kb"
	"flowing/internal/model/monitor"
	"flowing/internal/repository"
	"fmt"

	"gorm.io/gorm"
)

const (
	CollectionNamePrefix = "flowingkb"
)

func KnowledgeBaseCollectionName(id int64) string {
	return fmt.Sprintf("%s%d", CollectionNamePrefix, id)
}

func CreateKnowledgeBase(ctx context.Context, req kb.CreateKnowledgeBaseReq) error {
	datasource, err := monitor.GetDatasource(ctx, req.DatasourceId)
	if err != nil {
		return global.NewError(500, "创建知识库失败", err)
	}
	if datasource.Type != monitor.DatasourceTypeMilvus {
		return global.NewError(500, "创建知识库失败, 数据源类型错误", nil)
	}
	embeddingModel, err := ai.GetProviderModel(ctx, req.EmbeddingModel)
	if err != nil {
		return global.NewError(500, "创建知识库失败", err)
	}
	if embeddingModel == nil || embeddingModel.ModelType != ai.ModelTypeEmbedding {
		return global.NewError(500, "创建知识库失败，模型类型错误", nil)
	}
	model := kb.KnowledgeBase{
		Name:           req.Name,
		Description:    req.Description,
		DatasourceId:   req.DatasourceId,
		EmbeddingModel: req.EmbeddingModel,
		Enable:         req.Enable,
	}
	return repository.Tx(ctx, func(c context.Context) error {
		if err := kb.CreateKnowledgeBase(c, &model); err != nil {
			return global.NewError(500, "创建知识库失败", err)
		}
		// 创建文件存储
		if err := repository.File().CreateBucket(c, fmt.Sprintf("%s%d", CollectionNamePrefix, model.Id)); err != nil {
			return global.NewError(500, "创建知识库失败", err)
		}
		// 获取向量库数据源连接
		vectorStore, err := repository.NewVectorStore(datasource)
		if err != nil {
			return global.NewError(500, "创建知识库失败", err)
		}
		// 创建向量库
		if err := vectorStore.CreateCollection(c, fmt.Sprintf("%s%d", CollectionNamePrefix, model.Id)); err != nil {
			return global.NewError(500, "创建知识库失败", err)
		}
		return nil
	})
}

func UpdateKnowledgeBase(ctx context.Context, req kb.UpdateKnowledgeBaseReq) error {
	err := kb.UpdateKnowledgeBase(ctx, kb.KnowledgeBase{
		BaseModel: common.BaseModel{
			Id: req.Id,
		},
		Name:        req.Name,
		Description: req.Description,
		Enable:      req.Enable,
	})
	if err != nil {
		return global.NewError(500, "更新知识库失败", err)
	}
	return nil
}

func ListKnowledgeBase(ctx context.Context, query kb.KnowledgeBaseQuery) ([]*kb.KnowledgeBase, int64, error) {
	result, total, err := kb.ListKnowledgeBase(ctx, query)
	if err != nil {
		return nil, 0, global.NewError(500, "获取知识库列表失败", err)
	}
	return result, total, nil
}

func DeleteKnowledgeBase(ctx context.Context, id int64) error {
	return repository.Tx(ctx, func(c context.Context) error {
		// 获取知识库详情
		base, err := kb.GetKnowledgeBase(c, id)
		if err != nil {
			return global.NewError(500, "删除知识库失败", err)
		}
		var docKeys []string
		// 获取知识库的所有文档存储key
		err = repository.DB(c).Model(&kb.Document{}).
			Where("knowledge_base_id = ?", id).
			Pluck("uri", &docKeys).
			Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return global.NewError(500, "删除知识库失败", err)
		}

		// 删除文档
		if err := repository.DB(c).Where("knowledge_base_id = ?", id).Delete(&kb.Document{}).Error; err != nil {
			return global.NewError(500, "删除知识库失败", err)
		}
		// 删除知识库
		if err := repository.DB(c).Delete(&kb.KnowledgeBase{}, "id = ?", id).Error; err != nil {
			return global.NewError(500, "删除知识库失败", err)
		}
		// 获取数据源
		datasource, err := monitor.GetDatasource(c, base.DatasourceId)
		if err != nil {
			return global.NewError(500, "删除知识库失败", err)
		}
		// 删除向量库
		vectorStore, err := repository.NewVectorStore(datasource)
		if err != nil {
			return global.NewError(500, "删除知识库失败", err)
		}

		collectionName := fmt.Sprintf("%s%d", CollectionNamePrefix, id)
		// 删除向量库
		if err := vectorStore.DropCollection(c, collectionName); err != nil {
			return global.NewError(500, "删除知识库失败", fmt.Errorf("删除向量存储失败: %w", err))
		}
		// 删除文档存储
		if err := repository.File().DeleteBucket(c, collectionName); err != nil {
			return global.NewError(500, "删除知识库失败", fmt.Errorf("删除文档存储失败: %w", err))
		}
		return nil
	})
}
