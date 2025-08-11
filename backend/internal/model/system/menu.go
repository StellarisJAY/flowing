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
	MenuName   string  `json:"menuName" gorm:"column:menu_name;type:varchar(50);not null;"`
	Type       int     `json:"type" gorm:"column:type;type:int;not null;"`
	Path       string  `json:"path" gorm:"column:path;type:varchar(255);not null;unique"`
	Component  string  `json:"component" gorm:"column:component;type:varchar(255);not null;"`
	ParentId   int64   `json:"parentId,string" gorm:"column:parent_id;type:int;not null;"`
	OrderNum   int     `json:"orderNum" gorm:"column:order_num;type:int;not null;"`
	Status     int     `json:"status" gorm:"column:status;type:int;default:1;"`
	ActionCode *string `json:"actionCode" gorm:"column:action_code;type:varchar(50);unique;default:null"` // 权限标识
	Children   []*Menu `json:"children" gorm:"-"`
	Key        int64   `json:"key,string" gorm:"-"`
}

func (m *Menu) TableName() string {
	return "sys_menu"
}

type RoleMenu struct {
	common.BaseModel
	RoleId int64 `json:"role_id" gorm:"column:role_id;type:int;not null; uniqueIndex: role_menu_index"`
	MenuId int64 `json:"menu_id" gorm:"column:menu_id;type:int;not null; uniqueIndex: role_menu_index"`
}

func (m *RoleMenu) TableName() string {
	return "sys_role_menu"
}

type RoleGrantedMenu struct {
	RoleMenuId int64 `json:"role_menu_id" gorm:"column:role_menu_id;"`
	Menu
}

type MenuQuery struct {
	MenuName string `json:"menuName" form:"menuName"`
}

type CreateMenuReq struct {
	MenuName   string  `json:"menuName" binding:"required"`
	Type       int     `json:"type" binding:"required"`
	Path       string  `json:"path" binding:"required"`
	Component  string  `json:"component"`
	ParentId   int64   `json:"parentId,string"`
	OrderNum   int     `json:"orderNum"`
	ActionCode *string `json:"actionCode"`
}

type UpdateMenuReq struct {
	Id         int64   `json:"id,string" binding:"required"`
	MenuName   string  `json:"menuName" binding:"required"`
	Type       int     `json:"type" binding:"required"`
	Path       string  `json:"path" binding:"required"`
	Component  string  `json:"component"`
	ParentId   int64   `json:"parentId,string"`
	OrderNum   int     `json:"orderNum"`
	ActionCode *string `json:"actionCode"`
}

type CreateRoleMenuReq struct {
	RoleId  int64   `json:"roleId,string" binding:"required"`
	MenuIds []int64 `json:"menuIds" binding:"required"`
}

type SaveRoleMenuReq struct {
	RoleId     int64   `json:"roleId,string" binding:"required"`
	NewMenuIds []int64 `json:"newMenuIds" binding:"required"`
	OldMenuIds []int64 `json:"oldMenuIds" binding:"required"`
}

func CreateMenu(ctx context.Context, menu *Menu) error {
	return repository.DB(ctx).Create(menu).Error
}

func GetMenu(ctx context.Context, menuId int64) (*Menu, error) {
	var menu Menu
	if err := repository.DB(ctx).First(&menu, menuId).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func ListMenu(ctx context.Context, query MenuQuery) ([]*Menu, error) {
	var menus []*Menu
	db := repository.DB(ctx).Model(&Menu{})
	if query.MenuName != "" {
		db = db.Where("menu_name LIKE ?", "%"+query.MenuName+"%")
	}
	if err := db.Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

func CreateRoleMenu(ctx context.Context, rms []RoleMenu) error {
	return repository.DB(ctx).CreateInBatches(rms, 64).Error
}

func GetUserMenus(ctx context.Context, userId int64) ([]*Menu, error) {
	var menus []*Menu
	if err := repository.DB(ctx).Table("sys_menu").
		Select("sys_menu.*").
		Joins("INNER JOIN sys_role_menu ON sys_menu.id = sys_role_menu.menu_id").
		Joins("INNER JOIN sys_user_role ON sys_user_role.role_id = sys_role_menu.role_id").
		Where("sys_user_role.user_id = ?", userId).
		Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

func UpdateMenu(ctx context.Context, menu Menu) error {
	return repository.DB(ctx).Model(&Menu{}).Where("id = ?", menu.Id).Updates(menu).Error
}

func BatchDeleteMenu(ctx context.Context, menuIds []int64) error {
	return repository.DB(ctx).Delete(&Menu{}, "id IN ?", menuIds).Error
}
