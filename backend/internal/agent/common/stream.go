package common

import (
	"errors"
	"flowing/internal/model/chat"
	"flowing/internal/model/common"
	"io"

	"github.com/cloudwego/eino/schema"
)

type StreamConsumerMeta struct {
	ConversationId int64
	AgentId        int64
	AgentRunId     int64
	MessageId      int64
}

func StreamConsumer(stream *schema.StreamReader[*schema.Message], messageChan chan chat.Message, meta StreamConsumerMeta) error {
	defer stream.Close()
	// 模型输出
	for {
		chunk, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}
		response := chat.Message{
			BaseModel:       common.BaseModel{Id: meta.MessageId},
			ConversationId:  meta.ConversationId,
			Content:         chunk.Content,
			Type:            chat.MessageTypeAssistant,
			AgentId:         meta.AgentId,
			AgentRunId:      meta.AgentRunId,
			ThinkingContent: chunk.ReasoningContent,
		}
		messageChan <- response
	}
	return nil
}
