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
      <a-menu-item key="2" @click="$router.push({ name: 'Workspace' }); selectedKeysSetting(2)">
        <template #icon>
          <CloudOutlined />
        </template>
        <span>{{ $t("label.workspace") }}</span>
      </a-menu-item>
      <a-menu-item key="3" @click="$router.push({ name: 'VirtualMachine' }); selectedKeysSetting(3)">
        <template #icon>
          <DesktopOutlined />
        </template>
        <span>{{ $t("label.vm") }}</span>
      </a-menu-item>
      <a-menu-item key="4" @click="$router.push({ name: 'Account' }); selectedKeysSetting(4)">
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
      <a-menu-item key="6">
        <!-- <a-menu-item key="6" @click="$router.push({ name: 'Audit' }); selectedKeysSetting(6)"> -->
        <template #icon>
          <BarChartOutlined />
        </template>
        <span>{{ $t("label.audit") }}</span>
      </a-menu-item>
      <!-- <a-menu-item key="7" @click="$router.push({ name: 'Community' }); selectedKeysSetting(7)">
        <template #icon>
          <CoffeeOutlined />
        </template>
        <span>{{ $t("label.community") }}</span>
      </a-menu-item> -->
    </a-menu>
  </div>
</template>
<script>
import { defineComponent, reactive, ref, watch } from "vue";
import Logo from "./Logo";
export default defineComponent({
  components: {
    Logo,
  },
  props: {
    collapsed: Boolean,
  },
  emits: ["changeToggleCollapsed"],
  setup(props) {
    const state = reactive({
      collapsed: ref(props.collapsed),
      selectedKeys: [""],
      openKeys: [""],
      preOpenKeys: [""],
    });
    watch(
      () => state.openKeys,
      (val, oldVal) => {
        console.log(val);
        console.log(oldVal);
        state.preOpenKeys = oldVal;
      }
    );
    return {
      state,
    };
  },
  created() {
    this.updateMenu();
  },
  methods: {
    updateMenu() {
      this.selectedKeys = [ "" ? "1" : sessionStorage.getItem("menukey") ];
    },
    // toggleCollapsed: function () {
    //   this.$emit("changeToggleCollapsed");
    // },
    selectedKeysSetting: function (key) {
      sessionStorage.setItem("menukey", key);
      //this.selectedKeys = [key];
    },
  },
});
</script>
