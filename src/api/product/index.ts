import request from '@/config/axios'
import type { ProductTableData } from './types'

export const getProductListApi = (params: any): Promise<IResponse<ProductTableData>> => {
  return request.get({ url: '/product/list', params })
}

export const createProductApi = (data: Partial<ProductTableData>): Promise<IResponse> => {
  return request.post({ url: '/product/save', data })
}

export const updateProductApi = (data: Partial<ProductTableData>): Promise<IResponse> => {
  console.log(data.offdate)
  return request.put({ url: '/product/update', data })
}
export const delTableListApi = (ids: string[] | number[]): Promise<IResponse> => {
  return request.post({ url: '/product/delete', data: { ids } })
}
