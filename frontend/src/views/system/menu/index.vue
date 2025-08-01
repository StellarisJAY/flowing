<template>
  <div style="height: 100%; width: 100%">
    <div class="toolbar">
      <Form v-model:form="queryForm" layout="inline">
        <Form.Item label="菜单名称" name="menuName">
          <Input v-model:value="queryForm.menuName" />
        </Form.Item>
        <Form.Item>
          <Button type="primary" @click="menuStore.queryMenuList">查询</Button>
        </Form.Item>
        <Form.Item>
          <Button @click="menuStore.clearQueryForm">重置</Button>
        </Form.Item>
      </Form>
      <Button
        type="primary"
        @click="
          () => {
            menuDrawer.setVisible(true);
          }
        "
        >新增菜单</Button>
    </div>
    <Table :columns="columns" :dataSource="menuList" :pagination="false">
      <template #bodyCell="{ column, record }">
        <div v-if="column.dataIndex === 'type'">
          {{ menuStore.getMenuTypeName(record.type) }}
        </div>
        <div v-else-if="column.dataIndex === 'action'">
          <Button type="link">编辑</Button>
          <Button type="link" danger>删除</Button>
        </div>
      </template>
    </Table>
    <MenuDrawer ref="menuDrawer" />
  </div>
</template>

<script lang="js" setup>
  import { computed, onMounted, reactive, ref } from 'vue';
  import { Table, Button, Input, Form } from 'ant-design-vue';
  import MenuDrawer from '@/views/system/menu/menuDrawer.vue';
  import { useMenuStore } from '@/views/system/menu/menuStore.js';

  const menuStore = useMenuStore();
  const menuDrawer = ref(null);
  const menuList = computed(() => menuStore.getMenuTree());
  const columns = computed(() => menuStore.getColumns());
  const queryForm = reactive(menuStore.queryForm);

  onMounted(async () => {
    await menuStore.queryMenuList();
  });
</script>

<style scoped>
  .toolbar {
    display: flex;
    justify-content: flex-start;
  }
</style>
