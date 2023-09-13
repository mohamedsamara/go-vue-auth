<script setup lang="ts">
import { ref, computed } from 'vue'
import { Bars3BottomRightIcon, XMarkIcon } from '@heroicons/vue/24/solid'

import { useAuth } from '@/stores'
import { Button } from '@/components/Shared'
import { LogoLink } from '@/components'

const open = ref(false)

const authStore = useAuth()

const mobileNavClass = computed(() => [
  'bg-white rounded-2xl shadow-md fixed right-4 left-4 overflow-hidden max-h-0 z-[100] transition-all duration-500',
  open.value ? 'max-h-[460px] border border-gray-200' : 'max-h-0'
])
</script>

<template>
  <nav class="sticky top-0 left-0 right-0 w-full z-[100] border bg-white">
    <div
      class="flex flex-row justify-between max-w-screen-xl gap-4 px-6 py-4 mx-auto lg:items-center lg:px-8 lg:py-6"
    >
      <LogoLink />
      <Button size="sm" class="bg-white rounded-full lg:hidden w-14 h-14" @click="open = !open">
        <XMarkIcon v-if="open" class="w-6 h-6" />
        <Bars3BottomRightIcon v-else class="w-6 h-6" />
      </Button>
      <div class="hidden lg:block">
        <div v-if="authStore.authenticated" class="flex items-center justify-center gap-4">
          <Button variant="primary" class="rounded-lg" @click="$router.push('account')"
            >Account</Button
          >
          <Button variant="secondary" class="rounded-lg" @click="authStore.logout">Logout</Button>
        </div>
        <div v-else class="flex items-center justify-center gap-4">
          <Button variant="primary" class="rounded-lg" @click="$router.push('login')"
            >Sign In</Button
          >
          <Button variant="secondary" class="rounded-lg" @click="$router.push('signup')"
            >Sign Up</Button
          >
        </div>
      </div>
    </div>
    <div :class="mobileNavClass">
      <div class="px-6 py-8 space-y-4">
        <div v-if="authStore.authenticated" class="flex flex-col justify-center gap-4">
          <Button variant="primary" class="rounded-lg" @click="$router.push('account')"
            >Account</Button
          >
          <Button variant="secondary" class="rounded-lg" @click="authStore.logout">Logout</Button>
        </div>
        <div v-else class="flex flex-col justify-center gap-4">
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
