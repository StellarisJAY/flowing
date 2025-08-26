<template>
  <ModalSelect
    title="选择知识库"
    :columns="columns"
    :queryFormSchema="queryFormSchema"
    :api="queryKnowledgeList"
    label-field="name"
    value-field="id"
    :multiple="false"
    v-model:value="value"
    :allowClear="allowClear"
  />
</template>

<script lang="js" setup>

import ModalSelect from '@/components/Select/ModalSelect.vue';
import { queryKnowledgeList } from '@/api/ai/knowledge.api.js';

const value = defineModel('value', {
  type: String,
  default: ''
});

defineProps({
  allowClear: {
    type: Boolean,
    default: true,
  },
});

const columns = [
  {
    title: '知识库名称',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '知识库描述',
    dataIndex: 'description',
    key: 'description',
  },
];

const queryFormSchema = [
  {
    label: '知识库名称',
    name: 'name',
    type: 'input',
    defaultValue: '',
  },
  {
    label: '我创建的',
    name: 'private',
    type: 'checkbox',
    defaultValue: false,
  }
];

</script>

<style scoped>

</style>
