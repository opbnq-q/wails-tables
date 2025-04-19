import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { Config } from 'primevue'
import Aura from '@primevue/themes/aura'
import 'primeicons/primeicons.css'
import { ru } from 'primelocale/js/ru.js'
import { createPinia } from 'pinia'
import Tooltip from 'primevue/tooltip'


createApp(App).directive('tooltip', Tooltip).use(createPinia()).use(Config, {
    theme: {
        preset: Aura,
    },
    locale: ru
}).mount('#app')
