<script setup lang="ts">
import { MY_WEBSITE } from '@/lib/constants'
</script>

<template>
  <div class="flex items-center justify-center py-8">
    <p className="text-gray-700">
      Made by
      <a :href="MY_WEBSITE" rel="nofollow noopener" target="_blank" class="focus:outline-blue-300">
        Mohamed Samara
      </a>
    </p>
  </div>
</template>
