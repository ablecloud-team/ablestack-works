<template>
  <a-layout style="height: 100%">
    <a-layout-sider
      class="admin-sider-layout"
      width="280"
      v-model:collapsed="state.collapsed" :trigger="null" collapsible
    >
      <AdminSider
        :collapsed="state.collapsed"
      />
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="admin-layout-header">
        <menu-unfold-outlined
          v-if="state.collapsed"
          class="trigger3"
          @click="() => (state.collapsed = !state.collapsed)"
        />
        <menu-fold-outlined v-else class="trigger3" @click="() => (state.collapsed = !state.collapsed)" />
        <AdminHeader />
      </a-layout-header>
      <a-layout-content class="admin-layout-content">
        <router-view />
      </a-layout-content>
      <a-layout-footer class="admin-layout-footer">
        <AdminFooter />
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
import {
  MenuUnfoldOutlined,
  MenuFoldOutlined,
} from '@ant-design/icons-vue';
import AdminSider from "./sider/AdminSider";
import AdminHeader from "./header/AdminHeader";
import AdminFooter from "./footer/AdminFooter";
import { defineComponent, reactive, ref } from "vue";

export default defineComponent({
  name: "AdminBaseLayout",
  components: {
    AdminSider,
    AdminHeader,
    AdminFooter,
    MenuUnfoldOutlined,
    MenuFoldOutlined,
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
    setToggleCollapsed: function () {
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
.trigger3 {
  font-size: 18px;
  line-height: 64px;
  padding: 0 24px;
  cursor: pointer;
  transition: color 0.3s;
}

.trigger3:hover {
  color: #1890ff;
}

</style>
