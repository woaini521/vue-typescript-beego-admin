/*
 * @Author: Sy.
 * @Create: 2019-11-07 22:58:31
 * @LastTime: 2019-11-18 22:10:26
 * @LastEdit: Sy.
 * @FilePath: \smallyan.admin\client\src\api\role.ts
 * @Description: 角色
 */
import request from '@/utils/request'
import { IAdminRole } from './types'

export const getRoles = (data: any = {}) =>
  request({
    url: '/role/list',
    method: 'post',
    data
  })

export const getRoleAuth = (id: number = 0) =>
  request({
    url: '/role/info',
    method: 'post',
    data: { id }
  })

export const delRole = (id: number) =>
  request({
    url: '/role/del',
    method: 'post',
    data: { id },
    showSuccess: 'toast'
  })

export const editRole = (data: IAdminRole) =>
  request({
    url: '/role/edit',
    method: 'post',
    data,
    showSuccess: 'toast'
  })
