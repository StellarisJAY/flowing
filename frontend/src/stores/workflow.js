import { defineStore } from 'pinia';
import { MarkerType } from '@vue-flow/core';
import { v4 as uuid } from 'uuid';
import { getDetail, saveConfig } from '@/api/ai/agent.api.js';
import { message } from 'ant-design-vue';

// 开始节点原型
const startNodeProto = {
  id: 'start',
  type: 'base',
  position: { x: 0, y: 0 },
  data: {
    type: 'start',
    label: '开始',
    description: '流程开始',
    config: {
      output: [
        {
          id: 'sys.query',
          name: 'query',
          type: 'string',
          description: '用户输入',
          fixed: true,
        },
        {
          id: 'sys.file',
          name: 'file',
          type: 'file',
          description: '用户上传文件',
          fixed: true,
        },
      ],
    },
    hideTargetHandle: true,
  },
};
// 模型节点原型
const modelNodeProto = {
  id: 'proto',
  type: 'base',
  position: { x: 0, y: 0 },
  data: {
    type: 'model',
    label: '大模型',
    description: '调用大模型',
    config: {
      model: '',
      prompt: '',
      output: [
        {
          id: 'model.text',
          name: 'text',
          type: 'string',
          description: '模型输出',
          fixed: true,
          outputChat: true,
        }
      ]
    },
  },
};
// 知识库节点原型
const knowledgeNodeProto = {
  id: 'proto',
  type: 'base',
  position: { x: 0, y: 0 },
  data: {
    type: 'knowledge',
    label: '知识库搜索',
    description: '从知识库中搜索信息',
    config: {
      knowledgeBaseId: '',
      query: '',
      searchType: 'vector',
      topK: 2,
      threshold: 0.5,
      hybridType: 'weight',
      weight: 0.5,
      input: [
        {
          id: 'knowledge.query',
          name: 'query',
          type: 'string',
          description: '搜索内容',
          fixed: true,
        }
      ],
      output: [
        {
          id: 'knowledge.text',
          name: 'text',
          type: 'string',
          description: '知识库输出',
          fixed: true,
          outputChat: false,
        }
      ]
    },
  },
};
// 条件节点原型
const conditionNodeProto = {
  id: 'proto',
  type: 'base',
  position: { x: 0, y: 0 },
  data: {
    type: 'condition',
    label: '条件判断',
    config: {},
    hideSourceHandle: true, // 隐藏条件节点的输出handle，输出handle在条件组定义
  },
};
// 回复消息节点原型
const replyNodeProto = {
  id: 'reply',
  type: 'base',
  position: { x: 0, y: 0 },
  data: {
    type: 'reply',
    label: '消息回复',
    description: '回复用户消息',
    config: {
      message: '',
    },
  },
};

// 节点原型列表，添加节点时从这个列表找到该类型节点原型并拷贝
export const nodePrototypes = [modelNodeProto, knowledgeNodeProto, conditionNodeProto, replyNodeProto];
// 判断节点类型是否允许添加input
export const allowAddInputVariable = (nodeType) => {
  return false;
};
// 判断节点类型是否允许添加output
export const allowAddOutputVariable = (nodeType) => {
  return nodeType === 'start';
};
// 生成变量id
export const genVariableId = (nodeId, varName) => {
  return `${nodeId}.${varName}`;
};

export const showSourceHandle = (nodeType) => {
  return true;
};

export const showTargetHandle = (nodeType) => {
  return nodeType !== 'start';
};

