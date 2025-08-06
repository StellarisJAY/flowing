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
    <template #bodyCell="{column, record}">
      <p v-if="column.dataIndex === 'type'">{{ menuStore.getMenuTypeName(record.type) }}</p>
      <Space v-if="column.dataIndex === 'action'">
        <Button type="link" size="small" @click="()=>openDrawer(true, record)">编辑</Button>
        <Button type="link" size="small" danger>删除</Button>
      </Space>
    </template>
  </Table>
  <MenuDrawer ref="menuDrawer" @submit-ok="refresh"/>
</template>

<script lang="js" setup>
  import { computed, ref } from 'vue';
  import { Button, Space } from 'ant-design-vue';
  import Table from '@/components/Table/index.vue';
  import MenuDrawer from '@/views/system/menu/menuDrawer.vue';
  import { useMenuStore } from '@/views/system/menu/menuStore.js';

  const menuStore = useMenuStore();
  const columns = computed(() => menuStore.columns);
  const records = computed(() => menuStore.menuList);
  const queryFormSchema = computed(()=>menuStore.queryFormSchema);
  const menuDrawer = ref();

  const refresh = async (e) => {
    await menuStore.queryMenuList(e);
  };

  const openDrawer = (isUpdate, record)=>{
    if(isUpdate){
      menuStore.menuForm = {...record};
    } else {
      menuStore.initMenuForm();
    }
    menuDrawer.value.setVisible(true);
  };
</script>

<style scoped></style>
