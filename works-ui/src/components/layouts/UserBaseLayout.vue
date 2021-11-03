<template>
  <a-layout style="height: 100%">
    <a-layout-sider
      v-model:collapsed="state.collapsed"
      collapsed-width="72"
      class="admin-sider-layout"
      width="240"
      breakpoint="lg"
      :trigger="null"
      collapsible
    >
      <UserSider :collapsed="state.collapsed" />
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="admin-layout-header">
        <UserHeader 
          :collapsed="state.collapsed"
          @setCollapsed="setCollapsed"
        />
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
import UserSider from "./sider/UserSider";
import UserHeader from "./header/UserHeader";
import Footer from "./footer/Footer";
import { defineComponent, reactive, ref } from "vue";

export default defineComponent({
  name: "UserBaseLayout",
  components: {
    UserSider,
    UserHeader,
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
  min-height: 700px;
  height: 100%;
}

.admin-layout-footer {
  background: #f0f2f5;
  text-align: center;
}
</style>
