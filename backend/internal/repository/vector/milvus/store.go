package milvus

import (
	"context"
	"errors"
	"flowing/internal/repository/vector"
	"fmt"
	"time"

	"github.com/milvus-io/milvus/client/v2/entity"
	"github.com/milvus-io/milvus/client/v2/index"
	"github.com/milvus-io/milvus/client/v2/milvusclient"
)

type Store struct {
	client *milvusclient.Client
}

func NewStore(address string, username string, password string, dbName string) (*Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := milvusclient.New(ctx, &milvusclient.ClientConfig{
		Address:  address,
		Username: username,
		Password: password,
		DBName:   dbName,
	})
	if err != nil {
		return nil, err
	}
	return &Store{client: client}, nil
}

func (s *Store) Ping() error {
	return nil
}
func (s *Store) Close() error {
	return s.client.Close(context.Background())
}

func (s *Store) CreateCollection(ctx context.Context, name string, denseDims int64) error {
	schema, indexOptions := getCollectionSchema(name, denseDims)
	option := milvusclient.NewCreateCollectionOption(name, schema).WithIndexOptions(indexOptions...)
	return s.client.CreateCollection(ctx, option)
}

func (s *Store) DropCollection(ctx context.Context, name string) error {
	return s.client.DropCollection(ctx, milvusclient.NewDropCollectionOption(name))
}

func (s *Store) Add(ctx context.Context, coll string, slices []vector.Slice) error {
	rows := make([]any, len(slices))
	for i, slice := range slices {
		rows[i] = map[string]any{
			"id":           slice.Id(),
			"doc_id":       slice.DocId(),
			"slice_id":     slice.Id(),
			"content":      slice.Content(),
			"dense_vector": slice.DenseVector(),
			"metadata":     slice.Metadata(),
		}
	}
	_, err := s.client.Insert(ctx, milvusclient.NewRowBasedInsertOption(coll, rows...))
	return err
}

func (s *Store) Delete(ctx context.Context, coll string, slices []vector.Slice) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) ListSlices(ctx context.Context, coll string, query vector.ListSliceQuery) ([]vector.QueriedSlice, int64, error) {

	option := milvusclient.NewQueryOption(coll).WithFilter(fmt.Sprintf("doc_id == %d", query.DocId))
	var total int64
	var err error
	var resultSet milvusclient.ResultSet

	_, err = s.client.LoadCollection(ctx, milvusclient.NewLoadCollectionOption(coll))
	if err != nil {
		return nil, 0, err
	}

	if query.Page {
		totalOption := option.WithOutputFields("id")
		totalResultSet, err := s.client.Query(ctx, totalOption)
		if err != nil {
			return nil, 0, err
		}
		total = int64(totalResultSet.ResultCount)
		option = option.WithLimit(int(query.PageSize)).
			WithOffset(int((query.PageNum-1)*query.PageSize)).
			WithOutputFields("id", "doc_id", "slice_id", "content")
		resultSet, err = s.client.Query(ctx, option)
		if err != nil {
			return nil, 0, err
		}
	} else {
		resultSet, err = s.client.Query(ctx, option)
		if err != nil {
			return nil, 0, err
		}
		total = int64(resultSet.ResultCount)
	}

	return fromResultSetToQueryResult(resultSet), total, nil
}

func fromResultSetToQueryResult(resultSet milvusclient.ResultSet) []vector.QueriedSlice {
	res := make([]vector.QueriedSlice, resultSet.Len())
	idColumn := resultSet.GetColumn("id")
	docIdColumn := resultSet.GetColumn("doc_id")
	sliceIdColumn := resultSet.GetColumn("slice_id")
	contentColumn := resultSet.GetColumn("content")
	scores := resultSet.Scores
	if len(scores) == 0 {
		scores = make([]float32, resultSet.Len())
	}
	for i := 0; i < resultSet.Len(); i++ {
		id, _ := idColumn.GetAsInt64(i)
		docId, _ := docIdColumn.GetAsInt64(i)
		sliceId, _ := sliceIdColumn.GetAsString(i)
		content, _ := contentColumn.GetAsString(i)
		res[i] = vector.QueriedSlice{
			Id:            id,
			DocId:         docId,
			SliceId:       sliceId,
			Content:       content,
			Score:         float64(scores[i]),
			VectorScore:   float64(scores[i]),
			FulltextScore: float64(scores[i]),
		}
	}
	return res
}

func (s *Store) Search(ctx context.Context, coll string, req vector.SearchReq) ([]vector.QueriedSlice, error) {
	var resultSet []milvusclient.ResultSet
	var err error
	switch req.Type {
	case vector.SearchTypeFulltext:
		resultSet, err = s.fulltextSearch(ctx, coll, req)
	case vector.SearchTypeVector:
		resultSet, err = s.vectorSearch(ctx, coll, req)
	case vector.SearchTypeHybrid:
		resultSet, err = s.hybridSearch(ctx, coll, req)
	default:
		return nil, errors.New("search type not supported")
	}
	if err != nil {
		return nil, err
	}
	res := fromResultSetToQueryResult(resultSet[0])
	return res, nil
}

