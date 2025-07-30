import http from '../index.js';


export const getUserMenus = ()=>http.get('/sys/menus');
