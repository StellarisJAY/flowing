package system

import "slices"

type UserPermission struct {
	Menus   []*Menu  `json:"menus"`
	Actions []string `json:"actions"`
}

func (u *UserPermission) HasPermission(code string) bool {
	return slices.Contains(u.Actions, code)
}
