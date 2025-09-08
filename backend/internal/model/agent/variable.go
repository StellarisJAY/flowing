package agent

type VariableType string

const (
	VariableTypeText   VariableType = "string"
	VariableTypeNumber VariableType = "number"
	VariableTypeSelect VariableType = "select"
	VariableTypeSwitch VariableType = "switch"
)

type Variable struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`        // 名称
	Description string       `json:"description"` // 描述
	Type        VariableType `json:"type"`        // 类型
	Fixed       bool         `json:"fixed"`       // 是否固定
}
