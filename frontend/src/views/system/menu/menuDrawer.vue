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
  const props = defineProps({
    isUpdate: {
      type: Boolean,
      default: false,
    },
  });
  const emit = defineEmits(['submit-ok', 'submit-fail']);

  const setVisible = (val) => {
    if (val) menuStore.getParentMenuOptions();
    visible.value = val;
  };

  const onSubmit = async (formState) => {
    if (!props.isUpdate) {
      const ok = await menuStore.addMenu(formState);
      if (ok) {
        setVisible(false);
        emit('submit-ok');
      } else {
        emit('submit-fail');
      }
    } else {
      // menuStore.updateMenu(formState);
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
  });
</script>

<style scoped></style>
