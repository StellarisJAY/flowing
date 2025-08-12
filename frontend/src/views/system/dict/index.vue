<template>
  <Table :columns="columns"
         :records="records"
         :query-form-schema="queryFormSchema"
         :query-form-rules="[]"
         :pagination="true"
         @refresh="refresh">
    <template #tool-buttons>
      <Button type="primary" @click="()=>openDictDrawer(false)">新增字典</Button>
    </template>
    <template #bodyCell="{column, record}">
      <Space v-if="column.dataIndex === 'actions'">
        <Button type="link" size="small" @click="()=>openDictDrawer(true, record)">编辑</Button>
        <Button type="link" size="small" @click="()=>openDictItemDrawer(record)">字典配置</Button>
        <Button type="link" size="small" danger>删除</Button>
      </Space>
    </template>
  </Table>
  <FormDrawer ref="dictFormDrawerRef"
              :form-schema="dictFormSchema"
              :form-state="dictForm"
              :form-rules="dictFormRules"
              :submit="saveDict"
              @close="refresh" />
  <DictItemDrawer ref="dictItemDrawerRef" />
</template>

<script lang="js" setup>
import { Button, Space } from 'ant-design-vue';
import Table from '@/components/Table/index.vue';
import { useDictStore } from './dict.store.js';
import { computed, ref } from 'vue';
import FormDrawer from '@/components/Drawer/FormDrawer.vue';
import DictItemDrawer from '@/views/system/dict/DictItemDrawer.vue';

const dictStore = useDictStore();

const columns = computed(() => dictStore.columns);
const records = computed(() => dictStore.dictList);
const queryFormSchema = computed(()=>dictStore.queryFormSchema);

const dictFormDrawerRef = ref();
const dictFormSchema = computed(()=>dictStore.dictFormSchema);
const dictForm = computed(()=>dictStore.dictForm);
const dictFormRules = computed(()=>dictStore.dictFormRules);

const dictItemDrawerRef = ref();

const openDictDrawer = (isUpdate, record) => {
  if (isUpdate) {
    dictStore.dictForm = record;
  }else {
    dictStore.initDictForm();
  }
  dictFormDrawerRef.value.open(isUpdate);
};

const refresh = async (e) => {
  await dictStore.fetchDictList(e);
}

const saveDict = async (data, isUpdate) => {
  return await dictStore.saveDict(data, isUpdate);
}

const openDictItemDrawer = (record) => {
  dictItemDrawerRef.value.open(record);
};
</script>

<style scoped></style>
