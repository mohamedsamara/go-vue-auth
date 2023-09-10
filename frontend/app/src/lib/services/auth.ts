import { API } from '../apis'
import { decodeToken } from '../utils'
import type { Auth } from '../types'

export class AuthClient {
  private readonly TOKEN_STORAGE: string = 'token'

  public async storeTokens(payload: Auth) {
    localStorage.setItem(this.TOKEN_STORAGE, JSON.stringify(payload))
  }

  public async getTokens(): Promise<Auth> {
    return JSON.parse(localStorage.getItem(this.TOKEN_STORAGE) ?? '')
  }

  public async removeTokens() {
    localStorage.removeItem(this.TOKEN_STORAGE)
  }

  public async verifyJWT() {
    try {
      const tokens = await this.getTokens()
      if (!tokens) return null
      const decodedJwt = decodeToken(tokens.accessToken)
      // TODO: check expiry on the server
      if (decodedJwt.exp * 1000 < Date.now()) {
        const result = await API.AUTH_CLIENT.getRefreshToken()
        if (!result.success) return null
        return result.data
      }
      return tokens
    } catch (error) {
      return null
    }
  }
}
