<template>
  <Drawer :open="visible" destroy-on-close size="large" @close="close" title="字典项配置">
    <Table
      :columns="columns"
      :records="records"
      :query-form-schema="queryFormSchema"
      :pagination="true"
      @refresh="refresh"
    >
      <template #tool-buttons>
        <Button type="primary" @click="() => openFormModal(false)">新增字典项</Button>
      </template>
      <template #bodyCell="{ column, record }">
        <Space v-if="column.dataIndex === 'actions'">
          <Button type="link" @click="() => openFormModal(true, record)">编辑</Button>
          <ConfirmButton text="删除" @confirm="async ()=> {
            await deleteDictItem(record.id);
            await refresh();
          }" />
        </Space>
      </template>
    </Table>

    <FormModal
      ref="formModalRef"
      ok-text="保存"
      title="字典项"
      :form-schema="dictItemFormSchema"
      :form-rules="dictItemFormRules"
      :form-state="dictItemForm"
      :submit="submit"
      @close="refresh"
    />
  </Drawer>
</template>

<script lang="js" setup>
  import { Drawer, Button, Space } from 'ant-design-vue';
  import {
    columns, deleteDictItem, dictItemFormRules,
    dictItemFormSchema,
    queryFormSchema,
    saveDictItem,
    useDictItemStore,
  } from '@/views/system/dict/dictItem.data.js';
  import Table from '@/components/Table/index.vue';
  import { computed, ref } from 'vue';
  import FormModal from '@/components/Modal/FormModal.vue';
  import ConfirmButton from '@/components/Button/ConfirmButton.vue';

  const visible = ref(false);
  const dict = ref();
  const dictItemStore = useDictItemStore();

  const formModalRef = ref();
  const dictItemForm = computed(() => dictItemStore.dictItemForm);
  const records = computed(()=>dictItemStore.dictItemList);

  const open = (record) => {
    visible.value = true;
    dict.value = record;
  };

  const refresh = async (query) => {
    console.log(query);
    await dictItemStore.getDictItemList({
      ...query,
      dictId: dict.value.id,
    });
  };

  const close = () => {
    visible.value = false;
  };

  const openFormModal = (isUpdate, record) => {
    if (isUpdate) {
      dictItemStore.dictItemForm = record;
    } else {
      dictItemStore.dictItemForm = {
        itemKey: '',
        itemValue: '',
        description: '',
      };
    }
    formModalRef.value.open(isUpdate);
  };

  const submit = async (data, isUpdate) => {
    return await saveDictItem(
      {
        ...data,
        dictId: dict.value.id,
      },
      isUpdate
    );
  };

  defineExpose({
    open,
  });
</script>

<style scoped></style>
