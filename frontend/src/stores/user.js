import { defineStore } from 'pinia';

export const userStore = defineStore('login_user', {
  state: () => ({
    menus: [],
    userInfo: {},
  }),
  actions: {
    loadMenus() {
    },
    setToken(token) {
      localStorage.setItem("flowing_access_token", token);
    },
    getToken() {
      return localStorage.getItem("flowing_access_token");
    }
  },
});
