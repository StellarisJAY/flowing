<template>
  <component :is="VueIcon[icon]" />
</template>

<script lang="js" setup>
import * as VueIcon from '@ant-design/icons-vue';
import { defineProps } from 'vue';
defineProps({
  icon: {
    type: String,
    default: ''
  }
});
</script>
