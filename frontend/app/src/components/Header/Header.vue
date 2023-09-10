<script setup lang="ts">
import { useAuth } from '@/stores'
import { Button } from '@/components/Shared'
import { LogoLink } from '@/components'

const authStore = useAuth()
</script>

<template>
  <nav class="sticky top-0 left-0 right-0 w-full z-[100] border bg-white">
    <div
      class="flex flex-row items-center justify-between max-w-screen-xl px-6 py-4 mx-auto lg:px-8 lg:py-6"
    >
      <LogoLink />
      <div class="flex items-center justify-between">
        <div v-if="authStore.authenticated" class="flex items-center gap-4 ml-2">
          <Button variant="primary" class="rounded-lg" @click="$router.push('account')"
            >Account</Button
          >
          <Button variant="secondary" class="rounded-lg" @click="authStore.logout">Logout</Button>
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
