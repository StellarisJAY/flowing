package migration

import (
	"context"
	"flowing/internal/model/agent"
	"flowing/internal/model/ai"
	"flowing/internal/model/kb"
	"flowing/internal/model/monitor"
	"flowing/internal/model/system"
	"flowing/internal/repository"
)

func MigrateDB() {
	db := repository.DB(context.Background())
	err := db.AutoMigrate(
		&system.User{},
		&system.Role{},
		&system.Dict{},
		&system.DictItem{},
		&system.Menu{},
		&system.UserRole{},
		&system.RoleMenu{},
		&ai.Provider{},
		&ai.ProviderModel{},
		&monitor.Datasource{},
		&kb.KnowledgeBase{},
		&kb.Document{},
		&kb.Task{},
		&agent.Agent{},
	)
	if err != nil {
		panic(err)
	}
	//if err := db.SetupJoinTable(&system.User{}, "Roles", &system.UserRole{}); err != nil {
	//	panic(err)
	//}
	//if err := db.SetupJoinTable(&system.Role{}, "Menus", &system.RoleMenu{}); err != nil {
	//	panic(err)
	//}
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
