import axios, {
  AxiosInstance,
  InternalAxiosRequestConfig,
  AxiosRequestHeaders,
  AxiosResponse,
  AxiosError
} from 'axios'

import qs from 'qs'

import { config } from './config'

import { ElMessage } from 'element-plus'

const { result_code, base_url } = config

export const PATH_URL = base_url[import.meta.env.VITE_API_BASE_URL]
axios.defaults.baseURL = import.meta.env.VITE_API_BASE_URL
console.log(import.meta.env.VITE_API_BASE_URL)
// 创建axios实例
const service: AxiosInstance = axios.create({
  baseURL: PATH_URL, // api 的 base_url
  timeout: config.request_timeout // 请求超时时间
})

// request拦截器
service.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    config.headers['uid'] = localStorage.getItem('uid')
    config.headers['token'] = localStorage.getItem('token')

    if (
      config.method === 'post' &&
      (config.headers as AxiosRequestHeaders)['Content-Type'] ===
        'application/x-www-form-urlencoded'
    ) {
      config.data = qs.stringify(config.data)
    }
    if (config.method === 'put') {
      console.log('put update the form data')
      console.log(config.data)
      config.data = config.data
    }
    if (
      config.method === 'post' &&
      (config.headers as AxiosRequestHeaders)['Content-Type'] === 'multipart/form-data'
    ) {
      console.log('kkkkkkkkkkkkkk')
      config.data = config.data
    }
    if (config.method === 'get' && config.params) {
      let url = config.url as string
      console.log(config.baseURL)
      url += '?'
      const keys = Object.keys(config.params)
      for (const key of keys) {
        if (config.params[key] !== void 0 && config.params[key] !== null) {
          url += `${key}=${encodeURIComponent(config.params[key])}&`
        }
      }
      url = url.substring(0, url.length - 1)
      config.params = {}
      config.url = url
    }
    return config
  },
  (error: AxiosError) => {
    // Do something with request error
    console.log(error) // for debug
    Promise.reject(error)
  }
)

// response 拦截器
service.interceptors.response.use(
  (response: AxiosResponse<any>) => {
    if (response.config.responseType === 'blob') {
      // 如果是文件流，直接过
      return response
    } else if (Number(response.data.code) === Number(result_code)) {
      return response.data
    } else {
      ElMessage.error(response.data.message)
    }
  },
  (error: AxiosError) => {
    console.log('err' + error) // for debug
    ElMessage.error(error.message)
    return Promise.reject(error)
  }
)

export { service }
