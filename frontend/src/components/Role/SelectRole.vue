<template>
  <ModalSelect
    title="选择角色"
    :api="queryRoleList"
    :query-form-schema="queryFormSchema"
    :columns="columns"
    label-field="roleName"
    value-field="id"
    v-model:value="value"
    :disabled="disabled"
    :multiple="multiple"
  />
</template>
<script setup lang="ts">
  import ModalSelect from '@/components/Select/ModalSelect.vue';
  import { queryRoleList } from '@/api/system/role.api';
  import { queryFormSchema } from '@/views/system/role/role.data';

  defineProps({
    disabled: {
      type: Boolean,
      default: false,
    },
    multiple: {
      type: Boolean,
      default: false,
    },
  });

  const columns = [
    {
      title: '角色名称',
      dataIndex: 'roleName',
      key: 'roleName',
    },
    {
      title: '角色编码',
      dataIndex: 'roleKey',
      key: 'roleKey',
    },
    {
      title: '角色描述',
      dataIndex: 'description',
      key: 'description',
    },
  ];

  const value = defineModel('value', {
    type: [Array, String],
    default: () => [],
  });
</script>
