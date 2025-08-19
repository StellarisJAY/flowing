<template>
  <Table
    ref="tableRef"
    :columns="columns"
    :records="records"
    :total="total"
    :query-form-schema="queryFormSchema"
    :query-form-rules="[]"
    :pagination="true"
    @refresh="refresh"
  >
    <template #tool-buttons>
      <IconButton
        type="primary"
        icon="PlusOutlined"
        @click="() => openDrawer(false)"
        title="新增数据源"
      />
    </template>
    <template #bodyCell="{ column, record }">
      <Space v-if="column.dataIndex === 'actions'">
        <Button type="link" size="small" @click="() => openDrawer(true, record)">编辑</Button>
        <ConfirmButton text="删除" @confirm="async () => {}" />
      </Space>
    </template>
  </Table>
  <FormDrawer
    ref="drawerRef"
    :form-schema="datasourceFormSchema"
    :form-rules="datasourceFormRules"
    :form-state="formState"
    title="数据源"
    :submit="submit"
    @close="triggerQuery"
  >
    <template #form-port="{ formState }">
      <Space>
        <InputNumber v-model:value="formState.port" :min="0" :max="65535" />
        <Button
          type="primary"
          @click="() => testConn(formState)"
          :loading="connLoading"
          :disabled="connLoading"
          >测试连接</Button
        >
        {{ connState }}
      </Space>
    </template>
  </FormDrawer>
</template>

<script lang="js" setup>
  import { Button, Space } from 'ant-design-vue';
  import Table from '@/components/Table/index.vue';
  import { computed, ref } from 'vue';
  import ConfirmButton from '@/components/Button/ConfirmButton.vue';
  import IconButton from '@/components/Button/IconButton.vue';
  import {
    columns,
    datasourceFormRules,
    datasourceFormSchema,
    queryFormSchema,
    testConnection,
    useDatasourceStore,
  } from '@/views/monitor/datasource/datasource.data.js';
  import FormDrawer from '@/components/Drawer/FormDrawer.vue';
  import { InputNumber } from 'ant-design-vue';
  const tableRef = ref();
  const drawerRef = ref();
  const datasourceStore = useDatasourceStore();
  const records = computed(() => datasourceStore.records);
  const total = computed(() => datasourceStore.total);
  const formState = computed(() => datasourceStore.formState);
  const connState = ref('');
  const connLoading = ref(false);

  const refresh = async (query) => {
    await datasourceStore.getDatasourceList(query);
  };

  const openDrawer = (isUpdate, record) => {
    connState.value = '';
    if (isUpdate === true) {
      datasourceStore.setFormState(record);
    } else {
      datasourceStore.initFormState();
    }
    drawerRef.value.open(isUpdate);
  };

  const submit = async (data, isUpdate) => {
    return await datasourceStore.save(data, isUpdate);
  };

  const testConn = async (data) => {
    connState.value = '';
    connLoading.value = true;
    connState.value = await testConnection(data);
    connLoading.value = false;
  };

  const triggerQuery = async () => {
    await tableRef.value.triggerQuery();
  };
</script>

<style scoped></style>
