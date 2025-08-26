import { defineStore } from 'pinia';
import { fetchEventStream } from '@/api/index.js';

const MESSAGE_HUB_API = '/chat';

export const useMessageHub = defineStore('flowing_message_hub', {
  state: () => ({
    messages: [],
    conversationId: '',
    agentInfo: {},
    eventSource: null,
    state: 'new',
  }),
  actions: {
    newConversation() {
      this.state = 'new';
      this.conversationId = '';
      this.messages = [];
    },
    switchAgent(agent) {
      this.newConversation();
      this.agentInfo = { ...agent };
    },
    // 调试模式发送消息，传agent配置，不传agentId
    async sendMessageDebugger(text, agentType, agentConfigs) {
      const data = {
        agentConfig: agentConfigs,
        content: text,
        conversationId: this.conversationId ? this.conversationId : '0',
        agentId: '0',
        mode: 'debugger',
        agentType,
      }
      this.state = 'sending';
      await fetchEventStream(MESSAGE_HUB_API+'/send', data, this.onSSEOpen, this.onSSEMessage, this.onSSEError, this.onSSEClose)
      this.state = 'new'
    },
    // 非调试模式发送消息，传agentId，不传agentConfigs
    async sendMessage(text) {
      const data = {
        message: text,
        conversationId: this.conversationId ? this.conversationId : '0',
        agentId: this.agentInfo ? this.agentInfo.id : '0',
        mode: 'normal',
      };
      this.state = 'sending';
      await fetchEventStream(MESSAGE_HUB_API+'/send', data, this.onSSEOpen, this.onSSEMessage, this.onSSEError, this.onSSEClose)
      this.state = 'new';
    },

    onSSEMessage(msg) {
      console.log(msg);
      const data = JSON.parse(msg);
      // 通过id找到未完成的上一条消息
      const lastMsg = this.messages.find((item) => item.id === data.id);
      if (lastMsg) {
        // 未完成的消息，追加内容
        lastMsg.content += data.content;
      } else {
        // 新消息，直接添加
        this.messages.push(data);
      }
      // 更新conversationId
      if (data.conversationId) {
        this.conversationId = data.conversationId;
      }
    },
    onSSEOpen() {
      console.log('sse open');
    },
    onSSEError(error) {
      console.log('error', error);
    },
    onSSEClose() {
      console.log('sse close');
      this.state = 'new';
    },
    isTransmitting() {
      return this.state === 'sending' || this.state === 'receiving';
    }
  },
});
