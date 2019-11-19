export interface IArticleData {
  id: number
  status: string
  title: string
  abstractContent: string
  fullContent: string
  sourceURL: string
  imageURL: string
  timestamp: string | number
  platforms: string[]
  disableComment: boolean
  importance: number
  author: string
  reviewer: string
  type: string
  pageviews: number
}

export interface IAdminUserData {
  id: number
  loginName: string
  realName: string
  phone: string
  email: string
  roleIds: string
  lastLogin: string
  lastIP: string[]
  status: number
  createTime: string
  updateTime: string
}

export interface IAdminRole {
  id: number
  roleName: string
  detail: string
  createTime?: string
  updateTime?: string
  nodesData?: string
  authIds?: number[]
}
