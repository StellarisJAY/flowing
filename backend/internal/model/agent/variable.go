package agent

type VariableType string

const (
	VariableTypeText   VariableType = "text"
	VariableTypeNumber VariableType = "number"
	VariableTypeSelect VariableType = "select"
	VariableTypeSwitch VariableType = "switch"
)

type Variable struct {
	Name        string   `json:"name"`        // 名称
	DisplayName string   `json:"displayName"` // 显示名称
	Type        string   `json:"type"`        // 类型
	Remark      string   `json:"remark"`      // 备注
	Options     []string `json:"options"`     // 选项
	Default     string   `json:"default"`     // 默认值
}
