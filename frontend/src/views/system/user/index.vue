<template>
  <Table
    :query-form-schema="searchUserFormSchema"
    :query-form-rules="[]"
    :columns="columns"
    :records="records"
    @refresh="refresh"
  >
    <template #tool-buttons>
      <Button type="primary">新增用户</Button>
    </template>
  </Table>
</template>

<script lang="js" setup>
  import Table from '@/components/Table/index.vue';
  import { searchUserFormSchema, useUserStore } from './userStore.js';
  import { Button } from 'ant-design-vue';
  import { computed } from 'vue';

  const userStore = useUserStore();
  const columns = computed(() => userStore.getColumns());
  const records = computed(() => userStore.getRecords());

  const refresh = async (e) => {
    await userStore.refresh(e);
  };
</script>

<style scoped></style>
