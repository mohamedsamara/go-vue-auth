<script setup lang="ts">
import { onMounted } from 'vue'

import { usePing } from '@/stores/ping'
import { Header, Footer } from '@/components'
import { Button, Heading } from '@/components/Shared'

const pingStore = usePing()

onMounted(() => {
  pingStore.fetchPing()
})
</script>
<template>
  <Header />
  <div class="absolute z-[-1] top-0 left-0 w-full h-[160vh] bg-white-lilacapprox" />
  <main class="h-full">
    <span>{{ pingStore.ping }}</span>
    <section class="relative py-24 overflow-hidden">
      <div class="flex flex-col justify-center max-w-screen-xl px-4 mx-auto sm:px-8">
        <div class="flex flex-col items-center justify-between w-full lg:flex-row">
          <div class="order-2 lg:w-1/2 lg:order-1">
            <div>
              <Heading as="h2" class="mb-3 leading-tight capitalize">
                Auth system platform
                <span class="mx-1">For</span> webstie
              </Heading>
            </div>
            <p class="mb-4 paragraph">
              Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
              incididunt ut labore et dolore magna aliqua.
            </p>
            <div>
              <Button variant="secondary" class="max-w-full rounded-3xl">Get Started</Button>
            </div>
          </div>
          <div class="order-1 lg:w-1/2 lg:order-2">
            <div id="hero-image">
              <img src="@/assets/hero.svg" alt="Hero" class="mx-auto h-[400px]" />
            </div>
          </div>
        </div>
      </div>
    </section>
  </main>
  <Footer />
</template>
