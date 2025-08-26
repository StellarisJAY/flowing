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
	SessionId int64       `json:"sessionId" gorm:"column:session_id;type:bigint;not null;"` // 会话ID
	Type      MessageType `json:"type" gorm:"column:type;type:varchar(16);not null;"`       // 角色
	Content   string      `json:"content" gorm:"column:content;type:text;not null;"`        // 内容
}

func (m *Message) TableName() string {
	return "chat_message"
}
