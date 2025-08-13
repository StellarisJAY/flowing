import http from "@/api/index.js";

export const queryUserList = (params)=>http.get("/user/list", {params: params});

export const addUser = (data) => http.post('/user/add', data);

export const updateUser = (data) => http.put('/user/update', data);

export const deleteUser = (data) => http.delete('/user/delete', {params: data});

export const getUserDetail = (query) => http.get('/user/detail', {params: query});
