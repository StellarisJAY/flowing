import { defineStore } from 'pinia';
import { MarkerType } from '@vue-flow/core';

const startNodeProto = {
  id: 'start',
  type: 'base',
  position: { x: 0, y: 0 },
  data: {
    type: 'start',
    label: '开始',
    description: '流程开始',
    config: {},
    hideTargetHandle: true,
  },
}

const modelNodeProto = {
  id: 'proto',
  type: 'base',
  position: { x: 0, y: 0 },
  data: {
    type: 'model',
    label: '大模型',
    description: '调用大模型',
    config: {}
  },
};

const knowledgeNodeProto = {
  id: 'proto',
  type: 'base',
  position: { x: 0, y: 0 },
  data: {
    type: 'knowledge',
    label: '知识库搜索',
    description: '从知识库中搜索信息',
    config: {}
  },
};

const conditionNodeProto = {
  id: 'proto',
  type: 'base',
  position: { x: 0, y: 0 },
  data: {
    type: 'condition',
    label: '条件判断',
    config: {},
    hideSourceHandle: true,
  },
};

export const nodePrototypes = [
  modelNodeProto,
  knowledgeNodeProto,
  conditionNodeProto,
];

export const useWorkflowStore = defineStore('workflow', {
  state: () => ({
    draft: {
      nodes: [
        startNodeProto,
      ],
      edges: [],
      viewport: {
        x: 0,
        y: 0,
        zoom: 1,
      },
    },
  }),
  actions: {
    addEdge(source, target, sourceHandle, targetHandle) {
      // 避免自连接
      if (source === target) return;
      // 避免反向连接
      if (!sourceHandle.includes("source") || !targetHandle.includes("target")) {
        return;
      }
      const id =  `${source}-${target}-${sourceHandle}-${targetHandle}`;
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
    addNode(node) {
      this.draft.nodes.push(node);
    },
    updateNodePosition(id, position) {
      const node = this.draft.nodes.find((node) => node.id === id);
      if (node) {
        node.position = position;
      }
    },
    updateViewport(viewport) {
      this.draft.viewport = viewport;
    }
  }
});
