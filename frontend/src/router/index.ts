import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import { AuthGuard } from '@/auth/AuthGuard'
import Dashboard from '@/views/Dashboard.vue'
import DefaultDashboardView from '@/views/DefaultDashboardView.vue'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Home',
    beforeEnter: AuthGuard,
    component: () => {
      return import('@/views/Home.vue')
    }
  },
  {
    path: '/dashboard',
    // name: 'Dashboard',
    component: Dashboard,
    beforeEnter: AuthGuard,
    children: [
      {
        path: 'objects',
        name: 'objects',
        component: () => import('@/views/Objects.vue'),
        alias: '/objects'
      },
      {
        path: 'profile',
        name: 'profile',
        component: () => import('@/views/Profile.vue'),
        alias: '/profile'
      },
      {
        path: 'settings',
        name: 'settings',
        component: () => import('@/views/Settings.vue'),
        alias: '/settings'
      },
      {
        path: '',
        name: 'default dashboard',
        component: DefaultDashboardView
      }
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
