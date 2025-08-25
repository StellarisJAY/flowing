package agent

// SimpleAgentConfig 简单智能体配置
type SimpleAgentConfig struct {
	ModelId         int64           `json:"modelId,string"`         // 模型ID
	Prompt          string          `json:"prompt"`                 // 提示词
	KnowledgeBaseId int64           `json:"knowledgeBaseId,string"` // 知识库ID
	Variables       []Variable      `json:"variables"`              // 变量
	KbSearchOption  *KbSearchOption `json:"kbSearchOption"`         // 知识库搜索选项
}

type KbSearchOption struct {
	TopK        int     `json:"topK"`
	Threshold   float64 `json:"threshold,string"`
	SearchType  string  `json:"searchType"`
	HybridType  string  `json:"hybridType"`
	Weight      float64 `json:"weight,string"`
	RerankModel int64   `json:"rerankModel,string"`
}

func DefaultSimpleAgentConfig() SimpleAgentConfig {
	return SimpleAgentConfig{
		ModelId:         0,
		Prompt:          "你是一个聊天助手",
		KnowledgeBaseId: 0,
		Variables:       []Variable{},
		KbSearchOption:  nil,
	}
}
