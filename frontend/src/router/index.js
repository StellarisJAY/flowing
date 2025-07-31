import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/sys',
      component: () => import('@/layouts/page/index.vue'),
      children: [
        {
          path: "login",
          component: ()=>import('@/views/sys/login/index.vue')
        }
      ]
    },
    {
      path: '/system',
      component: () => import('@/layouts/default/index.vue'),
      meta: {hideTab: true},
      children: [
        {
          path: "user",
          component: ()=>import('@/views/system/user/index.vue')
        },
        {
          path: "menu",
          component: ()=>import('@/views/system/user/index.vue')
        }
      ]
    }
  ],
});

export const setupRouterGuard = ()=>{

}

export default router
