<template>
  <!--后继节点列表-->
  <List :data-source="nextNodes">
    <template #renderItem="{ item }">
      <List.Item>
        {{ item.label }}
        <template #actions>
          <Button @click="() => disconnectNextNode(item.value)">断开连接</Button>
        </template>
      </List.Item>
    </template>
  </List>
</template>

<script lang="js" setup>
  import { List, Button } from 'ant-design-vue';
  import { useWorkflowStore } from '@/stores/workflow.js';
  import { computed } from 'vue';

  const props = defineProps({
    nodeId: {
      type: String,
      required: true,
    },
  });

  const workflowStore = useWorkflowStore();
  // 后继节点列表
  const nextNodes = computed(() => {
    const childIds = workflowStore.getChildNodes(props.nodeId);
    return childIds.map((child) => {
      const node = workflowStore.getNode(child);
      if (!node) return null;
      return {
        label: node.data.label,
        value: node.id,
        type: node.data.type,
      };
    });
  });

  const disconnectNextNode = (id) => {
    workflowStore.disconnect(props.nodeId, id);
  };
</script>

<style scoped></style>
