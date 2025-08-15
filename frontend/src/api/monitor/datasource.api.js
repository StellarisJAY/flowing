import http from '@/api/index.js';

export const queryDatasourceList = (query) => http.get('/monitor/datasource/list', {params: query});

export const createDatasource = (data) => http.post('/monitor/datasource/create', data);

export const updateDatasource = (data) => http.put('/monitor/datasource/update', data);

export const deleteDatasource = (query) => http.delete('/monitor/datasource/delete', {params: query});

export const pingDatasource = (data) => http.post('/monitor/datasource/ping', data);
