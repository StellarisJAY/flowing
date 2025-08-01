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
	if err := db.SetupJoinTable(&system.User{}, "Roles", &system.UserRole{}); err != nil {
		panic(err)
	}
	if err := db.SetupJoinTable(&system.Role{}, "Menus", &system.RoleMenu{}); err != nil {
		panic(err)
	}
	//password := hex.EncodeToString(sha256.New().Sum([]byte("123456")))
	//if err := system.CreateUser(context.Background(), &system.User{
	//	Username: "admin",
	//	NickName: "admin",
	//	Password: password,
	//	Email:    "admin@admin",
	//	Phone:    "",
	//	Status:   1,
	//}); err != nil {
	//	panic(err)
	//}
}
