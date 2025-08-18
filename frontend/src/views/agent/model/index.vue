<template>
  <div>
    <CardList
      :records="records"
      :query-form-schema="queryFormSchema"
      :search="search"
      use-add-card
      @add="() => openDrawer(false)"
      @item-click="(item) => openModelDrawer(item)"
    >
      <template #bodyCell="{ item }">
        {{ item.providerName }}
      </template>
      <template #actions="{ item }">
        <SettingOutlined @click.stop="() => openDrawer(true, item)" />
        <DeleteOutlined @click.stop="() => {}" />
      </template>
    </CardList>
    <FormDrawer
      :form-schema="providerFormSchema"
      :form-state="formState"
      :form-rules="providerFormRules"
      ref="providerDrawerRef"
      :submit="submit"
    />
    <ModelDrawer ref="modelDrawerRef" />
  </div>
</template>

<script lang="js" setup>
  import CardList from '@/components/CardList/CardList.vue';
  import FormDrawer from '@/components/Drawer/FormDrawer.vue';
  import { computed, ref } from 'vue';
  import {
    providerFormRules,
    providerFormSchema,
    queryFormSchema,
    useProviderStore,
  } from '@/views/agent/model/provider.data.js';
  import { DeleteOutlined, SettingOutlined } from '@ant-design/icons-vue';
  import ModelDrawer from '@/views/agent/model/ModelDrawer.vue';

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
</script>
