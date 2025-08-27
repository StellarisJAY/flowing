package agent

import (
	"context"
	"errors"
	"flowing/internal/agent/simple"
	"flowing/internal/model/agent"
	"flowing/internal/model/chat"
)

type Run interface {
	Run(ctx context.Context, input chat.Message)
	Receive() (chat.Message, error)
}

func NewAgentRun(ctx context.Context, a *agent.Agent, config any, conversationId int64) (Run, error) {
	switch a.Type {
	case agent.TypeSimple:
		return simple.NewAgentRun(ctx, a, config.(agent.SimpleAgentConfig), conversationId)
	default:
		return nil, errors.New("暂不支持的智能体类型")
	}
}
