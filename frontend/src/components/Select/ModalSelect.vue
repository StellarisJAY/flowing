<template>
  <div>
    <Space direction="horizontal">
      <Select
        :value="value"
        :disabled="disabled"
        :mode="multiple ? 'multiple' : 'default'"
        :options="options"
        style="width: 500px"
        :open="false"
        @change="handleChange"
      />
      <Button type="primary" @click="openModal">选择</Button>
    </Space>
    <FormItem>
      <Modal :title="title" :open="visible" :width="1000" @cancel="closeModal" @ok="handleOk">
        <Table
          :columns="columns"
          :pagination="true"
          :query-form-schema="queryFormSchema"
          @refresh="refresh"
          :total="total"
          :records="records"
          v-model:selected-keys="selectedKeys"
        />
      </Modal>
    </FormItem>
  </div>
</template>

<script setup lang="js">
  import Table from '@/components/Table/index.vue';
  import { Modal, Select, Button, Space, message, FormItem } from 'ant-design-vue';
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
      const res = await props.api(query);
      total.value = res.total;
      records.value = res.data;
      options.value = records.value.map((item) => ({
        label: item[props.labelField],
        value: item[props.valueField],
      }));
    } catch (err) {
      console.log(err);
      message.error(err);
    }
  };

  const openModal = () => {
    visible.value = true;
    if (value.value instanceof String) {
      selectedKeys.value = [value.value];
    } else {
      selectedKeys.value = value.value;
    }
  };

  const closeModal = () => {
    visible.value = false;
  };

  const handleOk = () => {
    value.value = selectedKeys.value;
    closeModal();
  };

  const handleChange = (e) => {
    selectedKeys.value = e;
    value.value = e;
  };
  onMounted(refresh);
</script>
