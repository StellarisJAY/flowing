import { createRouter, createWebHistory } from 'vue-router';
import { usePermissionStore } from '@/stores/permission.js';
import { LayoutMap, ViewMap } from '@/router/routes.js';
import { useUserStore } from '@/stores/user.js';

const router = createRouter({
  history: createWebHistory(),
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
    {
      path: '/',
      redirect: '/home',
    },
  ],
  strict: true,
  sensitive: true,
});

const buildRoutes = (menus) => {
  menus.forEach((menu) => {
    const item = {
      path: menu.path,
      name: menu['menuName'],
      meta: {
        title: menu['menuName'],
        icon: menu.icon,
      },
      children: [],
    };
    item.component = LayoutMap[menu['component']];
    menu.children?.forEach((child) => {
      const childItem = {
        name: child['menuName'],
      };
      childItem.path = child.path.replace(item.path, '').slice(1);
      childItem.component = ViewMap[child.component];
      item.children.push(childItem);
    });
    router.addRoute(item);
  });
};

export const setupRouterGuard = () => {
  router.beforeEach(async (to, from, next) => {
    const permissionStore = usePermissionStore();
    const userStore = useUserStore();
    const token = userStore.getToken();
    if (!token) {
      if (to.path === '/sys/login') {
        next();
        return;
      }
      next({ path: '/sys/login', replace: true });
      return;
    }
    if (permissionStore.isPermissionLoaded) {
      userStore.changeTabPanesOnRouting(to, false);
      next();
      return;
    }
    const menus = await permissionStore.getUserPermissions();
    if (menus) {
      buildRoutes(menus);
      permissionStore.setPermissionLoaded(true);
      permissionStore.setNavMenus(menus);
      // 重新加载路由后，需要重新导航到当前路由，否则会出现404
      next({ ...to, replace: true });
    } else {
      next({ path: '/sys/login', replace: true });
    }
  });
};

export default router;
