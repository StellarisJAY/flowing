<template>
  <div class="model-node-config">
    <PromptTextArea :node-id="nodeId" v-model:value="config.message" title="回复消息" />
  </div>
</template>

<script lang="js" setup>
import PromptTextArea from '@/components/workflow/config/prompt/PromptTextArea.vue';

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
</style>
