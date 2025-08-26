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
  "/views/agent/knowledge/document/index": ()=> import("@/views/agent/knowledge/document/index.vue"),
  "/views/agent/knowledge/detail/index": ()=> import("@/views/agent/knowledge/detail/index.vue"),
  "/views/agent/knowledge/retrieve/index": ()=> import("@/views/agent/knowledge/retrieve/index.vue"),
  "/views/agent/knowledge/document/chunks/index": ()=> import("@/views/agent/knowledge/document/chunks/index.vue"),
  "/views/agent/chat/index": ()=> import("@/views/agent/chat/index.vue"),
  "/views/agent/agent/index": ()=> import("@/views/agent/agent/index.vue"),
  "/views/agent/agent/configuration/index": ()=> import("@/views/agent/agent/configuration/index.vue"),
  "/views/agent/agent/configuration/chat/index": ()=> import("@/views/agent/agent/configuration/chat/index.vue"),
  "/views/agent/agent/configuration/workflow/index": ()=> import("@/views/agent/agent/configuration/workflow/index.vue"),
  // "/views/agent/knowledge/configuration/index": ()=> import("@/views/agent/knowledge/configuration/index.vue"),
};
