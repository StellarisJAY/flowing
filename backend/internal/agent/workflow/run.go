package workflow

import (
	"context"
	"flowing/internal/dag"
	"flowing/internal/model/agent"
	"flowing/internal/model/chat"
	"io"
)

// AgentRun 简单智能体运行
type AgentRun struct {
	Id            int64                 // 智能体运行id，如果是调试模式id为0
	AgentConfig   *agent.WorkflowConfig // 智能体配置
	Agent         *agent.Agent          // 智能体信息，调试模式为空
	configNodeMap map[string]*agent.Node
	configEdgeMap map[string]*agent.Edge

	conversationId int64             // 对话id
	messageChan    chan chat.Message // 消息通道，消息层通过该channel接受模型的回复
	errorChan      chan error        // 错误通道，智能体通过该channel返回错误
}

func NewAgentRun(ctx context.Context, a *agent.Agent, config *agent.WorkflowConfig, conversationId int64) (*AgentRun, error) {
	agentRun := &AgentRun{
		AgentConfig:    config,
		Agent:          a,
		conversationId: conversationId,
		configNodeMap:  make(map[string]*agent.Node),
		configEdgeMap:  make(map[string]*agent.Edge),
		messageChan:    make(chan chat.Message, 10), // TODO 消息通道大小
		errorChan:      make(chan error),
	}
	return agentRun, nil
}

// Run 运行简单智能体，单独开启goroutine运行
func (a *AgentRun) Run(ctx context.Context, input chat.Message) {
	graph, err := a.compileWorkflow(a.AgentConfig)
	if err != nil {
		panic(err)
	}
	graph.ForCompiledPath(func(_ dag.Node) {})
}

func (a *AgentRun) Receive() (chat.Message, error) {
	var msg chat.Message
	var err error
	var ok bool
	select {
	case msg, ok = <-a.messageChan:
	case err, ok = <-a.errorChan:
	}
	if !ok {
		return chat.Message{}, io.EOF
	}
	return msg, err
}
