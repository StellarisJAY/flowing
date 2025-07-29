package system

import (
	"context"
	"errors"
	sysmodel "flowing/internal/model/system"
)

func CreateMenu(ctx context.Context, menu sysmodel.CreateMenuReq) error {
	if menu.ParentId != 0 {
		if _, err := sysmodel.GetMenu(ctx, menu.ParentId); err != nil {
			return errors.New("父级菜单不存在")
		}
	}
	model := &sysmodel.Menu{
		MenuName:  menu.MenuName,
		Type:      menu.Type,
		Path:      menu.Path,
		Component: menu.Component,
		ParentId:  menu.ParentId,
		OrderNum:  menu.OrderNum,
	}
	return sysmodel.CreateMenu(ctx, model)
}

func buildMenuTree(menus []*sysmodel.Menu) []*sysmodel.Menu {
	menuMap := make(map[int64]*sysmodel.Menu)
	for _, menu := range menus {
		menuMap[menu.Id] = menu
	}
	var rootMenus []*sysmodel.Menu
	for _, menu := range menus {
		if menu.ParentId == 0 {
			rootMenus = append(rootMenus, menu)
			continue
		}
		parentMenu, ok := menuMap[menu.ParentId]
		if ok {
			parentMenu.Children = append(parentMenu.Children, menu)
		} else {
			rootMenus = append(rootMenus, menu)
		}
	}
	return rootMenus
}

func ListMenuTree(ctx context.Context, query sysmodel.MenuQuery) ([]*sysmodel.Menu, error) {
	menus, err := sysmodel.ListMenu(ctx, query)
	if err != nil {
		return nil, err
	}
	return buildMenuTree(menus), nil
}
