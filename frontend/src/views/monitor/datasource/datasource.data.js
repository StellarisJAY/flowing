import { defineStore } from 'pinia';
import {
  createDatasource,
  pingDatasource,
  queryDatasourceList,
  updateDatasource,
} from '@/api/monitor/datasource.api.js';
import { message } from 'ant-design-vue';

export const columns = [
  {
    title: '数据源名称',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '数据源编码',
    dataIndex: 'code',
    key: 'code',
  },
  {
    title: '数据源类型',
    dataIndex: 'type',
    key: 'type',
  },
  {
    title: '主机',
    dataIndex: 'host',
    key: 'host',
  },
  {
    title: '端口',
    dataIndex: 'port',
    key: 'port',
  },
  {
    title: '数据库名',
    dataIndex: 'database',
    key: 'database',
  },
  {
    title: '操作',
    dataIndex: 'actions',
    key: 'actions',
  },
];

export const queryFormSchema = [
  {
    label: '数据源名称',
    type: 'input',
    name: 'name',
    placeholder: '请输入数据源名称',
    defaultValue: '',
  },
  {
    label: '数据源编码',
    type: 'input',
    placeholder: '请输入数据源编码',
    defaultValue: '',
    name: 'code',
  },
  {
    label: '主机',
    type: 'input',
    placeholder: '请输入主机',
    defaultValue: '',
    name: 'host',
  },
  {
    label: '数据源类型',
    type: 'select',
    placeholder: '请输入数据源类型',
    defaultValue: '',
    name: 'type',
    options: () => datasourceTypeOptions,
  },
];

export const datasourceTypeOptions = [
  {
    label: 'MySQL',
    value: 'mysql',
  },
  {
    label: 'PostgreSQL',
    value: 'postgresql',
  },
  {
    label: 'Milvus',
    value: 'milvus',
  },
  {
    label: 'Redis',
    value: 'redis',
  },
];

export const datasourceFormSchema = [
  {
    label: '数据源名称',
    type: 'input',
    placeholder: '请输入数据源名称',
    defaultValue: '',
    name: 'name',
  },
  {
    label: '数据源编码',
    type: 'input',
    placeholder: '请输入数据源编码',
    defaultValue: '',
    name: 'code',
    disabled: (formState) => formState.id !== undefined,
  },
  {
    label: '数据源类型',
    type: 'select',
    placeholder: '请输入数据源类型',
    defaultValue: '',
    name: 'type',
    options: () => datasourceTypeOptions,
  },
  {
    label: '主机',
    type: 'input',
    placeholder: '请输入主机',
    defaultValue: '',
    name: 'host',
  },
  {
    label: '端口',
    slot: true,
    name: 'port',
  },
  {
    label: '数据库名',
    type: 'input',
    placeholder: '请输入数据库名',
    defaultValue: '',
    name: 'database',
  },
  {
    label: '用户名',
    type: 'input',
    placeholder: '请输入用户名',
    defaultValue: '',
    name: 'username',
  },
  {
    label: '密码',
    type: 'inputPassword',
    placeholder: '请输入密码',
    defaultValue: '',
    name: 'password',
  },
  {
    label: '备注',
    type: 'input',
    placeholder: '请输入备注',
    defaultValue: '',
    name: 'description',
  },
];

export const datasourceFormRules = {
  name: [{ required: true, message: '请输入数据源名称', trigger: 'blur' }],
  code: [{ required: true, message: '请输入数据源编码', trigger: 'blur' }],
  type: [{ required: true, message: '请输入数据源类型', trigger: 'blur' }],
  host: [{ required: true, message: '请输入主机', trigger: 'blur' }],
  port: [{ required: true, message: '请输入端口', trigger: 'blur' }],
  description: [{ required: true, message: '请输入备注', trigger: 'blur' }],
};

export const testConnection = async (req) => {
  try {
    const { data } = await pingDatasource(req);
    return `连接成功，延迟${data}ms`;
  } catch {
    return '连接失败';
  }
};

export const useDatasourceStore = defineStore('monitor_datasource', {
  state: () => ({
    records: [],
    total: 0,
    formState: {},
  }),
  actions: {
    async getDatasourceList(query) {
      try {
        const { data, total } = await queryDatasourceList(query);
        console.log(data);
        this.setRecords(data);
        this.setTotal(total);
      } catch {
        message.error('获取数据源列表失败');
      }
    },
    async save(data, isUpdate) {
      try {
        if (isUpdate) {
          await updateDatasource(data);
        } else {
          await createDatasource(data);
        }
        message.success('操作成功');
        return true;
      } catch {
        message.error('操作失败');
        return false;
      }
    },
    setRecords(records) {
      this.records = records;
    },
    setTotal(total) {
      this.total = total;
    },
    setFormState(formState) {
      this.formState = formState;
    },
    initFormState() {
      this.formState = {
        name: '',
        code: '',
        type: '',
        host: '',
        port: '',
        database: '',
        username: '',
        password: '',
        description: '',
      };
    },
  },
});
