<template>
  <Button :type="type" :danger="danger" :size="size" @click="() => emits('click')" :icon="h(icons[icon])">
    {{title}}
  </Button>
</template>

<script setup lang="js">
import { Button} from 'ant-design-vue';
import { h } from 'vue';
import { SearchOutlined, SaveOutlined, RedoOutlined, PlusOutlined } from '@ant-design/icons-vue';

const icons = {
  SearchOutlined,
  SaveOutlined,
  RedoOutlined,
  PlusOutlined,
};

defineProps({
  icon: {
    type: String,
    default: '',
  },
  title: {
    type: String,
    default: '',
  },
  type: {
    type: String,
    default: 'default',
  },
  danger: {
    type: Boolean,
    default: false,
  },
  size: {
    type: String,
    default: 'medium',
  },
});

const emits = defineEmits(['click']);
</script>
