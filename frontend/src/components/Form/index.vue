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
      </Form.Item>
    </div>
    <Form.Item>
      <Space>
        <Button type="primary" @click="submit">{{ submitBtnText }}</Button>
        <Button v-if="showResetBtn" @click="reset">重置</Button>
      </Space>
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
  });

  const formState = ref({ ...props.state });
  watch(() => props.state, (newVal) => {
    formState.value = { ...newVal };
  });
  const emit = defineEmits(['submit', 'reset']);
  const formRef = ref();

  const submit = async () => {
    try {
      await formRef.value.validate();
      emit('submit', formState.value);
    } catch (error) {
      console.log(error);
    }
  };

  const reset = () => {
    emit('reset');
  };
</script>

<style scoped></style>
