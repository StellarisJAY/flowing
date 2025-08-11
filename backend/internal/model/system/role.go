package system

import (
	"context"
	"flowing/internal/model/common"
	"flowing/internal/repository"
	"flowing/internal/repository/db"
)

type Role struct {
	common.BaseModel
	RoleName    string  `json:"roleName" gorm:"column:role_name;type:varchar(50);not null"`
	RoleKey     string  `json:"roleKey" gorm:"column:role_key;type:varchar(50);not null;unique"`
	Description string  `json:"description" gorm:"column:description;type:varchar(200)"`
	Menus       []*Menu `json:"menus" gorm:"many2many:sys_role_menu;"`
}

func (Role) TableName() string {
	return "sys_role"
}

type UserRole struct {
	common.BaseModel
	UserId int64 `json:"userId" gorm:"column:user_id;type:int;not null;uniqueIndex:idx_user_role_id"`
	RoleId int64 `json:"roleId" gorm:"column:role_id;type:int;not null;uniqueIndex:idx_user_role_id"`
}

func (UserRole) TableName() string {
	return "sys_user_role"
}

type CreateRoleReq struct {
	RoleName    string `json:"roleName" binding:"required"`
	RoleKey     string `json:"roleKey" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateRoleReq struct {
	Id          int64  `json:"id,string" binding:"required"`
	RoleName    string `json:"roleName" binding:"required"`
	RoleKey     string `json:"roleKey" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type RoleQuery struct {
	common.BaseQuery
	RoleName string `json:"roleName" form:"roleName"`
	RoleKey  string `json:"roleKey" form:"roleKey"`
}

type CreateUserRoleReq struct {
	RoleId int64 `json:"roleId,string" binding:"required"`
	UserId int64 `json:"userId,string" binding:"required"`
}

func CreateRole(ctx context.Context, role *Role) error {
	return repository.DB(ctx).Create(role).Error
}

func GetRole(ctx context.Context, id int) (*Role, error) {
	var role Role
	err := repository.DB(ctx).Where("id =?", id).Preload("Menus").First(&role).Error
	return &role, err
}

func CreateUserRole(ctx context.Context, user *UserRole) error {
	return repository.DB(ctx).Create(user).Error
}

func ListRole(ctx context.Context, query RoleQuery) ([]*Role, int64, error) {
	var roles []*Role
	var total int64
	d := repository.DB(ctx).Model(&Role{})
	if query.RoleName != "" {
		d = d.Where("role_name like ?", "%"+query.RoleName+"%")
	}
	if query.RoleKey != "" {
		d = d.Where("role_key like ?", "%"+query.RoleKey+"%")
	}
	if err := d.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := d.Scopes(db.Page(query.Page, query.PageNum, query.PageSize)).
		Preload("Menus").
		Find(&roles).
		Error
	return roles, total, err
}

func UpdateRole(ctx context.Context, role Role) error {
	return repository.DB(ctx).Model(&Role{}).Where("id = ?", role.Id).Updates(role).Error
}
