<template>
  <div class="retrieve-content">
    <div class="header">
      <div class="docs-scores">
        {{`得分: ${formatScore(item.score)} 全文检索: ${formatScore(item.fulltextScore)} 向量检索: ${formatScore(item.vectorScore)}`}}
      </div>
      <div class="docs-src">
        <div class="docs-src-title">源文档: </div>
        <div class="docs-src-link">{{item.documentName}}</div>
      </div>
    </div>
    <div class="content">
      {{item.content}}
    </div>
  </div>
</template>

<script lang="js" setup>

  defineProps({
    item: {
      type: Object,
      default: () => ({}),
    },
  });

  const formatScore = (score) => {
    return score.toFixed(2);
  }
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
    .docs-src {
      width: 50%;
      display: flex;
      justify-content: flex-end;
      gap: 5px;
      .docs-src-title {
        font-size: 14px;
        color: #666666;
      }
      .docs-src-link {
        font-size: 14px;
        color: #1890ff;
        text-decoration: underline;
      }
      .docs-src-link:hover{
        color: #40a9ff;
        cursor: pointer;
      }
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
