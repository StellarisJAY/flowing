<template>
  <Drawer title="模型列表" :open="visible" size="large" @close="close" destroy-on-close>
    <Table
      :columns="columns"
      :records="records"
      :query-form-schema="queryFormSchema"
      @refresh="search"
      :pagination="false"
    >
      <template #tool-buttons>
        <IconButton
          type="primary"
          title="新增模型"
          icon="PlusOutlined"
          @click="() => openFormModal(false)"
        />
      </template>
      <template #bodyCell="{ column, record }">
        <ConfirmButton
          v-if="column.dataIndex === 'actions'"
          title="删除"
          @confirm="() => handleDelete(record)"
        />
        <Switch
          v-if="column.dataIndex === 'enable'"
          :checked="record.enable"
          @change="(e) => changeEnable(e, record)"
        />
      </template>
    </Table>
    <FormModal
      :form-schema="modelFormSchema"
      :form-state="formState"
      :form-rules="modelFormRules"
      ref="formModalRef"
      :submit="submit"
      @close="search"
    />
  </Drawer>
</template>

<script lang="js" setup>
  import { Drawer, Switch } from 'ant-design-vue';
  import FormModal from '@/components/Modal/FormModal.vue';
  import Table from '@/components/Table/index.vue';
  import { computed, ref } from 'vue';
  import {
    columns,
    modelFormRules,
    modelFormSchema,
    queryFormSchema,
    useModelStore,
  } from '@/views/agent/model/model.data.js';
  import IconButton from '@/components/Button/IconButton.vue';
  import ConfirmButton from '@/components/Button/ConfirmButton.vue';

  const modelStore = useModelStore();
  const records = computed(() => modelStore.records);
  const formState = computed(() => modelStore.formState);
  const visible = ref(false);
  const providerId = ref('');
  const formModalRef = ref();

  const search = async (query) => {
    const data = {
      providerId: providerId.value,
      ...query,
    };
    await modelStore.list(data);
  };

  const open = (record) => {
    visible.value = true;
    providerId.value = record.id;
  };

  const close = () => {
    visible.value = false;
  };

  const openFormModal = (isUpdate, record) => {
    if (isUpdate) {
      modelStore.setFormState(record);
    } else {
      modelStore.resetFormState();
    }
    formModalRef.value.open(isUpdate);
  };

  const submit = async (record, isUpdate) => {
    const form = {
      ...record,
      providerId: providerId.value,
    };
    return await modelStore.save(form, isUpdate);
  };

  const changeEnable = async (e, record) => {
    record.enable = e;
    await modelStore.save(record, true);
  };

  const handleDelete = async (record) => {
    await modelStore.delete(record.id);
    await search();
  };

  defineExpose({
    open,
  });
</script>
