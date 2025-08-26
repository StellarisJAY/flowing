package messagehub

import (
	"context"
	"encoding/json"
	"flowing/global"
	"flowing/internal/model/agent"
	"flowing/internal/model/ai"
	"flowing/internal/model/chat"
	"flowing/internal/model/common"
	"flowing/internal/repository"
	"fmt"
	"github.com/cloudwego/eino/schema"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func MockHandleSendMessage(ctx context.Context, req chat.SendMessageReq) error {
	writer := ctx.Value(global.ContextKeySSEWriter).(gin.ResponseWriter)
	messageId := repository.Snowflake().Generate().Int64()
	conversationId := repository.Snowflake().Generate().Int64()
	userMessage := chat.Message{
		BaseModel:      common.BaseModel{Id: messageId},
		ConversationId: conversationId,
		Content:        req.Content,
		Type:           chat.MessageTypeUser,
	}

	messages := []chat.Message{
		userMessage,
		{
			BaseModel:      common.BaseModel{Id: messageId + 1},
			ConversationId: conversationId,
			Content:        "你好",
			Type:           chat.MessageTypeAssistant,
		},
		{
			BaseModel:      common.BaseModel{Id: messageId + 1},
			ConversationId: conversationId,
			Content:        "，我是",
			Type:           chat.MessageTypeAssistant,
		},
		{
			BaseModel:      common.BaseModel{Id: messageId + 1},
			ConversationId: conversationId,
			Content:        "flowing",
			Type:           chat.MessageTypeAssistant,
		},
		{
			BaseModel:      common.BaseModel{Id: messageId + 1},
			ConversationId: conversationId,
			Content:        " AI助手，",
			Type:           chat.MessageTypeAssistant,
		},
		{
			BaseModel:      common.BaseModel{Id: messageId + 1},
			ConversationId: conversationId,
			Content:        "有什么",
			Type:           chat.MessageTypeAssistant,
		},
		{
			BaseModel:      common.BaseModel{Id: messageId + 1},
			ConversationId: conversationId,
			Content:        "可以",
			Type:           chat.MessageTypeAssistant,
		},
		{
			BaseModel:      common.BaseModel{Id: messageId + 1},
			ConversationId: conversationId,
			Content:        "帮助",
			Type:           chat.MessageTypeAssistant,
		},
		{
			BaseModel:      common.BaseModel{Id: messageId + 1},
			ConversationId: conversationId,
			Content:        "你",
			Type:           chat.MessageTypeAssistant,
		},
		{
			BaseModel:      common.BaseModel{Id: messageId + 1},
			ConversationId: conversationId,
			Content:        "的吗？",
			Type:           chat.MessageTypeAssistant,
		},
	}

	// 模拟消息发送
	for _, msg := range messages {
		msgData, _ := json.Marshal(msg)
		_, err := fmt.Fprintf(writer, "data: %s\n\n", string(msgData))
		if err != nil {
			return global.NewError(500, "send message failed", err)
		}
		writer.Flush()
		time.Sleep(time.Millisecond * 200)
	}
	return nil
}

func HandleSendMessage(ctx context.Context, req chat.SendMessageReq) error {
	switch req.Mode {
	case chat.SendMessageModeDebugger:
		if req.AgentType == "simple" {
			return handleDebuggerModeSimple(ctx, req)
		}
		return global.NewError(500, "暂不支持工作流智能体", nil)
	default:
		return global.NewError(500, "暂不支持的聊天模式", nil)
	}
}

func handleDebuggerModeSimple(ctx context.Context, req chat.SendMessageReq) error {
	writer := ctx.Value(global.ContextKeySSEWriter).(gin.ResponseWriter)

	messageId := repository.Snowflake().Generate().Int64()
	if req.ConversationId == 0 {
		req.ConversationId = repository.Snowflake().Generate().Int64()
	}
	userMessage := chat.Message{
		BaseModel:      common.BaseModel{Id: messageId},
		ConversationId: req.ConversationId,
		Content:        req.Content,
		Type:           chat.MessageTypeUser,
	}

	var agentConfig agent.SimpleAgentConfig
	err := json.Unmarshal([]byte(req.AgentConfig), &agentConfig)
	if err != nil {
		return global.NewError(500, "无效的智能体配置", err)
	}

	pm, err := ai.GetProviderModelDetail(ctx, agentConfig.ModelId)
	if err != nil {
		return global.NewError(500, "无效的模型ID", err)
	}

	chatModel, err := GetChatModel(ctx, *pm)
	if err != nil {
		return global.NewError(500, "获取模型失败", err)
	}

	stream, err := chatModel.Stream(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: req.Content,
		},
	})
	if err != nil {
		return global.NewError(500, "模型调用失败", err)
	}
	defer stream.Close()

	// 回送用户消息
	userMessageData, _ := json.Marshal(userMessage)
	_, err = fmt.Fprintf(writer, "data: %s\n\n", string(userMessageData))
	if err != nil {
		return global.NewError(500, "send message failed", err)
	}
	writer.Flush()

	// 助手消息
	messageId = repository.Snowflake().Generate().Int64()
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return global.NewError(500, "模型调用失败", err)
		}
		message := chat.Message{
			BaseModel:      common.BaseModel{Id: messageId},
			ConversationId: req.ConversationId,
			Content:        chunk.Content,
			Type:           chat.MessageTypeAssistant,
		}
		messageData, _ := json.Marshal(message)
		_, err = fmt.Fprintf(writer, "data: %s\n\n", string(messageData))
		if err != nil {
			return global.NewError(500, "send message failed", err)
		}
		writer.Flush()
	}

	return nil
}
