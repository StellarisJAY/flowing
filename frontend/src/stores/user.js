import { defineStore } from 'pinia';
import {GetUserMenus} from '@/api/menu.js';

export const userStore = defineStore('login_user', {
  state: () => ({
    menus: [],
    userInfo: {}
  }),
  actions: {
    loadMenus() {
      GetUserMenus().then((res) => {
        this.menus = res.data;
      })
    }
  }
});
