package system

import (
	"context"
	"flowing/global"
	"flowing/internal/model/common"
	sysmodel "flowing/internal/model/system"
	"flowing/internal/repository"
	"strconv"
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
	return roles, total, nil
}

func CreateUserRole(ctx context.Context, req sysmodel.CreateUserRoleReq) error {
	return sysmodel.CreateUserRole(ctx, &sysmodel.UserRole{
		UserId: req.UserId,
		RoleId: req.RoleId,
	})
}

func SaveRoleMenus(ctx context.Context, req sysmodel.SaveRoleMenuReq) error {
	roleMenus := make([]*sysmodel.RoleMenu, len(req.MenuIds))
	for i, menuId := range req.MenuIds {
		id, err := strconv.ParseInt(menuId, 10, 64)
		if err != nil {
			return global.NewError(500, "修改角色菜单失败", err)
		}
		roleMenus[i] = &sysmodel.RoleMenu{
			RoleId: req.RoleId,
			MenuId: id,
		}
	}
	return repository.Tx(ctx, func(c context.Context) error {
		err := repository.DB(c).Delete(&sysmodel.RoleMenu{}, "role_id = ?", req.RoleId).Error
		if err != nil {
			return global.NewError(500, "修改角色菜单失败", err)
		}
		err = repository.DB(c).Save(roleMenus).Error
		if err != nil {
			return global.NewError(500, "修改角色菜单失败", err)
		}
		return nil
	})
}

func UpdateRole(ctx context.Context, req sysmodel.UpdateRoleReq) error {
	role := sysmodel.Role{
		BaseModel:   common.BaseModel{Id: req.Id},
		RoleName:    req.RoleName,
		RoleKey:     req.RoleKey,
		Description: req.Description,
	}
	return sysmodel.UpdateRole(ctx, role)
}

func DeleteRole(ctx context.Context, id int64) error {
	return repository.Tx(ctx, func(c context.Context) error {
		// 删除角色
		if err := repository.DB(c).Delete(&sysmodel.Role{}, "id = ?", id).Error; err != nil {
			return global.NewError(500, "删除角色失败", err)
		}
		// 删除角色-用户关联
		if err := repository.DB(c).Delete(&sysmodel.UserRole{}, "role_id = ?", id).Error; err != nil {
			return global.NewError(500, "删除角色失败", err)
		}
		// 删除角色-菜单关联
		if err := repository.DB(c).Delete(&sysmodel.RoleMenu{}, "role_id = ?", id).Error; err != nil {
			return global.NewError(500, "删除角色失败", err)
		}
		return nil
	})
}

func GetRoleMenus(ctx context.Context, roleId int64) (*sysmodel.RoleMenuResp, error) {
	var roleMenus []*sysmodel.RoleMenu
	err := repository.DB(ctx).Model(&sysmodel.RoleMenu{}).Where("role_id = ?", roleId).Find(&roleMenus).Error
	if err != nil {
		return nil, global.NewError(500, "获取角色菜单失败", err)
	}

	checkedKeys := make([]string, len(roleMenus))
	for i, menu := range roleMenus {
		checkedKeys[i] = strconv.FormatInt(menu.MenuId, 10)
	}

	// 获取系统所有菜单
	allMenus, err := sysmodel.ListMenu(ctx, sysmodel.MenuQuery{})
	if err != nil {
		return nil, global.NewError(500, "获取菜单失败", err)
	}
	menus := make([]*sysmodel.RoleMenuOption, len(allMenus))
	for i, menu := range allMenus {
		menus[i] = &sysmodel.RoleMenuOption{
			RoleId:   roleId,
			Name:     menu.MenuName,
			Key:      menu.Id,
			ParentId: *menu.ParentId,
		}
	}
	// 构建菜单树
	allMenuMap := make(map[int64]*sysmodel.RoleMenuOption)
	for _, menu := range menus {
		allMenuMap[menu.Key] = menu
	}
	result := make([]*sysmodel.RoleMenuOption, 0, len(menus))
	for _, menu := range menus {
		if menu.ParentId == 0 {
			result = append(result, menu)
			continue
		}
		parent, ok := allMenuMap[menu.ParentId]
		if !ok {
			result = append(result, menu)
		} else {
			parent.Children = append(parent.Children, menu)
		}
	}
	return &sysmodel.RoleMenuResp{
		Menus:       result,
		CheckedKeys: checkedKeys,
	}, nil
}
