package system

import "flowing/internal/model/common"

const (
	MenuTypeDefault int = 1 + iota
	MenuTypePage
	MenuTypeButton
)

type Menu struct {
	common.BaseModel
	MenuName  string `json:"menuName" gorm:"column:menu_name;type:varchar(50);not null;"`
	Type      int    `json:"type" gorm:"column:type;type:int;not null;"`
	Path      string `json:"path" gorm:"column:path;type:varchar(255);not null;"`
	Component string `json:"component" gorm:"column:component;type:varchar(255);not null;"`
	ParentId  int    `json:"parentId" gorm:"column:parent_id;type:int;not null;"`
	OrderNum  int    `json:"orderNum" gorm:"column:order_num;type:int;not null;"`
	Status    int    `json:"status" gorm:"column:status;type:int;default:1;"`
}

func (m *Menu) TableName() string {
	return "sys_menu"
}

type RoleMenu struct {
	common.BaseModel
	RoleId int64 `json:"role_id" gorm:"column:role_id;type:int;not null;"`
	MenuId int64 `json:"menu_id" gorm:"column:menu_id;type:int;not null;"`
}

func (m *RoleMenu) TableName() string {
	return "sys_role_menu"
}
