import http from "@/api/index.js";

export const queryMenuTree = (params)=>http.get("/menu/list", {params: params});

export const createMenu = (data)=>http.post("/menu/create", data)
