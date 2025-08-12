import http from '../index.js';

export const getDictItemsByCode = (params) => http.get(`/dict/item/code`, { params: params });

export const getDictItems = (params) => http.get(`/dict/item/list`, { params: params });

export const listDict = (params) => http.get('/dict/list', { params: params });

export const createDict = (data) => http.post('/dict/create', data);

export const updateDict = (data) => http.put('/dict/update', data);

export const createDictItem = (data) => http.post('/dict/item/create', data);

export const updateDictItem = (data) => http.put('/dict/item/update', data);
