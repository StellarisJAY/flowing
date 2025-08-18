import { defineStore } from 'pinia';
import { message } from 'ant-design-vue';
import { createModel, deleteModel, listModel, updateModel } from '@/api/ai/model.api.js';

export const modelTypeOptions = [
  {
    label: '大语言模型',
    value: 'llm',
  },
  {
    label: '文本嵌入模型',
    value: 'embedding',
  },
  {
    label: '重排序模型',
    value: 'rerank',
  },
];

export const columns = [
  {
    title: '模型名称',
    dataIndex: 'modelName',
    key: 'modelName',
  },
  {
    title: '模型类型',
    dataIndex: 'modelType',
    key: 'modelType',
  },
  {
    title: '是否启用',
    dataIndex: 'enable',
    key: 'enable',
  },
  {
    title: '操作',
    dataIndex: 'actions',
    key: 'actions',
  },
];

export const queryFormSchema = [
  {
    label: '模型名称',
    name: 'modelName',
    type: 'input',
  },
  {
    label: '模型类型',
    name: 'modelType',
    type: 'select',
    options: () => modelTypeOptions,
  },
];

export const modelFormSchema = [
  {
    label: '模型名称',
    name: 'modelName',
    type: 'input',
  },
  {
    label: '模型类型',
    name: 'modelType',
    type: 'select',
    options: () => modelTypeOptions,
  },
  {
    label: '是否启用',
    name: 'enable',
    type: 'switch',
  },
];

export const modelFormRules = {
  modelName: [{ required: true, message: '请输入模型名称' }],
  modelType: [{ required: true, message: '请选择模型类型' }],
  enable: [{ required: true, message: '请选择是否启用' }],
};

export const useModelStore = defineStore('ai_model', {
  state: () => ({
    records: [],
    formState: {},
  }),
  actions: {
    async list(query) {
      try {
        const { data } = await listModel(query);
        this.records = data;
      } catch {
        message.error('获取模型列表失败');
      }
    },
    async save(record, isUpdate) {
      try {
        if (isUpdate) {
          await updateModel(record);
        } else {
          await createModel(record);
        }
        message.success('操作成功');
        return true;
      } catch {
        message.error('操作失败');
        return false;
      }
    },
    async delete(id) {
      try {
        await deleteModel(id);
        message.success('删除成功');
        return true;
      } catch {
        message.error('删除失败');
        return false;
      }
    },
    setFormState(state) {
      this.formState = state;
    },
    resetFormState() {
      this.formState = {
        modelName: '',
        modelType: '',
        enable: false,
      };
    },
  },
});
