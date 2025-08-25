<template>
  <Form ref="formRef" :model="formState" :rules="rules" :layout="layout">
    <div v-for="item in formSchema" :key="item.name">
      <Form.Item
        v-if="!(item.hidden && item.hidden(formState))"
        :name="item.name"
        :label="item.label"
      >
        <slot v-if="item.slot === true" :name="`form-${item.name}`" :formState="formState">
        </slot>
        <div v-else>
          <Input
            v-if="item.type === 'input'"
            :placeholder="item.placeholder"
            v-model:value="formState[item.name]"
            :disabled="item.disabled && item.disabled(formState)"
          />
          <RadioGroup
            v-if="item.type === 'radioGroup'"
            v-model:value="formState[item.name]"
            :options="item.options ? item.options(formState) : []"
            option-type="button"
            button-style="solid"
            :disabled="item.disabled && item.disabled(formState)"
          />
          <Select
            v-if="item.type === 'select'"
            v-model:value="formState[item.name]"
            :placeholder="item.placeholder"
            :options="item.options ? item.options(formState) : []"
            :disabled="item.disabled && item.disabled(formState)"
          />
          <TreeSelect
            v-if="item.type === 'treeSelect'"
            v-model:value="formState[item.name]"
            :tree-data="item.options ? item.options(formState) : []"
            :disabled="item.disabled && item.disabled(formState)"
          />
          <InputNumber
            v-if="item.type === 'inputNumber'"
            v-model:value="formState[item.name]"
            :disabled="item.disabled && item.disabled(formState)"
          />
          <Switch
            v-if="item.type === 'switch'"
            v-model:checked="formState[item.name]"
            :disabled="item.disabled && item.disabled(formState)"
          />
          <Checkbox
            v-if="item.type === 'checkbox'"
            v-model:checked="formState[item.name]"
            :disabled="item.disabled && item.disabled(formState)"
          />
          <ApiSelect
            v-if="item.type === 'apiSelect'"
            v-model:value="formState[item.name]"
            :api="item.componentProps.api"
            :params="item.componentProps.params"
            :placeholder="item.componentProps.placeholder"
            :disabled="item.disabled && item.disabled(formState)"
          />
          <ApiTreeSelect
            v-if="item.type === 'apiTreeSelect'"
            v-model:value="formState[item.name]"
            :api="item.componentProps.api"
            :params="item.componentProps.params"
            :label-field="item.componentProps.labelField"
            :value-field="item.componentProps.valueField"
            :multiple="item.componentProps.multiple"
            :disabled="item.disabled && item.disabled(formState)"
          />
          <Input.Password
            v-if="item.type === 'inputPassword'"
            v-model:value="formState[item.name]"
            :disabled="item.disabled && item.disabled(formState)"
          />
          <SelectRole
            v-if="item.type === 'selectRole'"
            v-model:value="formState[item.name]"
            :disabled="item.disabled && item.disabled(formState)"
            :multiple="item.componentProps.multiple"
          />
          <ModelSelect
            v-if="item.type === 'modelSelect'"
            v-model:value="formState[item.name]"
            :modelType="item.componentProps.modelType"
            :disabled="item.disabled && item.disabled(formState)"
          />
          <DatasourceSelect
            v-if="item.type === 'datasourceSelect'"
            v-model:value="formState[item.name]"
            :datasourceType="item.componentProps.datasourceType"
            :disabled="item.disabled && item.disabled(formState)"
          />
        </div>
      </Form.Item>
    </div>
    <Form.Item>
      <Space v-if="!customButton || customButton === false">
        <IconButton :icon="submitBtnIcon" :title="submitBtnText" @click="submit" type="primary"/>
        <IconButton :icon="resetBtnIcon" title="重置" @click="reset" />
      </Space>
      <slot name="buttons" v-else></slot>
    </Form.Item>
  </Form>
</template>

<script lang="js" setup>
  import {
    Form,
    Input,
    Space,
    RadioGroup,
    InputNumber,
    Select,
    TreeSelect,
    Switch,
    Checkbox,
  } from 'ant-design-vue';
  import { ref, watch } from 'vue';
  import ApiTreeSelect from '@/components/Select/ApiTreeSelect.vue';
  import SelectRole from '@/components/Role/SelectRole.vue';
  import ApiSelect from '@/components/Select/ApiSelect.vue';
  import IconButton from '@/components/Button/IconButton.vue';
  import ModelSelect from '@/components/AIModel/ModelSelect.vue';
  import DatasourceSelect from '@/components/Datasource/DatasourceSelect.vue';

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
    submitBtnIcon: {
      type: String,
      default: 'SearchOutlined',
    },
    showResetBtn: {
      type: Boolean,
      default: true,
    },
    resetBtnIcon: {
      type: String,
      default: 'RedoOutlined',
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
