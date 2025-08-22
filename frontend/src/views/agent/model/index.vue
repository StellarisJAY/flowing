<template>
  <div style="width: 100%; height: 100%;">
    <Table
      ref="tableRef"
      :columns="columns"
      :records="records"
      :query-form-schema="queryFormSchema"
      @refresh="search"
      :pagination="false"
    >
      <template #tool-buttons>
        <IconButton icon="PlusOutlined" @click="() => openDrawer(false)" type="primary" title="新增模型提供方"/>
      </template>
      <template #bodyCell="{column, record}">
        <img :src="getProviderIcon(record.providerType)" width="32" height="32" alt="icon" v-if="column.dataIndex === 'icon'" />
        <Space v-if="column.dataIndex === 'actions'">
          <Button type="link" @click="() => openDrawer(true, record)">编辑</Button>
          <Button type="link" @click="() => openModelDrawer(record)">模型</Button>
          <ConfirmButton title="删除" />
        </Space>
      </template>
    </Table>
    <FormDrawer
      :form-schema="providerFormSchema"
      :form-state="formState"
      :form-rules="providerFormRules"
      ref="providerDrawerRef"
      :submit="submit"
      @close="triggerQuery"
    />
    <ModelDrawer ref="modelDrawerRef" />
  </div>
</template>

<script lang="js" setup>
  import Table from '@/components/Table/index.vue';
  import FormDrawer from '@/components/Drawer/FormDrawer.vue';
  import { computed, ref } from 'vue';
  import {
    columns, getProviderIcon,
    providerFormRules,
    providerFormSchema,
    queryFormSchema,
    useProviderStore,
  } from '@/views/agent/model/provider.data.js';
  import ModelDrawer from '@/views/agent/model/ModelDrawer.vue';
  import IconButton from '@/components/Button/IconButton.vue';
  import { Space, Button } from 'ant-design-vue';
  import ConfirmButton from '@/components/Button/ConfirmButton.vue';

  const tableRef = ref();
  const providerStore = useProviderStore();
  const records = computed(() => providerStore.records);
  const formState = computed(() => providerStore.formState);
  const providerDrawerRef = ref();
  const modelDrawerRef = ref();

  const search = async (query) => {
    await providerStore.list(query);
  };

  const openDrawer = (isUpdate, record) => {
    if (isUpdate) {
      providerStore.setFormState(record);
    } else {
      providerStore.initFormState();
    }
    providerDrawerRef.value.open(isUpdate);
  };

  const submit = async (record, isUpdate) => {
    return await providerStore.save(record, isUpdate);
  };

  const openModelDrawer = (record) => {
    modelDrawerRef.value.open(record);
  };

  const triggerQuery = async () => {
    await tableRef.value.triggerQuery();
  };
</script>
