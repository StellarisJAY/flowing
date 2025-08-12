<template>
  <Modal
    :open="visible"
    :ok-text="okText"
    :cancel-text="cancelText"
    :title="isUpdate ? `编辑${title}` : `新增${title}`"
    @cancel="close"
    @ok="handleOk"
  >
    <Form
      :custom-button="true"
      :form-schema="formSchema"
      :rules="formRules"
      :state="formState"
      layout="vertical"
      ref="formRef"
      :submit-func="async (state) => await submit(state, isUpdate)"
    />
  </Modal>
</template>

<script setup lang="js">
  import { Modal } from 'ant-design-vue';
  import { ref } from 'vue';
  import Form from '@/components/Form/index.vue';

  const formRef = ref();

  const props = defineProps({
    okText: {
      type: String,
      default: '保存',
    },
    cancelText: {
      type: String,
      default: '取消',
    },
    formSchema: {
      type: Array,
      default: () => [],
    },
    formRules: {
      type: Object,
      default: () => ({}),
    },
    formState: {
      type: Object,
      default: () => ({}),
    },
    submit: {
      type: Function,
      default: () => () => {},
    },
    title: {
      type: String,
      default: '',
    },
  });

  const emit = defineEmits(['close']);

  const visible = ref(false);
  const isUpdate = ref(false);

  const open = (update) => {
    visible.value = true;
    isUpdate.value = update;
  };

  const handleOk = async () => {
    const ok = await formRef.value.submit();
    if (ok) {
      close();
    }
  };

  const close = () => {
    visible.value = false;
    emit('close');
  };

  defineExpose({
    open,
  });
</script>

<style scoped></style>
