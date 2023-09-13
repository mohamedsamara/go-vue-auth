<script setup lang="ts">
import { DefaultLayout } from '@/components'
import { Button, Heading } from '@/components/Shared'
</script>
<template>
  <DefaultLayout show-footer>
    <main class="flex items-center justify-center flex-1">
      <section class="relative overflow-hidden">
        <div
          class="flex flex-col justify-center max-w-screen-xl px-4 py-24 mx-auto lg:py-0 sm:px-8"
        >
          <div class="flex flex-col items-center justify-between w-full lg:flex-row">
            <div class="order-2 lg:w-1/2 lg:order-1">
              <div>
                <Heading as="h2" class="mb-3 leading-tight capitalize">
                  Auth <span class="mx-1">system</span> platform for websites
                </Heading>
              </div>
              <p class="mb-4 paragraph">
                Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
                incididunt ut labore et dolore magna aliqua.
              </p>
              <div>
                <Button
                  variant="secondary"
                  class="max-w-full rounded-3xl"
                  @click="$router.push('login')"
                  >Get Started</Button
                >
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
  </DefaultLayout>
</template>