// 当前正在编辑或正在查看的流程store
export const useWorkflowStore = defineStore('flowing_workflow', {
  state: () => ({
    // 当前显示的流程图
    draft: {
      nodes: [startNodeProto],
      edges: [],
      viewport: {
        x: 0,
        y: 0,
        zoom: 1,
      },
    },
  }),
  actions: {
    // 获取节点
    getNode(nodeId) {
      return this.draft.nodes.find((node) => node.id === nodeId);
    },
    // 获取节点的所有父节点id
    getParentNodes(nodeId) {
      const node = this.getNode(nodeId);
      if (!node) return null;
      const edges = this.draft.edges.filter((item) => item.target === nodeId);
      if (edges.length === 0) return null;
      return edges.map((item) => item.source);
    },
    // 获取节点的所有子节点id
    getChildNodes(nodeId) {
      const node = this.getNode(nodeId);
      if (!node) return [];
      const edges = this.draft.edges.filter(edge=>edge.source === nodeId);
      if (edges.length === 0) return [];
      return edges.map((item) => item.target);
    },
    // 获取从一个节点出发，能够到达的所有前驱节点id
    getConnectedPrevNodes(nodeId) {
      const parentNodes = this.getParentNodes(nodeId);
      if (!parentNodes) return null;
      const grandParentNodes = parentNodes.flatMap(
        (item) => this.getConnectedPrevNodes(item) || []
      );
      return [...parentNodes, ...grandParentNodes];
    },
    // 添加连接,source和target是节点id，sourceHandle和targetHandle是连接的端口id（条件节点可能有多个出口）
    addEdge(source, target, sourceHandle, targetHandle) {
      // 避免自连接
      if (source === target) return;
      // 避免反向连接
      if (!sourceHandle.includes('source') || !targetHandle.includes('target')) {
        return;
      }
      const id = `${source}-${target}-${sourceHandle}-${targetHandle}`;
      // 排除已存在的连接
      const edge = this.draft.edges.find((edge) => edge.id === id);
      if (edge) {
        return;
      }
      this.draft.edges.push({
        id: `${source}-${target}-${sourceHandle}-${targetHandle}`,
        source,
        target,
        sourceHandle,
        targetHandle,
        type: 'base',
        markerEnd: MarkerType.ArrowClosed,
        data: {
          hello: 'world',
        },
      });
    },
    // 添加节点
    addNode(type) {
      const nodeProto = nodePrototypes.find((item) => item.data.type === type);
      if (!nodeProto) return;
      // 拷贝一个节点原型，深拷贝
      const node = JSON.parse(JSON.stringify(nodeProto));
      // 生成一个唯一的id
      node.id = uuid();
      // 设置节点的变量id = 节点id.变量名
      node.data.config.input?.forEach((item) => {
        item.id = genVariableId(node.id, item.name);
      });
      node.data.config.output?.forEach((item) => {
        item.id = genVariableId(node.id, item.name);
      });
      const startNode = this.getNode('start');
      // 加入到节点列表
      this.draft.nodes.push({
        ...node,
        position: {
          x: startNode.position.x + Math.random() * 200, // TODO 节点初始坐标
          y: startNode.position.y + Math.random() * 200,
        },
      });
    },
    // 更新节点位置
    updateNodePosition(id, position) {
      const node = this.getNode(id);
      if (node) {
        node.position = position;
      }
    },
    // 更新viewport
    updateViewport(viewport) {
      this.draft.viewport = viewport;
    },
    // 更新节点基本信息和配置
    updateNodeInfo(id, label, description, config) {
      const node = this.getNode(id);
      if (node) {
        node.data.label = label;
        node.data.description = description;
        node.data.config = config;
      }
    },
    // 获取节点可以引用的变量列表
    getRefVariables(nodeId) {
      const node = this.getNode(nodeId);
      if (!node) return [];
      const prevNodes = this.getConnectedPrevNodes(nodeId);
      if (!prevNodes) return [];
      return prevNodes.map((item) => {
        const prevNode = this.getNode(item);
        return {
          label: prevNode.data.label,
          children: prevNode.data.config.output.map((variable) => ({
            label: variable.name,
            value: variable.id,
            type: variable.type,
          })),
        };
      });
    },
    // 断开两个节点的连接
    disconnect(source, target) {
      const idx = this.draft.edges.findIndex(edge => edge.source === source && edge.target === target);
      if (idx !== -1) {
        this.draft.edges.splice(idx, 1);
      }
    },
    // 保存
    async save(id) {
      try {
        await saveConfig({
          id,
          config: JSON.stringify(this.draft),
        });
        message.success('保存成功');
      } catch (err) {
        console.log(err);
        message.error('保存失败');
      }
    },
    // 加载
    async load(id) {
      try {
        const { data } = await getDetail(id);
        this.draft = JSON.parse(data.config);
      }catch (err) {
        console.log(err);
        message.error('加载流程失败');
      }
    },
  },
});
