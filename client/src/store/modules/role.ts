import { VuexModule, Module, Action, Mutation, getModule } from 'vuex-module-decorators'
import { getRoles as APIGetRoles } from '@/api/role'
import { getToken, setToken, removeToken, getRoles, setRoles } from '@/utils/cookies'
import store from '@/store'
import router, { AuthRouters } from '@/router'
import { RouteConfig } from 'vue-router'

interface IRole {
  id: number
  title: string
}
export interface IRoleState {
  token: string
  name: string
  avatar: string
  introduction: string
  roles: IRole[]
}

@Module({ dynamic: true, store, name: 'role' })
class Role extends VuexModule implements IRoleState {
  public token = getToken() || ''
  public name = ''
  public avatar = ''
  public introduction = ''
  public roles: IRole[] = []

  @Mutation
  private SET_ROLES(token: string) {
    this.token = token
  }

  @Mutation
  private SET_NAME(name: string) {
    this.name = name
  }

  @Mutation
  private SET_AVATAR(avatar: string) {
    this.avatar = avatar
  }

  @Mutation
  private SET_INTRODUCTION(introduction: string) {
    this.introduction = introduction
  }

  @Action
  public async LogOut() {
    if (this.token === '') {
      throw Error('LogOut: token is undefined!')
    }
  }
}

function GetNodeList(list: any[], first: boolean) {
  if (!Array.isArray(list)) {
    return []
  }
  let arr: any[] = []

  list.forEach((v: any) => {
    let router: RouteConfig = {
      path: v.path,
      meta: {
        title: v.title,
        icon: v.icon
      }
    }
    v.redirect && (router.redirect = v.redirect)
    v.children && (router.children = GetNodeList(v.children, false))
    // const subList = GetNodeList(v.children || [], false)
    // if (subList.length > 0) {
    //   router.children = subList
    // }
    arr.push(router)
  })
  return arr
}

export const UserModule = getModule(Role)
