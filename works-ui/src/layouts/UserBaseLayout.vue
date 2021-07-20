<template>
  <a-layout>
    <a-layout-sider class="user-sider-layout">
      <UserSider/>
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="user-layout-header">
        <UserHeader />
      </a-layout-header>
      <a-layout-content class="user-layout-content">
        <router-view/>
      </a-layout-content>
      <a-layout-footer class="user-layout-footer">
        <UserFooter />
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
import UserSider from "./sider/UserSider";
import UserHeader from "./header/UserHeader";
import UserFooter from "./footer/UserFooter";

export default ({
  name: 'UserBaseLayout',
  components:{
    UserSider,
    UserHeader,
    UserFooter,
  }
})
</script>

<style>
.user-sider-layout{
  flex: 0 0 256px !important;
  max-width: 256px !important;
  min-width: 256px !important;
  width: 256px !important;
  background: white;
}

.user-layout-header{
  background: white;
}

.user-layout-content{
  background: white;
}

.user-layout-footer{
  background: #f0f2f5;;
  text-align: center;
}
</style>
