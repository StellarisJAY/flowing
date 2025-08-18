export const LayoutMap = {
  "/layouts/default/index": ()=> import("@/layouts/default/index.vue"),
  "/layouts/page/index": ()=> import("@/layouts/page/index.vue"),
};

export const ViewMap = {
  "/views/monitor/performance/index": ()=> import("@/views/monitor/performance/index.vue"),
  "/views/system/user/index": ()=> import("@/views/system/user/index.vue"),
  "/views/system/menu/index": ()=> import("@/views/system/menu/index.vue"),
  "/views/system/role/index": ()=> import("@/views/system/role/index.vue"),
  "/views/system/dict/index": ()=> import("@/views/system/dict/index.vue"),
  "/views/agent/knowledge/index": () => import("@/views/agent/knowledge/index.vue"),
  "/views/monitor/datasource/index": ()=> import("@/views/monitor/datasource/index.vue"),
  "/views/agent/model/index": ()=>import("@/views/agent/model/index.vue"),
};
