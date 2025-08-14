import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router, { setupRouterGuard } from './router';

const app = createApp(App)
async function bootstrap() {
  app.use(createPinia());
  setupRouterGuard();
  app.use(router);
  // TODO 获取用户权限，构建菜单路由
  app.mount('#app');
}

await bootstrap()
