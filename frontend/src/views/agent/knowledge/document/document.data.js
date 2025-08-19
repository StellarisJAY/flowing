import { defineStore } from 'pinia';
import { message } from 'ant-design-vue';
import { getDownloadUrl, listDocument, uploadDocument } from '@/api/ai/document.api.js';

export const queryFormSchema = [
  {
    label: '文档名称',
    name: 'name',
    type: 'input',
    placeholder: '请输入文档名称',
  },
];

export const columns = [
  {
    title: '文档名称',
    dataIndex: 'originalName',
    key: 'originalName',
  },
  {
    title: '文档类型',
    dataIndex: 'type',
    key: 'type',
  },
  {
    title: '文档大小',
    dataIndex: 'size',
    key: 'size',
  },
  {
    title: '操作',
    dataIndex: 'actions',
    key: 'actions',
  },
];

export const upload = async (formData) => {
  await uploadDocument(formData);
};

export const formatFileSize = (size) => {
  if (size < 1024) {
    return `${size} B`;
  }
  if (size < 1024 * 1024) {
    return `${(size / 1024).toFixed(2)} KB`;
  }
  if (size < 1024 * 1024 * 1024) {
    return `${(size / 1024 / 1024).toFixed(2)} MB`;
  }
  return `${(size / 1024 / 1024 / 1024).toFixed(2)} GB`;
};

export const download = async (record) => {
  try {
    const {data} = await getDownloadUrl(record.id);
    window.open(data, record.originalName);
  } catch {
    message.error('下载失败');
  }
};

export const useDocumentStore = defineStore('kb_document', {
  state: () => ({
    records: [],
    total: 0,
    formState: {},
  }),
  actions: {
    async list(query) {
      try {
        const res = await listDocument(query);
        this.records = res.data;
        this.total = res.total;
      } catch {
        message.error('获取文档列表失败');
      }
    },
    setFormState(state) {
      this.formState = state;
    },
    initFormState() {
      this.formState = {};
    },
  },
});
