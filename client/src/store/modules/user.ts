import { VuexModule, Module, Action, Mutation, getModule } from 'vuex-module-decorators'
import { login, logout, getUserInfo } from '@/api/users'
import { getToken, setToken, removeToken, getRoles, setRoles } from '@/utils/cookies'
import store from '@/store'
import router, { AuthRouters } from '@/router'
import { authComponents } from '@/router/authRouter'
import { RouteConfig } from 'vue-router'

export interface IUserState {
  token: string
  name: string
  avatar: string
  introduction: string
  roles: string[]
}

@Module({ dynamic: true, store, name: 'user' })
class User extends VuexModule implements IUserState {
  public token = getToken() || ''
  public name = ''
  public avatar = require('@/icons/avatar.gif')
  public introduction = ''
  public roles: string[] = []

  @Mutation
  private SET_TOKEN(token: string) {
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

  @Mutation
  private SET_ROLES(roles: string[]) {
    this.roles = roles
  }

  @Action
  public async Login(userInfo: { username: string; password: string }) {
    try {
      let { username, password } = userInfo
      username = username.trim()
      const { data } = await login({ username, password })
      setToken(data.auth)
      this.SET_TOKEN(data.auth)

      this.SET_ROLES(data.roles)
      setRoles(data.roles.split(','))
      this.SET_NAME(data.loginName)
      await this.GetUserInfo()
      // this.SET_AVATAR(avatar)
      // this.SET_INTRODUCTION(introduction)
    } catch (error) {
      //
    }
  }

  @Action
  public ResetToken() {
    removeToken()
    this.SET_TOKEN('')
    this.SET_ROLES([])
  }

  @Action
  public async GetUserInfo() {
    this.SET_ROLES(getRoles().split(','))

    // if (this.token === '') {
    //   throw Error('GetUserInfo: token is undefined!')
    // }
    const { data } = await getUserInfo({
      /* Your params here */
    })
    if (!data) {
      throw Error('Verification failed, please Login again.')
    }

    const { list = [] } = data
    try {
      let newAuthRouters = [AuthRouters[0], ...GetNodeList(list, true)]
      let constRoute: any = (router as any).options.routes
      ;(router as any).options.routes = { ...constRoute, ...newAuthRouters }
      console.error(newAuthRouters)
      router.addRoutes(newAuthRouters)
    } catch (error) {
      console.error(error)
    }

    // const { roles, name, avatar, introduction } = data.user
    // // roles must be a non-empty array
    // if (!roles || roles.length <= 0) {
    //   throw Error('GetUserInfo: roles must be a non-null array!')
    // }
    // this.SET_ROLES(roles)
    // this.SET_NAME(name)
    // this.SET_AVATAR(avatar)
    // this.SET_INTRODUCTION(introduction)
  }

  @Action
  public async LogOut() {
    if (this.token === '') {
      throw Error('LogOut: token is undefined!')
    }
    await logout()
    removeToken()
    this.SET_TOKEN('')
    this.SET_ROLES([])
  }
}

function GetNodeList(list: any[], first: boolean) {
  if (!Array.isArray(list)) {
    return []
  }
  let arr: any[] = []

  list.forEach((v: any) => {
    try {
      let router: RouteConfig = {
        path: v.path,
        component: first ? AuthRouters[0].component : v.component && authComponents[v.component].component,
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
    } catch (error) {}
  })
  return arr
}

export const UserModule = getModule(User)
