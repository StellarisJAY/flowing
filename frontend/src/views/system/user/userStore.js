import { defineStore } from 'pinia';
import { queryUserList } from '@/views/system/user/api.js';

export const searchUserFormSchema = [
  {
    name: "username",
    label: "用户名",
    type: 'input',
    placeholder: '请输入用户名',
    defaultValue: ''
  },
  {
    name: "nickName",
    label: "昵称",
    type: 'input',
    placeholder: '请输入昵称',
    defaultValue: ''
  },
  {
    name: "phone",
    label: "手机号",
    type: 'input',
    placeholder: '请输入手机号',
    defaultValue: ''
  }
];

export const useUserStore = defineStore('userTable', {
  state: ()=>({
    columns: [
      {
        title: "用户名",
        key: "username",
        dataIndex: "username"
      },
      {
        title: "昵称",
        key: "nickName",
        dataIndex: "nickName"
      },
      {
        title: "手机号",
        key: "phone",
        dataIndex: "phone"
      },
      {
        title: "邮箱",
        key: "email",
        dataIndex: "email"
      }
    ],
    records: []
  }),
  actions: {
    async refresh(query) {
      const res = await queryUserList(query);
      console.log(res);
      this.records = res;
    },
    getRecords() {
      return this.records;
    },
    getColumns() {
      return this.columns;
    }
  }
});
