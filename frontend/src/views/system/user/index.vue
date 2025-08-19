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
      <IconButton type="primary" @click="openDrawer(false)" icon="PlusOutlined" title="新增用户" />
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
    @close="triggerQuery"
  />
</template>

<script lang="js" setup>
  import Table from '@/components/Table/index.vue';
  import {
    columns, getDetail, saveUser,
    searchUserFormSchema,
    userFormRules,
    userFormSchema,
    useUserStore,
  } from './user.data.js';
  import { Button, Space } from 'ant-design-vue';
  import { computed, ref } from 'vue';
  import FormDrawer from '@/components/Drawer/FormDrawer.vue';
  import ConfirmButton from '@/components/Button/ConfirmButton.vue';
  import IconButton from '@/components/Button/IconButton.vue';

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
      getDetail(record.id).then((data)=>userStore.setUserForm(data));
    } else {
      userStore.initUserForm();
    }
    formDrawerRef.value.open(isUpdate);
  };

  const triggerQuery = async () => {
    await tableRef.value.triggerQuery();
  };
</script>

<style scoped></style>
