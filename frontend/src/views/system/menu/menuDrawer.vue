<template>
  <Drawer :open="visible" @close="() => setVisible(false)" size="large" destroy-on-close>
    <Form
      :form-schema="menuStore.getMenuFormSchema()"
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
  import { computed, onMounted, ref } from 'vue';
  import { useMenuStore } from '@/views/system/menu/menuStore.js';
  import Form from '@/components/Form/index.vue';

  const menuStore = useMenuStore();
  const visible = ref(false);
  const formState = computed(()=>menuStore.menuForm);
  const formRules = computed(()=>menuStore.menuFormRules);
  const emit = defineEmits(['submit-ok', 'submit-fail']);
  const isUpdate = ref(false);

  const setVisible = (val) => {
    if (val) menuStore.getParentMenuOptions();
    visible.value = val;
  };

  const setUpdate = (val) => {
    isUpdate.value = val;
  };

  const onSubmit = async (formState) => {
    if (isUpdate.value === false) {
      const ok = await menuStore.addMenu(formState);
      if (ok) {
        setVisible(false);
        emit('submit-ok');
      } else {
        emit('submit-fail');
      }
    } else {
      const ok = await menuStore.updateMenu(formState);
      if (ok) {
        setVisible(false);
        emit('submit-ok');
      } else {
        emit('submit-fail');
      }
    }
  };

  const onReset = () => {
    menuStore.initMenuForm();
  };

  onMounted(()=>{
    console.log(menuStore.menuForm);
  })

  defineExpose({
    setVisible,
    setUpdate,
  });
</script>

<style scoped></style>
