<template>
  <Drawer :title="title" :open="visible" @close="close" size="large">
    <div class="chat-container">
      <div class="chat-messages">
        <Message :type="message.type" :message="message.message" v-for="message in mockMessages" :key="message.id" />
      </div>
      <div class="chat-input">
        <Input/>
        <Button type="primary">发送</Button>
      </div>
    </div>
  </Drawer>
</template>

<script lang="js" setup>
import {Drawer, Input, Button} from 'ant-design-vue';
import { ref } from 'vue';
import Message from '@/components/Chat/Message.vue';

defineProps({
  title: {
    type: String,
    default: 'AI助手'
  }
});

const visible = ref(false);
const mockMessages = ref([
  {
    id: 1,
    type: 'user',
    message: 'Ant Design Vue 设计规范和技术上支持灵活的样式定制，以满足业务和品牌上多样化的视觉需求，包括但不限于全局样式（主色、圆角、边框）和指定组件的视觉定制。在 4.0 版本的 Ant Design Vue 中，我们提供了一套全新的定制主题方案。不同于 3.x 版本的 less 和 CSS 变量，有了 CSS-in-JS 的加持后，动态主题的能力也得到了加强，包括但不限于：支持动态切换主题；支持同时存在多个主题；支持针对某个/某些组件修改主题变量；Ant Design Vue 设计规范和技术上支持灵活的样式定制，以满足业务和品牌上多样化的视觉需求，包括但不限于全局样式（主色、圆角、边框）和指定组件的视觉定制。在 4.0 版本的 Ant Design Vue 中，我们提供了一套全新的定制主题方案。不同于 3.x 版本的 less 和 CSS 变量，有了 CSS-in-JS 的加持后，动态主题的能力也得到了加强，包括但不限于：支持动态切换主题；支持同时存在多个主题；支持针对某个/某些组件修改主题变量；Ant Design Vue 设计规范和技术上支持灵活的样式定制，以满足业务和品牌上多样化的视觉需求，包括但不限于全局样式（主色、圆角、边框）和指定组件的视觉定制。在 4.0 版本的 Ant Design Vue 中，我们提供了一套全新的定制主题方案。不同于 3.x 版本的 less 和 CSS 变量，有了 CSS-in-JS 的加持后，动态主题的能力也得到了加强，包括但不限于：支持动态切换主题；支持同时存在多个主题；支持针对某个/某些组件修改主题变量；'
  },
  {
    id: 2,
    type: 'assistant',
    message: 'Ant Design Vue 设计规范和技术上支持灵活的样式定制，以满足业务和品牌上多样化的视觉需求，包括但不限于全局样式（主色、圆角、边框）和指定组件的视觉定制。在 4.0 版本的 Ant Design Vue 中，我们提供了一套全新的定制主题方案。不同于 3.x 版本的 less 和 CSS 变量，有了 CSS-in-JS 的加持后，动态主题的能力也得到了加强，包括但不限于：支持动态切换主题；支持同时存在多个主题；支持针对某个/某些组件修改主题变量；Ant Design Vue 设计规范和技术上支持灵活的样式定制，以满足业务和品牌上多样化的视觉需求，包括但不限于全局样式（主色、圆角、边框）和指定组件的视觉定制。在 4.0 版本的 Ant Design Vue 中，我们提供了一套全新的定制主题方案。不同于 3.x 版本的 less 和 CSS 变量，有了 CSS-in-JS 的加持后，动态主题的能力也得到了加强，包括但不限于：支持动态切换主题；支持同时存在多个主题；支持针对某个/某些组件修改主题变量；Ant Design Vue 设计规范和技术上支持灵活的样式定制，以满足业务和品牌上多样化的视觉需求，包括但不限于全局样式（主色、圆角、边框）和指定组件的视觉定制。在 4.0 版本的 Ant Design Vue 中，我们提供了一套全新的定制主题方案。不同于 3.x 版本的 less 和 CSS 变量，有了 CSS-in-JS 的加持后，动态主题的能力也得到了加强，包括但不限于：支持动态切换主题；支持同时存在多个主题；支持针对某个/某些组件修改主题变量；'
  }
]);

const open = () => {
  visible.value = true;
};

const close = () => {
  visible.value = false;
};

defineExpose({
  open
})
</script>

<style scoped>
.chat-container {
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}
.chat-messages {
  height: 420%;
  overflow: auto;
  padding: 10px;
  background-color: #f5f5f5;
  border-radius: 10px;
  margin-bottom: 10px;
  display: flex;
  flex-direction: column;
}
.chat-input {
  height: 20%;
  display: flex;
  align-items: center;
  background-color: #fff;
  border-radius: 10px;
  gap: 10px;
}
</style>
