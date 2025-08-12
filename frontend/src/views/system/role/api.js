import http from '@/api/index.js';

export const queryRoleList = (query) => {
  return http.get('/role/list', {params: query});
}

export const addRole = (data) => {
  return http.post('/role/create', data);
}

export const updateRole = (data) => {
  return http.put('/role/update', data);
}

export const getRoleMenus = (query) => {
  return http.get('/role/menus', {params: query});
}

export const saveRoleMenus = (data) => {
  return http.post('/role/menus', data);
}
