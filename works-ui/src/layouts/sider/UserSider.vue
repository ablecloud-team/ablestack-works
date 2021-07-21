<template>
  <div style="width: 256px">
    <Logo />
    <a-menu
        mode="inline"
        theme="light"
    >
      <a-menu-item
          key="1"
          @click="$router.push({name: 'UserDashboard'})"
      >
        <template #icon>
          <CloudOutlined />
        </template>
        <span>가상 데스크톱/앱목록</span>
      </a-menu-item>
      <a-menu-item key="2">
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
  CloudOutlined,
  CoffeeOutlined,
} from '@ant-design/icons-vue';
import Logo from "./Logo";
export default defineComponent({
  components: {
    Logo,
    CloudOutlined,
    CoffeeOutlined,
  },
});
</script>
