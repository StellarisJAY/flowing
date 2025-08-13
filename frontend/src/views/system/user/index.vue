<template>
  <Table
    ref="tableRef"
    :query-form-schema="searchUserFormSchema"
    :query-form-rules="[]"
    :columns="columns"
    :records="records"
    :total="total"
    @refresh="refresh"
  >
    <template #tool-buttons>
      <Button type="primary" @click="openDrawer(false)">新增用户</Button>
    </template>
    <template #bodyCell="{ column, record }">
      <Space v-if="column.dataIndex === 'actions'">
        <Button type="link" @click="openDrawer(true, record)">编辑</Button>
        <ConfirmButton text="删除" />
      </Space>
    </template>
  </Table>
  <FormDrawer
    ref="formDrawerRef"
    title="用户"
    :form-schema="userFormSchema"
    :form-rules="userFormRules"
    :form-state="formState"
    :submit="saveUser"
    @close="refresh"
  />
</template>

<script lang="js" setup>
  import Table from '@/components/Table/index.vue';
  import {
    columns, saveUser,
    searchUserFormSchema,
    userFormRules,
    userFormSchema,
    useUserStore,
  } from './user.data.js';
  import { Button, Space } from 'ant-design-vue';
  import { computed, ref } from 'vue';
  import FormDrawer from '@/components/Drawer/FormDrawer.vue';
  import ConfirmButton from '@/components/Button/ConfirmButton.vue';

  const tableRef = ref();
  const formDrawerRef = ref();
  const userStore = useUserStore();
  const records = computed(() => userStore.records);
  const total = computed(() => userStore.total);
  const formState = computed(() => userStore.userForm);

  const refresh = async (e) => {
    await userStore.refresh(e);
  };

  const openDrawer = (isUpdate, record) => {
    if (isUpdate) {
      userStore.setUserForm(record);
    } else {
      userStore.initUserForm();
    }
    formDrawerRef.value.open(isUpdate);
  };
</script>

<style scoped></style>
