package mysql

import (
	"flowing/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(c *config.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(c.MySQL.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
