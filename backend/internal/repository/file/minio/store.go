package minio

import (
	"context"
	"flowing/internal/config"
	"io"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Store struct {
	client *minio.Client
	bucket string
}

func NewStore(config *config.Config) *Store {
	client, err := minio.New(config.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Minio.AccessKey, config.Minio.SecretKey, ""),
		Secure: true,
	})
	if err != nil {
		panic(err)
	}
	return &Store{
		client: client,
		bucket: config.Minio.Bucket,
	}
}

func (s *Store) Upload(ctx context.Context, content io.Reader) (string, error) {
	filename := uuid.New().String()
	info, err := s.client.PutObject(ctx, s.bucket, filename, content, -1, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return info.Key, nil
}

func (s *Store) Download(ctx context.Context, key string) (io.Reader, error) {
	info, err := s.client.GetObject(ctx, s.bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (s *Store) Delete(ctx context.Context, key string) error {
	return s.client.RemoveObject(ctx, s.bucket, key, minio.RemoveObjectOptions{
		ForceDelete: true,
	})
}
