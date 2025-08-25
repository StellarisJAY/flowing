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
    setUserInfo(userInfo) {
      this.userInfo = userInfo;
      localStorage.setItem('flowing_user_info', JSON.stringify(userInfo));
    },
    getUserInfo() {
      if (this.userInfo.id) {
        return this.userInfo;
      }
      const userInfo = JSON.parse(localStorage.getItem('flowing_user_info'));
      if (userInfo) {
        return userInfo;
      }
      return null;
    },
    getUserId() {
      const user = this.getUserInfo();
      if (user) {
        return user.id;
      }
      return null;
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
    logout() {
      this.userInfo = {};
      this.activeTab = '';
      localStorage.clear();
    },
  },
});
