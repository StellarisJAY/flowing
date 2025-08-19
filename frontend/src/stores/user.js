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
    changeTabPanesOnRouting(to) {
      if (to.meta.hideTab) {
        return;
      }
      const tabPanes = this.tabPanes;
      const index = tabPanes.findIndex((item) => item.key === to.path);
      if (index !== -1) {
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
    deleteTabPane(key) {
      const tabPanes = this.tabPanes;
      const index = tabPanes.findIndex((item) => item.key === key);
      if (index !== -1) {
        tabPanes.splice(index, 1);
        if (this.activeTab === key) {
          this.activeTab = tabPanes[0].key;
        }
      }
    },
  },
});
