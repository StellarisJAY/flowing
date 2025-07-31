package system

type UserPermission struct {
	Menus   []*Menu  `json:"menus"`
	Actions []string `json:"actions"`
}
