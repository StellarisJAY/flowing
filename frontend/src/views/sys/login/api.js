import http from '@/api/index.js';

export const login = (data) => http.post('/login', data);

export const getCaptcha = () => http.get('/captcha');
