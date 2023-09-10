<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from 'vue'

import { useAuth } from '@/stores'
import { Button } from '@/components/Shared'
import { LogoLink } from '@/components'

const sticky = ref(false)

const authStore = useAuth()

onMounted(() => {
  window.addEventListener('scroll', handleStickyNavbar)
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleStickyNavbar)
})

const handleStickyNavbar = () => {
  if (window.scrollY >= 80) {
    sticky.value = true
  } else {
    sticky.value = false
  }
}

const navClass = computed(() => [
  'sticky top-0 left-0 right-0 w-full z-[100] border rounded-full bg-transparent transform origin-center transition-all duration-300 ease-in-out',
  sticky.value
    ? 'z-[100] bg-white border-gray-200 scale-75 shadow-md'
    : 'bg-transparent border-transparent scale-100'
])
</script>

<template>
  <nav :class="navClass">
    <div
      class="flex flex-row items-center justify-between max-w-screen-xl px-6 py-4 mx-auto lg:px-8 lg:py-6"
    >
      <LogoLink />
      <div class="flex items-center justify-between">
        <div v-if="authStore.authenticated">
          <Button variant="secondary" class="rounded-lg" @click="$router.push('account')"
            >Account</Button
          >
        </div>
        <div v-else class="flex items-center gap-4 ml-2">
          <Button variant="primary" class="rounded-lg" @click="$router.push('login')"
            >Sign In</Button
          >
          <Button variant="secondary" class="rounded-lg" @click="$router.push('signup')"
            >Sign Up</Button
          >
        </div>
      </div>
    </div>
  </nav>
</template>
