import request from '@/config/axios'
import type { ContractTableData } from './types'

export const getContractListApi = (params: any): Promise<IResponse<ContractTableData>> => {
  return request.get({ url: '/contract/list', params })
}

export const createContractApi = (data: Partial<ContractTableData>): Promise<IResponse> => {
  console.log(data)
  console.log('hhhhhhhhhhhh')
  return request.post({ url: '/contract/create', data })
}
export const updateContractApi = (data: Partial<ContractTableData>): Promise<IResponse> => {
  return request.put({ url: '/contract/update', data })
}
// export const updateContractApi = (data: Partial<ContractTableData>): Promise<IResponse> => {
//   console.log(data)
//   console.log('cggggggggggg')
//   return request.put({ url: '/contract/update', data })
// }

export const getTableDetApi = (id: string): Promise<IResponse<ContractTableData>> => {
  return request.get({ url: '/example/detail', params: { id } })
}

export const delTableListApi = (ids: string[] | number[]): Promise<IResponse> => {
  return request.post({ url: '/contract/delete', data: { ids } })
}
