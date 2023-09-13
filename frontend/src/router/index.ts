import { createRouter, createWebHistory } from 'vue-router'

import { useAuth } from '@/stores'
import Home from '../views/HomeView.vue'
import Login from '../views/LoginView.vue'
import Signup from '../views/SignupView.vue'
import Account from '../views/AccountView.vue'
import NotFound from '../views/NotFoundView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
      meta: {
        protected: true
      }
    },
    {
      path: '/signup',
      name: 'signup',
      component: Signup,
      meta: {
        protected: true
      }
    },
    {
      path: '/account',
      name: 'account',
      component: Account,
      meta: {
        private: true
      }
    },
    { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound }
  ]
})

router.beforeEach(async (to, from, next) => {
  const authStore = useAuth()
  const authenticated = await authStore.canAccess()

  const privateAuth = to.matched.some((x) => x.meta.private)
  const protectedAuth = to.matched.some((x) => x.meta.protected)

  if (privateAuth && !authenticated) {
    next('/login')
  } else if (protectedAuth && authenticated) {
    next('/account')
  } else {
    next()
  }
})

export default router
