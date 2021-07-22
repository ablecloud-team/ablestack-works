import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/auth/Login'
import AdminApp from '../layouts/AdminApp.vue'
import AdminBaseLayout from '../layouts/AdminBaseLayout';
import Dashboard from '../views/admin/dashboard/Dashboard';
import Workspaces from '../views/admin/workSpace/WorkSpace';
import UserBaseLayout from '../layouts/UserBaseLayout';
import Favorites from '../views/user/favorites/Favorites';
import WorkspacesDetail from '../components/TableContent';
import UserDesktop from '../views/user/desktopApp/DesktopApp';
import A from "../views/admin/dashboard/A";

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/adminApp',
    name: 'AdminApp',
    component: AdminApp
  },
  {
    path: '/a',
    name: 'A',
    component: A
  },
  {
    path: '/',
    name: 'home',
    component: AdminBaseLayout,
    meta: { icon: 'home' },
    redirect: '/dashboard',
    children: [
      {
        path: '/dashboard',
        name: 'Dashboard',
        component: Dashboard
      },
      {
        path: '/workspaces',
        name: 'Workspaces',
        component: Workspaces
      },
      {
        path: '/workspacesDetail',
        name: 'WorkspacesDetail',
        component: WorkspacesDetail
      }
    ]
  },
  {
    path: '/user',
    name: 'User',
    component: UserBaseLayout,
    meta: { icon: 'home' },
    redirect: '/favorites',
    children: [
      {
        path: '/favorites',
        name: 'Favorites',
        component: Favorites
      },
      {
        path: '/userDesktop',
        name: 'UserDesktop',
        component: UserDesktop
      }
    ]
  }
]

const index = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default index
