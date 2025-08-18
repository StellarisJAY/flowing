<template>
  <ModalSelect
    title="选择数据源"
    :api="queryDatasourceList"
    :params="{ type: datasourceType }"
    :columns="columns"
    label-field="name"
    value-field="id"
    v-model:value="value"
    :query-form-schema="queryFormSchema"
    :disabled="disabled"
    :multiple="false"
  />
</template>
<script setup lang="ts">
  import ModalSelect from '@/components/Select/ModalSelect.vue';
  import { queryDatasourceList } from '@/api/monitor/datasource.api';

  const columns = [
    {
      title: '数据源名称',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: '数据源编码',
      dataIndex: 'code',
      key: 'code',
    },
    {
      title: '数据源类型',
      dataIndex: 'type',
      key: 'type',
    },
    {
      title: '主机',
      dataIndex: 'host',
      key: 'host',
    },
    {
      title: '端口',
      dataIndex: 'port',
      key: 'port',
    },
    {
      title: '数据库名',
      dataIndex: 'database',
      key: 'database',
    },
  ];

  const queryFormSchema = [
    {
      label: '数据源名称',
      type: 'input',
      name: 'name',
      placeholder: '请输入数据源名称',
      defaultValue: '',
    },
    {
      label: '数据源编码',
      type: 'input',
      placeholder: '请输入数据源编码',
      defaultValue: '',
      name: 'code',
    },
    {
      label: '主机',
      type: 'input',
      placeholder: '请输入主机',
      defaultValue: '',
      name: 'host',
    },
  ];

  defineProps({
    disabled: {
      type: Boolean,
      default: false,
    },
    datasourceType: {
      type: String,
      default: '',
    },
  });

  const value = defineModel('value', {
    type: String,
    default: '',
  });
</script>
