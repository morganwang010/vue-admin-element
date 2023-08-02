import request from '@/config/axios'

// 获取所有字典
export const getDictApi = (): Promise<IResponse> => {
  return request.get({ url: '/dict/list' })
}
/** 卡片列表 */
export const getCardList = (data?: object): Promise<IResponse> => {
  return request.post({ url: '/getCardList', data })
}
