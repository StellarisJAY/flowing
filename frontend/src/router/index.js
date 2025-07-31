import { createRouter, createWebHistory } from 'vue-router';
import { usePermission } from '@/stores/permission.js';

const modules = import.meta.glob('@/**/**/*.vue');

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/sys',
      component: () => import('@/layouts/page/index.vue'),
      children: [
        {
          name: '登录',
          path: 'login',
          component: () => import('@/views/sys/login/index.vue'),
        },
      ],
    },
  ],
});

const buildRoutes = (menus) => {
  console.log(modules)
  menus.forEach((menu) => {
    const item = {
      path: menu.path,
      name: menu.menuName,
      meta: {
        title: menu.menuName,
        icon: menu.icon,
      },
      children: []
    };
    item.component = modules[`${menu.component}`];
    menu.children.forEach((child)=>{
      const childItem = {
        path: child.path,
        name: child.menuName,
        meta: {
          title: child.menuName,
          icon: child.icon,
        }
      };
      childItem.component = modules[`/src${child.component}`];
      item.children.push(childItem);
    })
    router.addRoute(item);
  });
  console.log(router.getRoutes());
};

export const setupRouterGuard = () => {
  router.beforeEach(async (to, from, next) => {
    const permissionStore = usePermission();
    if (to.path === '/sys/login') {
      next();
      return;
    }
    if (permissionStore.isPermissionLoaded) {
      next();
      return;
    }
    const menus = await permissionStore.getUserPermissions();
    buildRoutes(menus);
    next();
  });
};

export default router;
