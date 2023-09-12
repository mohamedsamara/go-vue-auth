import { ref } from 'vue'
import { defineStore } from 'pinia'
import { useToast } from 'vue-toastification'

import { API } from '@/lib/apis'
import { handleError } from '@/lib/utils'
import type { User } from '@/lib/types'

const toast = useToast()

export const useUser = defineStore('auth', () => {
  const user = ref<User | null>(null)

  async function fetchUser() {
    try {
      const result = await API.USER_CLIENT.fetchUser()
      user.value = result.data
    } catch (error) {
      const msg = handleError(error)
      console.log('msg', msg)

      // toast.error(msg)
    }
  }

  return { user, fetchUser }
})
