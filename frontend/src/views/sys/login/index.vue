<template>
  <div class="flowing-login-bg">
    <div class="flowing-login-card">
      <div class="flowing-login-side"></div>
      <div class="flowing-login-form">
        <Form layout="vertical" ref="formRef" :rules="rules" :model="loginForm">
          <Form.Item label="用户名" name="username">
            <Input v-model:value="loginForm.username" />
          </Form.Item>
          <Form.Item label="密码" name="password">
            <Input.Password v-model:value="loginForm.password" />
          </Form.Item>
          <Form.Item label="验证码" name="captcha">
            <div class="flowing-captcha">
              <Input v-model:value="loginForm.captcha" class="flowing-captcha-input" />
              <img
                :src="captchaImage"
                alt="验证码"
                class="flowing-captcha-image"
                @click="refreshCaptcha"
              />
            </div>
          </Form.Item>
          <Form.Item>
            <Button type="primary" class="flowing-login-button" @click="onSubmit">登录</Button>
          </Form.Item>
        </Form>
      </div>
    </div>
  </div>
</template>

<script lang="js" setup>
  import { Button, Form, Input, message} from 'ant-design-vue';
  import { onMounted, reactive, ref } from 'vue';
  import { getCaptcha, login } from './api.js';
  import { useUserStore } from '@/stores/user.js';
  import {useRouter} from 'vue-router';

  const loading = ref(false);
  const formRef = ref();
  const rules = {
    username: [
      {
        required: true,
        message: '请输入用户名',
        trigger: 'submit',
      },
    ],
    password: [
      {
        required: true,
        message: '请输入密码',
        trigger: 'submit',
      },
    ],
    captcha: [
      {
        required: true,
        message: '请输入验证码',
        trigger: 'submit',
      },
    ],
  };
  const loginForm = reactive({
    username: '',
    password: '',
    captcha: '',
    captchaKey: '',
  });
  const captchaImage = ref('');
  const store = useUserStore();
  const router = useRouter();

  onMounted(() => {
    refreshCaptcha();
  });

  const onSubmit = () => {
    formRef.value
      .validate()
      .then(() => {
        loading.value = true;
        return login(loginForm);
      })
      .then((res) => {
        store.setToken(res.data.token);
        store.setUserInfo(res.data.user);
        message.success("登录成功").then();
        router.push("/system");
      })
      .catch((err) => {
        message.error(err).then();
      })
      .finally(() => {
        loading.value = false;
      });
  };

  const refreshCaptcha = () => {
    getCaptcha()
      .then((res) => {
        const data = res.data;
        loginForm.captchaKey = data['key'];
        captchaImage.value = data['captcha'];
      })
      .catch(() => {
        message.error("获取验证码失败").then();
      });
  };
</script>

<style scoped>
  .flowing-login-bg {
    background-color: #0096ff;
    height: 100%;
    width: 100%;
    display: flex;
    justify-content: center;
  }
  .flowing-login-card {
    width: 800px;
    height: 500px;
    margin: auto;
    background-color: white;
    border-radius: 5px;
    display: flex;
    justify-content: flex-start;
  }

  .flowing-login-side {
    width: 50%;
    height: 100%;
    background-color: blue;
    border-top-left-radius: 5px;
    border-bottom-left-radius: 5px;
  }

  .flowing-login-form {
    width: 50%;
    height: 100%;
    background-color: transparent;
    padding: 20px 10px;
  }

  .flowing-login-button {
    width: 100%;
  }

  .flowing-captcha {
    display: flex;
    justify-content: flex-start;
    width: 100%;
  }

  .flowing-captcha-input {
    width: 60%;
  }

  .flowing-captcha-image {
    width: 40%;
  }
</style>
