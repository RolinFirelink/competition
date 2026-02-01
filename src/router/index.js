import { createRouter, createWebHistory } from 'vue-router'
import CompetitionList from '../components/CompetitionList.vue'
import Resume from '../components/Resume.vue'

const routes = [
  {
    path: '/',
    name: 'CompetitionList',
    component: CompetitionList
  },
  {
    path: '/20260131',
    name: 'Competition20260131',
    component: Resume
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router 