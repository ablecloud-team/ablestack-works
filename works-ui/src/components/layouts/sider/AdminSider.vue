<template>
  <div style="width: 100%; height: 100%">
    <Logo 
      @click="$router.push({ name: 'Dashboard' }); selectedKeysSetting(1)" 
    />
    <a-menu
      mode="inline"
      theme="light"
      inline="true"
      v-model:selectedKeys="state.selectedKeys"
      style="padding-top: 14px"
    >
      <a-menu-item key="1" @click="$router.push({ name: 'Dashboard' }); selectedKeysSetting(1)">
        <template #icon>
          <DashboardOutlined />
        </template>
        <span>{{ $t("label.dashboard") }}</span>
      </a-menu-item>
      <a-menu-item key="2" @click="$router.push({ name: 'Workspaces' }); selectedKeysSetting(2)">
        <template #icon>
          <CloudOutlined />
        </template>
        <span>{{ $t("label.workspace") }}</span>
      </a-menu-item>
      <a-menu-item key="3" @click="$router.push({ name: 'VirtualMachines' }); selectedKeysSetting(3)">
        <template #icon>
          <DesktopOutlined />
        </template>
        <span>{{ $t("label.vm") }}</span>
      </a-menu-item>
      <a-menu-item key="4" @click="$router.push({ name: 'Users' }); selectedKeysSetting(4)">
        <template #icon>
          <TeamOutlined />
        </template>
        <span>{{ $t("label.users") }}</span>
      </a-menu-item>
      <a-menu-item key="5" @click="$router.push({ name: 'GroupPolicy' }); selectedKeysSetting(5)">
        <template #icon>
          <ReconciliationOutlined />
        </template>
        <span>{{ $t("label.group.policy") }}</span>
      </a-menu-item>
      <a-menu-item key="6" @click="$router.push({ name: 'Audit' }); selectedKeysSetting(6)">
        <template #icon>
          <BarChartOutlined />
        </template>
        <span>{{ $t("label.audit") }}</span>
      </a-menu-item>
      <a-menu-item key="7" @click="$router.push({ name: 'Community' }); selectedKeysSetting(7)">
        <template #icon>
          <CoffeeOutlined />
        </template>
        <span>{{ $t("label.community") }}</span>
      </a-menu-item>
    </a-menu>
  </div>
</template>
<script>
import { defineComponent, reactive, ref, watch } from "vue";
import {
  DashboardOutlined,
  CloudOutlined,
  DesktopOutlined,
  TeamOutlined,
  ReconciliationOutlined,
  BarChartOutlined,
  CoffeeOutlined,
} from "@ant-design/icons-vue";
import Logo from "./Logo";
export default defineComponent({
  components: {
    Logo,
    DashboardOutlined,
    CloudOutlined,
    DesktopOutlined,
    TeamOutlined,
    ReconciliationOutlined,
    BarChartOutlined,
    CoffeeOutlined,
  },
  props: {
    collapsed: Boolean,
  },
  emits: ["changeToggleCollapsed"],
  setup(props) {
    const state = reactive({
      collapsed: ref(props.collapsed),
      selectedKeys: [ "" ? "1" : localStorage.getItem("menukey")],
      openKeys: ["sub1"],
      preOpenKeys: ["sub1"],
    });
    watch(
      () => state.openKeys,
      (val, oldVal) => {
        console.log(val)
        console.log(oldVal)

        state.preOpenKeys = oldVal;
      }
    );
    return {
      state,
    };
  },
  methods: {
    // toggleCollapsed: function () {
    //   this.$emit("changeToggleCollapsed");
    // },
    selectedKeysSetting: function (key) {
      localStorage.setItem("menukey",key);
      //this.selectedKeys = [key];
    },
  },
});
</script>