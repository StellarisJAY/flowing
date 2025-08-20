<template>
  <div class="retrieve-content">
    <div class="header">
      <div class="docs-scores">
        {{`得分: ${item.scores.hybrid} 全文检索: ${item.scores.fulltext} 向量检索: ${item.scores.vector}`}}
      </div>
      <div class="doc-src">
        <Button type="link">{{item.srcDocName}}</Button>
      </div>
    </div>
    <div class="content">
      {{item.content}}
    </div>
  </div>
</template>

<script lang="js" setup>
  import {Button} from 'ant-design-vue';

  defineProps({
    item: {
      type: Object,
      default: () => ({
        srcDocName: '',
        srcDocId: '',
        content: '',
        scores: {
          hybrid: 0.0,
          fulltext: 0.0,
          vector: 0.0,
        },
      }),
    },
  });
</script>

<style scoped>
  .retrieve-content {
    height: 100%;
    width: 100%;
    background-color: white;
    border-radius: 2px;
    padding: 10px;
    display: flex;
    flex-direction: column;
    gap: 5px;
  }
  .header {
    height: 10%;
    width: 100%;
    display: flex;
    justify-content: flex-start;
    gap: 10px;
    .docs-scores {
      font-size: 14px;
      color: #666666;
      width: 50%;
    }
    .doc-src {
      width: 50%;
      text-align: right;
    }
  }
  .content {
    height: 90%;
    width: 100%;
    overflow: auto;
    border: 1px solid #bfbfbf;
    border-radius: 2px;
    padding: 10px;
  }
  .content::-webkit-scrollbar {
    width: 0;
    height: 0;
  }

</style>
