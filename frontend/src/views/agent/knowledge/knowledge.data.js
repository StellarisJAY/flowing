import { defineStore } from 'pinia';
import { addKnowledge, deleteKnowledge, queryKnowledgeList, updateKnowledge } from '@/api/ai/knowledge.api.js';
import { message } from 'ant-design-vue';

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
    name: 'datasourceId',
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

export const deleteKb = async (id) => {
  try {
    await deleteKnowledge(id);
    message.success('删除成功');
    return true;
  } catch  {
    message.error('删除失败');
    return false;
  }
}

export const useKnowledgeStore = defineStore('agent_knowledge', {
  state: () => ({
    records: [],
    total: 0,
    knowledgeForm: {},
  }),
  actions: {
    async list(query) {
      try {
        const res = await queryKnowledgeList(query);
        this.records = res.data;
        this.total = res.total;
      } catch  {
        message.error('获取知识库列表失败');
      }
    },
    async saveKnowledgeBase(record, isUpdate) {
      try {
        if (isUpdate) {
          await updateKnowledge(record);
        } else {
          await addKnowledge(record);
        }
        message.success('操作成功');
        return true;
      } catch  {
        message.error('操作失败');
        return false;
      }
    },
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
