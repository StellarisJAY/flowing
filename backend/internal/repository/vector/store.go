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
	Id            int64   `json:"id"`            // 主键ID
	DocId         int64   `json:"docId"`         // 文档ID
	SliceId       string  `json:"sliceId"`       // 切片ID
	Content       string  `json:"content"`       // 切片内容
	Score         float64 `json:"score"`         // 相似度得分
	VectorScore   float64 `json:"vectorScore"`   // 向量相似度得分
	FulltextScore float64 `json:"fulltextScore"` // 全文搜索相似度得分
}

type SearchType string

const (
	SearchTypeFulltext SearchType = "fulltext"
	SearchTypeVector   SearchType = "vector"
	SearchTypeHybrid   SearchType = "hybrid"
)

type HybridType string

const (
	HybridTypeWeight HybridType = "weight"
	HybridTypeRerank HybridType = "rerank"
)

type SearchReq struct {
	Text       string     // 查询文本
	TopK       int        // TopK
	Type       SearchType // 搜索类型: fulltext, vector, hybrid
	Threshold  float64    // 相似度阈值
	HybridType HybridType // 混合搜索类型: weight, rerank
	Weight     float64    // 混合搜索向量权重
	Embedding  []float32  // 查询文本向量
}

type Store interface {
	Ping() error                                                                                      // 检查连接
	Close() error                                                                                     // 关闭连接
	CreateCollection(ctx context.Context, name string, denseDims int64) error                         // 创建集合
	DropCollection(ctx context.Context, name string) error                                            // 删除集合
	Add(ctx context.Context, coll string, slices []Slice) error                                       // 添加切片
	Delete(ctx context.Context, coll string, slices []Slice) error                                    // 删除切片
	ListSlices(ctx context.Context, coll string, query ListSliceQuery) ([]QueriedSlice, int64, error) // 列表切片
	Search(ctx context.Context, coll string, req SearchReq) ([]QueriedSlice, error)
}
