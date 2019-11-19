import router, { AuthRouters } from './router'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import { Message } from 'element-ui'
import { Route } from 'vue-router'
import { UserModule } from '@/store/modules/user'

NProgress.configure({ showSpinner: false })

const whiteList = ['/login']

router.beforeEach(async(to: Route, _: Route, next: any) => {
  // Start progress bar
  NProgress.start()

  // Determine whether the user has logged in
  if (UserModule.token) {
    if (to.path === '/login') {
      // If is logged in, redirect to the home page
      next({ path: '/' })
      NProgress.done()
    } else {
      // const routesLen: number = (router as any).options.routes.length || 0
      // if (routesLen === 3) {
      //   try {
      //     // Get user info, including roles
      //     await UserModule.GetUserInfo()
      //     // Set the replace: true, so the navigation will not leave a history record
      //     next({ ...to, replace: true })
      //   } catch (err) {
      //     // Remove token and redirect to login page
      //     UserModule.ResetToken()
      //     Message.error(err || 'Has Error')
      //     next(`/login?redirect=${to.path}`)
      //     NProgress.done()
      //   }
      //   // AuthRouters.splice(1, 2)
      //   // AuthRouters.splice(1, 0, {
      //   //   path: '/',
      //   //   component: AuthRouters[0].component,
      //   //   redirect: '/auth/tree',
      //   //   meta: {
      //   //     title: '权限管理',
      //   //     icon: 'example'
      //   //   },
      //   //   children: [
      //   //     {
      //   //       path: '/tree',
      //   //       component: () =>
      //   //         import(/* webpackChunkName: "tree" */ '@/views/auth/index.vue'),
      //   //       meta: {
      //   //         title: '树',
      //   //         icon: 'tree'
      //   //       }
      //   //     },
      //   //     {
      //   //       path: '/auth',
      //   //       component: () =>
      //   //         import(
      //   //           /* webpackChunkName: "table" */ '@/views/table/index.vue'
      //   //         ),
      //   //       meta: {
      //   //         title: '权限',
      //   //         icon: 'table'
      //   //       }
      //   //     }
      //   //   ]
      //   // })
      //   // let constRoute: any = (router as any).options.routes;
      //   // (router as any).options.routes = { ...constRoute, ...AuthRouters }
      //   // router.addRoutes(AuthRouters)
      // }
      // Check whether the user has obtained his permission roles
      if (UserModule.roles.length === 0) {
        try {
          // Get user info, including roles
          await UserModule.GetUserInfo()
          console.error({ ...to })
          // Set the replace: true, so the navigation will not leave a history record
          next({ ...to, path: to.redirectedFrom, replace: true })
        } catch (err) {
          // Remove token and redirect to login page
          UserModule.ResetToken()
          Message.error(err || 'Has Error')
          next(`/login?redirect=${to.path}`)
          NProgress.done()
        }
      } else {
        next()
      }
    }
  } else {
    // Has no token
    if (whiteList.indexOf(to.path) !== -1) {
      // In the free login whitelist, go directly
      next()
    } else {
      console.error(to.path)
      // Other pages that do not have permission to access are redirected to the login page.
      let redirect = to.path
      redirect === '/404' && (redirect = '/')
      next(`/login?redirect=${redirect}`)
      NProgress.done()
    }
  }
})

router.afterEach((to: Route) => {
  // Finish progress bar
  NProgress.done()

  // set page title
  to.meta.title && (document.title = to.meta.title)
})
