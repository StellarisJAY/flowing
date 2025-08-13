<template>
  <Drawer
    :open="visible"
    @close="() => setVisible(false)"
    size="large"
    destroy-on-close
    title="修改角色权限"
  >
    <div class="menu-tree">
      <Tree
        checkable
        :checked-keys="checkedKeys"
        :selectable="false"
        :tree-data="treeData"
        :check-strictly="true"
        @check="handleCheck"
      />
    </div>
    <div class="actions">
      <Space>
        <Button type="primary" @click="handleSubmit">保存</Button>
        <Button @click="() => handleSubmit(true)">保存并关闭</Button>
      </Space>
    </div>
  </Drawer>
</template>

<script lang="js" setup>
  import { Drawer, Tree, Space, Button } from 'ant-design-vue';
  import { computed, ref } from 'vue';
  import { useRoleStore } from '@/views/system/role/role.data.js';
  import { saveRoleMenus } from '@/api/system/role.api.js';

  const roleStore = useRoleStore();
  const visible = ref(false);
  const role = ref();

  const emit = defineEmits(['submit-ok', 'submit-fail']);
  const treeData = computed(() => roleStore.roleMenus);
  const checkedKeys = computed(() => roleStore.checkedKeys);

  const setVisible = (val, record) => {
    if (val) {
      role.value = record;
      roleStore.getRoleMenus({ id: record.id });
    }
    visible.value = val;
  };

  const handleCheck = (e) => {
    roleStore.setCheckedKeys(e);
  };
  const handleSubmit = async (close) => {
    try {
      await saveRoleMenus({
        roleId: role.value.id,
        menuIds: roleStore.checkedKeys.checked
          ? roleStore.checkedKeys.checked
          : roleStore.checkedKeys,
      });
      if (close === true) setVisible(false);
      emit('submit-ok');
    } catch {
      emit('submit-fail');
    }
  };

  defineExpose({
    setVisible,
  });
</script>

<style scoped>
  .menu-tree {
    height: 90%;
    overflow: auto;
    margin-bottom: 20px;
  }
  .actions {
    text-align: right;
  }
</style>
