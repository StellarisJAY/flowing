import { defineStore } from 'pinia';

export const useGlobalStore = defineStore('flowing_global', {
  state: () => ({
    loading: false,
  }),
  actions: {
    setLoading(loading) {
      this.loading = loading;
    },
  }
});
