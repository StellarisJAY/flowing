package file

import (
	"context"
	"flowing/internal/config"
	"flowing/internal/repository/file/minio"
	"io"
)

type Store interface {
	// CreateBucket 创建存储桶
	CreateBucket(ctx context.Context, name string) error
	// DeleteBucket 删除存储桶
	DeleteBucket(ctx context.Context, name string) error
	// Upload 上传文件, 返回文件路径
	Upload(ctx context.Context, bucket string, content io.Reader) (string, error)
	// Download 下载文件, 返回文件内容
	Download(ctx context.Context, bucket, key string) (io.Reader, error)
	// Delete 删除文件
	Delete(ctx context.Context, bucket, key string) error
	// TempDownloadURL 生成临时下载URL
	TempDownloadURL(ctx context.Context, bucket, key string) (string, error)
}

func NewStore(config *config.Config) Store {
	switch config.FileStore {
	case "minio":
		return minio.NewStore(config)
	default:
		panic("file store not supported")
	}
}
