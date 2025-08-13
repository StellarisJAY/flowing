import { defineStore } from 'pinia';
import { apiDeleteDictItem, createDictItem, getDictItems, updateDictItem } from '@/api/system/dict.api.js';
import { message } from 'ant-design-vue';

export const queryFormSchema = [
  {
    label: '键',
    name: 'itemKey',
    type: 'input',
    defaultValue: '',
    placeholder: '请输入键',
  },
];

export const columns = [
  {
    title: '键',
    dataIndex: 'itemKey',
    key: 'itemKey',
  },
  {
    title: '值',
    dataIndex: 'itemValue',
    key: 'itemValue',
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

export const dictItemFormSchema = [
  {
    label: '键',
    name: 'itemKey',
    type: 'input',
    defaultValue: '',
    placeholder: '请输入键',
  },
  {
    label: '值',
    name: 'itemValue',
    type: 'input',
    defaultValue: '',
    placeholder: '请输入值',
  },
  {
    label: '描述',
    name: 'description',
    type: 'input',
    defaultValue: '',
    placeholder: '请输入描述',
  },
  {
    label: '排序',
    name: 'sort',
    type: 'inputNumber',
    defaultValue: '0',
    placeholder: '请输入排序',
  },
  {
    label: '状态',
    name: 'enable',
    type: 'switch',
    defaultValue: true,
  },
];

export const dictItemFormRules = {
  itemKey: [{ required: true, message: '请输入键' }],
  itemValue: [{ required: true, message: '请输入值' }],
};

export const saveDictItem = async (data, isUpdate) => {
  try {
    if (isUpdate) {
      await updateDictItem(data);
    } else {
      await createDictItem(data);
    }
    message.success('操作成功');
    return true;
  } catch (err) {
    console.log(err);
    return false;
  }
};

export const deleteDictItem = async (id) => {
  try {
    await apiDeleteDictItem({id: id});
    message.success('删除成功');
  }catch (err) {
    message.error(err);
  }
}

export const useDictItemStore = defineStore('sys_dict_item', {
  state: () => ({
    dictItemList: [],
    total: 0,

    dictItemForm: {
      itemKey: '',
      itemValue: '',
      description: '',
    },
  }),
  actions: {
    async getDictItemList(query) {
      try {
        const { data, total } = await getDictItems(query);
        this.dictItemList = data;
        this.total = total;
      } catch (error) {
        console.log(error);
      }
    },
  },
});
