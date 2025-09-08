<template>
  <div class="vue-flow__node-default base-node">
    <Handle
      type="target"
      :position="Position.Left"
      class="base-handle"
      :id="`${id}-target`"
      v-if="showTargetHandle(data.type)"
    />
    <Handle
      type="source"
      :position="Position.Right"
      class="base-handle"
      :id="`${id}-source`"
      v-if="showSourceHandle(data.type)"
    />
    <div class="base-node-header">
      <div class="icon-title">
        <img :src="getNodeIconForType()" height="100%" alt="icon" />
        {{ data.label }}
      </div>
      <Button
        type="link"
        style="
          height: 100%;
          align-items: center;
          padding: 0;
          display: flex;
          justify-content: center;
        "
      >
        <template #icon>
          <MoreOutlined style="width: 10px; height: 10px; font-size: 10px" />
        </template>
      </Button>
    </div>
    <div class="base-node-description">
      {{ data.description }}
    </div>
    <div class="base-node-body"> </div>
  </div>
</template>

<script setup>
  import { Button } from 'ant-design-vue';
  import { Position, Handle } from '@vue-flow/core';
  import { MoreOutlined } from '@ant-design/icons-vue';
  import robotIcon from '@/assets/svg/avatar_robot.svg';
  import knowledgeIcon from '@/assets/svg/ext_file_generic_icon.svg';
  import { computed } from 'vue';
  import { showSourceHandle, showTargetHandle } from '@/stores/workflow.js';

  const props = defineProps({
    id: {
      type: String,
      required: true,
    },
    position: {
      type: Object,
      required: true,
    },
    data: {
      type: Object,
      required: true,
    },
  });

  const type = computed(() => props.data.type);

  const getNodeIconForType = () => {
    switch (type.value) {
      case 'model':
        return robotIcon;
      case 'knowledge':
        return knowledgeIcon;
      default:
        return robotIcon;
    }
  };
</script>

<style scoped>
  .base-node {
    width: 200px;
    border: none;
    border-radius: 10px;
    padding: 10px;
    display: flex;
    justify-content: flex-start;
    flex-direction: column;
  }

  .base-node:hover {
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  }

  .base-node-header {
    height: 30px;
    background-color: white;
    display: flex;
    align-items: center;
    flex-direction: row;
    justify-content: space-between;
  }
  .base-node-description {
    width: 100%;
    height: 20px;
    font-size: 12px;
    color: #999;
    text-align: left;
  }
  .base-node-body {
    flex: 1;
    background-color: lightblue;
  }
  .icon-title {
    display: flex;
    height: 100%;
    font-size: 14px;
    align-items: center;
  }

  .base-handle {
    width: 10px;
    height: 10px;
    background-color: #0096ff;
    border-radius: 50%;
    z-index: 1;
  }
</style>
