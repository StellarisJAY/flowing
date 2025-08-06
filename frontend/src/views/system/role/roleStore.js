import { defineStore } from 'pinia';
import { addRole, queryRoleList } from '@/views/system/role/api.js';

export const useRoleStore = defineStore('roleList', {
  state: ()=>({
    queryFormSchema: [
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
    ],
    roleForm: {
      roleName: "",
      roleKey: "",
      description: ""
    },
    roleList: [],
    total: 0,
    columns: [
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
    ],
    roleFormSchema: [
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
    ],
    roleFormRules: {
      roleName: [{ required: true, message: '请输入角色名称' }],
      roleKey: [{ required: true, message: '请输入角色Key' }],
      description: [{ required: true, message: '请输入描述' }],
    },
  }),
  actions: {
    async getRoleList(query){
      try {
        const res = await queryRoleList(query);
        this.roleList = res.data;
        this.total = res.total;
      }catch (err) {
        console.log(err);
      }
    },
    async createRole(data) {
      try {
        await addRole(data);
        return true;
      }catch (err) {
        console.log(err);
        return false;
      }
    },
    resetRoleForm() {
      this.roleForm = {
        roleName: "",
        roleKey: "",
        description: ""
      }
    },
    getRoleDetail(id) {
      return this.roleList.find(item=>item.id === id);
    }
  },
});
