package simple

import (
	"context"
	agentcommon "flowing/internal/agent/common"
	"flowing/internal/docprocess"
	"flowing/internal/model/agent"
	"flowing/internal/model/ai"
	"flowing/internal/model/chat"
	"flowing/internal/model/kb"
	"flowing/internal/repository"
	"flowing/internal/util"
	"fmt"
	"io"
	"log/slog"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
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
	chatModel, err := util.GetChatModel(ctx, *pm)
	if err != nil {
		return nil, err
	}
	agentRun.chatModel = chatModel
	return agentRun, nil
}

// Run 运行简单智能体，单独开启goroutine运行
func (a *AgentRun) Run(ctx context.Context, input chat.Message) {
	go func() {
		// 运行错误处理，由于单独开启goroutine运行，通过channel将错误发出
		defer func() {
			if err := recover(); err != nil {
				a.errorChan <- err.(error)
			}
			close(a.messageChan)
			close(a.errorChan)
		}()

		// 流程的全局状态，包含用户输入和变量列表
		// TODO 变量列表
		localState := compose.WithGenLocalState[map[string]any](func(c context.Context) (state map[string]any) {
			return map[string]any{
				"agentId": a.Agent.Id,
				"query":   input.Content,
			}
		})

		// 创建简单智能体链，包含知识库retriever节点和模型节点
		// 输入 map[string]any，输出 *schema.Message
		chain := compose.NewChain[map[string]any, *schema.Message](localState)

		// 模型输出消息的id
		messageId := repository.Snowflake().Generate().Int64()

		// 知识库retriever节点
		if a.KnowledgeBase != nil {
			retriever, err := docprocess.NewVectorRetriever(ctx, a.KnowledgeBase)
			if err != nil {
				panic(err)
			}
			// 知识库读取结果转换成messages的lambda节点
			docToJson := compose.InvokableLambda(func(ctx context.Context, input []*schema.Document) (output map[string]any, err error) {
				refMessage, err := agentcommon.GetKnowledgeRefMessage(ctx, input, messageId, a.conversationId, a.Agent.Id, 0)
				if err != nil {
					slog.Error("创建知识库引用消息失败", "err", err)
				}
				if refMessage != nil {
					a.messageChan <- *refMessage
				}
				output, err = agentcommon.DocumentsToMessages("context", schema.System)(ctx, input)
				return
			})
			// 知识库检索，输入query，输出context
			chain = chain.AppendRetriever(retriever,
				compose.WithInputKey("query"),
				compose.WithNodeName("retriever"))
			// 输入[]*schema.Document，输出map[string][]*schema.Message
			chain = chain.AppendLambda(docToJson,
				compose.WithNodeName("docToJson"))
		}

		// 定义历史消息和提示词
		promptMessages := prompt.FromMessages(schema.FString,
			schema.SystemMessage(a.AgentConfig.Prompt),
			schema.MessagesPlaceholder("context", true),
			schema.UserMessage("{query}"),
		)
		chain = chain.AppendChatTemplate(promptMessages, compose.WithStatePreHandler(agentcommon.MergeMapStateAndInput))
		chain = chain.AppendChatModel(a.chatModel)

		// 编译智能体流程，创建流式运行
		runnableChain, err := chain.Compile(ctx)
		if err != nil {
			panic(fmt.Errorf("failed to compile chain: %w", err))
		}
		// 将填充conversationId和messageId的用户输入推送到messageChan
		a.messageChan <- input

		stream, err := runnableChain.Stream(ctx, map[string]any{
			"query": input.Content,
		})
		if err != nil {
			panic(fmt.Errorf("failed to start stream: %w", err))
		}

		// 消费模型输出，转换格式后推送到messageChan
		if err := agentcommon.StreamConsumer(stream, a.messageChan, agentcommon.StreamConsumerMeta{
			ConversationId: a.conversationId,
			AgentId:        a.Agent.Id,
			AgentRunId:     0,
			MessageId:      messageId,
		}); err != nil {
			panic(err)
		}
	}()
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
