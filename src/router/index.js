import { createRouter, createWebHistory } from 'vue-router'
import CompetitionList from '../components/CompetitionList.vue'
import Resume from '../components/Resume.vue'
import CompetitionLanding from '../components/CompetitionLanding.vue'

const routes = [
  {
    path: '/',
    name: 'CompetitionLanding',
    component: CompetitionLanding
  },
  {
    path: '/competition',
    name: 'CompetitionList',
    component: CompetitionList
  },
  {
    path: '/competition/20260131',
    name: 'Competition20260131',
    component: Resume
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
