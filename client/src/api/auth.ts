/*
 * @Author: Sy.
 * @Create: 2019-11-01 20:54:15
 * @LastTime: 2019-11-18 22:12:08
 * @LastEdit: Sy.
 * @FilePath: \smallyan.admin\client\src\api\users.ts
 * @Description: 用户
 */
import request from '@/utils/request'

export const getNodes = () =>
  request({
    url: '/auth/list',
    method: 'get'
  })

export const editNodes = (data: any) =>
  request({
    url: '/auth/edit',
    method: 'post',
    data,
    showSuccess: 'toast'
  })

export const delNode = (id: number) =>
  request({
    url: '/auth/del',
    method: 'get',
    data: { id }
  })
