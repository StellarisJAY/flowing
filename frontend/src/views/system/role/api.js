import http from '@/api/index.js';

export const queryRoleList = (query) => {
  return http.get('/role/list', {params: query});
}

export const addRole = (data) => {
  return http.post('/role/create', data);
}
