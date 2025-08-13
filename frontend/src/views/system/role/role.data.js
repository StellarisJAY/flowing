import { defineStore } from 'pinia';
import { addRole, queryRoleList, getRoleMenus, updateRole, apiDeleteRole } from '@/api/system/role.api.js';
import { message } from 'ant-design-vue';

export const queryFormSchema = [
  {
    label: '角色名称',
    name: 'roleName',
    type: 'input',
    defaultValue: '',
    placeholder: '请输入角色名称',
  },
  {
    label: '角色Key',
    name: 'roleKey',
    type: 'input',
    defaultValue: '',
    placeholder: '请输入角色Key',
  },
];

export const columns = [
  {
    title: '角色名称',
    dataIndex: 'roleName',
    key: 'roleName',
  },
  {
    title: '角色Key',
    dataIndex: 'roleKey',
    key: 'roleKey',
  },
  {
    title: '描述',
    dataIndex: 'description',
    key: 'description',
  },
  {
    title: '操作',
    dataIndex: 'action',
    key: 'action',
  },
];

export const roleFormSchema = [
  {
    label: '角色名称',
    name: 'roleName',
    type: 'input',
    defaultValue: '',
    placeholder: '请输入角色名称',
  },
  {
    label: '角色Key',
    name: 'roleKey',
    type: 'input',
    defaultValue: '',
    placeholder: '请输入角色Key',
  },
  {
    label: '描述',
    name: 'description',
    type: 'input',
    defaultValue: '',
    placeholder: '请输入描述',
  },
];

export const roleFormRules = {
  roleName: [{ required: true, message: '请输入角色名称' }],
  roleKey: [{ required: true, message: '请输入角色Key' }],
  description: [{ required: true, message: '请输入描述' }],
};

export const saveRole = async (data, isUpdate) => {
  try {
    if (isUpdate) {
      await updateRole(data);
    } else {
      await addRole(data);
    }
    message.success('保存成功');
    return true;
  } catch (err) {
    console.log(err);
    return false;
  }
};

export const deleteRole = async (id) => {
  try {
    await apiDeleteRole({id: id});
    message.success('删除成功');
    return true;
  } catch (err) {
    console.log(err);
    return false;
  }
}

export const useRoleStore = defineStore('roleList', {
  state: () => ({
    roleForm: {
      roleName: '',
      roleKey: '',
      description: '',
    },
    roleList: [],
    total: 0,
    roleMenus: [],
    checkedKeys: [],
  }),
  actions: {
    async getRoleList(query) {
      try {
        const res = await queryRoleList(query);
        this.roleList = res.data;
        this.total = res.total;
      } catch (err) {
        console.log(err);
      }
    },
    async getRoleMenus(query) {
      const { data } = await getRoleMenus(query);
      const { menus, checkedKeys } = data;
      this.roleMenus = menus;
      this.checkedKeys = checkedKeys;
    },
    resetRoleForm() {
      this.roleForm = {
        roleName: '',
        roleKey: '',
        description: '',
      };
    },
    getRoleDetail(id) {
      return this.roleList.find((item) => item.id === id);
    },
    setCheckedKeys(keys) {
      this.checkedKeys = keys;
    },
  },
});
