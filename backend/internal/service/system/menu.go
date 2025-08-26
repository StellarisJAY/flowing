package system

import (
	"context"
	"errors"
	"flowing/global"
	"flowing/internal/model/common"
	sysmodel "flowing/internal/model/system"
	"flowing/internal/repository"
)

func CreateMenu(ctx context.Context, menu sysmodel.CreateMenuReq) error {
	if menu.ParentId != 0 {
		if _, err := sysmodel.GetMenu(ctx, menu.ParentId); err != nil {
			return errors.New("父级菜单不存在")
		}
	}
	if menu.Type != sysmodel.MenuTypeButton {
		var count int64
		repository.DB(ctx).Model(&sysmodel.Menu{}).Where("path = ?", menu.Path).Count(&count)
		if count > 0 {
			return global.NewError(500, "路径已存在", nil)
		}
	}
	model := &sysmodel.Menu{
		BaseModel: common.BaseModel{
			CreateBy: ctx.Value(global.ContextKeyUser).(sysmodel.User).Id,
		},
		MenuName:   menu.MenuName,
		Type:       menu.Type,
		Path:       menu.Path,
		Component:  menu.Component,
		ParentId:   &menu.ParentId,
		OrderNum:   menu.OrderNum,
		ActionCode: menu.ActionCode,
		ShowInNav:  menu.ShowInNav,
		HideTab:    menu.HideTab,
		Icon:       menu.Icon,
	}
	return sysmodel.CreateMenu(ctx, model)
}

func buildMenuTree(menus []*sysmodel.Menu, excludeButtons bool) []*sysmodel.Menu {
	menuMap := make(map[int64]*sysmodel.Menu)
	for _, menu := range menus {
		menuMap[menu.Id] = menu
		menu.Key = menu.Id
	}
	var rootMenus []*sysmodel.Menu
	for _, menu := range menus {
		if excludeButtons && menu.Type == sysmodel.MenuTypeButton {
			continue
		}
		if *menu.ParentId == 0 {
			rootMenus = append(rootMenus, menu)
			continue
		}
		parentMenu, ok := menuMap[*menu.ParentId]
		if ok {
			parentMenu.Children = append(parentMenu.Children, menu)
		} else {
			rootMenus = append(rootMenus, menu)
		}
	}
	return rootMenus
}

func getAllChildMenuIds(menus []*sysmodel.Menu, parentId int64) []int64 {
	var childIds []int64
	for _, menu := range menus {
		if *menu.ParentId == parentId {
			childIds = append(childIds, menu.Id)
			childIds = append(childIds, getAllChildMenuIds(menus, menu.Id)...)
		}
	}
	return childIds
}

func ListMenuTree(ctx context.Context, query sysmodel.MenuQuery) ([]*sysmodel.Menu, error) {
	menus, err := sysmodel.ListMenu(ctx, query)
	if err != nil {
		return nil, err
	}
	return buildMenuTree(menus, false), nil
}

func UpdateMenu(ctx context.Context, req sysmodel.UpdateMenuReq) error {
	if req.ParentId != 0 {
		if _, err := sysmodel.GetMenu(ctx, req.ParentId); err != nil {
			return errors.New("父级菜单不存在")
		}
	}
	return sysmodel.UpdateMenu(ctx, sysmodel.Menu{
		BaseModel:  common.BaseModel{Id: req.Id},
		MenuName:   req.MenuName,
		Type:       req.Type,
		Path:       req.Path,
		Component:  req.Component,
		ParentId:   &req.ParentId,
		OrderNum:   req.OrderNum,
		ActionCode: req.ActionCode,
		ShowInNav:  req.ShowInNav,
		HideTab:    req.HideTab,
		Icon:       req.Icon,
	})
}

func DeleteMenu(ctx context.Context, id int64) error {
	return repository.Tx(ctx, func(c context.Context) error {
		// 获取所有菜单
		menus, err := sysmodel.ListMenu(c, sysmodel.MenuQuery{})
		if err != nil {
			return global.NewError(500, "删除菜单失败", err)
		}
		// 找到所有子菜单
		childIds := getAllChildMenuIds(menus, id)
		menuIds := append(childIds, id)
		// 删除菜单和子菜单
		if err := sysmodel.BatchDeleteMenu(c, menuIds); err != nil {
			return global.NewError(500, "删除菜单失败", err)
		}
		// 删除角色-菜单关联
		if err := repository.DB(c).Delete(&sysmodel.RoleMenu{}, "menu_id IN ?", menuIds).Error; err != nil {
			return global.NewError(500, "删除菜单失败", err)
		}
		return nil
	})
}
