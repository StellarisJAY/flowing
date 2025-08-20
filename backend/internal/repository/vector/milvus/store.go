package milvus

import (
	"context"
	"flowing/internal/repository/vector"
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

func (s *Store) CreateCollection(ctx context.Context, name string) error {
	schema, indexOptions := getCollectionSchema(name)
	option := milvusclient.NewCreateCollectionOption(name, schema).WithIndexOptions(indexOptions...)
	return s.client.CreateCollection(ctx, option)
}

func (s *Store) DropCollection(ctx context.Context, name string) error {
	return s.client.DropCollection(ctx, milvusclient.NewDropCollectionOption(name))
}

func (s *Store) Add(ctx context.Context, slices []vector.Slice) error {
	// TODO implement me
	panic("implement me")
}

func (s *Store) Delete(ctx context.Context, slices []vector.Slice) error {
	//TODO implement me
	panic("implement me")
}

// getCollectionSchema 获取集合模式
func getCollectionSchema(name string) (*entity.Schema, []milvusclient.CreateIndexOption) {
	// 主键
	pkField := entity.NewField().WithName("id").WithDataType(entity.FieldTypeInt64).WithIsPrimaryKey(true).
		WithIsAutoID(true)
	// 文档ID
	docIdField := entity.NewField().WithName("doc_id").WithDataType(entity.FieldTypeInt64)
	// 切片ID
	sliceIdField := entity.NewField().WithName("slice_id").WithDataType(entity.FieldTypeInt64)
	// 切片内容 (使用jieba分词器支持中文)
	contentField := entity.NewField().WithName("content").WithDataType(entity.FieldTypeVarChar).WithMaxLength(1024).
		WithEnableAnalyzer(true).WithAnalyzerParams(map[string]any{
		"tokenizer": "jieba",
	})
	// TODO 向量模型维度
	denseField := entity.NewField().WithName("dense_vector").WithDataType(entity.FieldTypeFloatVector).WithDim(1536)
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
