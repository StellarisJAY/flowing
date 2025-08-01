<template>
  <Drawer :open="visible" @close="() => setVisible(false)">
    <Form v-model:form="formState" layout="vertical" :rules="rules">
      <Form.Item label="菜单名称" name="menuName">
        <Input v-model:value="formState.menuName" />
      </Form.Item>
      <Form.Item label="菜单类型" name="type">
        <RadioGroup
          optionType="button"
          v-model:value="formState.type"
          :options="menuStore.getMenuTypeOptions()"
        />
      </Form.Item>
      <Form.Item v-if="formState.type !== 1" label="父菜单" name="parentId">
        <Select v-model:value="formState.parentId" />
      </Form.Item>
      <Form.Item v-if="formState.type !== 3" label="路径" name="path">
        <Input v-model:value="formState.path" />
      </Form.Item>
      <Form.Item v-if="formState.type === 3" label="权限标识" name="actionCode">
        <Input v-model:value="formState.actionCode" />
      </Form.Item>
      <Form.Item v-if="formState.type !== 3" label="组件" name="component">
        <Input v-model:value="formState.component" />
      </Form.Item>
      <Form.Item v-if="formState.type !== 3" label="图标" name="icon">
        <Input v-model:value="formState.icon" />
      </Form.Item>
      <Form.Item v-if="formState.type !== 3" label="排序" name="sort">
        <InputNumber v-model:value="formState.orderNum" />
      </Form.Item>
      <Form.Item>
        <Button type="primary">保存</Button>
      </Form.Item>
    </Form>
  </Drawer>
</template>

<script lang="js" setup>
  import { Drawer, Form, Input, InputNumber, Select, RadioGroup, Button } from 'ant-design-vue';
  import { computed, reactive, ref } from 'vue';
  import { useMenuStore } from '@/views/system/menu/menuStore.js';

  const menuStore = useMenuStore();
  const formState = reactive(menuStore.menuForm);
  const rules = computed(()=>menuStore.menuRules);
  const visible = ref(false);

  const setVisible = (val) => {
    visible.value = val;
  };

  defineExpose({
    setVisible,
  });
</script>

<style scoped></style>
