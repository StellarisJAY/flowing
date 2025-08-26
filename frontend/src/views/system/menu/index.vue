<template>
  <Table
    ref="tableRef"
    :columns="tableColumns"
    :records="records"
    :query-form-schema="tableQueryFormSchema"
    :query-form-rules="[]"
    :pagination="false"
    @refresh="refresh"
  >
    <template #tool-buttons>
      <IconButton type="primary" @click="() => openDrawer(false)" icon="PlusOutlined" title="新增菜单" />
    </template>
    <template #bodyCell="{ column, record }">
      <p v-if="column.dataIndex === 'type'">{{ getMenuTypeName(record.type) }}</p>
      <Space v-if="column.dataIndex === 'action'">
        <Button type="link" size="small" @click="() => openDrawer(true, record)">编辑</Button>
        <ConfirmButton text="删除" @confirm="async () => {
          await deleteMenu(record.id);
          await triggerQuery();
        }" />
      </Space>
      <span v-if="column.dataIndex === 'path'">
        {{record.type !== 3 ? record.path : record.actionCode}}
      </span>
    </template>
  </Table>
  <FormDrawer
    ref="menuDrawer"
    :form-schema="drawerFormSchema"
    :form-state="drawerFormState"
    :form-rules="drawerFormRules"
    :submit="saveMenu"
    @close="triggerQuery"
    title="菜单"
  />
</template>

<script lang="js" setup>
  import { computed, ref } from 'vue';
  import { Button, Space } from 'ant-design-vue';
  import FormDrawer from '@/components/Drawer/FormDrawer.vue';
  import Table from '@/components/Table/index.vue';
  import {
    menuFormRules,
    menuFormSchema,
    useMenuStore,
    queryFormSchema,
    getMenuTypeName,
    columns, deleteMenu,
  } from '@/views/system/menu/menu.data.js';
  import ConfirmButton from '@/components/Button/ConfirmButton.vue';
  import IconButton from '@/components/Button/IconButton.vue';

  const menuStore = useMenuStore();
  const tableColumns = columns;
  const records = computed(() => menuStore.menuList);
  const tableQueryFormSchema = queryFormSchema;
  const drawerFormSchema = menuFormSchema;
  const drawerFormState = computed(() => menuStore.menuForm);
  const drawerFormRules = menuFormRules;
  const menuDrawer = ref();
  const tableRef = ref();

  const refresh = async (e) => {
    await menuStore.queryMenuList(e);
  };

  const openDrawer = (isUpdate, record) => {
    if (isUpdate) {
      menuStore.menuForm = { ...record };
    } else {
      menuStore.initMenuForm();
    }
    menuDrawer.value.open(isUpdate);
  };

  const saveMenu = (e, isUpdate) => {
    if (e.type !== 3) {
      e.actionCode = null;
    }
    if (isUpdate) {
      return menuStore.updateMenu(e);
    } else {
      return menuStore.addMenu(e);
    }
  };

  const triggerQuery = async () => {
    await tableRef.value.triggerQuery();
  };
</script>

<style scoped></style>
