<template>
  <div style="height: 100%; width: 100%">
    <Table
      :columns="columns"
      :records="records"
      :query-form-schema="queryFormSchema"
      @refresh="search"
      :pagination="true"
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
          <Button type="link">重命名</Button>
          <ConfirmButton title="删除" />
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
  } from '@/views/agent/knowledge/document/document.data.js';
  import IconButton from '@/components/Button/IconButton.vue';
  import { useRoute } from 'vue-router';
  import UploadModal from '@/components/Modal/UploadModal.vue';
  import { Button } from 'ant-design-vue';
  import ConfirmButton from '@/components/Button/ConfirmButton.vue';

  const knowledgeBaseId = ref('');
  const route = useRoute();
  knowledgeBaseId.value = route.query['knowledgeBaseId'];

  const documentStore = useDocumentStore();
  const records = computed(() => documentStore.records);
  const uploadModalRef = ref();

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
</script>
