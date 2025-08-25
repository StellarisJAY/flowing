<template>
  <div class="flowing-chat-configure">
    <div class="chat-options">
      <Form class="form" layout="vertical" :model="configForm" :rules="formRules" no-style ref="formRef">
        <div class="form-section">
          <Form.Item label="提示词" name="prompt">
            <textarea v-model="configForm.prompt" />
          </Form.Item>
          <Form.Item label="大模型" name="modelId">
            <ModelSelect model-type="llm" v-model:value="configForm.modelId" />
          </Form.Item>
        </div>

        <div class="form-section">
          <Form.Item label="知识库" name="knowledgeBaseId">
            <Select />
          </Form.Item>
          <Form.Item label="知识库-检索设置" name="kbSearchOption" v-if="configForm.knowledgeBaseId">
            <Form.Item label="检索类型" name="searchType">
              <Select :options="retrieveTypeOptions" v-model:value="configForm.kbSearchOption.searchType" />
            </Form.Item>
            <Form.Item label="topK" name="topK">
              <InputNumber v-model:value="configForm.kbSearchOption.topK" :min="1" />
            </Form.Item>
            <Form.Item label="相似度阈值" name="threshold">
              <Slider v-model:value="configForm.kbSearchOption.threshold" :min="0" :max="1" :step="0.1" />
            </Form.Item>
            <Form.Item label="混合检索类型" name="hybridType">
              <Select :options="hybridRerankOptions" v-model:value="configForm.kbSearchOption.hybridType" />
            </Form.Item>
            <Form.Item label="权重" name="weight">
              <Slider v-model:value="configForm.kbSearchOption.weight" :min="0" :max="1" :step="0.1" />
            </Form.Item>
            <Form.Item label="重排序模型" name="rerankModel">
              <ModelSelect model-type="rerank" />
            </Form.Item>
          </Form.Item>
        </div>

        <div class="form-section">
          <Form.Item label="变量列表" name="variables">

          </Form.Item>
        </div>

        <IconButton icon="SaveOutlined" title="保存" type="primary" @click="handleSubmit"/>

      </Form>

    </div>
    <div class="chat-preview">

    </div>
  </div>
</template>

<script lang="js" setup>

import ModelSelect from '@/components/AIModel/ModelSelect.vue';
import { Form, InputNumber, message, Select, Slider } from 'ant-design-vue';
import { ref } from 'vue';
import IconButton from '@/components/Button/IconButton.vue';
import {useRoute} from 'vue-router';
import {saveConfig} from '@/api/ai/agent.api.js';

const route = useRoute();
const formRef = ref(null);

const retrieveTypeOptions = [
  {
    label: '全文检索',
    value: 'fulltext'
  },
  {
    label: '向量检索',
    value: 'vector'
  },
  {
    label: '混合检索',
    value: 'hybrid'
  }
];

const hybridRerankOptions = [
  {
    label: '权重',
    value: 'weight'
  },
  {
    label: '重排序模型',
    value: 'rerank'
  }
];

const configForm = ref({
  prompt: '你是一个智能助手',
  knowledgeBaseId: '',
  modelId: '',
  kbSearchOption: {
    searchType: 'fulltext',
    topK: 10,
    threshold: 0.5,
    weight: 0.5,
    hybridType: 'weight',
    rerankModel: ''
  },
  variables: [],
});

const formRules = {
  prompt: [
    { required: true, message: '请输入提示词' }
  ],
  modelId: [
    { required: true, message: '请选择模型' }
  ],
};

const handleSubmit = async () => {
  await formRef.value.validate();
  try {
    await saveConfig({
      id: route.query.id,
      config: JSON.stringify(configForm.value),
    });
    message.success('保存成功');
  }catch {
    message.error('保存失败');
  }
};

</script>

<style scoped>
.flowing-chat-configure {
  width: 100%;
  height: 100%;
  display: flex;
  gap: 10px;
  flex-direction: row;
}

.chat-options {
  width: 50%;
  height: 100%;
  background-color: white;
  padding: 10px;
  display: flow;
  flex-direction: column;
  gap: 5px;
  overflow: auto;
}

.chat-options::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}
.chat-options::-webkit-scrollbar-track {
  background-color: transparent;
}
.chat-options::-webkit-scrollbar-thumb {
  background-color: #bfbfbf;
  border-radius: 4px;
}
.chat-options::-webkit-scrollbar-thumb:hover {
  background-color: #A1A1A1FF;
}

.chat-preview {
  width: 50%;
  height: 100%;
  background-color: white;
  padding: 10px;
}

.form {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

textarea {
  width: 100%;
  height: 100%;
  resize: none;
  border: #0096ff 2px solid;
  border-radius: 5px;
  padding: 5px
}

.form-section {
  background-color: #f0f0f0;
  width: 100%;
  overflow: auto;
  padding: 5px;
}

.form-section::-webkit-scrollbar {
  width: 0;
  height: 0;
}

</style>
