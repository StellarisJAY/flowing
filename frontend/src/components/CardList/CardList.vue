<template>
  <div style="height: 90%">
    <div class="query-container">
      <Form
        :form-schema="queryFormSchema"
        v-model:state="queryForm"
        submit-btn-text="查询"
        :submit-func="search"
      />
    </div>
    <div class="card-container">
      <div class="card-cell" v-if="useAddCard">
        <Card
          class="card-item"
          :body-style="{
            height: '100%',
            width: '100%',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
          }"
          @click="(e) => emits('add')"
        >
          <div class="add-card-content">
            <PlusOutlined />
          </div>
        </Card>
      </div>
      <div class="card-cell" v-for="item in records" :key="item.id">
        <Card class="card-item" :body-style="itemBodyStyle" @click="() => emits('item-click', item)">
          <slot name="bodyCell" :item="item"/>
          <template #actions>
            <slot name="actions" :item="item" />
          </template>
        </Card>
      </div>
    </div>
  </div>
</template>

<script lang="js" setup>
  import { Card } from 'ant-design-vue';
  import Form from '@/components/Form/index.vue';
  import { PlusOutlined } from '@ant-design/icons-vue';
  import { onMounted, ref } from 'vue';

  const props = defineProps({
    useAddCard: {
      type: Boolean,
      default: false,
    },
    addCardTitle: {
      type: String,
      default: '',
    },
    records: {
      type: Array,
      default: () => [],
    },
    queryFormSchema: {
      type: Array,
      default: () => [],
    },
    search: {
      type: Function,
      default: () => () => {},
    },
  });
  const emits = defineEmits(['add', 'item-click']);

  const itemBodyStyle = {
    height: '80%',
    width: '100%',
  };
  const queryForm = ref({
    page: false,
    pageSize: 10,
    pageNum: 1,
  });

  onMounted(async () => {
    await props.search(queryForm.value);
  });
</script>

<style scoped>
  .query-container {
    background-color: white;
    width: 100%;
    max-height: 10%;
    padding: 10px;
    display: flex;
    justify-content: flex-start;
    margin-bottom: 20px;
  }
  .card-container {
    padding-top: 10px;
    background-color: transparent;
    width: 100%;
    height: 100%;
    overflow: auto;
    display: flex;
    justify-content: flex-start;
    flex-wrap: wrap; /* 允许换行 */
    gap: 20px; /* 设置卡片之间的间距 */
  }
  .card-cell {
    width: 256px;
    height: 256px;
    margin-bottom: 10px;
    flex-shrink: 0;
  }
  .card-item {
    width: calc(100% - 20px);
    height: 100%;
  }
  .card-item:hover {
    box-shadow: 0 0 8px rgba(97, 97, 97, 0.5);
    cursor: pointer;
  }
  .add-card-content {
    font-size: 64px;
    color: #999;
  }
</style>
