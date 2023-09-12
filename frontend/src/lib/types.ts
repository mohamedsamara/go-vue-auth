/* UI - Forms */
export interface SignupPayload {
  name: string
  email: string
  password: string
}
export interface LoginPayload {
  name: string
  email: string
  password: string
}
/* UI - Forms */

/* Stores & APIS */
export interface User {
  id: string
  name: string
  emil: string
  createdAt: number
  updatedAt: number
}

export interface Auth {
  accessToken: string
  refreshToken: string
}

export interface APIResponse {
  status: string
  success: string
  message: string
}
export interface AuthAPIResponse extends APIResponse {
  data: Auth
}
export interface UserAPIResponse extends APIResponse {
  data: User
}
/* Stores & APIS */
