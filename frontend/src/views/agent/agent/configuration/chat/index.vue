<template>
  <div class="flowing-chat-configure">
    <div class="header">
      <div class="title"> 智能体配置 </div>
      <div class="save-button">
        <Button type="primary" @click="handleSubmit">保存</Button>
      </div>
    </div>
    <div class="content">
      <div class="chat-options">
        <Form
          class="form"
          layout="vertical"
          :model="configForm"
          :rules="formRules"
          no-style
          ref="formRef"
        >
          <div class="form-section">
            <Form.Item label="提示词" name="prompt" class="form-item-prompt">
              <textarea v-model="configForm.prompt" class="textarea-prompt" />
            </Form.Item>
            <Form.Item label="大模型" name="modelId">
              <ModelSelect model-type="llm" v-model:value="configForm.modelId" />
            </Form.Item>
            <Form.Item label="知识库" name="knowledgeBaseId">
              <KnowledgeBaseSelect v-model:value="configForm.knowledgeBaseId" :allowClear="true" />
            </Form.Item>
          </div>

          <div class="form-section" v-if="configForm.knowledgeBaseId">
            <Form.ItemRest>
              <Form.Item label="检索类型" name="searchType">
                <Select
                  :options="retrieveTypeOptions"
                  v-model:value="configForm.kbSearchOption.searchType"
                />
              </Form.Item>
              <Form.Item label="topK" name="topK">
                <InputNumber v-model:value="configForm.kbSearchOption.topK" :min="1" />
              </Form.Item>
              <Form.Item
                label="相似度阈值"
                name="threshold"
                v-if="configForm.kbSearchOption.searchType !== 'fulltext'"
              >
                <Slider
                  v-model:value="configForm.kbSearchOption.threshold"
                  :min="0"
                  :max="1"
                  :step="0.1"
                />
              </Form.Item>
              <Form.Item
                label="混合检索类型"
                name="hybridType"
                v-if="configForm.kbSearchOption.searchType === 'hybrid'"
              >
                <Select
                  :options="hybridRerankOptions"
                  v-model:value="configForm.kbSearchOption.hybridType"
                />
              </Form.Item>
              <Form.Item
                label="权重"
                name="weight"
                v-if="
                  configForm.kbSearchOption.searchType === 'hybrid' &&
                  configForm.kbSearchOption.hybridType === 'weight'
                "
              >
                <Slider
                  v-model:value="configForm.kbSearchOption.weight"
                  :min="0"
                  :max="1"
                  :step="0.1"
                />
              </Form.Item>
              <Form.Item
                label="重排序模型"
                name="rerankModel"
                v-if="
                  configForm.kbSearchOption.searchType === 'hybrid' &&
                  configForm.kbSearchOption.hybridType === 'rerank'
                "
              >
                <ModelSelect model-type="rerank" />
              </Form.Item>
            </Form.ItemRest>
          </div>

          <!--        <div class="form-section">-->
          <!--          <Form.Item label="变量列表" name="variables">-->

          <!--          </Form.Item>-->
          <!--        </div>-->
        </Form>
      </div>
      <div class="chat-preview">
        <div class="preview-header">
          <div class="preview-header">
            <div class="preview-title">调试预览</div>
            <div class="preview-button">
              <IconButton icon="PlusOutlined" title="新建对话" type="link" />
            </div>
          </div>
        </div>
        <div class="preview-messages">
          <Empty v-if="!messages || messages.length === 0" description="" class="message-empty" />
          <Message
            v-else
            v-for="item in messages"
            :key="item.id"
            :type="item.type"
            :message="item.content"
          />
        </div>
        <div class="preview-input">
          <div class="input-section">
            <input
              placeholder="与智能体聊天..."
              class="textarea-chat"
              v-model="chatInput"
              :disabled="sendLoading"
              @keyup.enter="sendMessage"
            />
            <IconButton
              icon="SendOutlined"
              shape="circle"
              class="send-button"
              type="primary"
              @click="sendMessage"
              :loading="sendLoading"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="js" setup>
  import ModelSelect from '@/components/AIModel/ModelSelect.vue';
  import { Form, InputNumber, message, Button, Select, Slider, Empty } from 'ant-design-vue';
  import { computed, onMounted, ref } from 'vue';
  import { useRoute } from 'vue-router';
  import Message from '@/components/Chat/Message.vue';
  import { getDetail, saveConfig } from '@/api/ai/agent.api.js';
  import KnowledgeBaseSelect from '@/components/Knowledge/KnowledgeBaseSelect.vue';
  import IconButton from '@/components/Button/IconButton.vue';
  import { useMessageHub } from '@/stores/MessageHub.js';

  const route = useRoute();
  const formRef = ref(null);
  const messages = computed(() => messageHub.messages);
  const messageHub = useMessageHub();
  const chatInput = ref('');
  const sendLoading = computed(() => messageHub.isTransmitting());

  const getAgentDetail = async () => {
    try {
      const res = await getDetail(route.query.id);
      configForm.value = JSON.parse(res.data.config);
      console.log(configForm.value);
    } catch (err) {
      message.error(err);
    }
  };

  onMounted(() => {
    messageHub.newConversation();
    getAgentDetail();
  });

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

  const configForm = ref({});

  const formRules = {
    prompt: [{ required: true, message: '请输入提示词' }],
    modelId: [{ required: true, message: '请选择模型' }],
  };

  const handleSubmit = async () => {
    await formRef.value.validate();
    try {
      await saveConfig({
        id: route.query.id,
        config: JSON.stringify(configForm.value),
      });
      message.success('保存成功');
    } catch {
      message.error('保存失败');
    }
  };

  const sendMessage = async () => {
    await messageHub.sendMessageDebugger(chatInput.value, JSON.stringify(configForm.value));
    chatInput.value = '';
  };
