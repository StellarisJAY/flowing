<template>
  <TreeSelect
    :tree-data="options"
    v-model:value="value"
    :multiple="multiple"
    :disabled="disabled"
  />
</template>

<script lang="js" setup>
import { TreeSelect } from 'ant-design-vue';
import { onMounted, ref } from 'vue';

const props = defineProps({
  api: {
    type: Function,
    default: ()=> async () => ([]),
  },
  params: {
    type: Object,
    default: () => ({}),
  },
  labelField: {
    type: String,
    default: 'label',
  },
  valueField: {
    type: String,
    default: 'value',
  },
  multiple: {
    type: Boolean,
    default: false,
  },
  disabled: {
    type: Boolean,
    default: false,
  },
});

const value = defineModel('value', {
  type: [Array, String],
  default: () => [],
});

const options = ref([]);

const getOptions = async () => {
  try {
    const data = await props.api(props.params);
    const buildOptions = (item) => {
      return {
        label: item[props.labelField],
        value: item[props.valueField],
        children: item.children?.map(buildOptions),
      };
    }
    options.value = data.map(buildOptions);
  } catch (error) {
    console.log(error);
    options.value = [];
  }
};

onMounted(async () => {
  await getOptions();
});
</script>
