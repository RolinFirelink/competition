import { createRouter, createWebHistory } from 'vue-router'
import CompetitionList from '../components/CompetitionList.vue'
import Resume from '../components/Resume.vue'
import Competition20260328 from '../components/Competition20260328.vue'

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
  },
  {
    path: '/20260328',
    name: 'Competition20260328',
    component: Competition20260328
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router