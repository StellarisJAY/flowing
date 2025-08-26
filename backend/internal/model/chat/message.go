package chat

import "flowing/internal/model/common"

type MessageType string

const (
	MessageTypeUser      MessageType = "user"      // 用户
	MessageTypeAssistant MessageType = "assistant" // 助手
	MessageTypeLog       MessageType = "log"       // 日志
)

type Message struct {
	common.BaseModel
	ConversationId int64       `json:"conversationId,string" gorm:"type:bigint unsigned auto_increment;default:0"`
	Content        string      `json:"content" gorm:"column:content;type:text;not null;"`
	Type           MessageType `json:"type" gorm:"column:type;type:varchar(16);not null;"`
	AgentId        int64       `json:"agentId,string" gorm:"type:bigint;default:0"`
	AgentRunId     int64       `json:"agentRunId,string" gorm:"type:bigint;default:0"`
}

func (m *Message) TableName() string {
	return "chat_message"
}

type SendMessageMode string

const (
	SendMessageModeAgent    SendMessageMode = "agent"    // 智能体聊天模式
	SendMessageModeChat     SendMessageMode = "chat"     // 模型聊天模式
	SendMessageModeDebugger SendMessageMode = "debugger" // 调试模式
)

type SendMessageReq struct {
	ConversationId int64           `json:"conversationId,string"`
	Content        string          `json:"content"`
	AgentId        int64           `json:"agentId,string"`
	AgentConfig    string          `json:"agentConfig"`
	AgentType      string          `json:"agentType"`
	Mode           SendMessageMode `json:"mode"`
	Files          []string        `json:"files"`
	Variables      string          `json:"variables"`
}
