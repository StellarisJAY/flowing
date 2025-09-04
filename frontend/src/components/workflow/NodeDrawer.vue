<template>
  <Drawer :open="visible" title="配置节点" @close="close" destroy-on-close size="large">
    <!--基础信息配置-->
    <Form :model="nodeBaseForm.state" :rules="nodeBaseForm.rules" layout="vertical">
      <Form.Item label="节点名称" name="label">
        <Input v-model:value="nodeBaseForm.state.label" />
      </Form.Item>
      <Form.Item label="节点描述" name="description">
        <Input v-model:value="nodeBaseForm.state.description" />
      </Form.Item>
    </Form>
    <Divider />
    <!--节点类型配置-->
    <div class="config-section" v-if="nodeData.type !== 'start'">
      <!-- 模型配置 -->
      <ModelConfig :node-id="nodeId" v-model:config="config" v-if="nodeData.type === 'model'" />
      <!-- 回复配置 -->
      <ReplyNodeConfig :node-id="nodeId" v-model:config="config" v-if="nodeData.type === 'reply'" />
      <!-- 知识库配置 -->
      <KnowledgeNodeConfig :node-id="nodeId" v-model:config="config" v-if="nodeData.type === 'knowledge'" />
      <Divider />
    </div>
    <!--输入输出变量-->
    <Form layout="vertical">
      <Form.Item label="节点输入" v-if="(config.input && config.input.length) || canAddInput">
        <VariableList :node-id="nodeId" :allow-add="canAddInput" is-input v-model:variables="config.input" />
      </Form.Item>
      <Form.Item :label="nodeData.type === 'start' ? '全局变量' : '节点输出'" v-if="(config.output && config.output.length) || canAddOutput">
        <VariableList :node-id="nodeId" :allow-add="canAddOutput" v-model:variables="config.output"/>
      </Form.Item>
    </Form>
    <Divider/>
    <!--下一步节点列表-->
    <h4>下一步</h4>
    <NextStep :node-id="nodeId" />
  </Drawer>
</template>

<script setup lang="js">
  import { Drawer, Form, Input, Divider } from 'ant-design-vue';
  import { computed, ref } from 'vue';
  import { allowAddInputVariable, allowAddOutputVariable, useWorkflowStore } from '@/stores/workflow';
  import VariableList from '@/components/workflow/config/variable/VariableList.vue';
  import NextStep from '@/components/workflow/config/NextStep.vue';
  import ModelConfig from '@/components/workflow/config/model/ModelConfig.vue'
  import ReplyNodeConfig from '@/components/workflow/config/reply/ReplyNodeConfig.vue'
  import KnowledgeNodeConfig from '@/components/workflow/config/knowledge/KnowledgeNodeConfig.vue'

  const workflowStore = useWorkflowStore();
  const visible = ref(false);
  // 节点配置信息
  const config = ref({});
  // 节点id
  const nodeId = ref('');
  // 节点基础信息
  const nodeData = ref({type:''});
  // 该节点类型是否允许添加输入
  const canAddInput = computed(()=>allowAddInputVariable(nodeData.value.type));
  // 该节点类型是否允许添加输出
  const canAddOutput = computed(()=>allowAddOutputVariable(nodeData.value.type));

  // 基础信息form
  const nodeBaseForm = ref({
    state: {
      label: '',
      description: '',
    },
    rules: {
      label: [{ required: true, message: '请输入节点名称', trigger: 'blur' }],
    },
  });

  const open = (id, data, conf) => {
    visible.value = true;
    nodeId.value = id;
    config.value = conf;
    nodeData.value = data;
    nodeBaseForm.value.state.label = data.label;
    nodeBaseForm.value.state.description = data.description;
  };

  const close = () => {
    // 关闭时更新节点基本信息和配置
    workflowStore.updateNodeInfo(
      nodeId.value,
      nodeBaseForm.value.state.label,
      nodeBaseForm.value.state.description,
      config.value
    );
    visible.value = false;
    nodeBaseForm.value.state.label = '';
    nodeBaseForm.value.state.description = '';
    nodeId.value = '';
    config.value = {};
  };

  defineExpose({
    open,
    close,
  });
</script>

<style scoped>
.config-section {
  background-color: #fff;
  width: 100%;
  overflow: auto;
  padding: 5px;
  margin-bottom: 5px;
  border-radius: 8px;
}
</style>
