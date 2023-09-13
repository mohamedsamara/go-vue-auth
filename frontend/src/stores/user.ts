import { ref } from 'vue'
import { defineStore } from 'pinia'

import { API } from '@/lib/apis'
import { handleError } from '@/lib/utils'
import type { User } from '@/lib/types'

export const useUser = defineStore('user', () => {
  const user = ref<User | null>(null)

  async function fetchUser() {
    try {
      const result = await API.USER_CLIENT.fetchUser()
      user.value = result.data
    } catch (error) {
      handleError(error)
    }
  }

  return { user, fetchUser }
})
