package agent

import (
	"context"
	"encoding/json"
	"flowing/global"
	"flowing/internal/model/agent"
	"flowing/internal/model/common"
	"flowing/internal/model/system"
)

func CreateAgent(ctx context.Context, req agent.CreateAgentReq) error {
	a := agent.Agent{
		BaseModel: common.BaseModel{
			CreateBy: ctx.Value(global.ContextKeyUser).(system.User).Id,
		},
		Name:        req.Name,
		Type:        req.Type,
		Description: req.Description,
		Public:      req.Public,
	}
	var config any
	switch a.Type {
	case agent.TypeSimple:
		config = agent.DefaultSimpleAgentConfig()
	default:
		return global.NewError(500, "不支持的智能体类型", nil)
	}
	configData, _ := json.Marshal(config)
	a.Config = string(configData)
	if err := agent.CreateAgent(ctx, &a); err != nil {
		return global.NewError(500, "创建智能体失败", err)
	}
	return nil
}

func ListAgent(ctx context.Context, query agent.ListAgentQuery) ([]*agent.Agent, int64, error) {
	query.CreateBy = ctx.Value(global.ContextKeyUser).(system.User).Id
	res, total, err := agent.ListAgent(ctx, query)
	if err != nil {
		return nil, 0, global.NewError(500, "查询智能体失败", err)
	}
	return res, total, nil
}

func GetAgentDetail(ctx context.Context, id int64) (*agent.Agent, error) {
	a, err := agent.GetAgentDetail(ctx, id)
	if err != nil {
		return nil, global.NewError(500, "查询智能体失败", err)
	}
	return a, nil
}

func UpdateConfig(ctx context.Context, req agent.UpdateAgentConfigReq) error {
	// TODO 校验配置
	if err := agent.UpdateConfig(ctx, req.Id, req.Config); err != nil {
		return global.NewError(500, "更新智能体配置失败", err)
	}
	return nil
}
