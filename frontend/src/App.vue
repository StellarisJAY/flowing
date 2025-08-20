<script setup>
import { RouterView } from 'vue-router'
import { Spin } from 'ant-design-vue';
import { useGlobalStore } from '@/stores/global.js';
import { computed } from 'vue';
const globalStore = useGlobalStore();
const loading = computed(() => globalStore.loading);
</script>

<template>
  <Spin :spinning="loading">
    <RouterView />
  </Spin>
</template>

<style scoped>
</style>
