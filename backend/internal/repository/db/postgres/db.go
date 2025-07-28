package postgres

import (
	"flowing/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(c *config.Config) (*gorm.DB, error) {
	// TODO 更多数据库设置
	return gorm.Open(postgres.Open(c.Postgres.DSN), &gorm.Config{})
}
