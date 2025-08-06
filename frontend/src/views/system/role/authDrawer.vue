<template>
  <Drawer :open="visible" @close="() => setVisible(false)" size="large" destroy-on-close>
    <div class="menu-tree">
      <Tree checkable :checked-keys="checkedKeys" :tree-data="treeData"></Tree>
    </div>
    <div class="actions">
      <Space>
        <Button type="primary">保存</Button>
        <Button>保存并关闭</Button>
      </Space>
    </div>
  </Drawer>
</template>

<script lang="js" setup>
import { Drawer, Tree, Space, Button} from 'ant-design-vue';
import { computed, ref } from 'vue';
import { useMenuStore } from '@/views/system/menu/menuStore.js';
import {cloneDeep} from 'lodash';

const menuStore = useMenuStore();
const visible = ref(false);
const role = ref();

const emit = defineEmits(['submit-ok', 'submit-fail']);

const setVisible = (val, record) => {
  if (val) {
    role.value = record;
    menuStore.queryMenuList().then(mergeMenuTrees);
  }
  visible.value = val;
};

const treeData = ref([]);
const checkedKeys = ref([]);

const mergeMenuTrees = ()=>{
  const menuTrees = cloneDeep(menuStore.menuList);
  const roleMenus = role.value.menus;
  const roleMenuMap = {};
  checkedKeys.value = [];
  const storeRoleMenu = (menus)=>{
    menus.forEach(item=>{
      roleMenuMap[item.id] = item;
      if (item.children) {
        storeRoleMenu(item.children);
      }
    });
  };

  const setMenuChecked = (menus)=>{
    menus.forEach(item=>{
      item.checked = roleMenuMap[item.id] !== undefined;
      item.title = item.menuName;
      item.key = item.id;
      if (item.checked) {
        checkedKeys.value.push(item.id);
      }
      if (item.children) {
        setMenuChecked(item.children);
      }
    });
  };
  if (roleMenus && menuTrees) {
    storeRoleMenu(roleMenus);
    setMenuChecked(menuTrees);
  }
  treeData.value = menuTrees;
}

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
