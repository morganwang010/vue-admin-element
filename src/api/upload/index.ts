import request from '@/config/axios'
import type { TableData } from './types'

export const uploadImageApi = (formData: any): Promise<IResponse> => {
  console.log('hhhhhhhhhhhhhhhh')
  return request.post({
    url: '/common/file/upload',
    data: formData,
    headersType: 'multipart/form-data'
    // headers: {
    //   'Content-Type': 'multipart/form-data'
    // }
  })
}
