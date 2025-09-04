<template>
  <div class="variable-list">
    <Table :columns="dynamicColumns" :data-source="variables" :pagination="false">
      <template #bodyCell="{ column, record }">
        <Input v-if="!record.fixed && column.dataIndex === 'name'" v-model:value="record.name"
               @change="onVarNameChange(record)"/>
        <Input v-if="!record.fixed && column.dataIndex === 'description'" v-model:value="record.description" />
        <Select
          v-if="!record.fixed && column.dataIndex === 'type'"
          v-model:value="record.type"
          :options="variableTypeOptions"
        />

        <Space v-if="column.dataIndex === 'actions'">
          <Button type="link" v-if="!record.fixed">删除</Button>
        </Space>
      </template>
    </Table>
    <Button type="primary" @click="addVariable" v-if="allowAdd">添加变量</Button>
  </div>
</template>

<script lang="js" setup>
  import { Table, Button, Input, Select, Space } from 'ant-design-vue';
  import { computed } from 'vue';
  import { v4 as uuid } from 'uuid';
  import { genVariableId } from '@/stores/workflow.js';

  const variableTypeOptions = [
    {
      label: 'String',
      value: 'string',
    },
    {
      label: 'File',
      value: 'file',
    },
  ];

  const props = defineProps({
    nodeId: {
      type: String,
      required: true,
    },
    allowAdd: {
      type: Boolean,
      default: false,
    },
    isInput: {
      type: Boolean,
      default: false,
    }
  });

  const dynamicColumns = computed(()=>{
    const columns = [
      {
        title: '变量名称',
        dataIndex: 'name',
        key: 'name',
      },
      {
        title: '变量类型',
        dataIndex: 'type',
        key: 'type',
      },
      {
        title: '变量描述',
        dataIndex: 'description',
        key: 'description',
      },
    ];
    // 只有input变量允许引用
    if (props.isInput) {
      columns.push({
        title: '引用',
        dataIndex: 'ref',
        key: 'ref',
      });
    }
    columns.push({
        title: '操作',
        dataIndex: 'actions',
        key: 'actions',
    });
    return columns;
  });

  const variables = defineModel('variables', {
    default: () => [],
    type: Array,
  });

  // 添加新变量
  const addVariable = () => {
    // 临时变量名
    const name = uuid();
    const id = genVariableId(props.nodeId, name);
    variables.value.push({
      id,
      name,
      type: 'string',
      description: '新增变量',
    });
  };

  // 变量名变化，修改变量的id
  const onVarNameChange = (record) => {
    record.id = genVariableId(props.nodeId, record.name);
    console.log(record);
  };

</script>

<style scoped></style>
