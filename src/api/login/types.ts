export type UserLoginType = {
  username: string
  password: string
}

export type UserType = {
  username: string
  token: string
  role: string
  uid: string
  permissions: string | string[]
}
