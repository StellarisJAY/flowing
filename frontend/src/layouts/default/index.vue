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
      <div class="flowing-header-avatar"></div>
    </Layout.Header>

    <Layout.Content class="flowing-content">
      <Tabs
        :activeKey="activeTab"
        type="editable-card"
        @edit="editTabs"
        :hideAdd="true"
        @tabClick="onTabClick"
      >
        <Tabs.TabPane
          v-for="panel in tabPanes"
          :key="panel.key"
          :tab="panel.title"
          :closable="panel.closable"
        >
          <router-view v-slot="{ Component }">
            <transition>
              <keep-alive>
                <component :is="Component" />
              </keep-alive>
            </transition>
          </router-view>
        </Tabs.TabPane>
      </Tabs>
    </Layout.Content>
  </Layout>
</template>

<script lang="js" setup>
  import { Layout, Menu, Tabs } from 'ant-design-vue';
  import { computed, watch } from 'vue';
  import { useRouter } from 'vue-router';
  import { usePermissionStore } from '@/stores/permission.js';
  import { useUserStore } from '@/stores/user.js';

  const router = useRouter();
  const permissionStore = usePermissionStore();
  const userStore = useUserStore();
  const menuItems = computed(() => permissionStore.navMenus);
  const tabPanes = computed(() => userStore.tabPanes);
  const activeTab = computed(() => userStore.activeTab);

  watch(activeTab, (newVal) => {
    router.replace(newVal);
  });

  const editTabs = (path) => {
    userStore.changeTabPanesOnRouting({ path: path }, true);
  };

  const onMenuItemClick = (item) => {
    router.push(item.key);
  };

  const onTabClick = (key) => {
    router.replace(key);
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
    background-color: white;
    padding: 10px;
  }

  .flowing-header-logo {
    width: 25%;
    height: 100%;
  }

  .flowing-header-avatar {
    width: 25%;
    height: 100%;
  }

  .flowing-header-menu {
    background-color: transparent;
    width: 50%;
    color: white;
    justify-content: center;
  }
</style>
