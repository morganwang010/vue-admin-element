import request from '@/config/axios'
import type { ContractTableData } from './types'

export const getContractListApi = (params: any): Promise<IResponse<ContractTableData>> => {
  return request.get({ url: '/contract/list', params })
}

export const saveTableApi = (data: Partial<ContractTableData>): Promise<IResponse> => {
  return request.post({ url: '/example/save', data })
}

export const getTableDetApi = (id: string): Promise<IResponse<ContractTableData>> => {
  return request.get({ url: '/example/detail', params: { id } })
}

export const delTableListApi = (ids: string[] | number[]): Promise<IResponse> => {
  return request.post({ url: '/example/delete', data: { ids } })
}
