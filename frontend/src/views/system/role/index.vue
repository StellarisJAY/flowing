<template>
  <Table :columns="columns" :records="records" :query-form-schema="queryFormSchema" @refresh="refresh" :total="total">
    <template #tool-buttons>
      <Button type="primary" @click="()=>openRoleDrawer(false)">新增角色</Button>
    </template>
    <template #bodyCell="{column, record}">
      <Space v-if="column.dataIndex === 'action'">
        <Button type="link" size="small" @click="()=>openRoleDrawer(true, record)">编辑</Button>
        <Button type="link" size="small" @click="()=>openAuthDrawer(record.id)">授权</Button>
        <Button type="link" size="small" danger>删除</Button>
      </Space>
    </template>
  </Table>
  <RoleDrawer ref="roleDrawerRef" @submit-ok="refresh"/>
  <AuthDrawer ref="authDrawerRef"></AuthDrawer>
</template>

<script lang="js" setup>
  import Table from '@/components/Table/index.vue';
  import { useRoleStore } from '@/views/system/role/roleStore.js';
  import { computed, ref } from 'vue';
  import { Button, Space } from 'ant-design-vue';
  import RoleDrawer from '@/views/system/role/roleDrawer.vue';
  import AuthDrawer from '@/views/system/role/authDrawer.vue';

  const roleStore = useRoleStore();
  const columns = computed(() => roleStore.columns);
  const records = computed(() => roleStore.roleList);
  const queryFormSchema = computed(() => roleStore.queryFormSchema);
  const total = computed(() => roleStore.total);

  const roleDrawerRef = ref();
  const openRoleDrawer = (isUpdate, record)=>{
    roleDrawerRef.value.setVisible(true);
    if(isUpdate){
      roleStore.roleForm = record;
    }else {
      roleStore.resetRoleForm();
    }
  }

  const refresh = async (query)=>{
    await roleStore.getRoleList(query);
  }

  const authDrawerRef = ref();
  const openAuthDrawer = (id)=>{
    const role = roleStore.getRoleDetail(id);
    authDrawerRef.value.setVisible(true, role);
  }
</script>

<style scoped></style>
