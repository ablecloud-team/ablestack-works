<template>
  <div style="width: 100%; height: 100%">
    <Logo
      @click="
        $router.push({ name: 'UserDesktop' });
        selectedKeysSetting(2);
      "
    />
    <a-menu
      v-model:selectedKeys="state.selectedKeys"
      mode="inline"
      theme="light"
      inline="true"
      style="padding-top: 14px"
    >
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
        <span>데스크탑</span>
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
import { defineComponent, reactive, ref } from "vue";
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
      selectedKeys: [sessionStorage.getItem("menukey")],
    });
    return {
      state,
    };
  },
  created() {
    this.updateMenu();
  },
  methods: {
    updateMenu() {
      this.state.selectedKeys =
        this.state.selectedKeys == [""]
          ? ["1"]
          : [sessionStorage.getItem("menukey")];
    },
    // toggleCollapsed: function () {
    //   this.$emit("changeToggleCollapsed");
    // },
    selectedKeysSetting: function (key) {
      sessionStorage.setItem("menukey", key);
      this.state.selectedKeys = [key];
    },
  },
});
</script>
