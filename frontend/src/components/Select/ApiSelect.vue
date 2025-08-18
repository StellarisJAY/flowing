<template>
  <Select
    :options="options"
    :loading="loading"
    v-model:value="value"
    :placeholder="placeholder"
    :multiple="multiple"
    :disabled="disabled"
  />
</template>

<script lang="js" setup>
  import { Select } from 'ant-design-vue';
  import { onMounted, ref } from 'vue';

  const props = defineProps({
    api: {
      type: Function,
      default: () => async () => {},
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
    placeholder: {
      type: String,
      default: '请选择',
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
  const loading = ref(false);

  const getOptions = async () => {
    loading.value = true;
    try {
      const data = await props.api(props.params);
      options.value = data.map((item) => ({
        label: item[props.labelField],
        value: item[props.valueField],
      }));
    } catch (error) {
      console.log(error);
      options.value = [];
    } finally {
      loading.value = false;
    }
  };

  onMounted(async () => {
    await getOptions();
  });
</script>
