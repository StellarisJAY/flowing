package system

import (
	"context"
	"flowing/global"
	sysmodel "flowing/internal/model/system"
	"flowing/internal/repository"
)

func CreateRole(ctx context.Context, role sysmodel.CreateRoleReq) error {
	model := sysmodel.Role{
		RoleName:    role.RoleName,
		RoleKey:     role.RoleKey,
		Description: role.Description,
	}
	return sysmodel.CreateRole(ctx, &model)
}

func ListRole(ctx context.Context, query sysmodel.RoleQuery) ([]*sysmodel.Role, int64, error) {
	roles, total, err := sysmodel.ListRole(ctx, query)
	if err != nil {
		return nil, 0, err
	}
	for _, role := range roles {
		if len(role.Menus) == 0 {
			continue
		}
		role.Menus = buildMenuTree(role.Menus, false)
	}
	return roles, total, nil
}

func CreateUserRole(ctx context.Context, req sysmodel.CreateUserRoleReq) error {
	return sysmodel.CreateUserRole(ctx, &sysmodel.UserRole{
		UserId: req.UserId,
		RoleId: req.RoleId,
	})
}

func CreateRoleMenu(ctx context.Context, req sysmodel.CreateRoleMenuReq) error {
	rms := make([]sysmodel.RoleMenu, 0, len(req.MenuIds))
	for _, menuId := range req.MenuIds {
		rms = append(rms, sysmodel.RoleMenu{
			RoleId: req.RoleId,
			MenuId: menuId,
		})
	}
	return sysmodel.CreateRoleMenu(ctx, rms)
}

func SaveRoleMenus(ctx context.Context, req sysmodel.SaveRoleMenuReq) error {
	oldMenus := make(map[int64]struct{})
	for _, menuId := range req.OldMenuIds {
		oldMenus[menuId] = struct{}{}
	}
	newMenus := make(map[int64]struct{})
	for _, menuId := range req.NewMenuIds {
		newMenus[menuId] = struct{}{}
	}
	toCreate := make([]sysmodel.RoleMenu, 0)
	for _, menuId := range req.NewMenuIds {
		if _, ok := oldMenus[menuId]; !ok {
			toCreate = append(toCreate, sysmodel.RoleMenu{
				RoleId: req.RoleId,
				MenuId: menuId,
			})
		}
	}
	toDelete := make([]int64, 0)
	for _, menuId := range req.OldMenuIds {
		if _, ok := newMenus[menuId]; !ok {
			toDelete = append(toDelete, menuId)
		}
	}

	err := repository.DB().WithContext(ctx).Delete(sysmodel.RoleMenu{}, "role_id = ? and menu_id in ?", req.RoleId, toDelete).Error
	if err != nil {
		return global.NewError(500, "删除角色菜单失败", err)
	}
	err = repository.DB().WithContext(ctx).CreateInBatches(toCreate, 64).Error
	if err != nil {
		return global.NewError(500, "新增角色菜单失败", err)
	}
	return nil
}
