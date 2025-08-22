package kb

import (
	"context"
	"errors"
	"flowing/global"
	"flowing/internal/docprocess"
	"flowing/internal/model/ai"
	"flowing/internal/model/common"
	"flowing/internal/model/kb"
	"flowing/internal/model/monitor"
	"flowing/internal/repository"
	"flowing/internal/repository/vector"
	"fmt"
	"slices"

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
		collName := KnowledgeBaseCollectionName(model.Id)
		// 创建文件存储
		if err := repository.File().CreateBucket(c, collName); err != nil {
			return global.NewError(500, "创建知识库失败", err)
		}
		// 获取向量库数据源连接
		vectorStore, err := repository.NewVectorStore(datasource)
		if err != nil {
			return global.NewError(500, "创建知识库失败", err)
		}
		// 创建向量库
		// TODO 向量维度
		if err := vectorStore.CreateCollection(c, collName, 1024); err != nil {
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
		// 删除任务
		if err := kb.DeleteTasksInKb(c, id); err != nil {
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

func GetEmbeddingModel(ctx context.Context, id int64) (*ai.ProviderModelDetail, error) {
	var detail *ai.ProviderModelDetail
	err := repository.DB(ctx).Table("ai_provider_model apm").
		Joins("JOIN ai_provider ap on ap.id = apm.provider_id").
		Select("apm.*, ap.provider_name, ap.provider_type, ap.provider_config").
		Where("apm.id = ?", id).
		First(&detail).
		Error
	if err != nil {
		return nil, err
	}
	return detail, nil
}

func Search(ctx context.Context, req kb.KnowledgeQueryReq) ([]*kb.QueriedSlice, error) {
	// 获取知识库详情
	base, err := kb.GetKnowledgeBase(ctx, req.KnowledgeBaseId)
	if err != nil {
		return nil, global.NewError(500, "搜索知识库失败", err)
	}
	// 获取数据源
	datasource, err := monitor.GetDatasource(ctx, base.DatasourceId)
	if err != nil {
		return nil, global.NewError(500, "搜索知识库失败", err)
	}
	// 获取向量库
	vectorStore, err := repository.NewVectorStore(datasource)
	if err != nil {
		return nil, global.NewError(500, "搜索知识库失败", err)
	}
	// 获取嵌入模型
	modelDetail, err := GetEmbeddingModel(ctx, base.EmbeddingModel)
	if err != nil {
		return nil, global.NewError(500, "搜索知识库失败", err)
	}

	var embedding []float32
	searchType := vector.SearchType(req.SearchType)
	hybridType := vector.HybridType(req.HybridType)

	// 生成查询向量
	if searchType == vector.SearchTypeVector || searchType == vector.SearchTypeHybrid {
		embedding64, err := docprocess.EmbedQuery(ctx, req.QueryText, &docprocess.EmbedOption{
			ProviderType: modelDetail.ProviderType,
			ModelName:    modelDetail.ModelName,
			Config:       modelDetail.ProviderConfig,
		})
		if err != nil {
			return nil, global.NewError(500, "搜索知识库失败", err)
		}
		// 转换为float32
		embedding = make([]float32, len(embedding64))
		for i, v := range embedding64 {
			embedding[i] = float32(v)
		}
	}

	// 搜索
	collectionName := fmt.Sprintf("%s%d", CollectionNamePrefix, req.KnowledgeBaseId)
	chunks, err := vectorStore.Search(ctx, collectionName, vector.SearchReq{
		Text:       req.QueryText,
		Embedding:  embedding,
		Type:       searchType,
		HybridType: hybridType,
		TopK:       req.TopK,
		Weight:     req.Weight,
		Threshold:  req.Threshold,
	})
	if err != nil {
		return nil, err
	}

	// 获取文档名称
	docIds := make([]int64, 0)
	for _, chunk := range chunks {
		docIds = append(docIds, chunk.DocId)
	}
	// 去重id
	docIds = slices.Compact(docIds)
	docNames, err := kb.GetDocumentNames(ctx, docIds)
	if err != nil {
		return nil, err
	}
	// 文档名称map
	docNameMap := make(map[int64]string)
	for _, doc := range docNames {
		docNameMap[doc.Id] = doc.Name
	}

	result := make([]*kb.QueriedSlice, 0, len(chunks))

	// 文档名称
	for _, chunk := range chunks {
		result = append(result, &kb.QueriedSlice{
			SliceId:       chunk.SliceId,
			DocId:         chunk.DocId,
			Content:       chunk.Content,
			Score:         chunk.VectorScore,
			VectorScore:   chunk.VectorScore,
			FulltextScore: chunk.FulltextScore,
			DocumentName:  docNameMap[chunk.DocId],
		})
	}
	// TODO 重排序
	return result, nil
}
