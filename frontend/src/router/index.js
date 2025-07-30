import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/system',
      component: () => import('@/layouts/index.vue'),
      meta: {hideTab: true},
      children: [
        {
          path: "user",
          component: ()=>import('@/views/system/user/index.vue')
        }
      ]
    }
  ],
})

export default router
