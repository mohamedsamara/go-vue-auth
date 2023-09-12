<script setup lang="ts">
import { useForm } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/yup'
import * as Yup from 'yup'
import { useRouter } from 'vue-router'

import { useAuth } from '@/stores'
import type { LoginPayload } from '@/lib/types'
import { Button, Heading, Input, ErrorFeedback, Spinner } from '@/components/Shared'
import { LogoLink } from '@/components'

const router = useRouter()
const authStore = useAuth()

const { handleSubmit, isSubmitting, errors, defineInputBinds } = useForm<LoginPayload>({
  validationSchema: toTypedSchema(
    Yup.object().shape({
      name: Yup.string().required('Name is required'),
      email: Yup.string().required('Email is required'),
      password: Yup.string().required('Password is required')
    })
  )
})

const onSubmit = handleSubmit((values, actions) => {
  authStore.signup(values).then((done) => {
    if (done) {
      actions.resetForm()
      router.push('account')
    }
  })
})

const name = defineInputBinds('name')
const email = defineInputBinds('email')
const password = defineInputBinds('password')
</script>
<template>
  <main class="h-full">
    <section class="relative grid h-full grid-cols-12 lg:overflow-hidden gap-x-6">
      <div class="col-span-12 lg:col-span-6 bg-link-water">
        <div class="flex flex-col items-center justify-center h-full">
          <div class="px-8 py-8 space-y-2 text-center">
            <Heading as="h4">Try STEALTH free</Heading>
            <p class="text-gray-400">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
          </div>
          <div data-aos="fade-right" data-aos-delay="150" class="hidden px-4 py-8 lg:block sm:px-8">
            <div class="w-full">
              <img src="@/assets/signup-hero.svg" alt="Signup Hero" class="mx-auto h-[400px]" />
            </div>
          </div>
        </div>
      </div>
      <div class="col-span-12 py-20 lg:overflow-y-scroll lg:col-span-6">
        <div data-aos="fade-left" data-aos-delay="150" class="px-4 sm:px-6">
          <div class="flex flex-col justify-center">
            <div class="mb-6 text-center">
              <LogoLink />
              <Heading as="h3" class="mt-10 mb-4">Sign Up</Heading>
              <p>Already have an account? <RouterLink to="/login">Sign In</RouterLink></p>
            </div>
            <form @submit="onSubmit">
              <div class="mb-8 space-y-2">
                <label for="name">Name</label>
                <Input
                  type="text"
                  name="text"
                  id="name"
                  v-bind="name"
                  placeholder="Name"
                  :has-error="!!errors.name"
                />
                <ErrorFeedback>{{ errors.name }}</ErrorFeedback>
              </div>
              <div class="mb-8 space-y-2">
                <label for="email">Email</label>
                <Input
                  type="email"
                  name="email"
                  id="email"
                  v-bind="email"
                  placeholder="Email"
                  :has-error="!!errors.email"
                />
                <ErrorFeedback>{{ errors.email }}</ErrorFeedback>
              </div>
              <div class="mb-4 space-y-2">
                <label for="password">Password</label>
                <Input
                  type="password"
                  name="password"
                  id="password"
                  v-bind="password"
                  placeholder="Password"
                  :has-error="!!errors.password"
                />
                <ErrorFeedback>{{ errors.password }}</ErrorFeedback>
              </div>
              <div class="mt-8">
                <Button
                  type="submit"
                  variant="primary"
                  class="w-full rounded-lg"
                  :disabled="isSubmitting"
                >
                  <Spinner v-show="isSubmitting" class="mr-3" />
                  Sign Up
                </Button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </section>
  </main>
</template>
