package chat

import "flowing/internal/model/common"

type SessionType string

const (
	SessionTypeChat  SessionType = "agent"
	SessionTypeAgent SessionType = "agent"
)

type Session struct {
	common.BaseModel
	Title      string      `json:"title" gorm:"column:title;type:varchar(255);not null"`            // 会话标题
	Type       SessionType `json:"type" gorm:"column:type;type:varchar(16);not null;"`              // 会话类型
	AgentId    *int64      `json:"agentId" gorm:"column:agent_id;type:bigint;default:null;"`        // 智能体ID
	AgentRunId *int64      `json:"agentRunId" gorm:"column:agent_run_id;type:bigint;default:null;"` // 智能体运行ID
	ModelId    *int64      `json:"modelId" gorm:"column:model_id;type:bigint;default:null;"`        // 模型ID
}

func (s *Session) TableName() string {
	return "chat_session"
}
