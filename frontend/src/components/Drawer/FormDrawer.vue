<template>
  <Drawer
    :open="visible"
    @close="close"
    size="large"
    destroy-on-close
    :title="isUpdate === true ? `编辑${title}` : `新增${title}`"
  >
    <Form
      :form-schema="formSchema"
      :rules="formRules"
      :state="formState"
      layout="vertical"
      custom-button
      ref="formRef"
      :submit-func="async (state) => await submit(state, isUpdate)"
    >
      <template #buttons>
        <Space>
          <Button type="primary" @click="handleSave" v-if="isUpdate === true">保存</Button>
          <Button @click="handleSaveAndClose" v-if="isUpdate === true">保存并关闭</Button>
          <Button @click="handleSaveAndClose" v-if="isUpdate === false" type="primary">保存</Button>
        </Space>
      </template>
      <!-- 透传所有form item slots -->
      <template v-for="(_, slotName) in $slots" #[slotName]="slotData">
        <slot :name="slotName" v-bind="slotData" />
      </template>
    </Form>
  </Drawer>
</template>

<script lang="js" setup>
  import { Drawer, Space, Button } from 'ant-design-vue';
  import { ref } from 'vue';
  import Form from '@/components/Form/index.vue';

  const isUpdate = ref(false);
  const formRef = ref();
  const visible = ref(false);
  const emit = defineEmits(['submit', 'reset', 'open', 'close']);

  defineProps({
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

  const open = (update) => {
    visible.value = true;
    isUpdate.value = update;
    emit('open');
  };

  const close = () => {
    visible.value = false;
    emit('close');
  };

  const handleSave = async () => {
    await formRef.value.submit();
  };

  const handleSaveAndClose = async () => {
    const ok = await formRef.value.submit();
    if (ok) {
      close();
    }
  };

  defineExpose({
    open,
    close,
  });
</script>

<style scoped></style>
