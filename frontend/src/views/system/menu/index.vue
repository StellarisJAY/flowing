<template>
  <Table
    :columns="columns"
    :records="records"
    :query-form-schema="queryFormSchema"
    :query-form-rules="[]"
    :pagination="false"
    @refresh="refresh"
  >
    <template #tool-buttons>
      <Button type="primary" @click="openDrawer(false)">新增菜单</Button>
    </template>
    <template #bodyCell="{ column, record }">
      <p v-if="column.dataIndex === 'type'">{{ menuStore.getMenuTypeName(record.type) }}</p>
      <Space v-if="column.dataIndex === 'action'">
        <Button type="link" size="small" @click="() => openDrawer(true, record)">编辑</Button>
        <Button type="link" size="small" danger>删除</Button>
      </Space>
    </template>
  </Table>
  <FormDrawer
    ref="menuDrawer"
    :form-schema="drawerFormSchema"
    :form-state="drawerFormState"
    :form-rules="drawerFormRules"
    :submit="saveMenu"
    @close="refresh"
    @open="menuStore.getParentMenuOptions"
  />
</template>

<script lang="js" setup>
  import { computed, ref } from 'vue';
  import { Button, Space } from 'ant-design-vue';
  import FormDrawer from '@/components/Drawer/FormDrawer.vue';
  import Table from '@/components/Table/index.vue';
  import { useMenuStore } from '@/views/system/menu/menuStore.js';

  const menuStore = useMenuStore();
  const columns = computed(() => menuStore.columns);
  const records = computed(() => menuStore.menuList);
  const queryFormSchema = computed(() => menuStore.queryFormSchema);
  const drawerFormSchema = computed(() => menuStore.getMenuFormSchema());
  const drawerFormState = computed(() => menuStore.menuForm);
  const drawerFormRules = computed(() => menuStore.menuFormRules);
  const menuDrawer = ref();

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
    if (isUpdate) {
      return menuStore.updateMenu(e);
    } else {
      return menuStore.addMenu(e);
    }
  };
</script>

<style scoped></style>
