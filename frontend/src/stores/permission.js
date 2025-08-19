import { defineStore } from 'pinia';
import { getUserAllPermissions } from '@/api/sys/permission.js';
import { h } from 'vue';
import Icon from '@/components/Icon/Icon.vue';

export const usePermissionStore = defineStore('flowing_user_permission', {
  state: () => ({
    permissionLoaded: false,
    navMenus: [],
    actions: [],
  }),
  actions: {
    isPermissionLoaded() {
      return this.permissionLoaded;
    },
    setPermissionLoaded(value) {
      this.permissionLoaded = value;
    },
    async getUserPermissions() {
      try {
        const {data} = await getUserAllPermissions();
        return data.menus;
      } catch (err) {
        console.error(err);
      }
      return null;
    },
    setNavMenus(menus) {
      const navMenus = [];
      console.log(menus);
      menus.forEach((menu) => {
        if (menu.showInNav === false) return;
        const navMenu = {
          key: menu['path'],
          label: menu['menuName'],
          path: menu['path'],
          icon: ()=>h(Icon, {icon: menu['icon']}),
          order: menu['orderNum'],
          children: [],
        };
        navMenus.push(navMenu);
        menu.children?.forEach((child) => {
          if (child.showInNav === false) return;
          const navChild = {
            key: child['path'],
            label: child['menuName'],
            path: child['path'],
            icon: ()=>h(Icon, {icon: child['icon']}),
            order: child['orderNum'],
          };
          navMenu.children.push(navChild);
        });
      });
      this.navMenus = navMenus;
    },
  },
});
