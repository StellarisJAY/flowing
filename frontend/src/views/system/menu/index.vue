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
      <Button type="primary">新增菜单</Button>
    </template>
    <template #bodyCell="{column, record}">
      <p v-if="column.dataIndex === 'type'">{{ menuStore.getMenuTypeName(record.type) }}</p>
      <Space v-if="column.dataIndex === 'action'">
        <Button type="link" size="small">编辑</Button>
        <Button type="link" size="small" danger>删除</Button>
      </Space>
    </template>
  </Table>
  <MenuDrawer ref="menuDrawer" />
</template>

<script lang="js" setup>
  import { computed, ref } from 'vue';
  import { Button, Space } from 'ant-design-vue';
  import Table from '@/components/Table/index.vue';
  import MenuDrawer from '@/views/system/menu/menuDrawer.vue';
  import { useMenuStore, queryFormSchema } from '@/views/system/menu/menuStore.js';

  const menuStore = useMenuStore();
  const columns = computed(() => menuStore.getColumns());
  const records = computed(() => menuStore.getMenuTree());

  const menuDrawer = ref(null);

  const refresh = async (e) => {
    console.log(e);
    await menuStore.queryMenuList(e);
  };
</script>

<style scoped></style>
