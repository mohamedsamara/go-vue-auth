import { createApp } from 'vue'
import { createPinia } from 'pinia'
import Toast from 'vue-toastification'
import 'vue-toastification/dist/index.css'
import type { PluginOptions } from 'vue-toastification/dist/types/src/types'

import './styles/main.css'
import App from './App.vue'
import router from './router'

const app = createApp(App)

const pinia = createPinia()

app.use(pinia)
app.use(router)
const options: PluginOptions = {}

app.use(Toast, options)

app.mount('#app')
