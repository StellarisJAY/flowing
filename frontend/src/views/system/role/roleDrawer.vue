<template>
  <Drawer :open="visible" @close="() => setVisible(false)" size="large" destroy-on-close>
    <Form
      :form-schema="formSchema"
      :rules="formRules"
      :state="formState"
      @submit="onSubmit"
      @reset="onReset"
      layout="vertical"
    />
  </Drawer>
</template>

<script lang="js" setup>
import { Drawer } from 'ant-design-vue';
import { computed, ref } from 'vue';
import { useRoleStore } from '@/views/system/role/roleStore.js';
import Form from '@/components/Form/index.vue';

const roleStore = useRoleStore();
const visible = ref(false);
const formSchema = computed(()=>roleStore.roleFormSchema);
const formState = computed(()=>roleStore.roleForm);
const formRules = computed(()=>roleStore.roleFormRules);
defineProps({
  isUpdate: {
    type: Boolean,
    default: false,
  },
});
const emit = defineEmits(['submit-ok', 'submit-fail']);

const setVisible = (val) => {
  visible.value = val;
};

const onSubmit = async (formState) => {
  const ok = await roleStore.createRole(formState);
  if(ok){
    emit('submit-ok');
    setVisible(false);
  }else {
    emit('submit-fail');
  }
};

const onReset = () => {
  roleStore.resetRoleForm();
};

defineExpose({
  setVisible,
});
</script>

<style scoped></style>
