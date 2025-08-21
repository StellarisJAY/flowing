import { defineStore } from 'pinia';
import { message } from 'ant-design-vue';
import {
  cancelParse,
  deleteDocument,
  getDownloadUrl,
  listChunks,
  listDocument,
  parseDocument,
  renameDocument,
  uploadDocument,
} from '@/api/ai/document.api.js';

import iconPdf from '@/assets/svg/ext_pdf_icon.svg';
import iconFile from '@/assets/svg/ext_file_generic_icon.svg';
import iconDoc from '@/assets/svg/ext_doc_icon.svg';
import iconDocx from '@/assets/svg/ext_docx_icon.svg';
import iconMd from '@/assets/svg/ext_md_icon.svg';
import iconTxt from '@/assets/svg/ext_txt_icon.svg';

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
    title: '',
    dataIndex: 'icon',
    key: 'icon',
  },
  {
    title: '文档名称',
    dataIndex: 'originalName',
    key: 'originalName',
  },
  {
    title: '文档大小',
    dataIndex: 'size',
    key: 'size',
  },
  {
    title: '切片数量',
    dataIndex: 'sliceCount',
    key: 'sliceCount',
  },
  {
    title: '解析状态',
    dataIndex: 'task',
    key: 'task',
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
    const { data } = await getDownloadUrl(record.id);
    window.open(data, record.originalName);
  } catch {
    message.error('下载失败');
  }
};

export const renameFormSchema = [
  {
    label: '文档名称',
    name: 'originalName',
    type: 'input',
    placeholder: '请输入文档名称',
  },
];

export const renameFormRules = {
  originalName: [{ required: true, message: '请输入文档名称', trigger: 'submit' }],
};

export const deleteDoc = async (record) => {
  try {
    await deleteDocument(record.id);
    message.success('删除成功');
  } catch {
    message.error('删除失败');
  }
};

export const parse = async (data) => {
  try {
    await parseDocument(data);
  } catch {
    message.error('开始解析失败');
  }
};

export const cancel = async (id) => {
  try {
    await cancelParse(id);
  } catch {
    message.error('取消解析失败');
  }
};

export const isTaskFailed = (task) => {
  return task?.status === 'failed';
};

export const isTaskSuccess = (task) => {
  return task?.status === 'success';
};

export const isTaskRunning = (task) => {
  return task?.status === 'new' || task?.status === 'slicing' || task?.status === 'embedding';
};

export const getDocumentSuffix = (record) => {
  return record.originalName.split('.').pop();
};

export const getDocumentIcon = (record) => {
  const suffix = getDocumentSuffix(record);
  switch (suffix) {
    case 'pdf':
      return iconPdf;
    case 'doc':
      return iconDoc;
    case 'docx':
      return iconDocx;
    case 'md':
      return iconMd;
    case 'txt':
      return iconTxt;
    default:
      return iconFile;
  }
};

export const listDocChunks = async (query) => {
  try {
    const res = await listChunks(query);
    console.log(res);
  } catch {
    message.error('获取文档切片失败');
  }
};

export const useDocumentStore = defineStore('kb_document', {
  state: () => ({
    records: [],
    total: 0,
    formState: {},
    renameFormState: {},
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
    async rename(formState) {
      try {
        await renameDocument(formState);
        message.success('重命名成功');
        return true;
      } catch {
        message.error('重命名失败');
        return false;
      }
    },
    setRenameFormState(state) {
      this.renameFormState = state;
    },
    setFormState(state) {
      this.formState = state;
    },
    initFormState() {
      this.formState = {};
    },
  },
});
