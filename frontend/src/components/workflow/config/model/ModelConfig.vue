<template>
  <div class="model-node-config">
    <Space>
      <h4>选择大模型:</h4>
      <ModelSelect model-type="llm" v-model:value="config.model" style="width: 300px"/>
    </Space>
    <PromptTextArea :node-id="nodeId" v-model:value="config.prompt" />
  </div>
</template>

<script lang="js" setup>
import ModelSelect from '@/components/AIModel/ModelSelect.vue';
import PromptTextArea from '@/components/workflow/config/prompt/PromptTextArea.vue';
import { Space } from 'ant-design-vue';

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
.model-node-config {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 10px;
}
</style>
