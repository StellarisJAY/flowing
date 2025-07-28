package system

import (
	"context"
	"flowing/internal/model/common"
	"flowing/internal/repository"
	"flowing/internal/repository/db"
)

type Role struct {
	common.BaseModel
	RoleName    string `json:"role_name" gorm:"column:role_name;type:varchar(50);not null;"`
	Description string `json:"description" gorm:"column:description;type:varchar(200)"`
	Menus       []Menu `json:"menus" gorm:"many2many:sys_role_menu;"`
}

func (Role) TableName() string {
	return "sys_role"
}

type UserRole struct {
	common.BaseModel
	UserId int64 `json:"user_id" gorm:"column:user_id;type:int;not null;"`
	RoleId int64 `json:"role_id" gorm:"column:role_id;type:int;not null;"`
}

func (UserRole) TableName() string {
	return "sys_user_role"
}

type RoleQuery struct {
	common.BaseQuery
	RoleName string `json:"role_name" form:"role_name"`
}

func CreateRole(ctx context.Context, role *Role) error {
	return repository.DB().WithContext(ctx).Create(role).Error
}

func GetRole(ctx context.Context, id int) (*Role, error) {
	var role Role
	err := repository.DB().WithContext(ctx).Where("id =?", id).Preload("menus").First(&role).Error
	return &role, err
}

func CreateUserRole(ctx context.Context, user *UserRole) error {
	return repository.DB().WithContext(ctx).Create(user).Error
}

func ListRole(ctx context.Context, query RoleQuery) ([]Role, int64, error) {
	var roles []Role
	var total int64
	d := repository.DB().WithContext(ctx)
	if query.RoleName != "" {
		d = d.Where("role_name like ?", "%"+query.RoleName+"%")
	}
	err := d.Scopes(db.Page(query.Page, query.PageNum, query.PageSize, &total)).
		Preload("menus").
		Find(&roles).
		Error
	return roles, total, err
}
