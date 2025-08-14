import http from '../index.js';

export const getSystemPerformance = () => http.get('/monitor/performance');
