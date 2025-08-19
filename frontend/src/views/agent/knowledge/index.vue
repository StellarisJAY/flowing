<template>
  <CardList
    ref="cardListRef"
    :query-form-schema="searchFormSchema"
    :use-add-card="true"
    :records="records"
    :search="search"
    add-card-title="新增知识库"
    @add="() => openModal(false)"
    @item-click="(item) => toDocument(item.id)"
  >
    <template #bodyCell="{ item }">
      {{ item.name }}
    </template>
    <template #actions="{ item }">
      <SettingOutlined @click.stop="() => openModal(true, item)" />
      <DeleteOutlined />
    </template>
  </CardList>
  <FormModal
    ref="formModalRef"
    :form-schema="knowledgeFormSchema"
    :form-state="formState"
    :form-rules="knowledgeFormRules"
    title="知识库"
    :submit="submit"
    @close="triggerQuery"
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
  import { useRouter } from 'vue-router';

  const cardListRef = ref();
  const knowledgeStore = useKnowledgeStore();
  const formState = computed(() => knowledgeStore.knowledgeForm);
  const formModalRef = ref();
  const records = computed(() => knowledgeStore.records);
  const router = useRouter();

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
  };

  const toDocument = (id) => {
    router.push({
      path: `/agent/knowledge/documents?knowledgeBaseId=${id}`,
      query: {
        knowledgeBaseId: id,
      },
    });
  };

  const triggerQuery = async () => {
    await cardListRef.value.triggerQuery();
  };
</script>
