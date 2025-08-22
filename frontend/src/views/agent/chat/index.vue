<template>
  <div class="flowing-chat">
    <div class="chat-sidebar">
      <div class="sidebar-back">
        <IconButton
          type="primary"
          icon="ArrowLeftOutlined"
          title="返回"
          @click="handleBack"
          style="width: 100%"
        />
      </div>
      <div class="sidebar-history"></div>
      <div class="sidebar-history"></div>
    </div>
    <div class="chat-content">
      <div class="content-header"></div>
      <div class="content-messages"></div>
      <div :class="['content-input-section', !mode ? '' : 'bordered']">
        <div class="toolbar" v-if="mode">
          <div class="mode-options">
            <span v-if="mode === 'chat'">选择模型: </span>
            <ModelSelect v-if="mode === 'chat'" model-type="llm" style="min-width:30%;max-width: 50%"/>
            <span v-if="mode === 'agent'">选择智能体: </span>
          </div>
          <div class="close-mode">
            <IconButton
              type="link"
              icon="CloseOutlined"
              @click="setMode('')"
            />
          </div>
        </div>
        <div class="toolbar" v-else>
          <IconButton title="智能体" shape="round" type="default" icon="RobotOutlined" @click="setMode('agent')" class="mode-button" />
          <IconButton title="聊天" shape="round" type="default" icon="MessageOutlined" @click="setMode('chat')" class="mode-button" />
        </div>
        <div :class="['section-input-and-bottom', mode ? '' : 'bordered']">
          <textarea class="input-textarea" />
          <IconButton
            type="primary"
            icon="SendOutlined"
            shape="circle"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="js" setup>
  import IconButton from '@/components/Button/IconButton.vue';
  import { useRouter } from 'vue-router';
  import ModelSelect from '@/components/AIModel/ModelSelect.vue';
  import { ref } from 'vue';

  const router = useRouter();
  const handleBack = () => router.back();
  const mode = ref('1');

  const setMode = (value) => {
    mode.value = value;
  };
</script>

<style scoped>
  .flowing-chat {
    height: 100%;
    width: 100%;
    display: flex;
  }
  .chat-sidebar {
    height: 100%;
    width: 20%;
    background-color: white;
    padding-top: 10px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    .sidebar-back {
      height: 5%;
      width: 100%;
    }
    .sidebar-history {
      height: 45%;
      width: 100%;
      background-color: #cccccc;
      overflow: auto;
    }
  }
  .chat-content {
    height: 100%;
    width: 80%;
    background-color: white;
    display: flex;
    flex-direction: column;
    gap: 10px;
    padding: 10px;
    justify-content: center;
    align-items: center;

    .content-header {
      height: 5%;
      width: 90%;
    }
    .content-messages {
      height: 75%;
      width: 90%;
      background-color: #f5f5f5;
      overflow: auto;
      display: flex;
      gap: 10px;
    }

    .bordered {
      border: #0096ff 2px solid;
      border-radius: 10px;
    }

    .content-input-section {
      height: 20%;
      width: 90%;
      display: flex;
      flex-direction: column;
      gap: 1px;
      border-radius: 10px;
    }

    .toolbar {
      height: 30%;
      width: 100%;
      display: flex;
      flex-direction: row;
      gap: 5px;
      justify-content: flex-start;
      padding: 5px;
    }

    .mode-button {
      height: 80%;
    }

    .mode-options {
      height: 100%;
      width: 90%;
    }

    .close-mode {
      height: 100%;
      width: 10%;
      display: flex;
      flex-direction: row;
      justify-content: flex-end;
    }

    .section-input-and-bottom {
      height: 70%;
      width: 100%;
      display: flex;
      flex-direction: row;
      border-radius: 10px;
      gap: 5px;
      padding: 5px;
      justify-content: center;
      align-items: center;
    }

    .input-textarea {
      height: 100%;
      width: 100%;
      resize: none;
      border: none;
      outline: none;
      font-size: 16px;
      font-family: 'SFMono-Regular', Consolas, 'Liberation Mono', Menlo, Courier, monospace;
      background-color: #efefef;
      border-radius: 5px;
      padding: 5px;
    }
  }
</style>
