// 导入组件
import { ElNotification, ElMessage, ElMessageBox } from 'element-plus'

export const useMessage = () => {
  // 消息类型图标
  const iconMap = {
    success: 'SuccessFilled',
    warning: 'WarningFilled',
    info: 'InfoFilled',
    error: 'ErrorFilled'
  }

  const open1 = () => {
    ElMessage('this is a message.')
  }
  const open2 = () => {
    ElMessage({
      message: 'Congrats, this is a success message.',
      type: 'success'
    })
  }
  const warn = (warnMessage) => {
    ElMessage({
      message: warnMessage,
      type: 'warning'
    })
  }
  const open4 = () => {
    ElMessage.error('Oops, this is a error message.')
  }
  // 创建确认框
  const createConfirm = (options) => {
    return ElMessageBox.confirm(options)
  }

  // 创建消息提示
  const createNotification = (options) => {
    const iconComponent = iconMap[options.type]
    return ElNotification({
      ...options,
      icon: iconComponent
    })
  }

  return {
    warn,
    // 确认框
    confirm: createConfirm,
    // 消息框
    message: ElMessage
  }
}
