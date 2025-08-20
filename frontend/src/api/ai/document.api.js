import http from '@/api/index.js';

export const listDocument = (query) => http.get('/kb/doc/list', { params: query });

export const uploadDocument = (formData) => http.postForm('/kb/doc/upload', formData);

export const getDownloadUrl = (id) => http.get(`/kb/doc/download`, { params: { id: id } });

export const renameDocument = (data) => http.put('/kb/doc/rename', data);

export const deleteDocument = (id) => http.delete('/kb/doc/delete', { params: { id: id } });

export const parseDocument = (data) => http.post('/kb/doc/parse', data);

export const cancelParse = (id) => http.post(`/kb/doc/cancel?id=${id}`);
