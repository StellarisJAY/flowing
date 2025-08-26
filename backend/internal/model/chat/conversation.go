package chat

import "flowing/internal/model/common"

type Conversation struct {
	common.BaseModel
	Title      string `json:"title" gorm:"column:title;type:varchar(255);not null;default:''"`
	AgentId    int64  `json:"agentId,string" gorm:"column:agent_id;type:bigint;not null;default:0"`
	AgentRunId int64  `json:"agentRunId,string" gorm:"column:agent_run_id;type:bigint;not null;default:0"`
}

func (Conversation) TableName() string {
	return "chat_conversation"
}
