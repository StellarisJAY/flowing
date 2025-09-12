package simple

import (
	"context"
	agentcommon "flowing/internal/agent/common"
	"flowing/internal/dag"
	"flowing/internal/docprocess"
	"flowing/internal/model/agent"
	"flowing/internal/model/ai"
	"flowing/internal/model/chat"
	"flowing/internal/model/kb"
	"flowing/internal/repository"
	"flowing/internal/util"
	"io"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

// AgentRun 简单智能体运行
type AgentRun struct {
	Id            int64                   // 智能体运行id，如果是调试模式id为0
	AgentConfig   agent.SimpleAgentConfig // 智能体配置
	Agent         *agent.Agent            // 智能体信息，调试模式为空
	KnowledgeBase *kb.KnowledgeBase       // 关联知识库
	ProviderModel *ai.ProviderModelDetail // 聊天模型信息

	chatModel      model.BaseChatModel // 聊天模型实现
	conversationId int64               // 对话id
	messageChan    chan chat.Message   // 消息通道，消息层通过该channel接受模型的回复
	errorChan      chan error          // 错误通道，智能体通过该channel返回错误
}

func NewAgentRun(ctx context.Context, agent *agent.Agent, config agent.SimpleAgentConfig, conversationId int64) (*AgentRun, error) {
	// 获取模型详情
	pm, err := ai.GetProviderModelDetail(ctx, config.ModelId)
	if err != nil {
		return nil, err
	}
	agentRun := &AgentRun{
		AgentConfig:    config,
		Agent:          agent,
		ProviderModel:  pm,
		conversationId: conversationId,
		messageChan:    make(chan chat.Message, 10), // TODO 消息通道大小
		errorChan:      make(chan error),
	}
	// 获取知识库详情
	if config.KnowledgeBaseId != 0 {
		knowledgeBase, err := kb.GetKnowledgeBase(ctx, config.KnowledgeBaseId)
		if err != nil {
			return nil, err
		}
		agentRun.KnowledgeBase = knowledgeBase
	}
	// 加载聊天模型实现
	chatModel, err := util.GetChatModel(ctx, *pm, true, false)
	if err != nil {
		return nil, err
	}
	agentRun.chatModel = chatModel
	return agentRun, nil
}

// Run 运行简单智能体，单独开启goroutine运行
func (a *AgentRun) Run(ctx context.Context, input chat.Message) {
	// 新建DAG流程图
	chain := dag.NewChain()
	if a.KnowledgeBase != nil {
		// 如果有知识库检索配置，则创建知识库检索节点
		chain.AddNode(dag.NewNode("retriever", "retriever", a.retrieverNodeFunc))
	}
	// 聊天模型节点
	chain.AddNode(dag.NewNode("chat", "chat", a.chatNodeFunc))
	_ = chain.Compile()

	// 模型输出消息的id
	messageId := repository.Snowflake().Generate().Int64()
	// 流程初始参数
	// TODO 前端配置的变量列表
	variables := map[string]any{
		"query":          input,
		"messageId":      messageId,
		"conversationId": a.conversationId,
	}
	runner := dag.NewChainRun(chain)
	a.messageChan <- input
	// 非组赛模式运行dag，由于是单线流程，所以并行度为1
	if err := runner.Run(dag.WithNonBlocking(),
		dag.WithContext(ctx),
		dag.WithCallback(a.callback),
		dag.WithParallelNum(1),
		dag.WithPanicHandler(a.panicHandler),
		dag.WithVariables(variables)); err != nil {
		panic(err)
	}
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

func (a *AgentRun) callback(event dag.CallbackEvent) {
	if event.Type == dag.EventTypeEnd {
		close(a.messageChan)
		close(a.errorChan)
	}
}

// chatNodeFunc 聊天模型节点
func (a *AgentRun) chatNodeFunc(ctx context.Context, _ dag.Node) (result dag.NodeFuncReturn) {
	contextMessages := ctx.Value("context")
	messageId := ctx.Value("messageId").(int64)
	vs := map[string]any{
		"query": ctx.Value("query").(chat.Message).Content,
	}
	if contextMessages != nil {
		vs["context"] = contextMessages.([]*schema.Message)
	}
	// 生成提示词和历史消息
	promptMessages := prompt.FromMessages(schema.FString,
		schema.SystemMessage(a.AgentConfig.Prompt),
		schema.MessagesPlaceholder("context", true),
		schema.UserMessage("{query}"),
	)
	prompts, err := promptMessages.Format(ctx, vs)
	if err != nil {
		result.Error = err
		return
	}
	// 创建聊天流
	stream, err := a.chatModel.Stream(ctx, prompts)
	if err != nil {
		result.Error = err
		return
	}
	// 消费聊天流
	err = agentcommon.StreamConsumer(stream, a.messageChan, agentcommon.StreamConsumerMeta{
		ConversationId: a.conversationId,
		AgentId:        a.Agent.Id,
		AgentRunId:     0,
		MessageId:      messageId,
	})
	if err != nil {
		result.Error = err
		return
	}
	return
}

func (a *AgentRun) retrieverNodeFunc(ctx context.Context, _ dag.Node) (result dag.NodeFuncReturn) {
	query := ctx.Value("query").(chat.Message)
	messageId := ctx.Value("messageId").(int64)

	retriever, err := docprocess.NewVectorRetriever(ctx, a.KnowledgeBase, a.AgentConfig.KbSearchOption)
	if err != nil {
		result.Error = err
		return
	}
	docs, err := retriever.Retrieve(ctx, query.Content)
	if err != nil {
		result.Error = err
		return
	}
	refMessage, err := agentcommon.GetKnowledgeRefMessage(ctx, docs, messageId, a.conversationId, a.Agent.Id, 0)
	if err != nil {
		result.Error = err
		return
	}
	if refMessage != nil {
		a.messageChan <- *refMessage
	}
	output, err := agentcommon.DocumentsToMessages("context", schema.System)(ctx, docs)
	if err != nil {
		result.Error = err
		return
	}
	result.Output = output
	return
}

func (a *AgentRun) panicHandler(err error) {
	a.errorChan <- err
	close(a.messageChan)
	close(a.errorChan)
}
