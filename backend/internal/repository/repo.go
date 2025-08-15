package repository

import (
	"context"
	"flowing/internal/config"
	"flowing/internal/repository/db"
	rdb "flowing/internal/repository/redis"

	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Data struct {
	db        *gorm.DB
	snowflake *snowflake.Node
	redis     *redis.Client
	config    *config.Config
}

var data *Data

func DB(ctx context.Context) *gorm.DB {
	d := ctx.Value("tx_db")
	if d == nil {
		d = data.db
	}
	return d.(*gorm.DB)
}

func Tx(ctx context.Context, fn func(context.Context) error) error {
	return data.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		c := context.WithValue(ctx, "tx_db", tx)
		return fn(c)
	})
}

func Redis() *redis.Client {
	return data.redis
}

func Snowflake() *snowflake.Node {
	return data.snowflake
}

func Config() *config.Config {
	return data.config
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
		config:    c,
	}
}

func PingMySQL(dsn string) error {
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return d.Exec("SELECT 1").Error
}
