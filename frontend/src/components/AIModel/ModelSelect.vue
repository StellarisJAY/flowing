<template>
  <ApiSelect v-model:value="value" :api="list" label-field="modelName" value-field="id" :multiple="false"/>
</template>
<script setup lang="ts">
import { listModel } from '@/api/ai/model.api.js';
import ApiSelect from '@/components/Select/ApiSelect.vue';

const props = defineProps ({
  modelType: {
    type: String,
    default: '',
  },
});

const value = defineModel('value', {
  type: String,
  default: '',
});

const list = async (params) => {
  params.modelType = props.modelType;
  const { data } = await listModel(params);
  return data;
};
</script>
