import { createRouter, createWebHistory } from "vue-router";
import { worksApi } from "../api";
import { message } from "ant-design-vue";
import store from "../store/index";
import Login from "../views/auth/Login.vue";
import AdminApp from "../components/layouts/AdminApp.vue";
import AdminBaseLayout from "../components/layouts/AdminBaseLayout.vue";
import Dashboard from "../views/dashboard/Dashboard.vue";
import Workspace from "../views/workSpace/WorkSpace.vue";
import UserBaseLayout from "../components/layouts/UserBaseLayout.vue";
import WorkspaceDetail from "../views/workSpace/WorkSpaceDetail.vue";
import VirtualMachineDetail from "../views/virtualMachine/VirtualMachineDetail.vue";
import VirtualMachine from "../views/virtualMachine/VirtualMachine.vue";
import Favorites from "../views/favorites/Favorites.vue";
import UserDesktop from "../views/desktopApplication/DesktopApplication.vue";
import Account from "../views/account/Account.vue";
import AccountDetail from "../views/account/AccountDetail.vue";
import GroupPolicy from "../views/groupPolicy/GroupPolicy.vue";
import GroupPolicyDetail from "../views/groupPolicy/GroupPolicyDetail.vue";
import Audit from "../views/audit/Audit.vue";
import AuditDetail from "../views/audit/AuditDetail.vue";
import Community from "../views/community/Community.vue";
import CommunityDetail from "../views/community/CommunityDetail.vue";
import axios from "axios";

const requireAuth = (to, from, next) => {
  //console.log("-----------------------------------");
  //console.log("to.name : " + to.name + ", from.name : " + from.name);

  let menukey = "1";
  //let menuName = to.name.toLowerCase();
  if (to.name.includes("Dashboard")) { menukey = "1"; }
  if (to.name.includes("Workspace")) { menukey = "2"; }
  if (to.name.includes("VirtualMachine")) { menukey = "3"; }
  if (to.name.includes("Account")) { menukey = "4"; }
  if (to.name.includes("GroupPolicy")) { menukey = "5"; }
  if (to.name.includes("Audit")) { menukey = "6"; }
  if (to.name.includes("Community")) { menukey = "7"; }
  if (to.name.includes("Favorite")) { menukey = "8"; }
  if (to.name.includes("UserDesktop")) { menukey = "9"; }

  //console.log(menukey);
  sessionStorage.setItem("menukey", menukey);

  const isAuth = sessionStorage.getItem("token");
  if (isAuth && isAuth !== "") {
    if (to.name === "Dashboard" && from.name === "Login") {
      next();
      // setTimeout(() => {
      //   location.reload(); //강제 리로드 필요함. 버그인지 모르겠음. =>(정상적인 토큰이 localstorage에 있어도 토큰체크시 response status값이 9998로 받음)
      // }, 0);
    } else {
      worksApi
        .get("/api/v1/token")
        .then((response) => {
          //console.log(response);
          if (response.status === 200) {
            //this.userDataInfo = response.data.result.vmInfo;
            next();
          } else {
            message.error("정상적인 토큰값이 아닙니다. 다시 로그인 해주세요.");
            sessionStorage.setItem("token", "");
            next({ name: "Login" });
          }
        })
        .catch(function (error) {
          message.error("정상적인 토큰값이 아닙니다. 다시 로그인 해주세요.");
          sessionStorage.setItem("token", "");
          next({ name: "Login" });
        });
    }
  } else {
    message.error("정상적인 토큰값이 아닙니다. 다시 로그인 해주세요.");
    sessionStorage.clear();
    next({ name: "Login" });
  }
};

const routes = [
  {
    path: '/:catchAll(.*)', 
    redirect: "/login",
    name: 'NotFound'
  },// 정의된 routes값 외 path 요청이 올 경우 자동 로그인 페이지로 이동
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
    path: "/",
    name: "home",
    component: AdminBaseLayout,
    redirect: "/dashboard",
    // beforeEnter: (to, from, failure) => {},
    children: [
      {
        path: "/dashboard",
        name: "Dashboard",
        component: Dashboard,
        beforeEnter: requireAuth,
      },
      {
        path: "/workspace",
        name: "Workspace",
        component: Workspace,
        beforeEnter: requireAuth,
      },
      {
        path: "/workspaceDetail/:workspaceUuid/:workspaceName",
        name: "WorkspaceDetail",
        component: WorkspaceDetail,
        beforeEnter: requireAuth,
        props: true,
      },
      {
        path: "/virtualMachine",
        name: "VirtualMachine",
        component: VirtualMachine,
        beforeEnter: requireAuth,
      },
      {
        path: "/virtualMachineDetail/:vmUuid/:vmName",
        name: "VirtualMachineDetail",
        component: VirtualMachineDetail,
        beforeEnter: requireAuth,
        props: true,
      },
      {
        path: "/account",
        name: "Account",
        component: Account,
        beforeEnter: requireAuth,
      },
      {
        path: "/accountDetail/:userName",
        name: "AccountDetail",
        component: AccountDetail,
        beforeEnter: requireAuth,
        props: true,
      },
      {
        path: "/groupPolicy",
        name: "GroupPolicy",
        component: GroupPolicy,
        beforeEnter: requireAuth,
      },
      {
        path: "/groupPolicyDetail/:groupName",
        name: "GroupPolicyDetail",
        component: GroupPolicyDetail,
        beforeEnter: requireAuth,
        props: true,
      },
      // {
      //   path: "/audit",
      //   name: "Audit",
      //   component: Audit,
      //   beforeEnter: requireAuth,
      // },
      // {
      //   path: "/auditDetail",
      //   name: "AuditDetail",
      //   component: AuditDetail,
      //   beforeEnter: requireAuth,
      //   props: true,
      // },
      // {
      //   path: "/community",
      //   name: "/Community",
      //   component: Community,
      //   beforeEnter: requireAuth,
      // },
      // {
      //   path: "/community",
      //   name: "CommunityDetail",
      //   component: CommunityDetail,
      //   beforeEnter: requireAuth,
      //   props: true,
      // },
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
        beforeEnter: requireAuth,
      },
      {
        path: "/userDesktop",
        name: "UserDesktop",
        component: UserDesktop,
        beforeEnter: requireAuth,
      },
    ],
  },
];

const index = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  mode: 'history',
  routes,
});

export default index;
