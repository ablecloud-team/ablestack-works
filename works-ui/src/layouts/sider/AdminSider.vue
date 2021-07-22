<template>
  <div style="width: 256px; height: 100%;">
    <Logo />
    <a-menu
        mode="inline"
        theme="light"
        style="padding-top: 14px"
    >
      <a-menu-item
          key="1"
          @click="$router.push({name: 'Dashboard'})"
      >
        <template #icon>
          <DashboardOutlined />
        </template>
        <span>{{ $t('label.Dashboard') }}</span>
      </a-menu-item>
      <a-menu-item key="2"
                   @click="$router.push({name: 'Workspaces'})">
        <template #icon>
          <CloudOutlined />
        </template>
        <span>{{ $t('label.Workspace') }}</span>
      </a-menu-item>
      <a-menu-item key="3"
                   @click="$router.push({name: 'Virtualmachine'})">
        <template #icon>
          <DesktopOutlined />
        </template>
        <span>{{ $t('label.VM') }}</span>
      </a-menu-item>
      <a-menu-item key="4">
        <template #icon>
          <TeamOutlined />
        </template>
        <span>{{ $t('label.User') }}</span>
      </a-menu-item>
      <a-menu-item key="5">
        <template #icon>
          <ReconciliationOutlined />
        </template>
        <span>{{ $t('label.GroupPolicy') }}</span>
      </a-menu-item>
      <a-menu-item key="6">
        <template #icon>
          <BarChartOutlined />
        </template>
        <span>{{ $t('label.Audit') }}</span>
      </a-menu-item>
      <a-menu-item key="7">
        <template #icon>
          <CoffeeOutlined />
        </template>
        <span>{{ $t('label.Community') }}</span>
      </a-menu-item>
    </a-menu>
  </div>
</template>
<script>
import { defineComponent } from 'vue';
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
});
</script>
