package kb

import (
	"context"
	"flowing/global"
	"flowing/internal/model/ai"
	"flowing/internal/model/common"
	"flowing/internal/model/kb"
	"flowing/internal/model/monitor"
	"flowing/internal/repository"
)

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
	}
	return repository.Tx(ctx, func(c context.Context) error {
		if err := kb.CreateKnowledgeBase(c, model); err != nil {
			return global.NewError(500, "创建知识库失败", err)
		}
		// TODO 创建向量库
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
