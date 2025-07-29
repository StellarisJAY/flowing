package redis

import (
	"context"
	"flowing/internal/config"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func Init(c *config.Config) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port),
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := cli.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	return cli
}
