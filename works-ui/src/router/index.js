import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/auth/Login.vue')
  },
  {
    path: '/dashboardLayout',
    name: 'DashboardLayout',
    components: () => import('../views/admin/dashboard/Dashboard.vue')
  },
  {
    path: '/',
    name: 'Home',
    component: () => import('../layouts/AdminApp.vue'),
    children: [
      {
        path: '/dashboard',
        name: 'Dashboard',
        components: () => import('../views/admin/dashboard/Dashboard.vue')
      },
      {
        path: '/workSpace',
        name: 'WorkSpace',
        components: () => import('../views/admin/workSpace/WorkSpace.vue')
      },
    ]
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
