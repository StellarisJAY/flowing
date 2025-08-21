package vector

import (
	"context"
)

// Slice 文档切片
// 文档切片是文档的一个片段，用于存储在向量数据库中
type Slice interface {
	Content() string             // 切片内容
	DocId() int64                // 文档ID
	Id() string                  // 切片ID
	DenseVector() []float32      // 稠密向量
	Metadata() map[string]string // 元数据
}

type StoreDatasource interface {
	GetType() string
	GetAddr() string
	GetUsername() string
	GetPassword() string
	GetDatabase() string
}

type ListSliceQuery struct {
	DocId    int64 `json:"docId" form:"docId"`
	Page     bool  `json:"page" form:"page"`
	PageNum  int64 `json:"pageNum" form:"pageNum"`
	PageSize int64 `json:"pageSize" form:"pageSize"`
}

type QueriedSlice struct {
	Id      int64   `json:"id"`
	DocId   int64   `json:"docId"`
	SliceId string  `json:"sliceId"`
	Content string  `json:"content"`
	Score   float64 `json:"score"`
}

type Store interface {
	Ping() error                                                                                      // 检查连接
	Close() error                                                                                     // 关闭连接
	CreateCollection(ctx context.Context, name string, denseDims int64) error                         // 创建集合
	DropCollection(ctx context.Context, name string) error                                            // 删除集合
	Add(ctx context.Context, coll string, slices []Slice) error                                       // 添加切片
	Delete(ctx context.Context, coll string, slices []Slice) error                                    // 删除切片
	ListSlices(ctx context.Context, coll string, query ListSliceQuery) ([]QueriedSlice, int64, error) // 列表切片
}
