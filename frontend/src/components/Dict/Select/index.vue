<template>
  <Select
    v-model:value="value"
    :placeholder="placeholder"
    :multiple="multiple"
    :allowClear="allowClear"
    :options="options"
  />
</template>

<script lang="js" setup>
  import { Select } from 'ant-design-vue';
  import { onMounted, ref } from 'vue';
  import { getDictItemsByCode } from '@/api/system/dict.api.js';

  const props =defineProps({
    dictCode: {
      type: String,
      required: true,
    },
    multiple: {
      type: Boolean,
      default: false,
    },
    placeholder: {
      type: String,
      default: '请选择...',
    },
    allowClear: {
      type: Boolean,
      default: false,
    },
  });

  const value = defineModel('value', {
    type: Array | String,
    default: () => [],
  });

  const options = ref([]);

  onMounted(async ()=>{
    const {data} = await getDictItemsByCode({dictCode: props.dictCode});
    options.value = data.map(item=>({
      label: item.itemName,
      value: item.itemValue,
    }));
  });
</script>
