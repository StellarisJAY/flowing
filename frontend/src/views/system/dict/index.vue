<template>
  <Table
    ref="tableRef"
    :columns="columns"
    :records="records"
    :query-form-schema="queryFormSchema"
    :query-form-rules="[]"
    :pagination="true"
    @refresh="refresh"
  >
    <template #tool-buttons>
      <IconButton type="primary" @click="() => openDictDrawer(false)" icon="PlusOutlined" title="新增字典" />
    </template>
    <template #bodyCell="{ column, record }">
      <Space v-if="column.dataIndex === 'actions'">
        <Button type="link" size="small" @click="() => openDictDrawer(true, record)">编辑</Button>
        <Button type="link" size="small" @click="() => openDictItemDrawer(record)">字典配置</Button>
        <ConfirmButton
          text="删除"
          @confirm="
            async () => {
              await deleteDict(record.id);
              await triggerQuery();
            }
          "
        />
      </Space>
    </template>
  </Table>
  <FormDrawer
    ref="dictFormDrawerRef"
    :form-schema="dictFormSchema"
    :form-state="dictForm"
    :form-rules="dictFormRules"
    :submit="saveDict"
    @close="triggerQuery"
    title="字典"
  />
  <DictItemDrawer ref="dictItemDrawerRef" />
</template>

<script lang="js" setup>
  import { Button, Space } from 'ant-design-vue';
  import Table from '@/components/Table/index.vue';
  import {
    columns,
    deleteDict,
    dictFormRules,
    dictFormSchema,
    queryFormSchema,
    useDictStore,
  } from './dict.data.js';
  import { computed, ref } from 'vue';
  import FormDrawer from '@/components/Drawer/FormDrawer.vue';
  import DictItemDrawer from '@/views/system/dict/DictItemDrawer.vue';
  import ConfirmButton from '@/components/Button/ConfirmButton.vue';
  import IconButton from '@/components/Button/IconButton.vue';

  const tableRef = ref();

  const dictStore = useDictStore();

  const records = computed(() => dictStore.dictList);

  const dictFormDrawerRef = ref();
  const dictForm = computed(() => dictStore.dictForm);

  const dictItemDrawerRef = ref();

  const openDictDrawer = (isUpdate, record) => {
    if (isUpdate) {
      dictStore.dictForm = record;
    } else {
      dictStore.initDictForm();
    }
    dictFormDrawerRef.value.open(isUpdate);
  };

  const refresh = async (e) => {
    await dictStore.fetchDictList(e);
  };

  const saveDict = async (data, isUpdate) => {
    return await dictStore.saveDict(data, isUpdate);
  };

  const openDictItemDrawer = (record) => {
    dictItemDrawerRef.value.open(record);
  };

  const triggerQuery = async () => {
    await tableRef.value.triggerQuery();
  };
</script>

<style scoped></style>
