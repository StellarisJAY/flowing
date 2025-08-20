package vector

import (
	"context"
)

// Slice 文档切片
// 文档切片是文档的一个片段，用于存储在向量数据库中
type Slice interface {
	Content() []byte             // 切片内容
	DocId() int64                // 文档ID
	Id() int64                   // 切片ID
	DenseVector() []float64      // 稠密向量
	Metadata() map[string]string // 元数据
	SparseVector() []float64     // 稀疏向量
}

type StoreDatasource interface {
	GetType() string
	GetAddr() string
	GetUsername() string
	GetPassword() string
	GetDatabase() string
}

type Store interface {
	Ping() error                                             // 检查连接
	Close() error                                            // 关闭连接
	CreateCollection(ctx context.Context, name string) error // 创建集合
	DropCollection(ctx context.Context, name string) error   // 删除集合
	Add(ctx context.Context, slices []Slice) error           // 添加切片
	Delete(ctx context.Context, slices []Slice) error        // 删除切片
}
