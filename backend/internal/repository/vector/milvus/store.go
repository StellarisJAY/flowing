package milvus

import (
	"context"
	"time"

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
