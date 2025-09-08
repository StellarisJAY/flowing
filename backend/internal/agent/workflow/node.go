package workflow

import (
	"context"
	"flowing/internal/dag"
	model "flowing/internal/model/agent"
)

func (a *AgentRun) getNodeFunc(node *model.Node) dag.NodeFunc {
	switch node.Data.Type {
	case model.NodeTypeStart:
		return a.newStartNodeFunc(node)
	case model.NodeTypeReply:
		return a.newReplyNodeFunc(node)
	case model.NodeTypeKnowledge:
		return a.newKnowledgeNodeFunc(node)
	default:
		return nil
	}
}

func (a *AgentRun) newStartNodeFunc(_ *model.Node) dag.NodeFunc {
	return func(ctx context.Context, dagNode dag.Node) (result dag.NodeFuncReturn) {
		result.Output = map[string]any{
			"sys.conversationId": a.conversationId,
		}
		return
	}
}

func (a *AgentRun) newModelNodeFunc(node *model.Node, outputMessage bool, sectionNum int) dag.NodeFunc {
	return func(ctx context.Context, dagNode dag.Node) (result dag.NodeFuncReturn) {
		return
	}
}

func (a *AgentRun) newReplyNodeFunc(node *model.Node) dag.NodeFunc {
	return func(ctx context.Context, dagNode dag.Node) (result dag.NodeFuncReturn) {
		return
	}
}

func (a *AgentRun) newKnowledgeNodeFunc(node *model.Node) dag.NodeFunc {
	return func(ctx context.Context, dagNode dag.Node) (result dag.NodeFuncReturn) {
		return
	}
}
