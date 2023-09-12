import type { UserAPIResponse } from '../types'
import { HttpClient } from './api'

export class UserClient {
  private readonly ROOT_PATH: string = '/user'
  private readonly httpClient = new HttpClient()

  public async fetchUser(): Promise<UserAPIResponse> {
    const path = `${this.ROOT_PATH}/me`

    const response = await this.httpClient.get({ path, authorized: true })
    const result = await response.json()
    if (response.ok) {
      return result
    }
    throw new Error(`Failed to fetch user. ${result.message}`)
  }
}
