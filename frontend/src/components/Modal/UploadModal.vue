<template>
  <Modal
    :open="visible"
    :title="title"
    :ok-text="okText"
    :cancel-text="cancelText"
    @cancel="close"
    @ok="onOk"
  >
    <Spin :spinning="loading" tip="上传中...">
      <UploadDragger :multiple="multiple" :file-list="fileList" :before-upload="beforeUpload">
        <p class="ant-upload-drag-icon">
          <InboxOutlined />
        </p>
        <p class="ant-upload-text">点击或拖拽文件到此处上传</p>
        <p class="ant-upload-hint"> 严禁上传违规文件 </p>
      </UploadDragger>
    </Spin>
  </Modal>
</template>

<script lang="js" setup>
  import { Modal, UploadDragger, Spin, message } from 'ant-design-vue';
  import { InboxOutlined } from '@ant-design/icons-vue';
  import { ref } from 'vue';

  const props = defineProps({
    title: {
      type: String,
      default: '上传文档',
    },
    okText: {
      type: String,
      default: '确认',
    },
    cancelText: {
      type: String,
      default: '取消',
    },
    multiple: {
      type: Boolean,
      default: false,
    },
    doUpload: {
      type: Function,
      default: () => {},
    },
  });

  const emits = defineEmits(['close']);

  const visible = ref(false);
  const loading = ref(false);
  const fileList = ref([]);

  const beforeUpload = (file) => {
    fileList.value = [file];
    return false;
  };

  const onOk = async () => {
    loading.value = true;
    try {
      await props.doUpload(fileList.value);
      message.success('上传成功');
      close();
    } catch {
      message.error('上传失败');
    } finally {
      loading.value = false;
    }
  };

  const close = () => {
    fileList.value = [];
    visible.value = false;
    emits('close');
  };

  const open = () => {
    visible.value = true;
  };

  defineExpose({
    close,
    open,
  });
</script>
