import {defineStore} from 'pinia';
import {getUserAllPermissions} from '@/api/sys/permission.js';

export const usePermission = defineStore("flowing_user_permission", {
  state: () => ({
    isPermissionLoaded: false,
    menus: [],
    actions: [],
  }),
  getters: {
    isPermissionLoaded: (state) => state.isPermissionLoaded,
    menus: (state) => state.menus,
  },
  actions: {
    async getUserPermissions() {
      try {
        const { menus, actions } = await getUserAllPermissions();
        Reflect.set(this, "isPermissionLoaded", true);
        Reflect.set(this, "menus", menus);
        Reflect.set(this, "actions", actions);
        return menus;
      }catch (err) {
        console.error(err);
      }
    }
  }
})
