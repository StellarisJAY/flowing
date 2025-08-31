import { defineStore } from 'pinia';

export const useWorkflowStore = defineStore('workflow', {
  state: () => ({
    draft: {
      nodes: [
        {
          id: '1',
          type: 'base',
          position: { x: 400, y: 200 },
          data: {
            label: 'Special Node',
            hello: 'world',
          },
        },
      ],
      edges: [],
    },
  }),
  actions: {
  }
});
