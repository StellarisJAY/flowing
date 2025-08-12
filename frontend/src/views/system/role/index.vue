<template>
  <Table
    :columns="columns"
    :records="records"
    :query-form-schema="queryFormSchema"
    @refresh="refresh"
    :total="total"
  >
    <template #tool-buttons>
      <Button type="primary" @click="() => openRoleDrawer(false)">新增角色</Button>
    </template>
    <template #bodyCell="{ column, record }">
      <Space v-if="column.dataIndex === 'action'">
        <Button type="link" size="small" @click="() => openRoleDrawer(true, record)">编辑</Button>
        <Button type="link" size="small" @click="() => openAuthDrawer(record.id)">授权</Button>
        <Button type="link" size="small" danger>删除</Button>
      </Space>
    </template>
  </Table>
  <FormDrawer
     ref="roleFormDrawerRef"
    :form-state="roleForm"
    :form-schema="roleFormSchema"
    :form-rules="roleFormRules"
    :submit="saveRole"
    @close="refresh"
  />
  <AuthDrawer ref="authDrawerRef"></AuthDrawer>
</template>

<script lang="js" setup>
  import Table from '@/components/Table/index.vue';
  import { useRoleStore } from '@/views/system/role/roleStore.js';
  import { computed, ref } from 'vue';
  import { Button, Space } from 'ant-design-vue';
  import FormDrawer from '@/components/Drawer/FormDrawer.vue';
  import AuthDrawer from '@/views/system/role/authDrawer.vue';

  const roleStore = useRoleStore();
  const columns = computed(() => roleStore.columns);
  const records = computed(() => roleStore.roleList);
  const queryFormSchema = computed(() => roleStore.queryFormSchema);
  const total = computed(() => roleStore.total);

  const roleFormSchema = computed(() => roleStore.roleFormSchema);
  const roleForm = computed(() => roleStore.roleForm);
  const roleFormRules = computed(() => roleStore.roleFormRules);
  const roleFormDrawerRef = ref();

  const openRoleDrawer = (isUpdate, record) => {
    if (isUpdate) {
      roleStore.roleForm = record;
    } else {
      roleStore.resetRoleForm();
    }
    roleFormDrawerRef.value.open(isUpdate);
  };

  const saveRole = async (data, isUpdate) => {
    return await roleStore.saveRole(data, isUpdate);
  };

  const refresh = async (query) => {
    await roleStore.getRoleList(query);
  };

  const authDrawerRef = ref();
  const openAuthDrawer = (id) => {
    const role = roleStore.getRoleDetail(id);
    authDrawerRef.value.setVisible(true, role);
  };
</script>

<style scoped></style>
