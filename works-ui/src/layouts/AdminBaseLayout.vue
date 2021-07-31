<template>
  <a-layout style="height: 100%">
    <a-layout-sider
        class="admin-sider-layout"
        :collapsed="state.collapsed"
        :trigger="null"
        collapsible>
      <AdminSider
          :collapsed="state.collapsed"
          @changeToggleCollapsed="setToggleCollapsed"/>
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="admin-layout-header">
        <AdminHeader />
      </a-layout-header>
      <a-layout-content class="admin-layout-content">
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
import AdminFooter from "./footer/AdminFooter";
import {defineComponent, reactive, ref} from "vue";

export default defineComponent({
  name: 'AdminBaseLayout',
  components:{
    AdminSider,
    AdminHeader,
    AdminFooter,
  },
  setup(){
    const state = reactive({
      collapsed: ref(false)
    })
    return{
      state,
    }
  },
  methods:{
    setToggleCollapsed: function (){
      this.state.collapsed = !this.state.collapsed
    }
  }
})
</script>

<style>
.admin-sider-layout{
  flex: 0 0 256px !important;
  max-width: 256px !important;
  min-width: 256px !important;
  width: 256px !important;
  background: white;
}

.admin-layout-header{
  background: white;
  padding: 0;
}

.admin-layout-content{
  background: #f0f2f5;
  min-height: 700px;
  height: 100%;
}

.admin-layout-footer{
  background: #f0f2f5;;
  text-align: center;
}
</style>
