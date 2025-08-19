<template>
  <div style="height: 100%; width: 100%">
    <Table
      :columns="columns"
      :records="records"
      :query-form-schema="queryFormSchema"
      @refresh="search"
      :pagination="true"
      ref="tableRef"
    >
      <template #tool-buttons>
        <IconButton type="primary" title="上传文档" icon="PlusOutlined" @click="openUploadModal" />
      </template>
      <template #bodyCell="{ column, record }">
        <div
          v-if="column.dataIndex === 'actions'"
          style="display: flex; justify-content: flex-start; gap: 10px"
        >
          <Button type="link" @click="() => download(record)">下载</Button>
          <Button type="link" @click="() => openRenameModal(record)">重命名</Button>
          <ConfirmButton title="删除" @confirm="() => handleDelete(record)" />
        </div>
        <div v-if="column.dataIndex === 'size'">
          {{ formatFileSize(record.size) }}
        </div>
        <Button type="link" v-if="column.dataIndex === 'originalName'">{{
          record.originalName
        }}</Button>
      </template>
    </Table>

    <UploadModal ref="uploadModalRef" :do-upload="uploadFile" @close="search" />
    <FormModal
      ref="renameModalRef"
      :form-schema="renameFormSchema"
      :form-rules="renameFormRules"
      :form-state="renameFormState"
      :submit="handleRename"
      @close="triggerQuery"
    />
  </div>
</template>

<script lang="js" setup>
  import Table from '@/components/Table/index.vue';
  import { computed, ref } from 'vue';
  import {
    formatFileSize,
    useDocumentStore,
    columns,
    queryFormSchema,
    upload,
    download,
    renameFormSchema,
    renameFormRules,
    deleteDoc,
  } from '@/views/agent/knowledge/document/document.data.js';
  import IconButton from '@/components/Button/IconButton.vue';
  import { useRoute } from 'vue-router';
  import UploadModal from '@/components/Modal/UploadModal.vue';
  import { Button } from 'ant-design-vue';
  import ConfirmButton from '@/components/Button/ConfirmButton.vue';
  import FormModal from '@/components/Modal/FormModal.vue';

  const knowledgeBaseId = ref('');
  const route = useRoute();
  knowledgeBaseId.value = route.query['knowledgeBaseId'];

  const documentStore = useDocumentStore();
  const records = computed(() => documentStore.records);
  const uploadModalRef = ref();
  const renameModalRef = ref();
  const tableRef = ref();
  const renameFormState = computed(() => documentStore.renameFormState);

  const search = async (query) => {
    const data = {
      knowledgeBaseId: knowledgeBaseId.value,
      ...query,
    };
    await documentStore.list(data);
  };

  const openUploadModal = () => {
    uploadModalRef.value.open();
  };

  const uploadFile = async (fileList) => {
    const formData = new FormData();
    formData.append('file', fileList[0]);
    formData.append('knowledgeBaseId', knowledgeBaseId.value);
    await upload(formData);
  };

  const openRenameModal = (record) => {
    documentStore.setRenameFormState({
      id: record.id,
      originalName: record.originalName,
    });
    renameModalRef.value.open(true);
  };

  const handleRename = async (formState) => {
    return await documentStore.rename(formState);
  };

  const triggerQuery = async () => {
    await tableRef.value.triggerQuery();
  };

  const handleDelete = async (record) => {
    await deleteDoc(record);
    await triggerQuery();
  };
</script>
