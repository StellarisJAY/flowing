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
      <VueFlow :nodes="draft.nodes" :edges="draft.edges">
        <Background pattern-color="#81818a" />
        <template #node-base="specialNodeProps">
          <BaseNode v-bind="specialNodeProps" />
        </template>
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

  const workflowStore = useWorkflowStore();
  const draft = computed(() => workflowStore.draft);
</script>
<style>
  .workflow-pane {
    width: 100%;
    height: 100%;
  }
</style>
