<template>
  <div class="performance-container">
    <div class="performance-card">
      <Statistic :value="performance.cpuUsage" title="CPU使用率" :precision="2" suffix="%" />
      <Statistic :title="performance.cpuInfo" :value="performance.cpuCores" suffix="核心"/>
    </div>
    <div class="performance-card">
      <Statistic :value="performance.memTotal" title="内存总量" suffix="MB" />
      <Statistic :value="performance.memUsage" title="内存已用" suffix="MB" />
      <Statistic :value="performance.memFree" title="内存空闲" suffix="MB" />
      <Statistic :value="performance.memUsagePercent" title="内存使用率" :precision="2" suffix="%" />
    </div>
    <div class="performance-card">
      <Statistic :value="performance.diskTotal" title="磁盘总量" suffix="GB" />
      <Statistic :value="performance.diskUsage" title="磁盘已用" suffix="GB" />
      <Statistic :value="performance.diskFree" title="磁盘空闲" suffix="GB" />
      <Statistic :value="performance.diskUsagePercent" title="磁盘使用率" :precision="2" suffix="%" />
    </div>
    <div class="performance-card">
      <Statistic :value="performance.goroutines" title="Goroutine数量" />
      <Statistic :value="performance.numGC" title="GC次数" />
      <Statistic :value="performance.lastGC" title="上一次GC时间" />
      <Statistic :value="performance.heapSys" title="堆内存总量" suffix="MB" />
      <Statistic :value="performance.heapAlloc" title="堆内存已用" suffix="MB" />
    </div>
  </div>
</template>

<script lang="js" setup>
import { getSystemPerformance } from '@/api/monitor/performance.api.js';
import { onMounted, onUnmounted, ref } from 'vue';
import {Statistic} from 'ant-design-vue';
import { message } from 'ant-design-vue';
const performance = ref({});

const getPerformance = async () => {
  try {
    const { data } = await getSystemPerformance();
    performance.value = data;
  } catch {
    message.error("获取系统性能失败");
  }
};

const interval = ref(null);

onMounted(async () => {
  await getPerformance();
  interval.value = setInterval(() => {
    getPerformance();
  }, 5000);
});

onUnmounted(() => {
  clearInterval(interval.value);
});

</script>

<style scoped>
.performance-container {
  background-color: transparent;
  display: flex;
  justify-content: flex-start;
}
.performance-card {
  background-color: #fff;
  padding: 20px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  width: 30%;
  max-height: 90%;
  margin: 10px
}
.performance-card:hover {
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
}
</style>
