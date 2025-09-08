package workflow

import (
	"flowing/internal/dag"
	model "flowing/internal/model/agent"
)

// compileWorkflow 编译工作流，只编译流程图，不填充节点函数
func (a *AgentRun) compileWorkflow(config *model.WorkflowConfig) (*dag.Graph, error) {
	for _, node := range config.Nodes {
		a.configNodeMap[node.Id] = node
	}
	for _, edge := range config.Edges {
		a.configEdgeMap[edge.Id] = edge
	}

	g := dag.NewGraph()
	// 添加节点
	for _, node := range config.Nodes {
		dagNode := dag.NewNode(node.Id, node.Data.Label, nil)
		g.AddNode(dagNode)
	}
	// 添加连接
	for _, edge := range config.Edges {
		g.AddEdge(dag.NewEdge(edge.Id, edge.Source, edge.Target, nil))
	}
	// 编译得到拓扑排序
	if err := g.Compile(); err != nil {
		return nil, err
	}

	// TODO 为每个输出节点设置排序
	// TODO 为每个模型节点设置输出通道和排序
	sectionNum := new(int)
	*sectionNum = 1
	sortedReplyNodes := make([]*model.Node, 0)
	// 拿到拓扑排序后的回复节点序列
	g.ForCompiledPath(func(n dag.Node) {
		node := a.configNodeMap[n.Id()]
		if node != nil && node.Data.Type == model.NodeTypeReply {
			sortedReplyNodes = append(sortedReplyNodes, node)
		}
	})
	return g, nil
}