// fulltextSearch 全文搜索
func (s *Store) fulltextSearch(ctx context.Context, coll string, req vector.SearchReq) ([]milvusclient.ResultSet, error) {
	annSearchParams := index.NewCustomAnnParam()
	annSearchParams.WithExtraParam("drop_ratio_search", 0.2)
	option := milvusclient.NewSearchOption(coll, req.TopK, []entity.Vector{entity.Text(req.Text)}).
		WithANNSField("sparse_vector").                         //  ANN搜索字段
		WithAnnParam(annSearchParams).                          // ANN搜索参数
		WithOutputFields("id", "doc_id", "slice_id", "content") // 输出字段
	resultSets, err := s.client.Search(ctx, option)
	if err != nil {
		return nil, err
	}
	return resultSets, nil
}

// vectorSearch 向量搜索
func (s *Store) vectorSearch(ctx context.Context, coll string, req vector.SearchReq) ([]milvusclient.ResultSet, error) {
	annParam := index.NewCustomAnnParam()
	// 查询范围：req.Threshold ~ 1.0
	annParam.WithRangeFilter(1.0)
	annParam.WithRadius(req.Threshold)
	option := milvusclient.NewSearchOption(coll, req.TopK, []entity.Vector{entity.FloatVector(req.Embedding)}).
		WithAnnParam(annParam).
		WithANNSField("dense_vector").
		WithOutputFields("id", "doc_id", "slice_id", "content")
	resultSets, err := s.client.Search(ctx, option)
	if err != nil {
		return nil, err
	}
	return resultSets, nil
}

func (s *Store) hybridSearch(ctx context.Context, coll string, req vector.SearchReq) ([]milvusclient.ResultSet, error) {
	queryText := entity.Text(req.Text)
	queryVector := entity.FloatVector(req.Embedding)
	// 全文搜索参数
	ftParam := index.NewSparseAnnParam()
	ftParam.WithDropRatio(0.2)
	ftReq := milvusclient.NewAnnRequest("sparse_vector", req.TopK, queryText).
		WithAnnParam(ftParam)
	// 向量搜索参数
	vecParam := index.NewIvfAnnParam(10)
	vecParam.WithRangeFilter(1.0)
	vecParam.WithRadius(req.Threshold)
	vecReq := milvusclient.NewAnnRequest("dense_vector", req.TopK, queryVector).WithAnnParam(vecParam)
	// 混合搜索option
	option := milvusclient.NewHybridSearchOption(coll, req.TopK, ftReq, vecReq).
		WithOutputFields("id", "doc_id", "slice_id", "content")

	if req.HybridType == vector.HybridTypeWeight {
		// 权重混合排序
		reranker := milvusclient.NewWeightedReranker([]float64{1 - req.Weight, req.Weight})
		option = option.WithReranker(reranker)
	}
	// 混合搜索
	resultSets, err := s.client.HybridSearch(ctx, option)
	if err != nil {
		return nil, err
	}
	return resultSets, nil
}

// getCollectionSchema 获取集合模式
func getCollectionSchema(name string, denseDims int64) (*entity.Schema, []milvusclient.CreateIndexOption) {
	// 主键
	pkField := entity.NewField().WithName("id").WithDataType(entity.FieldTypeInt64).WithIsPrimaryKey(true).
		WithIsAutoID(true)
	// 文档ID
	docIdField := entity.NewField().WithName("doc_id").WithDataType(entity.FieldTypeInt64)
	// 切片ID
	sliceIdField := entity.NewField().WithName("slice_id").WithDataType(entity.FieldTypeVarChar).WithMaxLength(128)
	// 切片内容 (使用jieba分词器支持中文)
	contentField := entity.NewField().WithName("content").WithDataType(entity.FieldTypeVarChar).WithMaxLength(10240).
		WithEnableAnalyzer(true).WithAnalyzerParams(map[string]any{
		"tokenizer": "jieba",
	})
	// TODO 向量模型维度
	denseField := entity.NewField().WithName("dense_vector").WithDataType(entity.FieldTypeFloatVector).WithDim(denseDims)
	// 元数据
	metadataField := entity.NewField().WithName("metadata").WithDataType(entity.FieldTypeJSON)
	// 稀疏向量
	sparseField := entity.NewField().WithName("sparse_vector").WithDataType(entity.FieldTypeSparseVector)

	// 密集向量索引
	denseIndex := index.NewAutoIndex(entity.IP)
	// 稀疏向量索引
	sparseIndex := index.NewAutoIndex(entity.BM25)

	indexOptions := []milvusclient.CreateIndexOption{
		milvusclient.NewCreateIndexOption(name, "dense_vector", denseIndex),
		milvusclient.NewCreateIndexOption(name, "sparse_vector", sparseIndex), // TODO bm25索引参数
	}

	// 自动生成bm25向量
	bm25Func := entity.NewFunction().
		WithName("content_to_bm25").
		WithType(entity.FunctionTypeBM25).
		WithInputFields("content").
		WithOutputFields("sparse_vector")

	schema := entity.NewSchema().
		WithName(name).
		WithDescription("flowing vector store schema").
		WithField(pkField).
		WithField(contentField).
		WithField(denseField).
		WithField(metadataField).
		WithField(sparseField).
		WithField(docIdField).
		WithField(sliceIdField).
		WithFunction(bm25Func)

	return schema, indexOptions
}
