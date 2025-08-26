<template>
  <div style="width: 100%">
    <div style="display: flex; justify-content: flex-start; gap: 10px">
      <Select
        :value="value"
        :disabled="disabled"
        :mode="multiple ? 'multiple' : 'default'"
        :options="options"
        style="width: 500px"
        :open="false"
        @change="handleChange"
        :allowClear="allowClear"
      />
      <Button type="primary" @click="openModal" :disabled="disabled">选择</Button>
    </div>
    <Form v-if="!disabled">
      <Modal :title="title" :open="visible" @cancel="closeModal" @ok="handleOk" width="100%" destroy-on-close>
        <Table
          :columns="columns"
          :pagination="true"
          :query-form-schema="queryFormSchema"
          @refresh="refresh"
          :total="total"
          :records="records"
          v-model:selected-keys="selectedKeys"
          :select-multiple="multiple"
        />
      </Modal>
    </Form>
  </div>
</template>

<script setup lang="js">
  import Table from '@/components/Table/index.vue';
  import { Modal, Select, Button, message, Form } from 'ant-design-vue';
  import { onMounted, ref } from 'vue';

  const props = defineProps({
    title: {
      type: String,
      default: '选择',
    },
    columns: {
      type: Array,
      default: () => [],
    },
    api: {
      type: Function,
      default: () => {},
    },
    params: {
      type: Object,
      default: () => ({}),
    },
    queryFormSchema: {
      type: Array,
      default: () => [],
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
    allowClear: {
      type: Boolean,
      default: true,
    },
  });

  const value = defineModel('value', {
    type: [Array, String],
    default: () => [],
  });

  const records = ref([]);
  const options = ref([]);
  const visible = ref(false);
  const total = ref(0);
  const selectedKeys = ref([]);

  const refresh = async (query) => {
    try {
      const res = await props.api({ ...props.params, ...query });
      total.value = res.total;
      records.value = res.data;
      options.value = records.value.map((item) => ({
        label: item[props.labelField],
        value: item[props.valueField],
      }));
    } catch (err) {
      message.error(err);
    }
  };

  const openModal = () => {
    visible.value = true;
    if (value.value instanceof Array) {
      selectedKeys.value = value.value;
    } else {
      selectedKeys.value = [value.value];
    }
  };

  const closeModal = () => {
    visible.value = false;
  };

  const handleOk = () => {
    if (props.multiple) {
      value.value = selectedKeys.value;
    } else {
      value.value = selectedKeys.value[0];
    }
    closeModal();
  };

  const handleChange = (e) => {
    selectedKeys.value = e;
    value.value = e;
  };
  onMounted(refresh);
</script>
