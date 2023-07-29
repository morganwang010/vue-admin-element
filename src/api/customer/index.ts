import request from '@/config/axios'
import type { CustomerTableData } from './types'

export const getCustomerListApi = (params: any): Promise<IResponse<CustomerTableData>> => {
  return request.get({ url: '/customer/list', params })
}

export const saveTableApi = (data: Partial<CustomerTableData>): Promise<IResponse> => {
  return request.post({ url: '/example/save', data })
}

export const getTableDetApi = (id: string): Promise<IResponse<CustomerTableData>> => {
  return request.get({ url: '/example/detail', params: { id } })
}

export const delTableListApi = (ids: string[] | number[]): Promise<IResponse> => {
  return request.post({ url: '/example/delete', data: { ids } })
}