</script>

<style scoped>
  .flowing-chat-configure {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
  }

  .header {
    width: 100%;
    height: 10%;
    background-color: white;
    display: flex;
    flex-direction: row;
    padding: 5px;
    .title {
      width: 30%;
      height: 100%;
      font-size: 20px;
      font-weight: bold;
      line-height: 50px;
      display: flex;
      justify-content: flex-start;
      align-items: center;
    }
    .save-button {
      width: 70%;
      height: 100%;
      display: flex;
      justify-content: flex-end;
      align-items: center;
    }
  }

  .content {
    height: 90%;
    width: 100%;
    display: flex;
    flex-direction: row;
  }

  .chat-options {
    width: 50%;
    height: 100%;
    background-color: white;
    padding-left: 5px;
    padding-right: 5px;
  }

  .form {
    width: 100%;
    height: 100%;
    display: flow;
    flex-direction: column;
    overflow: auto;
  }

  .form::-webkit-scrollbar {
    width: 8px;
    height: 8px;
  }
  .form::-webkit-scrollbar-track {
    background-color: transparent;
  }
  .form::-webkit-scrollbar-thumb {
    background-color: #bfbfbf;
    border-radius: 4px;
  }
  .form::-webkit-scrollbar-thumb:hover {
    background-color: #a1a1a1ff;
  }

  .textarea-prompt {
    width: 100%;
    height: 200px;
    resize: none;
    border: #0096ff 2px solid;
    border-radius: 5px;
    padding: 5px;
  }

  .form-section {
    background-color: #f0f0f0;
    width: 100%;
    overflow: auto;
    padding: 5px;
    margin-bottom: 5px;
    border-radius: 8px;
  }

  .form-section::-webkit-scrollbar {
    width: 0;
    height: 0;
  }

  .chat-preview {
    width: 50%;
    height: 100%;
    background-color: #f0f0f0;
    padding: 5px;
    border-radius: 8px;
    display: flex;
    flex-direction: column;
  }

  .preview-header {
    width: 100%;
    height: 10%;
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;
    .preview-title {
      width: 50%;
      height: 100%;
      font-size: 16px;
      font-weight: bold;
      display: flex;
      justify-content: flex-start;
      align-items: center;
    }
    .preview-button {
      width: 50%;
      height: 100%;
      display: flex;
      justify-content: flex-end;
      align-items: center;
    }
  }

  .preview-messages {
    width: 100%;
    height: 70%;
    overflow: auto;
    display: flow;
    flex-direction: column;

    .message-empty {
      width: 100%;
      height: 100%;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
    }
  }

  .preview-input {
    width: 100%;
    height: 20%;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 20px;

    .input-section {
      background-color: white;
      width: 100%;
      height: 100%;
      box-shadow: #999999 0 0 10px;
      border-radius: 10px;
      display: flex;
      flex-direction: row;
      justify-content: flex-start;
      align-items: center;
      gap: 5px;
      padding: 5px;
    }

    .textarea-chat {
      width: 90%;
      height: 100%;
      resize: none;
      border: none;
      outline: none;
      font-size: 16px;
    }
  }
</style>
