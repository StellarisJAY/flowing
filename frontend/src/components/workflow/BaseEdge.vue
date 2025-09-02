<template>
  <SmoothStepEdge
    :source-position="sourcePosition"
    :target-position="targetPosition"
    :target-x="targetX"
    :target-y="targetY"
    :source-x="sourceX"
    :source-y="sourceY"
    :marker-end="markerEnd"
    class="base-edge"
  />
  <!--  <EdgeLabelRenderer>-->
  <!--    <div-->
  <!--      :style="{-->
  <!--        pointerEvents: 'all',-->
  <!--        position: 'absolute',-->
  <!--        transform: `translate(-50%, -50%) translate(${path[1]}px,${path[2]}px)`,-->
  <!--      }"-->
  <!--      class="nodrag nopan"-->
  <!--    >-->
  <!--    </div>-->
  <!--  </EdgeLabelRenderer>-->
</template>

<script setup>
  import { SmoothStepEdge } from '@vue-flow/core';

  const props = defineProps({
    sourceX: {
      type: Number,
      required: true,
    },
    sourceY: {
      type: Number,
      required: true,
    },
    targetX: {
      type: Number,
      required: true,
    },
    targetY: {
      type: Number,
      required: true,
    },
    sourcePosition: {
      type: String,
      required: true,
    },
    targetPosition: {
      type: String,
      required: true,
    },
    data: {
      type: Object,
      required: true,
    },
    markerEnd: {
      type: String,
      required: true,
    },
  });

  // const path = computed(() => getSmoothStepPath(props));
</script>

<script>
  export default {
    inheritAttrs: false,
  };
</script>

<style scoped>
  .base-edge {
    z-index: 9999 !important;
  }
</style>
