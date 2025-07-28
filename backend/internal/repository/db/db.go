package db

import (
	"errors"
	"flowing/internal/config"
	"flowing/internal/repository/db/postgres"
	"gorm.io/gorm"
)

func Page(page bool, pageNum, pageSize int, total *int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !page {
			return db
		}
		if pageNum == 0 {
			pageNum = 1
		}
		if pageSize == 0 {
			pageSize = 10
		}
		offset := (pageNum - 1) * pageSize
		db.Count(total)
		return db.Offset(offset).Limit(pageSize)
	}
}

func Init(c *config.Config) (*gorm.DB, error) {
	switch c.DB.Driver {
	case "postgres":
		return postgres.Init(c)
	default:
		return nil, errors.New("DB driver not supported")
	}
}
