package common

import (
	"context"
	"encoding/json"

	"github.com/cloudwego/eino/schema"
)

// MergeMapStateAndInput 合并input和state，input的优先级高，用于需要全局变量的流程节点
// 在节点配置中使用WithStatePreHandler合并input和state
func MergeMapStateAndInput(ctx context.Context, input map[string]any, state map[string]any) (map[string]any, error) {
	mergedInput := make(map[string]any)
	for k, v := range input {
		mergedInput[k] = v
	}
	for k, v := range state {
		mergedInput[k] = v
	}
	return mergedInput, nil
}

type DocToMessagesLambda func(context.Context, []*schema.Document) (map[string]any, error)

// DocumentsToMessages 将文档转换为消息，消息角色为system
// 用于知识库节点将文档转为消息，传递给之后的消息节点
func DocumentsToMessages(outputKey string, role schema.RoleType) DocToMessagesLambda {
	return func(ctx context.Context, input []*schema.Document) (map[string]any, error) {
		data, _ := json.Marshal(input)
		messages := []*schema.Message{
			{
				Content: string(data),
				Role:    role,
			},
		}
		return map[string]any{outputKey: messages}, nil
	}
}
