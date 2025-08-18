import http from '@/api/index.js';

export const createProvider = (data) => http.post('/ai/provider/create', data);

export const listProvider = (query) => http.get('/ai/provider/list', {params: query});

export const updateProvider = (data) => http.put('/ai/provider/update', data);
