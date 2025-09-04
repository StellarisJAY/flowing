<template>
  <div class="prompt-card">
    <div class="prompt-title">
      {{title}}
     <Dropdown>
       <Button type="link">{x}</Button>
       <template #overlay>
         <Menu :items="variables" @click="(e)=>onVarMenuClick(e.item)" />
       </template>
     </Dropdown>
    </div>
    <textarea class="input-text" v-model="value"/>
  </div>
</template>

<script lang="js" setup>
  import { Dropdown, Menu, Button} from 'ant-design-vue';
  import { computed } from 'vue';
  import { useWorkflowStore } from '@/stores/workflow';

  const props = defineProps({
    title: {
      type: String,
      default: '提示词'
    },
    nodeId: {
      type: String,
      required: true,
    },
    allowedVarTypes: {
      type: Array,
      default: ()=>(['string'])
    },
  });

  const value = defineModel('value', {
    type: String,
    required: true,
  });

  const workflowStore = useWorkflowStore();

  const variables = computed(()=>{
    const variables = workflowStore.getRefVariables(props.nodeId);
    return variables.filter(item=>{
      item.children = item.children.filter(variable=>props.allowedVarTypes.includes(variable.type));
      return item.children.length > 0;
    });
  });

  const onVarMenuClick = (item) => {
    value.value = value.value + `{{${item.value}}}`
  };
</script>

<style scoped>
.prompt-card {
  border-radius: 10px;
  padding: 10px;
  background-color: white;
  border: #999999 1px solid;

  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 10px;
  height: 200px;
}

.prompt-title {
  height: 20px;
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.input-text {
  resize: none;
  outline: none;
  height: 100%;
  overflow: auto;
  border: #007bff 2px solid;
  border-radius: 10px;
  padding: 10px;
  font-size: 14px;
}
</style>
