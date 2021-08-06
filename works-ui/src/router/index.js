import { createRouter, createWebHistory } from "vue-router";
import Login from "../views/auth/Login.vue";
import AdminApp from "../components/layouts/AdminApp.vue";
import AdminBaseLayout from "../components/layouts/AdminBaseLayout.vue";
import Dashboard from "../views/dashboard/Dashboard.vue";
import Workspaces from "../views/workSpace/WorkSpace.vue";
import UserBaseLayout from "../components/layouts/UserBaseLayout.vue";
import WorkspacesDetail from "../views/workSpace/WorkSpaceDetail.vue";
import VirtualMachineDetail from "../views/virtualMachine/VirtualMachineDetail.vue";
import VirtualMachine from "../views/virtualMachine/VirtualMachine.vue";
import Favorites from "../views/favorites/Favorites.vue";
import UserDesktop from "../views/desktopApplication/DesktopApplication.vue";
import A from "../views/dashboard/A.vue";
import Users from "../views/users/Users.vue";
import UserDetail from "../views/users/UserDetail.vue";
import GroupPolicy from "../views/groupPolicy/GroupPolicy.vue";
import GroupPolicyDetail from "../views/groupPolicy/GroupPolicyDetail.vue";

const routes = [
  {
    path: "/login",
    name: "Login",
    component: Login,
  },
  {
    path: "/adminApp",
    name: "AdminApp",
    component: AdminApp,
  },
  {
    path: "/a",
    name: "A",
    component: A,
  },
  {
    path: "/",
    name: "home",
    component: AdminBaseLayout,
    meta: { icon: "home" },
    redirect: "/dashboard",
    // beforeEnter: (to, from, failure) => {},
    children: [
      {
        path: "/dashboard",
        name: "Dashboard",
        component: Dashboard,
      },
      {
        path: "/workspaces",
        name: "Workspaces",
        component: Workspaces,
      },
      {
        path: "/workspacesDetail/",
        name: "WorkspacesDetail",
        component: WorkspacesDetail,
        props: true,
      },
      {
        path: "/virtualmachine",
        name: "VirtualMachine",
        component: VirtualMachine,
      },
      {
        path: "/vmdetail/",
        name: "VirtualMachineDetail",
        component: VirtualMachineDetail,
        props: true,
      },
      {
        path: "/users",
        name: "Users",
        component: Users,
      },
      {
        path: "/userdetail/",
        name: "UserDetail",
        component: UserDetail,
        props: true,
      },
      {
        path: "/groupPolicy",
        name: "GroupPolicy",
        component: GroupPolicy,
      },
      {
        path: "/groupPolicyDetail/",
        name: "GroupPolicyDetail",
        component: GroupPolicyDetail,
        props: true,
      },
    ],
  },
  {
    path: "/user",
    name: "User",
    component: UserBaseLayout,
    meta: { icon: "home" },
    redirect: "/favorites",
    children: [
      {
        path: "/favorites",
        name: "Favorites",
        component: Favorites,
      },
      {
        path: "/userDesktop",
        name: "UserDesktop",
        component: UserDesktop,
      },
    ],
  },
];

const index = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default index;
