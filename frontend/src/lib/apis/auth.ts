import type { AuthAPIResponse, LoginPayload, SignupPayload } from '../types'
import { HttpClient } from './api'
import { SERVICE } from '../services'

export class AuthClient {
  private readonly ROOT_PATH: string = '/auth'
  private readonly httpClient = new HttpClient()

  public async login(payload: LoginPayload): Promise<AuthAPIResponse> {
    const path = `${this.ROOT_PATH}/login`
    const body = JSON.stringify(payload)

    const response = await this.httpClient.post({ path, body })
    const result = (await response.json()) as AuthAPIResponse
    if (response.ok) {
      await SERVICE.AUTH_CLIENT.storeTokens(result.data)
      return result
    }
    throw new Error(`Failed to login. ${result.message}`)
  }

  public async signup(payload: SignupPayload): Promise<AuthAPIResponse> {
    const path = `${this.ROOT_PATH}/register`
    const body = JSON.stringify(payload)

    const response = await this.httpClient.post({ path, body })
    const result = await response.json()
    if (response.ok) {
      await SERVICE.AUTH_CLIENT.storeTokens(result.data)
      return result
    }
    throw new Error(`Failed to register. ${result.message}`)
  }

  public async getRefreshToken(): Promise<AuthAPIResponse> {
    const path = `${this.ROOT_PATH}/refresh`
    const tokens = await SERVICE.AUTH_CLIENT.getTokens()
    const body = JSON.stringify({
      refreshToken: tokens.refreshToken
    })
    const response = await this.httpClient.post({ path, body })
    const result = await response.json()
    if (response.ok) {
      await SERVICE.AUTH_CLIENT.storeTokens(result.data)
      return result
    }
    throw new Error(`Failed to get refresh token. ${result.message}`)
  }
}
