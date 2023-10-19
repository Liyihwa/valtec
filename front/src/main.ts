import { createApp } from 'vue'
import App from './App.vue'
import { createRouter, createWebHistory } from 'vue-router'
import routes from '~pages'
import { createPinia } from 'pinia'
import OpenLayersMap from 'vue3-openlayers'
import 'virtual:uno.css'
import '@unocss/reset/tailwind.css'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import 'vue3-openlayers/styles.css'

const pinia = createPinia()
const router = createRouter({
  history: createWebHistory(),
  routes,
})

createApp(App).use(router).use(OpenLayersMap).use(pinia).use(ElementPlus).mount('#app')
