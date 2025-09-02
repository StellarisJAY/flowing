<template>
  <div class="add-node-card">
    <div class="search">
      点击添加节点
    </div>
    <div class="node-list">
      <Menu :items="menuItems" :selectable="false" @click="addNode"></Menu>
    </div>
  </div>
</template>
<script setup lang="js">
import { Menu } from 'ant-design-vue';
import { nodePrototypes, useWorkflowStore } from '@/stores/workflow.js';
import { computed } from 'vue';
import { v4 as uuidv4 } from 'uuid';
import { useVueFlow } from '@vue-flow/core';

const { getViewport } = useVueFlow();
const workflowStore = useWorkflowStore();

const menuItems = computed(()=>{
  return nodePrototypes.map((item) => ({
    label: item.data.label,
    key: item.data.type,
  }))
});

const addNode = (e) => {
  const nodeProto = nodePrototypes.find(item=>item.data.type === e.key);
  if (!nodeProto) return;
  // 拷贝一个节点原型，深拷贝
  const node = JSON.parse(JSON.stringify(nodeProto));
  // 生成一个唯一的id
  node.id = uuidv4();
  const viewport = getViewport();
  // 加入到节点列表
  workflowStore.addNode({
    ...node,
    position: {
      x: viewport.x + 200,
      y: viewport.y + 200,
    }
  });
}
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
