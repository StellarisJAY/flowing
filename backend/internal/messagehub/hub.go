package messagehub

import (
	"context"
	"encoding/json"
	"errors"
	"flowing/global"
	runner "flowing/internal/agent"
	"flowing/internal/model/agent"
	"flowing/internal/model/chat"
	"flowing/internal/model/common"
	"flowing/internal/repository"
	"flowing/internal/util"
	"io"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func HandleSendMessage(ctx context.Context, req chat.SendMessageReq) error {
	switch req.Mode {
	case chat.SendMessageModeDebugger:
		if req.AgentType == "simple" {
			return handleDebuggerModeSimple(ctx, req)
		}
		return global.NewError(500, "暂不支持工作流智能体", nil)
	case chat.SendMessageModeAgent:
		return handleAgentMode(ctx, req)
	default:
		return global.NewError(500, "暂不支持的聊天模式", nil)
	}
}

func getAgentConfig(config string, agentType agent.Type) (any, error) {
	var err error
	var conf any
	switch agentType {
	case agent.TypeSimple:
		agentConfig := agent.SimpleAgentConfig{}
		err = json.Unmarshal([]byte(config), &agentConfig)
		conf = agentConfig
	default:
		err = errors.New("暂不支持的智能体类型")
	}
	return conf, err
}

func createAndRunAgent(ctx context.Context, agentDetail *agent.Agent, agentConfig any, req chat.SendMessageReq) error {
	writer := ctx.Value(global.ContextKeySSEWriter).(gin.ResponseWriter)
	// 创建智能体运行实例
	agentRun, err := runner.NewAgentRun(ctx, agentDetail, agentConfig, req.ConversationId)
	if err != nil {
		return global.NewError(500, "运行智能体失败", err)
	}
	// 为用户输入消息生成id
	messageId := repository.Snowflake().Generate().Int64()
	// 为会话生成id
	if req.ConversationId == 0 {
		req.ConversationId = repository.Snowflake().Generate().Int64()
	}
	// 创建用户输入消息
	inputMessage := chat.Message{
		BaseModel:      common.BaseModel{Id: messageId},
		ConversationId: req.ConversationId,
		Content:        req.Content,
		Type:           chat.MessageTypeUser,
		AgentId:        req.AgentId,
		AgentRunId:     0, // TODO 关联智能体运行实例
	}

	// 运行智能体
	agentRun.Run(ctx, inputMessage)
	for {
		msg, err := agentRun.Receive()
		if errors.Is(err, io.EOF) {
			slog.Info("智能体运行结束", "mode", req.Mode)
			break
		}
		if err != nil {
			return global.NewError(500, "运行智能体错误", err)
		}
		if err := util.SSESendMessage(msg, writer); err != nil {
			return global.NewError(500, "发送消息错误", err)
		}
		writer.Flush()
	}
	// TODO 保存聊天记录和会话记录
	return nil

}

// handleAgentMode 处理调用智能体的聊天
func handleAgentMode(ctx context.Context, req chat.SendMessageReq) error {
	// 获取智能体详情
	agentDetail, err := agent.GetAgentDetail(ctx, req.AgentId)
	if err != nil {
		return global.NewError(500, "获取智能体失败", err)
	}
	req.AgentConfig = agentDetail.Config
	// 解析智能体配置
	agentConfig, err := getAgentConfig(req.AgentConfig, agentDetail.Type)
	if err != nil {
		return global.NewError(500, "无效的智能体配置", err)
	}
	// 运行智能体
	return createAndRunAgent(ctx, agentDetail, agentConfig, req)
}

// handleDebuggerModeSimple 处理调试模式运行简单智能体，前端直接传智能体配置
func handleDebuggerModeSimple(ctx context.Context, req chat.SendMessageReq) error {
	var agentConfig agent.SimpleAgentConfig
	err := json.Unmarshal([]byte(req.AgentConfig), &agentConfig)
	if err != nil {
		return global.NewError(500, "无效的智能体配置", err)
	}
	tempAgent := agent.Agent{
		BaseModel: common.BaseModel{Id: 0},
		Name:      "调试运行",
		Type:      agent.TypeSimple,
		Config:    req.AgentConfig,
	}
	return createAndRunAgent(ctx, &tempAgent, agentConfig, req)
}
