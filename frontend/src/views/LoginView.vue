<script setup lang="ts">
import { useForm } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/yup'
import * as Yup from 'yup'
import { useRouter } from 'vue-router'

import { useAuth } from '@/stores'
import type { LoginPayload } from '@/lib/types'
import { Heading, Input, Button, ErrorFeedback, Spinner } from '@/components/Shared'
import { LogoLink } from '@/components'

const router = useRouter()
const authStore = useAuth()

const { handleSubmit, isSubmitting, errors, defineInputBinds } = useForm<LoginPayload>({
  validationSchema: toTypedSchema(
    Yup.object().shape({
      email: Yup.string().email('Please enter a valid email').required('Email is required'),
      password: Yup.string().min(6).required('Password is required')
    })
  )
})

const onSubmit = handleSubmit((values, actions) => {
  authStore.login(values).then((done) => {
    if (done) {
      actions.resetForm()
      router.push('account')
    }
  })
})

const email = defineInputBinds('email')
const password = defineInputBinds('password')
</script>
<template>
  <main class="h-full bg-link-water">
    <section
      class="relative flex flex-col justify-center h-full max-w-screen-md px-4 mx-auto overflow-hidden sm:px-8"
    >
      <div class="mb-6 text-center">
        <LogoLink />
      </div>
      <form @submit="onSubmit" class="px-8 py-12 bg-white border border-gray-200 rounded-md">
        <div class="mb-6 text-center">
          <Heading as="h3" class="mb-1">Welcome Back</Heading>
          <p class="text-gray-400">Enter your credentials to access your account</p>
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
            Login
          </Button>
          <p class="mt-6">
            Don’t have an account yet?<RouterLink to="/signup"> Sign Up</RouterLink>
          </p>
        </div>
      </form>
    </section>
  </main>
</template>
