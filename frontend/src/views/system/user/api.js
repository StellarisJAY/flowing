import http from "@/api/index.js";

export const queryUserList = (params)=>http.get("/user/list", {params: params});
