<script setup lang="ts">
import { useUser } from '@/stores'
import { DefaultLayout } from '@/components'
import { Heading } from '@/components/Shared'

const userStore = useUser()
</script>

<template>
  <DefaultLayout>
    <main class="h-full">
      <section
        class="relative flex flex-col justify-center h-full max-w-screen-md px-4 mx-auto overflow-hidden sm:px-8"
      >
        <div class="max-w-screen-xl px-8 py-12 bg-white border border-gray-200 rounded-md">
          <div class="text-center">
            <Heading as="h4" class="mb-1">Welcome Back, {{ userStore.user?.name }}!</Heading>
            <p class="text-gray-400">{{ userStore.user?.email }}</p>
          </div>
        </div>
      </section>
    </main>
  </DefaultLayout>
</template>
