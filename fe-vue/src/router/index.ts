import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '@/pages/home'
import GuidePage from '@/pages/guide'
import { defineComponent } from 'vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomePage
    },
    {
      path: '/guide',
      name: 'guide',
      component:GuidePage
    }
  ]
})

export default router
