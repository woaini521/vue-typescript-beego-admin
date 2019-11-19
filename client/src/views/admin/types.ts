export interface IAdminEditForm {
  id: number
  loginName: string
  realName: string
  phone: string
  email: string
  roleIds: number[]
}

export const AdminEditDefaultForm: IAdminEditForm = {
  id: -1,
  loginName: '',
  realName: '',
  phone: '',
  email: '',
  roleIds: []
}

export default {}
