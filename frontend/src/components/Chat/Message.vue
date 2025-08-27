<template>
  <div class="message-container" :class="`message-container-${type}`">
    <div class="avatar">
      <img :src="avatarRobot" alt="a" width="100%" height="100%" v-if="type === 'assistant'"/>
      <img :src="avatarHuman" alt="a" width="100%" height="100%" v-else/>
    </div>
    <div class="message-bubble" :class="`message-bubble-${type}`">
      <details v-if="message.thinkingContent">
        <summary>思考中...</summary>
        <div v-html="marked.parse(message.thinkingContent)" />
      </details>
      <div v-html="marked.parse(message.content)" v-if="type === 'assistant'" />
      <p v-else>{{message.content}}</p>
      <details v-if="message.knowledgeReferences">
        <summary>引用</summary>
        <details v-for="item in message.knowledgeReferences" :key="item.sliceId">
          <summary>{{item.documentName}}</summary>
          <p>{{item.content}}</p>
        </details>
      </details>
    </div>
  </div>
</template>

<script lang="js" setup>
import avatarRobot from '@/assets/svg/avatar_robot.svg';
import avatarHuman from '@/assets/svg/avatar_human.svg';
import { marked } from 'marked';

marked.use({
  // 开启异步渲染
  async: false,
  pedantic: false,
  gfm: true,
  mangle: false,
  headerIds: false
});



defineProps({
  type: {
    type: String,
    default: 'user'
  },
  message: {
    type: String,
    default: ''
  }
});
</script>

<style scoped>
/**
 * 消息组件
 */
.message-container {
  width: 100%;
}
/**
 * 消息组件-头像
 */
.avatar {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  background-color: #f0f0f0;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
}
/**
 * 消息组件-消息气泡
 */
.message-bubble {
  padding: 10px;
  border-radius: 10px;
  background-color: #f0f0f0;
  max-width: 90%;
  color: #000;
  font-size: 14px;
  overflow-y: auto;
  word-break: break-word;
  word-wrap: break-word;
}

/**
 * 消息组件-消息气泡-滚动条
 */
.message-bubble::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}
.message-bubble::-webkit-scrollbar-track {
  background-color: transparent;
}
.message-bubble::-webkit-scrollbar-thumb {
  background-color: #bfbfbf;
  border-radius: 4px;
}
.message-bubble::-webkit-scrollbar-thumb:hover {
  background-color: #A1A1A1FF;
}

/**
 * 消息组件-用户消息
 */
.message-container-user {
  display: flex;
  flex-direction: column;
  margin-bottom: 10px;
  gap: 10px;
  align-items: flex-end;
}
/**
 * 消息组件-助手消息
 */
.message-container-assistant {
  display: flex;
  flex-direction: column;
  margin-bottom: 10px;
  gap: 10px;
  align-items: flex-start;
}
.message-bubble-user {
  background-color: #007bff;
  color: #fff;
}
.message-bubble-assistant {
  background-color: white;
  color: black;
}
</style>
