<template>
  <div class="kb-retrieve">
    <div class="kb-retrieve-options">
      <div class="kb-retrieve-header">
        检索测试
      </div>
      <div class="kb-retrieve-form">
        <div class="form-item-text">
          <textarea class="form-item-textarea" placeholder="请输入检索文本"/>
        </div>
        <div class="form-item-settings">
          <div class="form-item">
            <div>检索类型</div>
            <Select :options="retrieveTypeOptions" v-model:value="formState.retrieveType" />
          </div>
          <div class="form-item" v-if="formState.retrieveType === 'hybrid'">
            <div>混合检索排序</div>
            <RadioGroup :options="hybridRerankOptions" v-model:value="formState.hybridRerank" />
          </div>
          <div class="form-item" v-if="formState.retrieveType !== 'fulltext'">
            <div>相似度阈值 (≥{{formState.threshold}})</div>
            <Slider v-model:value="formState.threshold" :min="0.01" :max="0.99" :step="0.01" />
          </div>

          <div class="form-item" v-if="formState.retrieveType === 'hybrid' && formState.hybridRerank === 'weight'">
            <div>
              权重 (向量：{{formState.vectorWeight}}, 全文: {{Math.round((1-formState.vectorWeight) * 10)/10}})
            </div>
            <Slider v-model:value="formState.vectorWeight" :min="0" :max="1" :step="0.1" />
          </div>
          <div class="form-item" v-if="formState.retrieveType === 'hybrid' && formState.hybridRerank === 'rerank'">
            <div>
              重排序模型
            </div>
            <ModelSelect model-type="rerank" />
          </div>
        </div>
        <div class="form-item form-item-search-btn">
          <IconButton icon="SearchOutlined" title="搜索" type="primary"/>
        </div>
      </div>
    </div>
    <div class="kb-retrieve-result">
      <div class="kb-retrieve-items">
        <div v-for="item in retrieveResultMock" :key="item.srcDocId" class="kb-retrieve-item">
          <RetrieveContent :item="item" />
        </div>
      </div>
      <div class="kb-retrieve-pagination">
        <Pagination />
      </div>
    </div>
  </div>
</template>

<script lang="js" setup>
import IconButton from '@/components/Button/IconButton.vue';
import { Select, Slider, RadioGroup, Pagination } from 'ant-design-vue';
import { ref } from 'vue';
import ModelSelect from '@/components/AIModel/ModelSelect.vue';
import RetrieveContent from '@/views/agent/knowledge/retrieve/RetrieveContent.vue';

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
  retrieveType: 'fulltext',
  threshold: 0.5,
  hybridRerank: 'weight',
  vectorWeight: 0.5,
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

const retrieveResultMock = [
  {
    srcDocName: '文档1',
    srcDocId: '1',
    content: '11',
    scores: {
      hybrid: 0.8,
      fulltext: 0.7,
      vector: 0.9,
    },
  },
  {
    srcDocName: '文档1',
    srcDocId: '1',
    content: '这是文档1的内容',
    scores: {
      hybrid: 0.8,
      fulltext: 0.7,
      vector: 0.9,
    },
  },
  {
    srcDocName: '文档1',
    srcDocId: '1',
    content: '这是文档1的内容',
    scores: {
      hybrid: 0.8,
      fulltext: 0.7,
      vector: 0.9,
    },
  },
  {
    srcDocName: '文档1',
    srcDocId: '1',
    content: '这是文档1的内容',
    scores: {
      hybrid: 0.8,
      fulltext: 0.7,
      vector: 0.9,
    },
  },
  {
    srcDocName: '文档1',
    srcDocId: '1',
    content: '这是文档1的内容',
    scores: {
      hybrid: 0.8,
      fulltext: 0.7,
      vector: 0.9,
    },
  },
  {
    srcDocName: '文档1',
    srcDocId: '1',
    content: '这是文档1的内容',
    scores: {
      hybrid: 0.8,
      fulltext: 0.7,
      vector: 0.9,
    },
  },
  {
    srcDocName: '文档1',
    srcDocId: '1',
    content: '这是文档1的内容',
    scores: {
      hybrid: 0.8,
      fulltext: 0.7,
      vector: 0.9,
    },
  },
]

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
  display: flex;
  flex-direction: column;
  gap: 10px;
  .kb-retrieve-items {
    width: 100%;
    height: 90%;
    overflow: auto;
    display: flow;
    .kb-retrieve-item {
      height: 50%;
      width: 100%;
      margin-bottom: 10px;
    }
  }
  .kb-retrieve-pagination {
    height: 10%;
    width: 100%;
    display: flex;
    justify-content: flex-end;
  }

  .kb-retrieve-items::-webkit-scrollbar {
    width: 8px;
    height: 8px;
  }
  .kb-retrieve-items::-webkit-scrollbar-track {
    background-color: transparent;
  }
  .kb-retrieve-items::-webkit-scrollbar-thumb {
    background-color: #bfbfbf;
    border-radius: 4px;
  }
  .kb-retrieve-items::-webkit-scrollbar-thumb:hover {
    background-color: #A1A1A1FF;
  }
}


</style>
