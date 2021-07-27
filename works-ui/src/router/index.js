import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/auth/Login'
import AdminApp from '../layouts/AdminApp.vue'
import AdminBaseLayout from '../layouts/AdminBaseLayout';
import Dashboard from '../views/dashboard/Dashboard';
import Workspaces from '../views/workSpace/WorkSpace';
import UserBaseLayout from '../layouts/UserBaseLayout';
import WorkspacesDetail from '../views/workSpace/WorkSpaceDetail'
import VirtualMachineDetail from "../views/virtualMachine/VirtualMachineDetail";
import VirtualMachine from "../views/virtualMachine/VirtualMachine";
import Favorites from '../views/favorites/Favorites';
import UserDesktop from '../views/desktopApplication/DesktopApplication';
import A from "../components/AddModal";

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
        path: '/workspacesDetail/',
        name: 'WorkspacesDetail',
        component: WorkspacesDetail,
        props: true
      },
      {
        path: '/virtualmachine',
        name: 'Virtualmachine',
        component: VirtualMachine
      },
      {
        path: '/vmdetail/',
        name: 'VirtualMachineDetail',
        component: VirtualMachineDetail,
        props: true
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
