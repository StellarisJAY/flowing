<template>
  <div class="kb-retrieve">
    <div class="kb-retrieve-options">
      <div class="kb-retrieve-header">
        检索测试
      </div>
      <div class="kb-retrieve-form">
        <div class="form-item-text">
          <textarea class="form-item-textarea" placeholder="请输入检索文本" v-model="formState.queryText" />
        </div>
        <div class="form-item-settings">
          <div class="form-item">
            <div>检索类型</div>
            <Select :options="retrieveTypeOptions" v-model:value="formState.searchType" />
          </div>
          <div class="form-item">
            <div>TopK</div>
            <InputNumber v-model:value="formState.topK" :min="1" :max="100" :step="1" />
          </div>
          <div class="form-item" v-if="formState.searchType === 'hybrid'">
            <div>混合检索排序</div>
            <RadioGroup :options="hybridRerankOptions" v-model:value="formState.hybridType" />
          </div>
          <div class="form-item" v-if="formState.searchType !== 'fulltext'">
            <div>相似度阈值 (≥{{formState.threshold}})</div>
            <Slider v-model:value="formState.threshold" :min="0.01" :max="0.99" :step="0.01" />
          </div>

          <div class="form-item" v-if="formState.searchType === 'hybrid' && formState.hybridType === 'weight'">
            <div>
              权重 (全文: {{Math.round((1-formState.weight) * 10)/10}} 向量：{{formState.weight}})
            </div>
            <Slider v-model:value="formState.weight" :min="0" :max="1" :step="0.1" />
          </div>
          <div class="form-item" v-if="formState.searchType === 'hybrid' && formState.hybridType === 'rerank'">
            <div>
              重排序模型
            </div>
            <ModelSelect model-type="rerank" />
          </div>
        </div>
        <div class="form-item form-item-search-btn">
          <IconButton icon="SearchOutlined" title="搜索" type="primary" @click="search" />
        </div>
      </div>
    </div>
    <div class="kb-retrieve-result">
      <div class="kb-retrieve-items" v-if="searchResult && searchResult.length > 0">
        <div v-for="item in searchResult" :key="item.sliceId" class="kb-retrieve-item">
          <RetrieveContent :item="item" />
        </div>
      </div>
      <div class="kb-retrieve-empty" v-else>
        <Empty />
      </div>
    </div>
  </div>
</template>

<script lang="js" setup>
import IconButton from '@/components/Button/IconButton.vue';
import { Select, Slider, RadioGroup, InputNumber, message, Empty } from 'ant-design-vue';
import { ref } from 'vue';
import ModelSelect from '@/components/AIModel/ModelSelect.vue';
import RetrieveContent from '@/views/agent/knowledge/retrieve/RetrieveContent.vue';
import { searchKnowledge } from '@/api/ai/knowledge.api.js';
import { useGlobalStore } from '@/stores/global.js';
import {useRoute} from 'vue-router';

const route = useRoute();
const knowledgeBaseId = route.query.knowledgeBaseId;
const globalStore = useGlobalStore();
const searchResult = ref([]);

const retrieveTypeOptions = [
  {
    label: '全文检索',
    value: 'fulltext'
  },
  {
    label: '向量检索',
    value: 'vector'
  },
  {
    label: '混合检索',
    value: 'hybrid'
  }
];

const formState = ref({
  knowledgeBaseId,
  queryText: '',
  searchType: 'fulltext',
  threshold: 0.5,
  topK: 10,
  hybridType: 'weight',
  weight: 0.5,
});

const hybridRerankOptions = [
  {
    label: '权重',
    value: 'weight'
  },
  {
    label: '重排序模型',
    value: 'rerank'
  }
];

const search = async () => {
  if (!formState.value.queryText) {
    message.error('请输入检索文本');
    return;
  }
  globalStore.setLoading(true);
  try {
    const { data } = await searchKnowledge(formState.value);
    searchResult.value = data;
  }catch {
    message.error('检索失败');
  }finally {
    globalStore.setLoading(false);
  }
};

</script>

<style scoped>
.kb-retrieve {
  height: 100%;
  width: 100%;
  display: flex;
  gap: 10px;
}

.kb-retrieve-header {
  font-size: 18px;
  font-weight: 500;
  margin-bottom: 10px;
}

.kb-retrieve-options {
  height: 100%;
  width: 40%;
  background-color: white;
  padding: 10px;
  .kb-retrieve-form {
    height: 90%;
    width: 100%;
    overflow: auto;
    display: flex;
    flex-direction: column;
    gap: 10px;

    .form-item {
      display: flex;
      flex-direction: column;
      gap: 5px;
      width: 100%;
      padding-left: 10px;
      padding-right: 10px;
    }

    .form-item-text {
      height: 50%;
      width: 100%;
    }
    .form-item-textarea {
      height: 100%;
      width: 100%;
      border: #0096ff 2px solid;
      resize: none;
      border-radius: 10px;
      padding: 10px;
    }
    .form-item-textarea:focus{
      height: 100%;
      width: 100%;
      border: #0081dc 2px solid;
    }
    .form-item-textarea:hover{
      height: 100%;
      width: 100%;
      border: #0081dc 2px solid;
    }
    textarea:focus {
      outline: none;
    }

    .form-item-settings {
      height: 40%;
      width: 100%;
      overflow: auto;
      display: flex;
      flex-direction: column;
      gap: 10px;
    }

    .form-item-settings::-webkit-scrollbar {
      width: 8px;
      height: 8px;
    }
    .form-item-settings::-webkit-scrollbar-track {
      background-color: transparent;
    }
    .form-item-settings::-webkit-scrollbar-thumb {
      background-color: #bfbfbf;
      border-radius: 4px;
    }
    .form-item-settings::-webkit-scrollbar-thumb:hover {
      background-color: #A1A1A1FF;
    }

    .form-item-search-btn {
      height: 5%;
      width: 100%;
    }
  }
}

.kb-retrieve-result {
  height: 100%;
  width: 60%;
  background-color: transparent;
  display: flow;
  flex-direction: column;
  overflow: auto;
  .kb-retrieve-empty {
    height: 100%;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .kb-retrieve-item {
    height: 50%;
    width: 100%;
    margin-bottom: 10px;
  }
  .kb-retrieve-result::-webkit-scrollbar {
    width: 8px;
    height: 8px;
  }
  .kb-retrieve-result::-webkit-scrollbar-track {
    background-color: transparent;
  }
  .kb-retrieve-result::-webkit-scrollbar-thumb {
    background-color: #bfbfbf;
    border-radius: 4px;
  }
  .kb-retrieve-items::-webkit-scrollbar-thumb:hover {
    background-color: #A1A1A1FF;
  }
}


</style>
