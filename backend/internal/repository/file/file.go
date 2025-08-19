package file

import (
	"context"
	"flowing/internal/config"
	"flowing/internal/repository/file/minio"
	"io"
)

type Store interface {
	// Upload 上传文件, 返回文件路径
	Upload(ctx context.Context, content io.Reader) (string, error)
	// Download 下载文件, 返回文件内容
	Download(ctx context.Context, key string) (io.Reader, error)
	// Delete 删除文件
	Delete(ctx context.Context, key string) error
	// TempDownloadURL 生成临时下载URL
	TempDownloadURL(ctx context.Context, key string) (string, error)
}

func NewStore(config *config.Config) Store {
	switch config.FileStore {
	case "minio":
		return minio.NewStore(config)
	default:
		panic("file store not supported")
	}
}
