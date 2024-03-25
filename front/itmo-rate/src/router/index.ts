import HomeView from '@/views/HomeView.vue'
import SubjectView from '@/views/SubjectView.vue'
import SubjectsView from '@/views/SubjectsView.vue'
import TeachersView from '@/views/TeachersView.vue'
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    component: HomeView,
  },
  {
    path: '/home',
    component: HomeView,
  },
  {
    path: '/subjects',
    component: SubjectsView,
  },
  {
    path: '/teachers',
    component: TeachersView,
  },
  {
    path: '/subject',
    component: SubjectView,
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes,
})

export default router
