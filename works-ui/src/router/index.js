import { createRouter, createWebHistory } from "vue-router";
import { axiosUserDetail } from "../api";
// import store from "../store/index";
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
    meta: { requiresAuth: false },
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
    meta: { requiresAuth: false },
  },
  {
    path: "/",
    name: "home",
    component: AdminBaseLayout,
    meta: { requiresAuth: true },
    redirect: "/dashboard",
    // beforeEnter: (to, from, failure) => {},
    children: [
      {
        path: "/dashboard",
        name: "Dashboard",
        component: Dashboard,
        meta: { requiresAuth: true },
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

index.beforeEach(async (to, from, next) => {
  // let isLogin = store.state.user.isLogin;
  // if (isLogin === undefined) isLogin = false;
  // // console.log("router index.beforeEach isLogin = " + isLogin);
  // if (!isLogin) {
  //   // console.log("router index.beforeEach if(!isLogin)");
  //   if (to.name === "Login" || to.name === "A") {
  //     // console.log(
  //     //   'router index.beforeEach (to.name === "Login" || to.name === "A")'
  //     // );
  //     return next();
  //   }
  //   next({ name: "Login" });
  // } else if (isLogin === to.meta.requiresAuth) {
  //   next();
  // } else {
  //   next({ name: "Login" });
  // }
  const res = await axiosUserDetail();
  if (res.status === 202) {
    if (to.name === "Login" || to.name === "A") {
      // console.log(
      //   'router index.beforeEach (to.name === "Login" || to.name === "A")'
      // );
      return next();
    }
    next({ name: "Login" });
  }
  next();
});

export default index;
