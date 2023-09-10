import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { useToast } from 'vue-toastification'
import { useRouter } from 'vue-router'

import { API } from '@/lib/apis'
import { SERVICE } from '@/lib/services'
import { handleError } from '@/lib/utils'
import type { LoginPayload, SignupPayload, Auth } from '@/lib/types'

const toast = useToast()

export const useAuth = defineStore('auth', () => {
  const auth = ref<Auth | null>(null)
  const authenticated = computed(() => auth.value)
  const router = useRouter()

  async function bootstrap() {
    const jwt = await SERVICE.AUTH_CLIENT.verifyJWT()
    if (jwt) {
      auth.value = jwt
    }
  }

  async function logout() {
    await SERVICE.AUTH_CLIENT.removeTokens()
    router.push('login')
    auth.value = null
  }

  async function login(payload: LoginPayload) {
    try {
      const result = await API.AUTH_CLIENT.login(payload)
      auth.value = result.data
      return true
    } catch (error) {
      const msg = handleError(error)
      toast.error(msg)
      return false
    }
  }

  async function signup(payload: SignupPayload) {
    try {
      const result = await API.AUTH_CLIENT.signup(payload)
      auth.value = result.data
      return true
    } catch (error) {
      const msg = handleError(error)
      toast.error(msg)
      return false
    }
  }

  return { bootstrap, authenticated, logout, login, signup }
})
