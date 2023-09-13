<script setup lang="ts">
import { Header, Footer } from '@/components'

interface Props {
  showFooter?: boolean
}

withDefaults(defineProps<Props>(), {
  showFooter: false
})
</script>
<template>
  <Header />
  <slot />
  <Footer v-if="showFooter" />
</template>
