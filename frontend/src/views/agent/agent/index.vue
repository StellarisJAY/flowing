<template>
  <div style="width: 100%; height: 100%">
    <CardList
      ref="cardListRef"
      :query-form-schema="queryFormSchema"
      :records="records"
      :use-add-card="true"
      :search="search"
      @add="()=>openModal(false)"
      @item-click="(item) => toConfiguration(item)"
    >
      <template #coverCell="{ item }">
        <img
          alt="example"
          src="https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"
        />
      </template>
      <template #bodyCell="{ item }">
        {{ item.name }}
      </template>
      <template #actions="{ item }">
        <SettingOutlined @click.stop="() => openModal(true, item)" v-if="item.createBy === userStore.getUserId()" />
        <Popconfirm
          title="确认删除？"
          @confirm="() => {}"
          @click.stop=""
          v-if="item.createBy === userStore.getUserId()"
        >
          <DeleteOutlined />
        </Popconfirm>
      </template>
    </CardList>
    <FormModal
      ref="formModalRef"
      :form-schema="agentFormSchema"
      :form-rules="agentFormRules"
      :form-state="formState"
      title="智能体"
      :submit="submit"
      @close="triggerQuery"
    />
  </div>
</template>
<script setup lang="js">
  import CardList from '@/components/CardList/CardList.vue';
  import { agentFormRules, agentFormSchema, queryFormSchema, useAgentStore } from '@/views/agent/agent/agent.data.js';
  import { computed, ref } from 'vue';
  import FormModal from '@/components/Modal/FormModal.vue';
  import { Popconfirm } from 'ant-design-vue';
  import { DeleteOutlined, SettingOutlined } from '@ant-design/icons-vue';
  import { useUserStore } from '@/stores/user.js';
  import { useRouter } from 'vue-router';

  const router = useRouter();
  const cardListRef = ref();
  const formModalRef = ref();
  const agentStore = useAgentStore();
  const records = computed(() => agentStore.records);
  const formState = computed(()=>agentStore.formState);
  const userStore = useUserStore();

  const search = async (query) => {
    await agentStore.list(query);
  };

  const openModal = (isUpdate, record) => {
    if (isUpdate) {
      agentStore.setFormState(record);
    } else {
      agentStore.initFormState();
    }
    formModalRef.value.open(isUpdate);
  };

  const toConfiguration = (record) => {
    router.push({
      path: `/agent/configuration${record.type === 'simple' ? '/chat' : '/workflow'}`,
      query: {
        id: record.id,
      },
    });
  };

  const submit = async (data, isUpdate) => {
    return await agentStore.save(data, isUpdate);
  };

  const triggerQuery = () => {
    cardListRef.value.triggerQuery();
  };
</script>
