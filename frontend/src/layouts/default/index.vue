<template>
  <Layout class="flowing-layout">
    <Layout.Header class="flowing-header">
      <div class="flowing-header-logo"></div>
      <Menu class="flowing-header-menu" mode="horizontal" :items="menuItems"/>
      <div class="flowing-header-avatar"></div>
    </Layout.Header>

    <Layout.Content class="flowing-content">
      <Tabs v-model:activeKey = "activeTab" type="editable-card" @edit="editTabs" :hideAdd="true">
        <Tabs.TabPane v-for="(panel) in panels" :key="panel.key" :tab="panel.title" :closable="panel.closable">
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
import {Layout, Menu, Tabs} from 'ant-design-vue';
import { ref, watch } from 'vue'
import { useRouter } from 'vue-router';

const router = useRouter();

const menuItems = ref([
  {
    key: "home",
    label: "主页",
    path: "/home",
  },
  {
    key: "system",
    label: "系统管理",
    path: "/system",
    children: [
      {
        key: "user",
        label: "用户管理",
        path: "/system/user",
      },
    ]
  },
  {
    key: "ai",
    label: "AI管理",
    path: "/ai",
    children: [
      {
        key: "agent",
        label: "AI Agent管理",
        path: "/ai/agent",
      },
      {
        key: "model",
        label: "模型管理",
        path: "/ai/model",
      },
      {
        key: "knowledge",
        label: "知识库管理",
        path: "/ai/knowledge",
      }
    ]
  },
  {
    key: "datasource",
    label: "数据源管理",
    path: "/datasource",
    children: [
      {
        key: "mysql",
        label: "数据库管理",
        path: "/datasource/database",
      },
      {
        key: "monitor",
        label: "系统监控",
        path: "/datasource/monitor",
      }
    ]
  }
]);

const panels = ref([
  {
    key: '/system/user',
    title: '用户管理',
    content: '',
    closable: true,
  },
  {
    key: '/system/menu',
    title: '菜单管理',
    content: '',
    closable: true,
  }
]);

const activeTab = ref('/system/user');

const editTabs = (e)=>{
};

watch(activeTab, (e)=>{
  router.replace(e);
})
</script>

<style scoped>
.flowing-layout {
  height: 100vh;
}
.flowing-header {
  width: 100%;
  background-color: #0096ff;
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
