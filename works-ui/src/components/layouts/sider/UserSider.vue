<template>
  <div style="width: 100%; height: 100%">
    <Logo
      :collapsed="collapsed"
      @click="
        $router.push({ name: 'UserDesktop' });
        selectedKeysSetting(2);
      "
    />
    <a-menu v-model:selectedKeys="menukey" mode="inline" theme="light" inline="true">
      <!-- <a-menu-item
        key="1"
        @click="
          $router.push({ name: 'Favorite' });
          selectedKeysSetting(1);
        "
      >
        <template #icon>
          <StarOutlined />
        </template>
        <span>즐겨찾기</span>
      </a-menu-item> -->
      <a-menu-item
        key="2"
        @click="
          $router.push({ name: 'UserDesktop' });
          selectedKeysSetting(2);
        "
      >
        <template #icon>
          <DesktopOutlined />
        </template>
        <span>{{ $t("label.desktop") }}</span>
      </a-menu-item>
      <!-- <a-menu-item
        key="3"
        @click="
          $router.push({ name: 'UserDesktop' });
          selectedKeysSetting(3);
        "
      >
        <template #icon>
          <AppstoreAddOutlined />
        </template>
        <span>애플리케이션</span>
      </a-menu-item>
      <a-menu-item
        key="4"
        @click="
          $router.push({ name: 'UserDesktop' });
          selectedKeysSetting(4);
        "
      >
        <template #icon>
          <CoffeeOutlined />
        </template>
        <span>커뮤니티</span>
      </a-menu-item> -->
    </a-menu>
  </div>
</template>
<script>
import { defineComponent, ref } from "vue";
import Logo from "./Logo";
export default defineComponent({
  components: {
    Logo,
  },
  props: {
    collapsed: {
      type: Boolean,
      required: false,
      default: false,
    },
  },
  watch: {
    menukey() {
      return [sessionStorage.getItem("menukey")];
    },
  },
  created() {},
  methods: {
    selectedKeysSetting(key) {
      sessionStorage.setItem("menukey", key);
    },
  },
});
</script>
