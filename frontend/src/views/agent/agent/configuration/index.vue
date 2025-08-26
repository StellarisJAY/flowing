<template>
  <div class="agent-configure">
    <div class="sidebar">
      <IconButton type="primary" icon="ArrowLeftOutlined" title="返回" @click="handleBack" style="width: 100%;"/>
    </div>
    <div class="content">
      <RouterView />
    </div>
  </div>
</template>

<script lang="js" setup>
import { RouterView } from 'vue-router';
import { computed} from 'vue';
import {useRoute, useRouter} from 'vue-router';
import IconButton from '@/components/Button/IconButton.vue';

const router = useRouter();
const route = useRoute();
const query = computed(()=>route.query);
const handleBack = () => {
  router.back();
};

</script>
<style scoped>
.agent-configure {
  height: 100%;
  width: 100%;
  display: flex;
  gap: 10px;
}

.sidebar {
  height: 100%;
  width: 20%;
  overflow: auto;
  background-color: white;
  padding: 10px;
}

.content {
  height: 100%;
  width: 80%;
  background-color: transparent;
  display: flex;
}
</style>
