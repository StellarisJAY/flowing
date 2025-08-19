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
    {
      path: '/error',
      component: () => import('@/layouts/default/index.vue'),
      children: [
        {
          name: '页面不存在',
          path: '404',
          component: () => import('@/views/sys/404/index.vue'),
        },
      ],
    },
  ],
  strict: true,
  sensitive: true,
});

const buildRoutes = (menus) => {
  const childBuilder = (parent, child) => {
    const childItem = {
      name: child['menuName'],
      children: [],
      meta: {
        showInNav: child.showInNav,
        hideTab: child.hideTab,
      },
    };
    childItem.path = child.path;
    childItem.component = ViewMap[child.component];
    parent.children.push(childItem);
    child.children?.forEach((grandchild) => {
      childBuilder(childItem, grandchild);
    });
  };
  menus.forEach((menu) => {
    const item = {
      path: menu.path,
      name: menu['menuName'],
      meta: {
        title: menu['menuName'],
        icon: menu.icon,
        showInNav: menu.showInNav,
        hideTab: menu.hideTab,
      },
      children: [],
    };
    item.component = LayoutMap[menu['component']];
    menu.children?.forEach((child) => {
      childBuilder(item, child);
    });
    router.addRoute(item);
  });
};

export const setupRouterGuard = () => {
  router.beforeEach(async (to, from) => {
    const permissionStore = usePermissionStore();
    const userStore = useUserStore();
    const token = userStore.getToken();
    // 没有token，跳转到登录页
    if (!token) {
      return { path: '/sys/login' };
    }
    // 权限已加载，直接判断路由中是否有目标路由
    if (permissionStore.isPermissionLoaded()) {
      // 判断路由中是否有目标路由
      if (router.hasRoute(to.name)) {
        // 添加tab标签
        userStore.changeTabPanesOnRouting(to);
        return;
      }
      // 没有目标路由，跳转到404
      return { path: '/error/404', replace: true };
    }
    // 加载用户菜单权限
    const menus = await permissionStore.getUserPermissions();
    if (menus) {
      // 构建路由
      buildRoutes(menus);
      permissionStore.setPermissionLoaded(true);
      permissionStore.setNavMenus(menus);
      return to;
    } else {
      // 没有菜单权限，跳转到登录页
      return { path: '/sys/login', replace: true };
    }
  });
};

export default router;
