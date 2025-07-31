import http from '../index.js';


export const getUserMenus = ()=>http.get('/sys/menus');

export const getUserAllPermissions = ()=>http.get('/sys/permissions');
