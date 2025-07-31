package system

import (
	"context"
	"flowing/global"
	sysmodel "flowing/internal/model/system"
)

func GetUserAllPermissions(ctx context.Context, userId int64) (*sysmodel.UserPermission, error) {
	menus, err := sysmodel.GetUserMenus(ctx, userId)
	if err != nil {
		return nil, global.NewError(500, "获取菜单失败", err)
	}
	menuTree := buildMenuTree(menus, true)
	var actions []string
	for _, menu := range menus {
		if menu.Type == sysmodel.MenuTypeButton {
			actions = append(actions, *menu.ActionCode)
		}
	}
	return &sysmodel.UserPermission{
		Menus:   menuTree,
		Actions: actions,
	}, nil
}
