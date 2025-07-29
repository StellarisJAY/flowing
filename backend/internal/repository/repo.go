package repository

import (
	"flowing/internal/config"
	"flowing/internal/repository/db"
	rdb "flowing/internal/repository/redis"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Data struct {
	db        *gorm.DB
	snowflake *snowflake.Node
	redis     *redis.Client
}

var data *Data

func DB() *gorm.DB {
	return data.db
}

func Redis() *redis.Client {
	return data.redis
}

func Snowflake() *snowflake.Node {
	return data.snowflake
}

func Init(c *config.Config) {
	database, err := db.Init(c)
	if err != nil {
		panic(err)
	}
	sf, _ := snowflake.NewNode(1)
	redisCli := rdb.Init(c)
	data = &Data{
		db:        database,
		snowflake: sf,
		redis:     redisCli,
	}
}
