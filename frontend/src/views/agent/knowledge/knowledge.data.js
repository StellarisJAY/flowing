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
    name: 'embeddingModel',
    label: '嵌入模型',
    type: 'select',
    placeholder: '请选择嵌入模型',
    defaultValue: '',
    options: () => [
      {
        label: 'OpenAI',
        value: 'openai',
      },
    ],
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
  name: [{ required: true, message: '请输入知识库名称', trigger: 'blur' }],
  description: [{ required: true, message: '请输入知识库描述', trigger: 'blur' }],
  embeddingModel: [{ required: true, message: '请选择嵌入模型', trigger: 'blur' }],
  enable: [{ required: true, message: '请选择是否启用', trigger: 'change' }],
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
