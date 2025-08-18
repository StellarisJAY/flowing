<template>
  <CardList
    :query-form-schema="searchFormSchema"
    :use-add-card="true"
    :records="records"
    :search="search"
    add-card-title="新增知识库"
    @add="() => openModal(false)"
  >
    <template #bodyCell="{ item }">
      {{ item.name }}
    </template>
    <template #actions="{ item }">
      <SettingOutlined @click="()=>openModal(true, item)"/>
      <DeleteOutlined/>
    </template>
  </CardList>
  <FormModal
    ref="formModalRef"
    :form-schema="knowledgeFormSchema"
    :form-state="formState"
    :form-rules="knowledgeFormRules"
    title="知识库"
    :submit="submit"
    @close="search"
  />
</template>
<script setup lang="js">
  import CardList from '@/components/CardList/CardList.vue';
  import FormModal from '@/components/Modal/FormModal.vue';
  import { computed, ref } from 'vue';
  import {
    knowledgeFormRules,
    knowledgeFormSchema,
    searchFormSchema,
    useKnowledgeStore,
  } from './knowledge.data.js';
  import { SettingOutlined, DeleteOutlined } from '@ant-design/icons-vue';

  const knowledgeStore = useKnowledgeStore();
  const formState = computed(() => knowledgeStore.knowledgeForm);
  const formModalRef = ref();
  const records = computed(()=>knowledgeStore.records);

  const openModal = (isUpdate, record) => {
    if (isUpdate) {
      knowledgeStore.setKnowledgeForm(record);
    } else {
      knowledgeStore.initKnowledgeForm();
    }
    formModalRef.value.open(isUpdate);
  };

  const search = async (query) => {
    await knowledgeStore.list(query);
  };

  const submit = async (data, isUpdate) => {
    return await knowledgeStore.saveKnowledgeBase(data, isUpdate);
  }
</script>
