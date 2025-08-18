package db

import (
	"errors"
	"flowing/internal/config"
	"flowing/internal/repository/db/mysql"
	"flowing/internal/repository/db/postgres"

	"gorm.io/gorm"
)

func Page(page bool, pageNum, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum == 0 {
			pageNum = 1
		}
		if pageSize == 0 {
			pageSize = 10
		}
		if !page {
			return db
		}
		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func Init(c *config.Config) (*gorm.DB, error) {
	switch c.DB.Driver {
	case "postgres":
		return postgres.Init(c)
	case "mysql":
		return mysql.Init(c)
	default:
		return nil, errors.New("DB driver not supported")
	}
}
