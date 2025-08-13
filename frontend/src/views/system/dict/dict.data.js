import { defineStore } from 'pinia';
import { apiDeleteDict, createDict, listDict, updateDict } from '@/api/system/dict.api.js';
import { message } from 'ant-design-vue';

export const columns = [
  {
    title: '字典名称',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '字典编码',
    dataIndex: 'code',
    key: 'code',
  },
  {
    title: '描述',
    dataIndex: 'description',
    key: 'description',
  },
  {
    title: '操作',
    dataIndex: 'actions',
    key: 'actions',
  },
];

export const queryFormSchema = [
  {
    label: '字典名称',
    type: 'input',
    name: 'name',
    placeholder: '请输入字典名称',
    defaultValue: '',
  },
  {
    label: '字典编码',
    type: 'input',
    placeholder: '请输入字典编码',
    defaultValue: '',
    name: 'code',
  },
];

export const dictFormSchema = [
  {
    label: '字典名称',
    type: 'input',
    placeholder: '请输入字典名称',
    defaultValue: '',
    name: 'name',
  },
  {
    label: '字典编码',
    type: 'input',
    placeholder: '请输入字典编码',
    defaultValue: '',
    name: 'code',
  },
  {
    label: '描述',
    type: 'input',
    placeholder: '请输入描述',
    defaultValue: '',
    name: 'description',
  },
];

export const dictFormRules = {
  name: [{ required: true, message: '请输入字典名称', trigger: 'blur' }],
  code: [{ required: true, message: '请输入字典编码', trigger: 'blur' }],
  description: [{ required: true, message: '请输入描述', trigger: 'blur' }],
};

export const deleteDict = async (id) => {
  try {
    await apiDeleteDict({ id });
    message.success('删除成功');
  } catch (err) {
    console.log(err);
    message.error('删除失败');
  }
}

export const useDictStore = defineStore('sys_dict', {
  state: () => ({
    dictList: [],
    dictForm: {
      name: '',
      code: '',
      description: '',
    },
  }),
  actions: {
    async fetchDictList(query) {
      const res = await listDict(query);
      this.dictList = res.data;
    },
    async saveDict(data, isUpdate) {
      try {
        if (isUpdate) {
          await updateDict(data);
        } else {
          await createDict(data);
        }
        message.success('保存成功');
        return true;
      } catch (err) {
        console.log(err);
        return false;
      }
    },
    initDictForm() {
      this.dictForm = {
        name: '',
        code: '',
        description: '',
      };
    },
  },
});
