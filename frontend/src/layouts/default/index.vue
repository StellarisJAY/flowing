<template>
  <Layout class="flowing-layout">
    <Layout.Header class="flowing-header">
      <div class="flowing-header-logo"></div>
      <Menu
        class="flowing-header-menu"
        mode="horizontal"
        :items="menuItems"
        :selectable="false"
        @click="onMenuItemClick"
      />
      <div class="flowing-header-avatar">
        <IconButton icon="MessageOutlined" shape="circle" type="primary" @click="openAiHelperDrawer" v-if="showChatButton"/>
        <Dropdown>
          <Avatar />
          <template #overlay>
            <Menu>
              <MenuItem @click="logout">退出登录</MenuItem>
            </Menu>
          </template>
        </Dropdown>
      </div>
    </Layout.Header>

    <Tabs
      v-if="!hideTabs"
      :activeKey="activeTab"
      type="editable-card"
      @edit="editTabs"
      :hideAdd="true"
      @tabClick="onTabClick"
      style="width: 100%"
    >
      <Tabs.TabPane
        v-for="panel in tabPanes"
        :key="panel.key"
        :tab="panel.title"
        :closable="panel.closable"
      />
    </Tabs>

    <Layout.Content class="flowing-content">
      <div class="flowing-tab-content">
        <router-view v-slot="{ Component }">
          <component :is="Component" :key="$route.fullPath" />
        </router-view>
      </div>
    </Layout.Content >
    <ChatDrawer ref="aiHelperDrawer" title="AI助手 (即将上线)" v-if="showChatButton" />
  </Layout>
</template>

<script lang="js" setup>
import { Layout, Menu, Tabs, FloatButton, Dropdown, MenuItem, Avatar, message } from 'ant-design-vue';
  import { computed, ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { usePermissionStore } from '@/stores/permission.js';
  import { useUserStore } from '@/stores/user.js';
  import { MessageOutlined } from '@ant-design/icons-vue';
  import ChatDrawer from '@/components/Chat/ChatDrawer.vue';
  import {useRoute} from 'vue-router';
  import { logoutAPI } from '@/api/sys/permission.js';
import IconButton from '@/components/Button/IconButton.vue';

  const route = useRoute();
  const router = useRouter();
  const permissionStore = usePermissionStore();
  const userStore = useUserStore();
  const menuItems = computed(() => permissionStore.navMenus);
  const tabPanes = computed(() => userStore.tabPanes);
  const activeTab = computed(() => userStore.activeTab);
  const hideTabs = computed(()=>route.meta.hideTab);
  const aiHelperDrawer = ref();

  const editTabs = (key) => {
    userStore.deleteTabPane(key);
    router.replace(activeTab.value);
  };

  const onMenuItemClick = (item) => {
    router.replace(item.key);
  };

  const onTabClick = (key) => {
    router.replace(key);
  };

  const openAiHelperDrawer = () => {
    aiHelperDrawer.value.open();
  };

  const showChatButton = computed(()=>{
    return route.name !== '聊天页面';
  });

  const logout = async () => {
    try {
      await logoutAPI();
      await userStore.logout();
      await router.replace('/sys/login');
    }catch {
      message.error('退出登录失败');
    }
  };
</script>

<style scoped>
  .flowing-layout {
    height: 100vh;
  }
  .flowing-header {
    width: 100%;
    background-color: #1677ff;
    margin: 0;
    padding: 0;
    display: flex;
    align-items: start;
    justify-content: flex-start;
  }
  .flowing-content {
    height: 90%;
  }

  .flowing-header-logo {
    width: 25%;
    height: 100%;
  }

  .flowing-header-avatar {
    width: 25%;
    height: 100%;
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    align-items: center;
    padding: 20px;
    gap: 10px;
  }

  .flowing-header-menu {
    background-color: transparent;
    width: 50%;
    color: white;
    justify-content: center;
  }

  .flowing-tab-content {
    height: 100%;
    width: 100%;
    padding-left: 10px;
    padding-right: 10px;
    margin: auto;
    left: 0;
    right: 0;
  }
</style>

<style>
  .ant-tabs-nav {
    background-color: white;
    padding-left: 10px;
    padding-right: 10px;
  }
  .ant-tabs-content {
    height: 100%;
  }
</style>
