<template>
  <Form ref="formRef" :model="formState" :rules="rules" :layout="layout">
    <div v-for="item in formSchema" :key="item.name">
      <Form.Item
        v-if="!(item.hidden && item.hidden(formState))"
        :name="item.name"
        :label="item.label"
      >
        <Input
          v-if="item.type === 'input'"
          :placeholder="item.placeholder"
          v-model:value="formState[item.name]"
        />
        <RadioGroup
          v-if="item.type === 'radioGroup'"
          v-model:value="formState[item.name]"
          :options="item.options ? item.options(formState) : []"
          option-type="button"
          button-style="solid"
        />
        <Select
          v-if="item.type === 'select'"
          v-model:value="formState[item.name]"
          :options="item.options ? item.options(formState) : []"
        />
        <TreeSelect
          v-if="item.type === 'treeSelect'"
          v-model:value="formState[item.name]"
          :tree-data="item.options ? item.options(formState) : []"
        />
        <InputNumber v-if="item.type === 'inputNumber'" v-model:value="formState[item.name]" />
        <Switch v-if="item.type === 'switch'" v-model:checked="formState[item.name]" />
      </Form.Item>
    </div>
    <Form.Item>
      <Space v-if="!customButton || customButton === false">
        <Button type="primary" @click="submit">{{ submitBtnText }}</Button>
        <Button v-if="showResetBtn" @click="reset">重置</Button>
      </Space>
      <slot name="buttons" v-else></slot>
    </Form.Item>
  </Form>
</template>

<script lang="js" setup>
  import {
    Button,
    Form,
    Input,
    Space,
    RadioGroup,
    InputNumber,
    Select,
    TreeSelect,
    Switch,
  } from 'ant-design-vue';
  import { ref, watch } from 'vue';

  const props = defineProps({
    formSchema: {
      type: Array,
      default: () => [],
    },
    rules: {
      type: Object,
      default: () => ({}),
    },
    submitBtnText: {
      type: String,
      default: '提交',
    },
    showResetBtn: {
      type: Boolean,
      default: true,
    },
    layout: {
      type: String,
      default: 'inline',
    },
    state: {
      type: Object,
      default: () => ({}),
    },
    customButton: {
      type: Boolean,
      default: false,
    },
    // 表单提交调用的方法，返回值为提交是否成功
    submitFunc: {
      type: Function,
      default: () => async () => {},
    },
    resetFunc: {
      type: Function,
      default: () => async () => {},
    },
  });

  const formState = ref({ ...props.state });
  watch(
    () => props.state,
    (newVal) => {
      formState.value = { ...newVal };
    }
  );
  const emit = defineEmits(['submit', 'reset']);
  const formRef = ref();

  const submit = async () => {
    try {
      await formRef.value.validate();
      return await props.submitFunc(formState.value);
    } catch (error) {
      console.log(error);
      return false;
    }
  };

  const reset = async () => {
    await props.resetFunc();
    emit('reset');
  };

  defineExpose({
    submit,
    reset,
  });
</script>

<style scoped></style>
