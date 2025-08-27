<template>
  <ModalSelect
    title="选择智能体"
    :api="queryAgentList"
    :query-form-schema="searchFormSchema"
    :columns="columns"
    label-field="name"
    value-field="id"
    v-model:value="value"
    :disabled="disabled"
    :multiple="false"
  />
</template>
<script setup lang="ts">
import ModalSelect from '@/components/Select/ModalSelect.vue';
import { queryAgentList } from '@/api/ai/agent.api';

defineProps({
  disabled: {
    type: Boolean,
    default: false,
  },
});

const searchFormSchema = [
  {
    name: 'name',
    label: '智能体名称',
    type: 'input',
    placeholder: '请输入智能体名称',
    defaultValue: '',
  },
  {
    name: 'type',
    label: '智能体类型',
    type: 'select',
    options: () => [
      {
        label: '聊天',
        value: 'simple',
      },
      {
        label: '工作流',
        value: 'workflow',
      }
    ],
  }
];

const columns = [
  {
    title: '智能体名称',
    key: 'name',
    dataIndex: 'name',
  },
  {
    title: '描述',
    key: 'description',
    dataIndex: 'description',
  },
];

const value = defineModel('value', {
  type: String,
  default: () => '',
});
</script>
