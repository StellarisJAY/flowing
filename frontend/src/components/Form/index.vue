<template>
  <Form ref="form" :model="formState" :rules="rules" layout="inline">
    <Form.Item v-for="item in formSchema" :key="item.name" :name="item.name" :label="item.label">
      <Input
        v-if="item.type === 'input'"
        :label="item.label"
        :placeholder="item.placeholder"
        v-model:value="formState[item.name]"
      />
    </Form.Item>
    <Form.Item>
      <Space>
        <Button type="primary" @click="submit">{{submitBtnText}}</Button>
        <Button v-if="showResetBtn" @click="reset">重置</Button>
      </Space>
    </Form.Item>
  </Form>
</template>

<script lang="js" setup>
import { Button, Form, Input, Space } from 'ant-design-vue';
import { onMounted, reactive, ref } from 'vue';

const props = defineProps({
  formSchema: {
    type: Array,
    default: ()=>{}
  },
  rules: {
    type: Array,
    default: ()=>[],
  },
  submitBtnText: {
    type: String,
    default: '查询'
  },
  showResetBtn: {
    type: Boolean,
    default: true
  }
})

const emit = defineEmits(['submit', 'reset']);

const formState = ref({});
const form = ref();

onMounted((()=>{
  initFormState();
}));

const initFormState = ()=>{
  for (const item of props.formSchema) {
    formState.value[item.name] = item.defaultValue;
  }
};

const submit = async () => {
  try {
    await form.value.validate();
    emit('submit', formState.value);
  } catch (error) {
    console.log(error);
  }
};

const reset = ()=>{
  initFormState();
  emit('reset');
}
</script>

<style scoped></style>
