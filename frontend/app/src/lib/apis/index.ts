import { PingClient } from './ping'
import { AuthClient } from './auth'
import { UserClient } from './user'

export const API = {
  AUTH_CLIENT: new AuthClient(),
  USER_CLIENT: new UserClient(),
  PING_CLIENT: new PingClient()
}
