import { defineStore } from 'pinia';
import { addUser, getUserDetail, queryUserList, updateUser } from '@/api/system/user.api.js';
import { message } from 'ant-design-vue';

export const searchUserFormSchema = [
  {
    name: 'username',
    label: '用户名',
    type: 'input',
    placeholder: '请输入用户名',
    defaultValue: '',
  },
  {
    name: 'nickName',
    label: '昵称',
    type: 'input',
    placeholder: '请输入昵称',
    defaultValue: '',
  },
  {
    name: 'phone',
    label: '手机号',
    type: 'input',
    placeholder: '请输入手机号',
    defaultValue: '',
  },
];

export const columns = [
  {
    title: '用户名',
    key: 'username',
    dataIndex: 'username',
  },
  {
    title: '昵称',
    key: 'nickName',
    dataIndex: 'nickName',
  },
  {
    title: '手机号',
    key: 'phone',
    dataIndex: 'phone',
  },
  {
    title: '邮箱',
    key: 'email',
    dataIndex: 'email',
  },
  {
    title: '操作',
    key: 'actions',
    dataIndex: 'actions',
  }
];

export const userFormSchema = [
  {
    name: 'username',
    label: '用户名',
    type: 'input',
    placeholder: '请输入用户名',
    defaultValue: '',
    disabled: (formState)=>formState.id !== undefined,
  },
  {
    name: 'nickName',
    label: '昵称',
    type: 'input',
    placeholder: '请输入昵称',
    defaultValue: '',
  },
  {
    name: 'password',
    label: '密码',
    type: 'inputPassword',
    placeholder: '请输入密码',
    defaultValue: '',
    hidden: (formState)=>formState.id !== undefined,
  },
  {
    name: 'phone',
    label: '手机号',
    type: 'input',
    placeholder: '请输入手机号',
    defaultValue: '',
  },
  {
    name: 'email',
    label: '邮箱',
    type: 'input',
    placeholder: '请输入邮箱',
    defaultValue: '',
  },
  {
    name: 'roleIds',
    label: '角色',
    type: 'selectRole',
    placeholder: '请选择角色',
    defaultValue: [],
    componentProps: {
      multiple: true,
    },
  },
];

export const userFormRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
  ],
  nickName: [
    { required: true, message: '请输入昵称', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
  ],
};

export const saveUser = async (data, isUpdate) => {
  try {
    if (isUpdate) {
      await updateUser(data);
    }else {
      await addUser(data);
    }
    message.success("操作成功");
    return true;
  }catch (err) {
    console.error(err);
    message.error("操作失败");
    return false;
  }
};

export const getDetail = async (id) => {
  try {
    const { data } = await getUserDetail({id});
    return data;
  }catch (err) {
    console.error(err);
    message.error("查询用户详情失败");
    return null;
  }
};

export const useUserStore = defineStore('userTable', {
  state: () => ({
    records: [],
    total: 0,
    userForm: {
      username: '',
      nickName: '',
      password: '',
      phone: '',
      email: '',
      roleIds: [],
    }
  }),
  actions: {
    async refresh(query) {
      try {
        const res = await queryUserList(query);
        this.records = res.data;
        this.total = res.total;
      }catch (err) {
        console.error(err);
        message.error("查询用户列表失败");
      }
    },
    initUserForm() {
      this.userForm = {
        username: '',
        nickName: '',
        password: '',
        phone: '',
        email: '',
      }
    },
    setUserForm(form) {
      this.userForm = form;
    }
  },
});
