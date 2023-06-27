<script setup lang="ts">
import { onMounted } from 'vue'
import { usePing } from '@/stores/ping'

const pingStore = usePing()

onMounted(() => {
  pingStore.fetchPing()
})
</script>

<template>
  <main>{{ pingStore.ping }}</main>
</template>
