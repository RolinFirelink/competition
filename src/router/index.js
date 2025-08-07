import { createRouter, createWebHistory } from 'vue-router'
import UserFront from '../components/UserFront.vue'
import AdminBack from '../components/AdminBack.vue'

const routes = [
  {
    path: '/',
    redirect: '/user'
  },
  {
    path: '/user',
    name: 'UserFront',
    component: UserFront
  },
  {
    path: '/meiriyidengdehoutai',
    name: 'AdminBack',
    component: AdminBack
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router 