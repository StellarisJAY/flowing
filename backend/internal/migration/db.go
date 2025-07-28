package migration

import (
	"flowing/internal/model/system"
	"flowing/internal/repository"
)

func MigrateDB() {
	db := repository.DB()
	err := db.AutoMigrate(
		&system.User{},
		&system.Role{},
		&system.Menu{},
		&system.UserRole{},
		&system.RoleMenu{})
	if err != nil {
		panic(err)
	}
}
