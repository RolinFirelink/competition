import { createRouter, createWebHistory } from 'vue-router'
import Resume from '../components/Resume.vue'

const routes = [
  {
    path: '/',
    name: 'Resume',
    component: Resume
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router 