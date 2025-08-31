<template>
  <div class="vue-flow__node-default base-node">
    <Handle type="target" :position="Position.Left" />
    <Handle type="source" :position="Position.Right" />
    <div class="base-node-header">
      <div class="icon-title">
        <img :src="robotIcon" height="100%" alt="icon" />
        {{ data.label }}
      </div>
      <Button type="link" style="height: 100%; align-items: center; padding: 0; display: flex; justify-content: center">
        <template #icon>
          <MoreOutlined style="width: 10px; height: 10px; font-size: 10px"/>
        </template>
      </Button>
    </div>
    <div class="base-node-body">

    </div>
  </div>
</template>

<script setup>
import { Button} from 'ant-design-vue';
import { Position, Handle } from '@vue-flow/core';
import {MoreOutlined} from '@ant-design/icons-vue';
import robotIcon from '@/assets/svg/avatar_robot.svg';
import { computed } from 'vue';

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
  }
});

const type = computed(() => props.data.type);

</script>

<style scoped>
.base-node {
  min-width: 200px;
  min-height: 200px;
  border: none;
  border-radius: 10px;
  padding: 10px;
  display: flex;
  flex-direction: column;
}
.base-node-header {
  height: 30px;
  background-color: white;
  display: flex;
  align-items: center;
  flex-direction: row;
  justify-content: space-between;
}
.base-node-body {
  flex: 1;
  background-color: lightblue;
}
.icon-title {
  display: flex;
  height: 100%;
  font-size: 14px;
}
</style>

