<template>
  <div style="height: 100%; width: 100%">
    <div class="search-bar">
      <Form
        :form-schema="queryFormSchema"
        :rules="queryFormRules"
        submit-btn-text="查询"
        show-reset-btn
        @submit="search"
      />
    </div>
    <div class="content">
      <Space style="margin-bottom: 10px">
        <slot name="tool-buttons"></slot>
      </Space>
      <Table
        class="table"
        :columns="columns"
        :data-source="records"
        :pagination="false"
        :scroll="{ y: 500 }"
        :row-selection="rowSelection"
        size="small"
        bordered
      >
        <template #bodyCell="{column, record}">
          <slot name="bodyCell" :column="column" :record="record"></slot>
        </template>
      </Table>
      <Pagination v-if="pagination" class="pagination" />
    </div>
  </div>
</template>

<script lang="js" setup>
  import { Table, Pagination, Space } from 'ant-design-vue';
  import Form from '@/components/Form/index.vue';
  import { onMounted, ref } from 'vue';

  const paginationForm = ref({
    paged: false,
    page: 1,
    pageSize: 10,
  });

  const rowSelection = ref({
    checkStrictly: false,
    selectedRowKeys: [],
    onChange: (selectedRowKeys, selectedRows) => {
      this.selectedRowKeys = selectedRowKeys;
    },
    onSelect: (record, selected, selectedRows) => {
      if (selected) {
        this.selectedRowKeys.push(record.id);
      } else {
        this.selectedRowKeys = this.selectedRowKeys.filter((key) => key !== record.id);
      }
    },
    onSelectAll: (selected, selectedRows, changeRows) => {
      if (selected) {
        this.selectedRowKeys = selectedRows.map((row) => row.id);
      } else {
        this.selectedRowKeys = [];
      }
    },
  });

  defineProps({
    columns: {
      type: Array,
      default: () => [],
    },
    records: {
      type: Array,
      default: () => [],
    },
    queryFormSchema: {
      type: Array,
      default: () => {},
    },
    queryFormRules: {
      type: Array,
      default: () => [],
    },
    pagination: {
      type: Boolean,
      default: true,
    },
  });

  const emit = defineEmits(['refresh']);

  const search = (e) => {
    const form = {
      ...paginationForm.value,
      ...e,
    };
    emit('refresh', form);
  };

  onMounted(() => search());
</script>

<style scoped>
  .search-bar {
    background-color: white;
    width: 100%;
    max-height: 10%;
    padding: 10px;
    display: flex;
    justify-content: flex-start;
  }
  .content {
    background-color: white;
    width: 100%;
    height: 85%;
    padding: 20px;
    margin-top: 20px;
  }
  .table {
    height: 600px;
  }
</style>
