<template>
  <div style="width: 256px; height: 100%;">
    <Logo />
    <a-menu
        mode="inline"
        theme="light"
        :inline-collapsed="collapsed"
        v-model:selectedKeys="selectedKeys"
        style="padding-top: 14px"
    >
      <a-menu-item
          key="1"
          @click="$router.push({name: 'Dashboard'})"
      >
        <template #icon>
          <DashboardOutlined />
        </template>
        <span>{{ $t('label.dashboard') }}</span>
      </a-menu-item>
      <a-menu-item key="2"
                   @click="$router.push({name: 'Workspaces'})">
        <template #icon>
          <CloudOutlined />
        </template>
        <span>{{ $t('label.workspace') }}</span>
      </a-menu-item>
      <a-menu-item key="3"
                   @click="$router.push({name: 'Virtualmachine'})">
        <template #icon>
          <DesktopOutlined />
        </template>
        <span>{{ $t('label.vm') }}</span>
      </a-menu-item>
      <a-menu-item key="4"
                   @click="$router.push({name: 'Users'})">
        <template #icon>
          <TeamOutlined />
        </template>
        <span>{{ $t('label.users') }}</span>
      </a-menu-item>
      <a-menu-item key="5"
                   @click="$router.push({name: 'GroupPolicy'})">
        <template #icon>
          <ReconciliationOutlined />
        </template>
        <span>{{ $t('label.group.policy') }}</span>
      </a-menu-item>
      <a-menu-item key="6">
        <template #icon>
          <BarChartOutlined />
        </template>
        <span>{{ $t('label.audit') }}</span>
      </a-menu-item>
      <a-menu-item key="7">
        <template #icon>
          <CoffeeOutlined />
        </template>
        <span>{{ $t('label.community') }}</span>
      </a-menu-item>
    </a-menu>
  </div>
</template>
<script>
import { defineComponent, reactive, toRefs, watch } from 'vue';
import {
  DashboardOutlined,
  CloudOutlined,
  DesktopOutlined,
  TeamOutlined,
  ReconciliationOutlined,
  BarChartOutlined,
  CoffeeOutlined,
} from '@ant-design/icons-vue';
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
  setup() {
    const state = reactive({
      collapsed: false,
      selectedKeys: ['1'],
      openKeys: ['sub1'],
      preOpenKeys: ['sub1'],
    });

    watch(
        () => state.openKeys,
        (val, oldVal) => {
          state.preOpenKeys = oldVal;
        },
    );
    const toggleCollapsed = () => {
      state.collapsed = !state.collapsed;
      state.openKeys = state.collapsed ? [] : state.preOpenKeys;
    };

    return {
      ...toRefs(state),
      toggleCollapsed,
    };
  },
});
</script>
