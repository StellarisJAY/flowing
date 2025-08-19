import http from '@/api/index.js';

export const listDocument = (query) => http.get('/kb/doc/list', {params: query});

export const uploadDocument = (formData) => http.postForm('/kb/doc/upload', formData);

export const getDownloadUrl = (id) => http.get(`/kb/doc/download`, {params: {id: id}});
