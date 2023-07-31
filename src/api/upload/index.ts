import request from '@/config/axios'
import type { TableData } from './types'

export const uploadImgApi = (params: any): Promise<IResponse> => {
  return request.post({
    url: '/common/file/upload',
    params,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}
