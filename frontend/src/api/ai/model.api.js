import http from '@/api/index.js';

export const listModel = (query) => http.get('/ai/model/list', {params: query});

export const createModel = (data) => http.post('/ai/model/create', data);

export const updateModel = (data) => http.put('/ai/model/update', data);

export const deleteModel = (id) => http.delete('/ai/model/delete', {params: {id: id}});
