<template>
  <div class="wrapper">
    <div class="content-wrapper" ref="contentWrapperRef">
      <textarea
        class="content"
        v-for="item in records"
        :key="item.id"
        :value="item.content"
        disabled
      />
    </div>
    <div class="pagination-wrapper">
      <Pagination :total="total" :page-size="pageSize" :current="current" @change="onPageChange" />
    </div>
  </div>
</template>

<script lang="js" setup>
  import { message, Pagination } from 'ant-design-vue';
  import { onMounted, ref } from 'vue';
  import { listChunks } from '@/api/ai/document.api.js';
  import { useRoute } from 'vue-router';
  import { useGlobalStore} from '@/stores/global.js'

  const globalStore = useGlobalStore();
  const route = useRoute();
  const docId = ref(route.query.id);
  const total = ref(0);
  const pageSize = ref(10);
  const current = ref(1);
  const records = ref([]);
  const contentWrapperRef = ref();

  const list = async () => {
    globalStore.setLoading(true);
    try {
      const res = await listChunks({
        docId: docId.value,
        page: true,
        pageSize: pageSize.value,
        pageNum: current.value,
      });
      total.value = res.total;
      records.value = res.data;
    } catch {
      message.error('获取文档分块失败');
    }finally {
      globalStore.setLoading(false);
    }
  };

  const onPageChange = async (page, size) => {
    current.value = page;
    pageSize.value = size;
    await list();
    contentWrapperRef.value.scrollTop = 0;
  };

  onMounted(() => {
    list();
  });
</script>

<style scoped>
  .wrapper {
    height: 100%;
    width: 100%;
    padding: 10px;
    display: flex;
    flex-direction: column;
    gap: 10px;
    .content-wrapper {
      height: 90%;
      width: 100%;
      overflow: auto;
      display: flow;
    }
    .content {
      height: 100%;
      width: 100%;
      border: #0096ff 2px solid;
      resize: none;
      border-radius: 10px;
      padding: 10px;
    }
    .content:focus {
      border: #0081dc 2px solid;
    }
    .content:hover {
      border: #0081dc 2px solid;
    }
    textarea:focus {
      outline: none;
    }
    textarea:disabled {
      background-color: white;
      outline: none;
    }
    .content::-webkit-scrollbar {
      width: 0;
      height: 0;
    }
    .pagination-wrapper {
      height: 10%;
      width: 100%;
      display: flex;
      justify-content: flex-end;
    }
  }
</style>
