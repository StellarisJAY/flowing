package system

import (
	"context"
	sysmodel "flowing/internal/model/system"
)

func CreateRole(ctx context.Context, role sysmodel.CreateRoleReq) error {
	model := sysmodel.Role{
		RoleName:    role.RoleName,
		RoleKey:     role.RoleKey,
		Description: role.Description,
	}
	return sysmodel.CreateRole(ctx, &model)
}

func ListRole(ctx context.Context, query sysmodel.RoleQuery) ([]sysmodel.Role, int64, error) {
	roles, total, err := sysmodel.ListRole(ctx, query)
	if err != nil {
		return nil, 0, err
	}
	for _, role := range roles {
		if len(role.Menus) == 0 {
			continue
		}
		role.Menus = buildMenuTree(role.Menus)
	}
	return roles, total, nil
}
