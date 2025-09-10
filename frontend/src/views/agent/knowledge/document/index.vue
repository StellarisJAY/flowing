<template>
  <div style="height: 100%; width: 100%">
    <Table
      :columns="columns"
      :records="records"
      :query-form-schema="queryFormSchema"
      @refresh="search"
      :total="total"
      :pagination="true"
      ref="tableRef"
    >
      <template #tool-buttons>
        <IconButton type="primary" title="上传文档" icon="PlusOutlined" @click="openUploadModal" />
      </template>
      <template #bodyCell="{ column, record }">
        <div v-if="column.dataIndex === 'icon'">
          <img :src="getDocumentIcon(record)" alt="icon" style="width: 32px; height: 32px" />
        </div>
        <div
          v-if="column.dataIndex === 'actions'"
          style="display: flex; justify-content: flex-start; gap: 10px"
        >
          <Button type="link" v-if="!record.task" @click="() => handleParse(record)"
            >开始解析</Button
          >
          <Button
            type="link"
            v-else-if="isTaskSuccess(record.task) || isTaskFailed(record.task)"
            @click="() => handleParse(record)"
            >重新解析</Button
          >
          <Button
            type="link"
            v-else-if="isTaskRunning(record.task)"
            @click="() => handleCancel(record)"
            >取消解析</Button
          >
          <Button type="link" @click="() => download(record)">下载</Button>
          <Button type="link" @click="() => openRenameModal(record)">重命名</Button>
          <ConfirmButton
            title="删除"
            v-if="!isTaskRunning(record.task)"
            @confirm="() => handleDelete(record)"
          />
        </div>
        <div v-if="column.dataIndex === 'size'">
          {{ formatFileSize(record.size) }}
        </div>
        <Button
          type="link"
          v-if="column.dataIndex === 'originalName'"
          @click="() => handlePreview(record)"
          >{{ record.originalName }}</Button
        >
        <div v-if="column.dataIndex === 'sliceCount'">{{
          record.task ? record.task.sliceCount : '-'
        }}</div>
        <div v-if="column.dataIndex === 'task'">
          <Tooltip title="解析状态，点击查看日志">
            <Tag v-if="!record.task">未解析</Tag>
            <Tag
              v-else-if="isTaskSuccess(record.task)"
              color="green"
              @click="() => openTaskLogModal(record)"
              >已解析</Tag
            >
            <Tag
              v-else-if="isTaskFailed(record.task)"
              color="red"
              @click="() => openTaskLogModal(record)"
              >解析失败</Tag
            >
            <Tag
              v-else-if="isTaskRunning(record.task)"
              color="default"
              @click="() => openTaskLogModal(record)"
              >解析中...</Tag
            >
          </Tooltip>
        </div>
      </template>
    </Table>

    <UploadModal ref="uploadModalRef" :do-upload="uploadFile" @close="triggerQuery" />
    <FormModal
      ref="renameModalRef"
      :form-schema="renameFormSchema"
      :form-rules="renameFormRules"
      :form-state="renameFormState"
      :submit="handleRename"
      @close="triggerQuery"
    />
    <TaskLogModal ref="taskLogModalRef" />
  </div>
</template>

<script lang="js" setup>
  import Table from '@/components/Table/index.vue';
  import { computed, onMounted, onUnmounted, ref } from 'vue';
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
    parse,
    cancel,
    isTaskSuccess,
    isTaskFailed,
    isTaskRunning,
    getDocumentIcon,
  } from '@/views/agent/knowledge/document/document.data.js';
  import IconButton from '@/components/Button/IconButton.vue';
  import { useRoute } from 'vue-router';
  import UploadModal from '@/components/Modal/UploadModal.vue';
  import { Button, Tag, Tooltip } from 'ant-design-vue';
  import ConfirmButton from '@/components/Button/ConfirmButton.vue';
  import FormModal from '@/components/Modal/FormModal.vue';
  import { useRouter } from 'vue-router';
  import TaskLogModal from '@/views/agent/knowledge/document/TaskLogModal.vue';

  const router = useRouter();
  const knowledgeBaseId = ref('');
  const route = useRoute();
  knowledgeBaseId.value = route.query['knowledgeBaseId'];

  const documentStore = useDocumentStore();
  const records = computed(() => documentStore.records);
  const total = computed(() => documentStore.total);
  const uploadModalRef = ref();
  const renameModalRef = ref();
  const tableRef = ref();
  const renameFormState = computed(() => documentStore.renameFormState);
  const interval = ref();
  const taskLogModalRef = ref();

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

  const handleParse = async (record) => {
    await parse({
      documentId: record.id,
    });
    await triggerQuery();
  };

  const handleCancel = async (record) => {
    await cancel(record.task.id);
    await triggerQuery();
  };

  const handlePreview = (record) => {
    router.push({
      path: '/agent/knowledge/document/chunks',
      query: {
        id: record.id,
      },
    });
  };

  const openTaskLogModal = (record) => {
    if (!record.task) return;
    taskLogModalRef.value.open(record.task);
  };

  onMounted(async () => {
    interval.value = setInterval(async () => {
      await triggerQuery();
    }, 5000);
  });

  onUnmounted(() => {
    clearInterval(interval.value);
  });
</script>
