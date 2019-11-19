import axios, { AxiosRequestConfig, AxiosInstance, AxiosPromise } from 'axios'
import qs from 'qs'
import { Message, MessageBox, Loading } from 'element-ui'
import { UserModule } from '@/store/modules/user'
import { ElLoadingComponent } from 'element-ui/types/loading'

const service = axios.create({
  // baseURL: process.env.VUE_APP_BASE_API,
  baseURL: '/api',
  timeout: 5000
})

// Request interceptors
service.interceptors.request.use(
  config => {
    config.headers['Content-Type'] = 'application/json;charset=UTF-8'
    // Add X-Access-Token header to every request, you can add other custom headers here
    if (config.method === 'post') {
      config.headers['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8'
      config.data = qs.stringify(config.data)
    }
    if (UserModule.token) {
      // config.params = { ...config.data, authAPI: UserModule.token }
      config.headers['X-Access-Token'] = UserModule.token
    }
    config.headers['Accept'] = 'application/json;charset=utf-8'
    return config
  },
  error => {
    Promise.reject(error)
  }
)

// Response interceptors
service.interceptors.response.use(
  response => {
    // Some example codes here:
    // code == 20000: success
    // code == 50001: invalid access token
    // code == 50002: already login in other place
    // code == 50003: access token expired
    // code == 50004: invalid user (user not exist)
    // code == 50005: username or password is incorrect
    // You can change this part for your own usage.
    const res = response.data
    if (res.code !== 0) {
      Message({
        message: res.message || 'Requst Error',
        type: 'error',
        duration: 5 * 1000
      })
      if (res.code === 5000) {
        MessageBox.confirm('你已被登出，可以取消继续留在该页面，或者重新登录', '确定登出', {
          confirmButtonText: '重新登录',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          UserModule.ResetToken()
          location.reload() // To prevent bugs from vue-router
        })
      }
      return Promise.reject(new Error(res.message || 'Requst Error'))
    } else {
      return response.data
    }
  },
  error => {
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)

interface AxiosExtends extends AxiosRequestConfig {
  hideLoading?: boolean
  showSuccess?: 'toast' | 'modal'
}
const request = (options: AxiosExtends): AxiosPromise<any> => {
  return new Promise(async(resolve, reject) => {
    // let loadingInstance: ElLoadingComponent | null = null
    // if (!options.hideLoading) {
    //   loadingInstance = Loading.service({
    //     lock: true,
    //     text: 'Loading...',
    //     spinner: 'el-icon-loading',
    //     background: 'rgba(0, 0, 0, 0.7)'
    //   })
    // }
    let res: any, err: any
    try {
      res = await service(options)
    } catch (error) {
      res = error
      err = error
    }

    // if (!options.hideLoading) {
    //   setTimeout(() => {
    //     loadingInstance && loadingInstance.close()
    //   }, 1.5e3)
    // }
    const { code = -1 } = res
    if (options.showSuccess && code === 0) {
      options.showSuccess === 'modal' &&
        MessageBox.alert('操作成功', '提示', {
          showCancelButton: false
        })
      options.showSuccess === 'toast' && Message.success('操作成功')
    }
    if (err) {
      return reject(err)
    }
    res && resolve(res)
  })
}
export default request
