<template>
  <div class="add-node-card">
    <div class="search"> 点击添加节点 </div>
    <div class="node-list">
      <Menu :items="menuItems" :selectable="false" @click="addNode"></Menu>
    </div>
  </div>
</template>
<script setup lang="js">
  import { Menu } from 'ant-design-vue';
  import { nodePrototypes, useWorkflowStore } from '@/stores/workflow.js';
  import { computed } from 'vue';

  const workflowStore = useWorkflowStore();

  const menuItems = computed(() => {
    return nodePrototypes.map((item) => ({
      label: item.data.label,
      key: item.data.type,
    }));
  });

  const addNode = (e) => {
    workflowStore.addNode(e.key);
  };
</script>

<style>
  .add-node-card {
    position: absolute;
    bottom: 20px;
    left: 20px;
    width: 150px;
    height: 300px;
    background-color: white;
    z-index: 10;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    border-radius: 4px;
    padding: 5px;
    display: flex;
    flex-direction: column;
  }

  .search {
    width: 100%;
    height: 10%;
    font-size: 14px;
    color: #007bff;
  }

  .node-list {
    height: 90%;
    display: flow;
    overflow: auto;
  }
</style>
