<script setup lang="ts">
import { MY_GITHUB } from '@/lib/constants'
</script>

<template>
  <div class="flex items-center justify-center py-8 text-base text-black">
    <span class="pr-1">🚀 Made by</span>
    <a
      :href="MY_GITHUB"
      rel="nofollow noopener"
      target="_blank"
      class="flex items-center focus:outline-blue-300"
    >
      <span class="pr-1">Mohamed Samara</span>
      <img src="@/assets/icons/github.svg" class="w-6 h-6" />
    </a>
  </div>
</template>
