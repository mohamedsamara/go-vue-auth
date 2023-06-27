import { API_BASE_URL } from '../constants'

export class HttpClient {
  private readonly domain: string = API_BASE_URL

  private async sendJSON(method: string, path: string, body: string | null): Promise<Response> {
    const request: RequestInit = {
      method: method,
      body: body
    }

    request.headers = {
      'Content-Type': 'application/json'
    }

    const endpoint = `${this.domain}${path}`
    return fetch(endpoint, request)
  }

  public async post(path: string, body: string | null): Promise<Response> {
    return this.sendJSON('POST', path, body)
  }

  public async patch(path: string, body: string | null): Promise<Response> {
    return this.sendJSON('PATCH', path, body)
  }

  public async put(path: string, body: string | null): Promise<Response> {
    return this.sendJSON('PUT', path, body)
  }

  public async get(path: string): Promise<Response> {
    return this.sendJSON('GET', path, null)
  }

  public async delete(path: string): Promise<Response> {
    return this.sendJSON('DELETE', path, null)
  }
}
