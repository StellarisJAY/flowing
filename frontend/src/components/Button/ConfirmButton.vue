<template>
  <Popconfirm
    :title="title"
    @cancel="() => emit('cancel')"
    @confirm="() => emit('confirm')"
    :ok-text="okText"
    :cancel-text="cancelText"
  >
    <Button :type="type" :disabled="disabled" :danger="danger">{{ text }}</Button>
  </Popconfirm>
</template>

<script lang="js" setup>
  import { Button, Popconfirm } from 'ant-design-vue';

  defineProps({
    type: {
      type: String,
      default: 'link',
    },
    text: {
      type: String,
      default: '删除',
    },
    title: {
      type: String,
      default: '确定删除吗？',
    },
    okText: {
      type: String,
      default: '确定',
    },
    cancelText: {
      type: String,
      default: '取消',
    },
    disabled: {
      type: Boolean,
      default: false,
    },
    danger: {
      type: Boolean,
      default: false,
    },
  });

  const emit = defineEmits(['confirm', 'cancel']);
</script>
