import { defineStore } from 'pinia';

export const knowledgeFormSchema = [
  {
    name: 'name',
    label: '知识库名称',
    type: 'input',
    placeholder: '请输入知识库名称',
    defaultValue: '',
  },
  {
    name: 'description',
    label: '知识库描述',
    type: 'input',
    placeholder: '请输入知识库描述',
    defaultValue: '',
  },
  {
    name: 'datasource',
    label: '数据源(配置后不可修改)',
    type: 'datasourceSelect',
    placeholder: '请选择数据源',
    defaultValue: '',
    componentProps: {
      datasourceType: 'milvus',
    },
    disabled: (formState) => formState.id !== undefined,
  },
  {
    name: 'embeddingModel',
    label: '嵌入模型(配置后不可修改)',
    type: 'modelSelect',
    placeholder: '请选择嵌入模型',
    defaultValue: '',
    componentProps: {
      modelType: 'embedding',
    },
    disabled: (formState) => formState.id !== undefined,
  },
  {
    name: 'enable',
    label: '是否启用',
    type: 'switch',
    placeholder: '请选择是否启用',
    defaultValue: false,
  },
];

export const knowledgeFormRules = {
  name: [{ required: true, message: '请输入知识库名称', trigger: 'submit' }],
  description: [{ required: true, message: '请输入知识库描述', trigger: 'submit' }],
  embeddingModel: [
    { required: true, message: '请选择嵌入模型(请先在模型管理中配置嵌入模型)', trigger: 'submit' },
  ],
  enable: [{ required: true, message: '请选择是否启用', trigger: 'submit' }],
  datasource: [
    { required: true, message: '请选择数据源(请先在数据源管理中配置向量库)', trigger: 'submit' },
  ],
};

export const searchFormSchema = [
  {
    name: 'name',
    label: '知识库名称',
    type: 'input',
    placeholder: '请输入知识库名称',
    defaultValue: '',
  },
];

export const useKnowledgeStore = defineStore('agent_knowledge', {
  state: () => ({
    records: [
      {
        id: 1,
        name: '知识库1',
        description: '知识库1描述',
        embeddingModel: 'openai',
        enable: true,
      },
      {
        id: 2,
        name: '知识库2',
        description: '知识库2描述',
        embeddingModel: 'openai',
        enable: true,
      },
      {
        id: 3,
        name: '知识库3',
        description: '知识库3描述',
        embeddingModel: 'openai',
        enable: true,
      },
    ],
    total: 0,
    knowledgeForm: {
      name: '',
      description: '',
      embeddingModel: '',
      enable: false,
    },
  }),
  actions: {
    initKnowledgeForm() {
      this.knowledgeForm = {
        name: '',
        description: '',
        embeddingModel: '',
        enable: false,
      };
    },
    setKnowledgeForm(form) {
      this.knowledgeForm = form;
    },
  },
});
