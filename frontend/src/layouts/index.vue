<template>
  <Layout class="flowing-layout">
    <Layout.Header class="flowing-header">

    </Layout.Header>
    <Layout.Content class="flowing-content">
      <RouterView />
    </Layout.Content>
  </Layout>
</template>

<script lang="js" setup>
import {Layout} from 'ant-design-vue'
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
}
.flowing-content {
  height: 90%;
  background-color: white;
  padding: 10px;
}
</style>
