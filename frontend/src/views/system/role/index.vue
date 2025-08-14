<template>
  <Table
    :columns="columns"
    :records="records"
    :query-form-schema="queryFormSchema"
    @refresh="refresh"
    :total="total"
  >
    <template #tool-buttons>
      <IconButton type="primary" @click="()=>openRoleDrawer(false)" icon="PlusOutlined" title="新增角色" />
    </template>
    <template #bodyCell="{ column, record }">
      <Space v-if="column.dataIndex === 'action'">
        <Button type="link" size="small" @click="() => openRoleDrawer(true, record)">编辑</Button>
        <Button type="link" size="small" @click="() => openAuthDrawer(record.id)">授权</Button>
        <ConfirmButton
          text="删除"
          @confirm="
            async () => {
              await deleteRole(record.id);
              await refresh();
            }
          "
        />
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
    title="角色"
  />
  <AuthDrawer ref="authDrawerRef"></AuthDrawer>
</template>

<script lang="js" setup>
  import Table from '@/components/Table/index.vue';
  import {
    columns,
    queryFormSchema,
    roleFormRules,
    roleFormSchema,
    saveRole,
    deleteRole,
    useRoleStore,
  } from '@/views/system/role/role.data.js';
  import { computed, ref } from 'vue';
  import { Button, Space } from 'ant-design-vue';
  import FormDrawer from '@/components/Drawer/FormDrawer.vue';
  import AuthDrawer from '@/views/system/role/authDrawer.vue';
  import ConfirmButton from '@/components/Button/ConfirmButton.vue';
  import IconButton from '@/components/Button/IconButton.vue';

  const roleStore = useRoleStore();
  const records = computed(() => roleStore.roleList);
  const total = computed(() => roleStore.total);

  const roleForm = computed(() => roleStore.roleForm);
  const roleFormDrawerRef = ref();

  const openRoleDrawer = (isUpdate, record) => {
    if (isUpdate) {
      roleStore.roleForm = record;
    } else {
      roleStore.resetRoleForm();
    }
    roleFormDrawerRef.value.open(isUpdate);
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
