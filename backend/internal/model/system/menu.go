package system

import (
	"context"
	"flowing/internal/model/common"
	"flowing/internal/repository"
)

const (
	MenuTypeDefault int = 1 + iota
	MenuTypePage
	MenuTypeButton
)

type Menu struct {
	common.BaseModel
	MenuName  string  `json:"menuName" gorm:"column:menu_name;type:varchar(50);not null;"`
	Type      int     `json:"type" gorm:"column:type;type:int;not null;"`
	Path      string  `json:"path" gorm:"column:path;type:varchar(255);not null;unique"`
	Component string  `json:"component" gorm:"column:component;type:varchar(255);not null;"`
	ParentId  int64   `json:"parentId" gorm:"column:parent_id;type:int;not null;"`
	OrderNum  int     `json:"orderNum" gorm:"column:order_num;type:int;not null;"`
	Status    int     `json:"status" gorm:"column:status;type:int;default:1;"`
	Children  []*Menu `json:"children" gorm:"-"`
}

func (m *Menu) TableName() string {
	return "sys_menu"
}

type RoleMenu struct {
	common.BaseModel
	RoleId int64 `json:"role_id" gorm:"column:role_id;type:int;not null; unique: role_menu_index"`
	MenuId int64 `json:"menu_id" gorm:"column:menu_id;type:int;not null; unique: role_menu_index"`
}

func (m *RoleMenu) TableName() string {
	return "sys_role_menu"
}

type MenuQuery struct {
	MenuName string `json:"menuName" form:"menuName"`
}

type CreateMenuReq struct {
	MenuName  string `json:"menuName" binding:"required"`
	Type      int    `json:"type" binding:"required"`
	Path      string `json:"path" binding:"required"`
	Component string `json:"component"`
	ParentId  int64  `json:"parentId"`
	OrderNum  int    `json:"orderNum"`
}

type CreateRoleMenuReq struct {
	RoleId int64 `json:"roleId,string" binding:"required"`
	MenuId int64 `json:"menuId,string" binding:"required"`
}

func CreateMenu(ctx context.Context, menu *Menu) error {
	return repository.DB().WithContext(ctx).Create(menu).Error
}

func GetMenu(ctx context.Context, menuId int64) (*Menu, error) {
	var menu Menu
	if err := repository.DB().WithContext(ctx).First(&menu, menuId).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func ListMenu(ctx context.Context, query MenuQuery) ([]*Menu, error) {
	var menus []*Menu
	db := repository.DB().WithContext(ctx).Model(&Menu{})
	if query.MenuName != "" {
		db = db.Where("menu_name LIKE ?", "%"+query.MenuName+"%")
	}
	if err := db.Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

func CreateRoleMenu(ctx context.Context, roleId, menuId int64) error {
	return repository.DB().WithContext(ctx).Create(&RoleMenu{
		RoleId: roleId,
		MenuId: menuId,
	}).Error
}

func GetUserMenus(ctx context.Context, userId int64) ([]*Menu, error) {
	var menus []*Menu
	if err := repository.DB().WithContext(ctx).Table("sys_menu").
		Select("sys_menu.*").
		Joins("INNER JOIN sys_role_menu ON sys_menu.id = sys_role_menu.menu_id").
		Joins("INNER JOIN sys_user_role ON sys_user_role.role_id = sys_role_menu.role_id").
		Where("sys_user_role.user_id = ?", userId).
		Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}
