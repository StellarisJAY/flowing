package mysql

import (
	"flowing/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(c *config.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(c.MySQL.DSN), &gorm.Config{})
}
