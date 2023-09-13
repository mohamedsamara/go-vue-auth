import { ref, computed, watch } from 'vue'
import { defineStore } from 'pinia'
import { useToast } from 'vue-toastification'
import { useRouter } from 'vue-router'

import type { LoginPayload, SignupPayload, Auth } from '@/lib/types'
import { API } from '@/lib/apis'
import { SERVICE } from '@/lib/services'
import { useUser } from '@/stores'
import { handleError } from '@/lib/utils'

const toast = useToast()

export const useAuth = defineStore('auth', () => {
  const auth = ref<Auth | null>(null)
  const authenticated = computed(() => (auth.value ? true : false))
  const router = useRouter()

  const userStore = useUser()

  watch(auth, () => {
    if (authenticated.value && !userStore.user) {
      userStore.fetchUser()
    }
  })

  async function canAccess() {
    const jwt = await SERVICE.AUTH_CLIENT.verifyJWT()
    if (jwt) {
      auth.value = jwt
      return true
    } else {
      await SERVICE.AUTH_CLIENT.removeTokens()
      auth.value = null
      return false
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

  return { canAccess, authenticated, logout, login, signup }
})
