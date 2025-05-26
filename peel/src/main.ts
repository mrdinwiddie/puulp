import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import ui from '@nuxt/ui/vue-plugin'

const app = createApp(App)

// app.use(PrimeVue, {
//     theme: {
//         preset: Aura,
//         options: {
//             darkModeSelector: '.app-dark'
//         }
//     },
//     inputVariant: "filled",
//     ripple: true
// })
app.use(createPinia())
app.use(router)
app.use(ui)
app.mount('#app')