<template>
  <div class="message-container" :class="`message-container-${type}`">
    <div class="avatar">
      {{ type }}
    </div>
    <div class="message-bubble" :class="`message-bubble-${type}`">
      <div class="message-content">
        {{ message }}
      </div>
    </div>
  </div>
</template>

<script lang="js" setup>

defineProps({
  type: {
    type: String,
    default: 'user'
  },
  message: {
    type: String,
    default: ''
  }
});
</script>

<style scoped>
.message-container {
  max-height: 50%;
}
.avatar {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  background-color: #f0f0f0;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
}
.message-bubble {
  padding: 10px;
  border-radius: 10px;
  background-color: #f0f0f0;
  max-width: 60%;
  color: #000;
  font-size: 14px;
  max-height: 90%;
  overflow: auto;
}

.message-bubble::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}
.message-bubble::-webkit-scrollbar-track {
  background-color: transparent;
}
.message-bubble::-webkit-scrollbar-thumb {
  background-color: #bfbfbf;
  border-radius: 4px;
}
.message-bubble::-webkit-scrollbar-thumb:hover {
  background-color: #A1A1A1FF;
}

.message-container-user {
  display: flex;
  flex-direction: column;
  margin-bottom: 10px;
  gap: 10px;
  align-items: flex-end;
}
.message-container-assistant {
  display: flex;
  flex-direction: column;
  margin-bottom: 10px;
  gap: 10px;
  align-items: flex-start;
}
.message-bubble-user {
  background-color: #007bff;
  color: #fff;
}
.message-bubble-assistant {
  background-color: #28a745;
  color: #fff;
}
</style>
