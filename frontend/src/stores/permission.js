import { defineStore } from 'pinia';
import { getUserAllPermissions } from '@/api/sys/permission.js';

export const usePermissionStore = defineStore('flowing_user_permission', {
  state: () => ({
    permissionLoaded: false,
    navMenus: [],
    actions: [],
  }),
  getters: {
    isPermissionLoaded: (state) => state.permissionLoaded,
  },
  actions: {
    setPermissionLoaded(value) {
      this.permissionLoaded = value;
    },
    async getUserPermissions() {
      try {
        const { menus } = await getUserAllPermissions();
        return menus;
      } catch (err) {
        console.error(err);
      }
      return null;
    },
    setNavMenus(menus) {
      const navMenus = [];
      menus.forEach((menu) => {
        const navMenu = {
          key: menu['path'],
          label: menu['menuName'],
          path: menu['path'],
          icon: menu['icon'],
          order: menu['orderNum'],
          children: [],
        };
        navMenus.push(navMenu);
        menu.children?.forEach((child) => {
          const navChild = {
            key: child['path'],
            label: child['menuName'],
            path: child['path'],
            icon: child['icon'],
            order: child['orderNum'],
          };
          navMenu.children.push(navChild);
        });
      });
      this.navMenus = navMenus;
    },
  },
});
