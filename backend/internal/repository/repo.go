package repository

import (
	"flowing/internal/config"
	"flowing/internal/repository/db"
	"gorm.io/gorm"
)

type Data struct {
	db *gorm.DB
}

var data *Data

func DB() *gorm.DB {
	return data.db
}

func Init(c *config.Config) {
	database, err := db.Init(c)
	if err != nil {
		panic(err)
	}
	data = &Data{
		db: database,
	}
}
