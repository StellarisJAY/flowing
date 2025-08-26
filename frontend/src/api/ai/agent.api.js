import http from '@/api/index.js';

export const queryAgentList = (query) => http.get('/agent/list', { params: query });

export const createAgent = (data) => http.post('/agent/create', data);

export const updateAgent = (data) => http.put('/agent/update', data);

export const saveConfig = (data) => http.put('/agent/config', data);

export const getDetail = (id) => http.get(`/agent/detail`, { params: { id: id } });
