<template>
  <div style="width: 256px">
    <Logo />
    <a-menu mode="inline" theme="light" style="padding-top: 14px">
      <a-menu-item key="1">
        <template #icon>
          <StarOutlined />
        </template>
        <span>즐겨찾기</span>
      </a-menu-item>
      <a-menu-item key="2" @click="$router.push({ name: 'UserDesktop' })">
        <template #icon>
          <DesktopOutlined />
        </template>
        <span>데스크탑</span>
      </a-menu-item>
      <a-menu-item key="3">
        <template #icon>
          <AppstoreAddOutlined />
        </template>
        <span>애플리케이션</span>
      </a-menu-item>
      <a-menu-item key="4">
        <template #icon>
          <CoffeeOutlined />
        </template>
        <span>커뮤니티</span>
      </a-menu-item>
    </a-menu>
  </div>
</template>
<script>
import { defineComponent } from "vue";
import Logo from "./Logo";
export default defineComponent({
  components: {
    Logo,
  },
});
</script>
