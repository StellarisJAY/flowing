import { defineStore } from 'pinia';
import { createProvider, listProvider, updateProvider } from '@/api/ai/provider.api.js';
import { message } from 'ant-design-vue';

export const providerTypeOptions = [
  {
    label: 'OpenAI',
    value: 'openai',
  },
  {
    label: '通义千问/阿里云',
    value: 'dashscope',
  },
];

export const queryFormSchema = [
  {
    label: '供应商名称',
    name: 'providerName',
    type: 'input',
  },
];

export const providerFormSchema = [
  {
    label: '供应商名称',
    name: 'providerName',
    type: 'input',
  },
  {
    label: '供应商类型',
    name: 'providerType',
    type: 'select',
    options: () => providerTypeOptions,
  },
  {
    label: 'API地址',
    name: 'baseUrl',
    type: 'input',
    hidden: (formState) => formState.providerType !== 'openai',
  },
  {
    label: 'API密钥',
    name: 'apiKey',
    type: 'input',
    hidden: (formState) => formState.providerType !== 'openai' && formState.providerType !== 'dashscope',
  },
];

export const providerFormRules = {
  providerName: [{ required: true, message: '请输入供应商名称' }],
  providerType: [{ required: true, message: '请选择供应商类型' }],
  baseUrl: [{ required: true, message: '请输入API地址' }],
  apiKey: [{ required: true, message: '请输入API密钥' }],
};

export const useProviderStore = defineStore('ai_provider', {
  state: () => ({
    records: [],
    formState: {},
  }),
  actions: {
    async list(query) {
      try {
        const { data } = await listProvider(query);
        this.records = data;
      } catch {
        message.error('获取供应商列表失败');
      }
    },
    setFormState(state) {
      const providerConfig = JSON.parse(state.providerConfig);
      this.formState = {
        ...state,
        ...providerConfig,
      };
    },
    initFormState() {
      this.formState = {
        type: 'openai',
        baseUrl: 'https://api.openai.com/v1',
        apiKey: '',
        name: '',
      };
    },
    async save(record, isUpdate) {
      const config = {
        baseUrl: record.baseUrl,
        apiKey: record.apiKey,
      };
      const data = {
        providerName: record.providerName,
        providerType: record.providerType,
        providerConfig: JSON.stringify(config),
      };
      try {
        if (isUpdate) {
          await updateProvider(data);
        } else {
          await createProvider(data);
        }
        message.success('操作成功');
        return true;
      } catch {
        message.error('操作失败');
        return false;
      }
    },
  },
});
