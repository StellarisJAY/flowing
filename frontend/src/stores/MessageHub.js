import { defineStore } from 'pinia';
import { fetchEventStream } from '@/api/index.js';

const MESSAGE_HUB_API = '/api/chat';

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
    async sendMessageDebugger(text, agentConfigs) {
      const data = {
        agentConfigs,
        message: text,
        conversationId: this.conversationId ? this.conversationId : '',
        agentId: '',
        mode: 'debugger',
      }
      this.state = 'sending';
      await fetchEventStream(MESSAGE_HUB_API, data, this.onSSEOpen, this.onSSEMessage, this.onSSEError, this.onSSEClose)
      this.state = 'new'
    },
    // 非调试模式发送消息，传agentId，不传agentConfigs
    async sendMessage(text) {
      const data = {
        message: text,
        conversationId: this.conversationId ? this.conversationId : '',
        agentId: this.agentInfo ? this.agentInfo.id : '',
        mode: 'normal',
      };
      this.state = 'sending';
      await fetchEventStream(MESSAGE_HUB_API, data, this.onSSEOpen, this.onSSEMessage, this.onSSEError, this.onSSEClose)
      this.state = 'new';
    },
    onSSEMessage(event) {
      const data = JSON.parse(event.data);
      this.messages.push(data);
      if (data.conversationId) {
        this.conversationId = data.conversationId;
      }
    },
    onSSEOpen(event) {},
    onSSEError(event) {},
    onSSEClose(event) {},
    isTransmitting() {
      return this.state === 'sending' || this.state === 'receiving';
    }
  },
});
