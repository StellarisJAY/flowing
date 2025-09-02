<template>
  <div class="workflow-pane">
    <div class="header">
      <div class="header-title"> 工作流配置 </div>
      <div class="header-toolbar">
        <Button>调试</Button>
        <Button type="primary">保存</Button>
      </div>
    </div>

    <div class="content">
      <!-- 节点添加卡片 -->
      <AddNodeCard />
      <!-- 工作流画布 -->
      <VueFlow :nodes="draft.nodes" :edges="draft.edges">
        <!-- 背景 -->
        <Background pattern-color="#81818a" />
        <!-- 自定义节点样式 -->
        <template #node-base="specialNodeProps">
          <BaseNode v-bind="specialNodeProps" />
        </template>
        <!-- 自定义边样式 -->
        <template #edge-base="specialEdgeProps">
          <BaseEdge v-bind="specialEdgeProps" />
        </template>
      </VueFlow>
    </div>
  </div>
</template>
<script setup lang="js">
  import './index.css';
  import { Button } from 'ant-design-vue';
  import { useWorkflowStore } from '@/stores/workflow.js';
  import { computed } from 'vue';
  import { VueFlow } from '@vue-flow/core';
  import { Background } from '@vue-flow/background';
  import BaseEdge from '@/components/workflow/BaseEdge.vue';
  import BaseNode from '@/components/workflow/BaseNode.vue';
  import { useVueFlow } from '@vue-flow/core';
  import AddNodeCard from '@/components/workflow/AddNodeCard.vue';

  const { onConnect, onNodeDragStop, onViewportChangeEnd } = useVueFlow();

  // 连接事件，更新draft的edges
  onConnect((params) => {
    workflowStore.addEdge(params.source, params.target, params.sourceHandle, params.targetHandle);
  });

  // 拖拽画布事件，更新draft的viewport
  onViewportChangeEnd((e) => {
    workflowStore.updateViewport(e);
  });

  // 节点拖拽事件，更新draft的节点位置
  onNodeDragStop((e) => {
    workflowStore.updateNodePosition(e.node.id, e.node.position);
  });

  const workflowStore = useWorkflowStore();
  // nodes和edges，只读，所有修改通过pinia action完成
  const draft = computed(() => workflowStore.draft);
</script>
<style>
  .workflow-pane {
    width: 100%;
    height: 100%;
  }
  .content {
    width: 100%;
    height: 100%;
    background-color: #f0f0f0;
    position: relative;
  }
</style>
