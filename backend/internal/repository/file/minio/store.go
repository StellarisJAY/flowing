package minio

import (
	"context"
	"flowing/internal/config"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Store struct {
	client *minio.Client
}

func NewStore(config *config.Config) *Store {
	client, err := minio.New(config.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Minio.AccessKey, config.Minio.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}
	return &Store{
		client: client,
	}
}

func (s *Store) Upload(ctx context.Context, bucket string, content io.Reader) (string, error) {
	filename := uuid.New().String()
	info, err := s.client.PutObject(ctx, bucket, filename, content, -1, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return info.Key, nil
}

func (s *Store) Download(ctx context.Context, bucket, key string) (io.Reader, error) {
	info, err := s.client.GetObject(ctx, bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (s *Store) Delete(ctx context.Context, bucket, key string) error {
	return s.client.RemoveObject(ctx, bucket, key, minio.RemoveObjectOptions{
		ForceDelete: true,
	})
}

func (s *Store) TempDownloadURL(ctx context.Context, bucket, key string) (string, error) {
	url, err := s.client.PresignedGetObject(ctx, bucket, key, time.Hour, nil)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

func (s *Store) CreateBucket(ctx context.Context, name string) error {
	err := s.client.MakeBucket(ctx, name, minio.MakeBucketOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteBucket(ctx context.Context, name string) error {
	objects := s.client.ListObjects(ctx, name, minio.ListObjectsOptions{
		Recursive: true,
	})
	chanErrors := s.client.RemoveObjects(ctx, name, objects, minio.RemoveObjectsOptions{})
	for err := range chanErrors {
		if err.Err != nil {
			return err.Err
		}
	}
	err := s.client.RemoveBucket(ctx, name)
	if err != nil {
		return err
	}
	return nil
}
