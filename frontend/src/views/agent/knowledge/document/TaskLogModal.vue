<template>
  <Modal
    title="解析日志"
    :open="visible"
    :width="800"
    @cancel="close"
  >
    <div>
      <ul>
        <li>
          <span>解析状态：
            <Tag v-if="isTaskSuccess(task)" color="green">完成</Tag>
            <Tag v-else-if="isTaskFailed(task)" color="red">解析失败</Tag>
            <Tag v-else-if="task.status === 'slicing'" color="yellow">切片中...</Tag>
            <Tag v-else-if="task.status === 'embedding'" color="yellow">嵌入中...</Tag>
          </span>
        </li>
        <li v-if="!isTaskFailed(task)">
          <span>切片开始时间：{{ task.slicingStartTime }}</span>
        </li>
        <li v-if="!isTaskFailed(task)">
          <span>切片结束时间：{{ task.slicingEndTime }}</span>
        </li>
        <li v-if="!isTaskFailed(task)">
          <span>嵌入开始时间：{{task.embeddingStartTime}}</span>
        </li>
        <li v-if="!isTaskFailed(task)">
          <span>嵌入结束时间：{{task.embeddingEndTime}}</span>
        </li>
        <li v-if="task.errorMessage !== ''">
          <span>错误信息：{{task.errorMessage}}</span>
        </li>
      </ul>
    </div>
  </Modal>
</template>

<script lang="js" setup>
import { Modal, Tag } from 'ant-design-vue';
import { ref } from 'vue';
import { isTaskFailed, isTaskSuccess } from '@/views/agent/knowledge/document/document.data.js';

const visible = ref(false);
const task = ref({});

const open = (data) => {
  visible.value = true;
  task.value = data;
};
const close = () => {
  task.value = {};
  visible.value = false;
};

defineExpose({
  open,
})
</script>
