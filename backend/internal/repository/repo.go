package repository

import (
	"flowing/internal/config"
	"flowing/internal/repository/db"
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

type Data struct {
	db        *gorm.DB
	snowflake *snowflake.Node
}

var data *Data

func DB() *gorm.DB {
	return data.db
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
	data = &Data{
		db:        database,
		snowflake: sf,
	}
}
