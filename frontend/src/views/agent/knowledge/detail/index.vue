<template>
  <div class="kb-detail">
    <div class="kb-sidebar">
      <div class="kb-header"></div>
      <div class="kb-menu">
        <Menu mode="vertical" :items="menuItems" :selected-keys="activeKey" @click="handleMenuClick" />
      </div>
    </div>
    <div class="kb-content">
      <RouterView />
    </div>
  </div>
</template>

<script lang="js" setup>
  import { Menu } from 'ant-design-vue';
  import { RouterView } from 'vue-router';
  import { computed, h } from 'vue';
  import { ReadOutlined, SearchOutlined, SettingOutlined } from '@ant-design/icons-vue';
  import {useRoute, useRouter} from 'vue-router';

  const router = useRouter();
  const route = useRoute();
  const query = computed(()=>route.query);
  const activeKey = computed(()=>[route.name]);

  const menuItems = [
    {
      label: '文档',
      key: '知识库-文档列表',
      icon: h(ReadOutlined),
    },
    {
      label: '检索测试',
      key: '知识库-检索测试',
      icon: h(SearchOutlined),
    },
    {
      label: '配置',
      key: '知识库-配置',
      icon: h(SettingOutlined),
    },
  ];

  const handleMenuClick = async (item) => {
    await router.replace({
      name: item.key,
      query: query.value,
    });
    activeKey.value = [item.key];
  }

</script>
<style scoped>
  .kb-detail {
    height: 100%;
    width: 100%;
    display: flex;
    gap: 10px;
  }

  .kb-sidebar {
    height: 100%;
    width: 20%;
    overflow: auto;
    background-color: white;
    padding: 10px;
  }

  .kb-content {
    height: 100%;
    width: 80%;
    overflow: auto;
    background-color: transparent;
  }
</style>
