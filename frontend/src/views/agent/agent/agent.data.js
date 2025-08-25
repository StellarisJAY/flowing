import { defineStore } from 'pinia';
import { createAgent, queryAgentList, updateAgent } from '@/api/ai/agent.api.js';
import { message } from 'ant-design-vue';

export const agentTypeOptions = [
  {
    label: '聊天',
    value: 'simple',
  },
  {
    label: '工作流',
    value: 'workflow',
  }
];

export const queryFormSchema = [
  {
    label: '智能体名称',
    name: 'name',
    type: 'input',
    placeholder: '请输入智能体名称',
  },
  {
    label: '智能体类型',
    name: 'type',
    type: 'select',
    placeholder: '请选择智能体类型',
    options: () => agentTypeOptions,
  },
  {
    label: '我创建的',
    name: 'private',
    type: 'checkbox',
    default: false,
  }
];

export const agentFormSchema = [
  {
    label: '智能体名称',
    name: 'name',
    type: 'input',
    placeholder: '请输入智能体名称',
  },
  {
    label: '智能体类型',
    name: 'type',
    type: 'select',
    placeholder: '请选择智能体类型',
    options: () => agentTypeOptions,
    disabled: (state) => state.id !== undefined,
  },
  {
    label: '智能体描述',
    name: 'description',
    type: 'input',
    placeholder: '请输入智能体描述',
  },
  {
    label: '公开',
    name: 'public',
    type: 'switch',
    default: false,
  }
];

export const agentFormRules = {
  name: [
    { required: true, message: '请输入智能体名称', trigger: 'blur' },
  ],
  type: [
    { required: true, message: '请选择智能体类型', trigger: 'blur' },
  ],
  description: [
    { required: true, message: '请输入智能体描述', trigger: 'blur' },
  ],
  public: [
    { required: true, message: '请选择公开状态', trigger: 'change' },
  ],
}

export const useAgentStore = defineStore('ai_agent', {
  state: () => ({
    records: [],
    formState: {},
  }),
  actions: {
    async list(query) {
      try {
        const { data } = await queryAgentList(query);
        this.records = data;
      } catch {
        message.error('查询智能体列表失败');
      }
    },
    async save(data, isUpdate) {
      try {
        if (isUpdate) {
          await updateAgent(data);
        } else {
          await createAgent(data);
        }
        message.success('操作成功');
        return true;
      } catch {
        message.error('操作失败');
        return false;
      }
    },
    initFormState() {
      this.formState = {
        id: undefined,
        name: '',
        type: 'simple',
        description: '',
        public: false,
      };
    },
    setFormState(record) {
      this.formState = record;
    },
  }
})
