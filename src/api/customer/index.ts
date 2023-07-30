import request from '@/config/axios'
import type { CustomerTableData } from './types'

export const getCustomerListApi = (params: any): Promise<IResponse<CustomerTableData>> => {
  return request.get({ url: '/customer/list', params })
}

export const createCustomerApi = (data: Partial<CustomerTableData>): Promise<IResponse> => {
  return request.post({ url: '/customer/save', data })
}

export const updateCustomerApi = (data: Partial<CustomerTableData>): Promise<IResponse> => {
  return request.put({ url: '/customer/update', data })
}
export const delTableListApi = (ids: string[] | number[]): Promise<IResponse> => {
  return request.post({ url: '/example/delete', data: { ids } })
}
