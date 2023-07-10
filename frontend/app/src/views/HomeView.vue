<script setup lang="ts">
import { onMounted } from 'vue'
import { usePing } from '@/stores/ping'

const pingStore = usePing()

onMounted(() => {
  pingStore.fetchPing()
})
</script>

<template>
  <main>
    <h4>Hello</h4>
    <span>{{ pingStore.ping }}</span>
  </main>
</template>
