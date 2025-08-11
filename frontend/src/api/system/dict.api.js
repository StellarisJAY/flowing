import http from '../index.js';

export const getDictItemsByCode = (params) => http.get(`/dict/item/code`, {params: params});

export const getDictItems = (params) => http.get(`/dict/item/list`, {params: params});

export const listDict = (params) => http.get('/dict/list', {params: params});
