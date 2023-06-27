import { ref } from 'vue'
import { defineStore } from 'pinia'

const pingPong = (): Promise<string> =>
  new Promise((resolve) => {
    setTimeout(() => resolve('pong'), 1000)
  })

export const usePing = defineStore('ping', () => {
  const ping = ref<string>('')

  async function fetchPing() {
    try {
      ping.value = await pingPong()
      return ping
    } catch (error) {
      return error
    }
  }

  return { ping, fetchPing }
})
