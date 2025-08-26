import axios from 'axios';
import { useUserStore } from '@/stores/user.js';

const http = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
  timeout: 10000,
});

http.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('flowing_access_token');
    if (token) {
      config.headers['X-Access-Token'] = token;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

http.interceptors.response.use(
  (response) => {
    const userStore = useUserStore();
    const res = response.data;
    const { code } = res;
    const msg = res.message;
    if (code === 200) {
      return response.data;
    }
    if (code === 401) {
      userStore.logout();
      window.location.href = '/sys/login';
      return;
    }
    if (code === 403) {
      return Promise.reject(msg || '无权限');
    }
    return Promise.reject(msg || 'Error');
  },
  (error) => {
    console.log(error);
  }
);

// fetchEventStream POST事件流请求
export const fetchEventStream = async (url, data, onOpen, onMessage, onError, onClose) => {
  const token = localStorage.getItem('flowing_access_token');
  const options = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'X-Access-Token': token, // token
    },
    body: JSON.stringify(data), // 请求数据JSON
  };
  try {
    const res = await fetch(import.meta.env.VITE_API_URL + url, options);
    onOpen();
    const reader = res.body.getReader();
    // 读取响应流
    while (true) {
      const { done, value } = await reader.read();
      if (done) {
        onClose();
        break;
      }
      const chunk = new TextDecoder().decode(value);
      chunk.split('\n\n').forEach((line) => {
        const idx = line.indexOf('data:');
        if (idx === -1) {
          return;
        }
        const message = line.substring(idx + 5).trim();
        onMessage(message);
      });
    }
  }catch (error) {
    onError(error);
    onClose();
  }
};

export default http;
