export const checkFileType = (file: File, accepts: string[]) => {
  const newTypes = accepts.join('|')
  // const reg = /\.(jpg|jpeg|png|gif|txt|doc|docx|xls|xlsx|xml)$/i;
  const reg = new RegExp('\\.(' + newTypes + ')$', 'i')

  return reg.test(file.name)
}

export const checkImgType = (file: File) => {
  return isImgTypeByName(file.name)
}

export const isImgTypeByName = (name: string) => {
  return /\.(jpg|jpeg|png|gif|webp)$/i.test(name)
}

export const checkVideoType = (file: File) => {
  return new Promise<{
    result: string
    file: File
  }>((resolve, reject) => {
    const reader = new FileReader()
    reader.readAsDataURL(file)
    reader.onload = () => resolve({ result: reader.result as string, file })
    reader.onerror = (error) => reject(error)
  })
}

export const getBase64WithFile = (file: File) => {
  return new Promise<{
    result: string
    file: File
  }>((resolve, reject) => {
    const reader = new FileReader()
    reader.readAsDataURL(file)
    reader.onload = () => resolve({ result: reader.result as string, file })
    reader.onerror = (error) => reject(error)
  })
}
