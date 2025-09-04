<template>
  <div class="node-config">
    <Space>
      <h4>选择知识库:</h4>
      <KnowledgeBaseSelect v-model:value="config.knowledgeBaseId" />
    </Space>
    <PromptTextArea :node-id="nodeId" v-model:value="config.query" title="搜索内容" />
    <Form layout="vertical">
      <Form.Item label="检索类型" name="searchType">
        <Select
          :options="retrieveTypeOptions"
          v-model:value="config.searchType"
        />
      </Form.Item>
      <Form.Item label="topK" name="topK">
        <InputNumber v-model:value="config.topK" :min="1" />
      </Form.Item>
      <Form.Item
        label="相似度阈值"
        name="threshold"
        v-if="config.searchType !== 'fulltext'"
      >
        <Slider v-model:value="config.threshold" :min="0" :max="1" :step="0.1" />
      </Form.Item>
      <Form.Item
        label="混合检索类型"
        name="hybridType"
        v-if="config.searchType === 'hybrid'"
      >
        <Select
          :options="hybridRerankOptions"
          v-model:value="config.hybridType"
        />
      </Form.Item>
      <Form.Item
        label="权重"
        name="weight"
        v-if="
          config.searchType === 'hybrid' &&
          config.hybridType === 'weight'
        "
      >
        <Slider v-model:value="config.weight" :min="0" :max="1" :step="0.1" />
      </Form.Item>
      <Form.Item
        label="重排序模型"
        name="rerankModel"
        v-if="
          config.searchType === 'hybrid' &&
          config.hybridType === 'rerank'
        "
      >
        <ModelSelect model-type="rerank" />
      </Form.Item>
    </Form>
  </div>
</template>

<script lang="js" setup>
  import PromptTextArea from '@/components/workflow/config/prompt/PromptTextArea.vue';
  import { InputNumber, Select, Slider, Form, Space } from 'ant-design-vue';
  import ModelSelect from '@/components/AIModel/ModelSelect.vue';
  import KnowledgeBaseSelect from '@/components/Knowledge/KnowledgeBaseSelect.vue';


  const retrieveTypeOptions = [
    {
      label: '全文检索',
      value: 'fulltext',
    },
    {
      label: '向量检索',
      value: 'vector',
    },
    {
      label: '混合检索',
      value: 'hybrid',
    },
  ];

  const hybridRerankOptions = [
    {
      label: '权重',
      value: 'weight',
    },
    {
      label: '重排序模型',
      value: 'rerank',
    },
  ];

  defineProps({
    nodeId: {
      type: String,
      required: true,
    },
  });

  const config = defineModel('config', {
    type: Object,
    required: true,
  });
</script>

<style scoped>
.node-config {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 10px;
}
</style>
