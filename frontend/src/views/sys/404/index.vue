<template>
  <div class="content">
    <div class="title">404</div>
    <div class="desc">页面不存在</div>
    <div class="btn">
      <Button type="link" @click="handleBack">返回主页</Button>
    </div>
  </div>
</template>
<script setup lang="js">
  import { Button } from 'ant-design-vue';
  import { useRouter } from 'vue-router';
  const router = useRouter();

  const handleBack = () => {
    router.replace('/');
  };
</script>

<style scoped>
  .content {
    height: 100%;
    background-color: white;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }
  .title {
    font-size: 40px;
    font-weight: bold;
    margin-bottom: 20px;
    color: #ff4d4f;
    text-align: center;
  }
  .desc {
    font-size: 18px;
    margin-bottom: 20px;
    color: #999;
    text-align: center;
  }
  .btn {
    margin-top: 20px;
  }
</style>
