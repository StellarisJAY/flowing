import http from '@/api/index.js';

export const queryKnowledgeList = (params) => http.get('/kb/list', { params: params });
export const addKnowledge = (data) => http.post('/kb/create', data);
export const updateKnowledge = (data) => http.put('/kb/update', data);
export const deleteKnowledge = (id) => http.delete('/kb/delete', { params: { id: id } });
export const searchKnowledge = (data) => http.post('/kb/search', data);
