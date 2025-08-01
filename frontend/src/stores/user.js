import { defineStore } from 'pinia';

export const useUserStore = defineStore('login_user', {
  state: () => ({
    userInfo: {},
    tabPanes: [],
    activeTab: '',
  }),
  actions: {
    setToken(token) {
      localStorage.setItem('flowing_access_token', token);
    },
    getToken() {
      return localStorage.getItem('flowing_access_token');
    },
    changeTabPanesOnRouting(to, del) {
      const tabPanes = this.tabPanes;
      const index = tabPanes.findIndex((item) => item.key === to.path);
      if (del === true && index === -1) return;
      if (index !== -1) {
        if (del) {
          tabPanes.splice(index, 1);
          this.activeTab = tabPanes[0].key;
          return;
        }
        this.activeTab = to.path;
        return;
      }
      tabPanes.push({
        key: to.path,
        title: to.name,
        closable: true,
      });
      this.activeTab = to.path;
    },
  },
});
