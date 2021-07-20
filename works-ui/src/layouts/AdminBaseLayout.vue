<template>
  <a-layout>
    <a-layout-sider class="admin-sider-layout">
      <AdminSider/>
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="admin-layout-header">
        <AdminHeader />
      </a-layout-header>
      <a-layout-content class="admin-layout-content">
        <Dashboard />
        <router-view/>
      </a-layout-content>
      <a-layout-footer class="admin-layout-footer">
        <AdminFooter />
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
import AdminSider from "./sider/AdminSider";
import AdminHeader from "./header/AdminHeader";
import Dashboard from "../views/admin/dashboard/Dashboard";
import AdminFooter from "./footer/AdminFooter";

export default ({
  name: 'AdminBaseLayout',
  components:{
    AdminSider,
    AdminHeader,
    AdminFooter,
    Dashboard,
  }
})
</script>

<style>
aside{

}
.admin-sider-layout{
  flex: 0 0 256px !important;
  max-width: 256px !important;
  min-width: 256px !important;
  width: 256px !important;
  background: cornflowerblue;
}

.admin-layout-header{
  background: darkorchid;
}

.admin-layout-content{
  background: chartreuse;
}

.admin-layout-footer{
  background: coral;
  text-align: center;
}
</style>
