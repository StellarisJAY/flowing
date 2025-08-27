package docprocess

import (
	"context"
	"flowing/internal/model/ai"
	"flowing/internal/model/kb"
	"flowing/internal/model/monitor"
	"flowing/internal/repository"
	"flowing/internal/repository/vector"
	"flowing/internal/util"
	"strconv"

	"github.com/cloudwego/eino/components/retriever"
	"github.com/cloudwego/eino/schema"
)

// VectorRetriever 向量检索器 实现eino的Retriever接口
type VectorRetriever struct {
	store          vector.Store
	knowledgeBase  *kb.KnowledgeBase
	datasource     *monitor.Datasource
	embeddingModel *ai.ProviderModelDetail
}

func NewVectorRetriever(ctx context.Context, knowledgeBase *kb.KnowledgeBase) (*VectorRetriever, error) {
	// 获取数据源详情
	datasource, err := monitor.GetDatasource(ctx, knowledgeBase.DatasourceId)
	if err != nil {
		return nil, err
	}
	// 获取嵌入模型详情
	embeddingModel, err := ai.GetProviderModelDetail(ctx, knowledgeBase.EmbeddingModel)
	if err != nil {
		return nil, err
	}
	// 创建向量存储
	store, err := repository.NewVectorStore(datasource)
	if err != nil {
		return nil, err
	}
	return &VectorRetriever{
		store:          store,
		knowledgeBase:  knowledgeBase,
		datasource:     datasource,
		embeddingModel: embeddingModel,
	}, nil
}

func (v *VectorRetriever) Retrieve(ctx context.Context, query string, opts ...retriever.Option) ([]*schema.Document, error) {
	// 创建嵌入模型
	embedder, err := util.GetEmbeddingModel(ctx, *v.embeddingModel)
	if err != nil {
		return nil, err
	}
	// 嵌入查询
	embeddings, err := embedder.EmbedStrings(ctx, []string{query})
	if err != nil {
		return nil, err
	}
	collName := util.KnowledgeBaseCollectionName(v.knowledgeBase.Id)
	// TODO 查询选项
	docs, err := v.store.Search(ctx, collName, vector.SearchReq{
		Text:      query,
		TopK:      10,
		Type:      vector.SearchTypeVector,
		Threshold: 0.5,
		Embedding: util.Float64EmbeddingTo32(embeddings[0]),
	})
	if err != nil {
		return nil, err
	}
	// 转换为文档
	var documents []*schema.Document
	for _, doc := range docs {
		documents = append(documents, &schema.Document{
			ID:      strconv.FormatInt(doc.Id, 10),
			Content: doc.Content,
			MetaData: map[string]any{
				"sliceId":         doc.SliceId,
				"docId":           doc.DocId,
				"score":           doc.Score,
				"vectorScore":     doc.VectorScore,
				"fulltextScore":   doc.FulltextScore,
				"knowledgeBaseId": v.knowledgeBase.Id,
			},
		})
	}
	return documents, nil
}
