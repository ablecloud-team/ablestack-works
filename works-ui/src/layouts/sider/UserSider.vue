<template>
  <div style="width: 256px">
    <Logo />
    <a-menu
        mode="inline"
        theme="light"
    >
      <a-menu-item
          key="1"
          @click="$router.push({name: 'Dashboard'})"
      >
        <template #icon>
          <DashboardOutlined />
        </template>
        <span>대시보드</span>
      </a-menu-item>
      <a-menu-item key="2">
        <template #icon>
          <CloudOutlined />
        </template>
        <span>워크스페이스</span>
      </a-menu-item>
      <a-menu-item key="3">
        <template #icon>
          <DesktopOutlined />
        </template>
        <span>가상머신</span>
      </a-menu-item>
      <a-menu-item key="4">
        <template #icon>
          <TeamOutlined />
        </template>
        <span>사용자</span>
      </a-menu-item>
      <a-menu-item key="5">
        <template #icon>
          <ReconciliationOutlined />
        </template>
        <span>정책</span>
      </a-menu-item>
      <a-menu-item key="6">
        <template #icon>
          <BarChartOutlined />
        </template>
        <span>로그기록</span>
      </a-menu-item>
      <a-menu-item key="7">
        <template #icon>
          <CoffeeOutlined />
        </template>
        <span>커뮤니티</span>
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
