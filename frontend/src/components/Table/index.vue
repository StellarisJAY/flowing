<template>
  <div style="height: 100%; width: 100%">
    <div class="search-bar">
      <Form
        :state="queryForm"
        :form-schema="queryFormSchema"
        :rules="queryFormRules"
        submit-btn-text="查询"
        show-reset-btn
        :submit-func="search"
        :reset-func="()=>{resetQueryForm(); search();}"
      />
    </div>
    <div class="content">
      <div class="tool-bar">
        <Space>
          <slot name="tool-buttons"></slot>
        </Space>
      </div>
      <div class="table-container">
        <Table
          :columns="columns"
          :data-source="records"
          :pagination="false"
          :row-selection="rowSelection"
          size="small"
          bordered
          :row-key="keyColumn"
        >
          <template #bodyCell="{ column, record }">
            <slot name="bodyCell" :column="column" :record="record"></slot>
          </template>
        </Table>
      </div>
      <div class="pagination">
        <Pagination
          v-if="pagination"
          :total="total"
          :page-size="paginationForm.pageSize"
          :current="paginationForm.pageNum"
          show-size-changer
          :show-total="(t) => `共${t}条`"
          @change="onPageChange"
          style="max-width:50%"
        >
          <template #buildOptionText="props">
            <span>{{ props.value }}条/页</span>
          </template>
        </Pagination>
      </div>
    </div>
  </div>
</template>

<script lang="js" setup>
  import { Table, Pagination, Space } from 'ant-design-vue';
  import Form from '@/components/Form/index.vue';
  import { computed, onMounted, ref, unref } from 'vue';

  const props = defineProps({
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
    total: {
      type: Number,
      default: 0,
    },
    keyColumn: {
      type: String,
      default: 'id',
    }
  });
  const emit = defineEmits(['refresh']);
  const selectedKeys = defineModel('selectedKeys', {
    type: Array,
    default: () => [],
  });

  const search = (e) => {
    const form = {
      ...paginationForm.value,
      ...e,
    };
    emit('refresh', form);
  };

  // 分页请求
  const paginationForm = ref({
    page: props.pagination,
    pageNum: 1,
    pageSize: 10,
  });
  const onPageChange = (pageNum, pageSize) => {
    paginationForm.value.pageNum = pageNum;
    paginationForm.value.pageSize = pageSize;
    search();
  };

  // 初始化搜索框表单
  const queryForm = ref({});
  const resetQueryForm = ()=>{
    queryForm.value = {};
    props.queryFormSchema.forEach((item)=>{
      queryForm.value[item.name] = item.defaultValue;
    });
  };
  resetQueryForm();

  const onSelectChange = (selectedRowKeys) => {
    selectedKeys.value = selectedRowKeys;
  };

  const rowSelection = computed(() => {
    return {
      selectedRowKeys: unref(selectedKeys),
      onChange: onSelectChange,
      hideDefaultSelections: true,
    };
  });

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

  .tool-bar {
    max-height: 10%;
    margin-bottom: 10px;
  }

  .table-container {
    height: 85%;
  }

  .pagination {
    display: flex;
    justify-content: end;
  }
</style>
