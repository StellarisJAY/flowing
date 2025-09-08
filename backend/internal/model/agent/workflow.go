package agent

type NodeType string

const (
	NodeTypeStart     NodeType = "start"
	NodeTypeModel     NodeType = "model"
	NodeTypeReply     NodeType = "reply"
	NodeTypeKnowledge NodeType = "knowledge"
)

var InitialWorkflowConfig = &WorkflowConfig{
	Edges: []*Edge{},
	Nodes: []*Node{
		{
			Id:   "start",
			Type: "base",
			Data: NodeData{
				Type:        NodeTypeStart,
				Label:       "开始",
				Description: "开始节点",
				Config: &NodeConfig{
					Output: []*Variable{
						{
							Id:          "sys.query",
							Name:        "query",
							Description: "用户聊天输入",
							Type:        VariableTypeText,
							Fixed:       true,
						},
					},
				},
			},
			Position: struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			}{
				X: 0,
				Y: 0,
			},
		},
	},
	Viewport: struct {
		X    float64 `json:"x"`
		Y    float64 `json:"y"`
		Zoom float64 `json:"zoom"`
	}{
		X:    0,
		Y:    0,
		Zoom: 1,
	},
}

type WorkflowConfig struct {
	Edges    []*Edge `json:"edges"`
	Nodes    []*Node `json:"nodes"`
	Viewport struct {
		X    float64 `json:"x"`
		Y    float64 `json:"y"`
		Zoom float64 `json:"zoom"`
	} `json:"viewport"`
}

type Edge struct {
	Id           string `json:"id"`
	Source       string `json:"source"`
	Target       string `json:"target"`
	SourceHandle string `json:"sourceHandle"`
	TargetHandle string `json:"targetHandle"`
}

type Node struct {
	Id       string   `json:"id"`
	Type     string   `json:"type"`
	Data     NodeData `json:"data"`
	Position struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"position"`
}

type NodeData struct {
	Type        NodeType    `json:"type"`
	Label       string      `json:"label"`
	Description string      `json:"description"`
	Config      *NodeConfig `json:"config"`
}

type NodeConfig struct {
	Output []*Variable `json:"output"`
	Input  []*Variable `json:"input"`
	*ModelNodeConfig
}

type ModelNodeConfig struct {
	ModelId int64  `json:"modelId,string"`
	Prompt  string `json:"prompt"`
}
