import Cookies from 'js-cookie'

// App
const sidebarStatusKey = 'sidebar_status'
export const getSidebarStatus = () => Cookies.get(sidebarStatusKey)
export const setSidebarStatus = (sidebarStatus: string) =>
  Cookies.set(sidebarStatusKey, sidebarStatus)

// User
const tokenKey = 'x_access_token'
export const getToken = () => Cookies.get(tokenKey)
export const setToken = (token: string) => Cookies.set(tokenKey, token)
export const removeToken = () => Cookies.remove(tokenKey)

const roleKey = 'roles'
export const getRoles = () => Cookies.get(roleKey) || ''
export const setRoles = (roles: string) => Cookies.set(roleKey, roles)
export const removeRoles = () => Cookies.remove(roleKey)
