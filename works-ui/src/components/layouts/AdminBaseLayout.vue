<template>
  <a-layout style="height: 100%">
    <a-layout-sider v-model:collapsed="state.collapsed" collapsed-width="72" class="admin-sider-layout" width="256" breakpoint="lg" :trigger="null" collapsible>
      <AdminSider :collapsed="state.collapsed" />
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="admin-layout-header">
        <AdminHeader :collapsed="state.collapsed" @setCollapsed="setCollapsed" />
      </a-layout-header>
      <a-layout-content class="admin-layout-content">
        <router-view />
      </a-layout-content>
      <a-layout-footer class="admin-layout-footer">
        <Footer />
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
import AdminSider from "./sider/AdminSider";
import AdminHeader from "./header/AdminHeader";
import Footer from "./footer/Footer";
import { defineComponent, reactive, ref } from "vue";

export default defineComponent({
  name: "AdminBaseLayout",
  components: {
    AdminSider,
    AdminHeader,
    Footer,
  },
  setup() {
    const state = reactive({
      collapsed: ref(false),
    });
    return {
      state,
    };
  },
  methods: {
    setCollapsed() {
      this.state.collapsed = !this.state.collapsed;
    },
  },
});
</script>

<style>
.ant-layout {
  overflow: auto;
}
.admin-sider-layout {
  min-width: 10px !important;
  background: white;
}

.admin-layout-header {
  background: white;
  padding: 0;
}

.admin-layout-content {
  background: #f0f2f5;
}

.ant-layout-content {
  min-height: fit-content !important;
}

.admin-layout-footer {
  background: #f0f2f5;
  text-align: center;
}
</style>
