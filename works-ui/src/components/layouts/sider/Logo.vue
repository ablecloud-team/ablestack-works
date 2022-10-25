<template>
  <div v-if="!collapsed" class="logo">
    <img alt="ABLESTACK Works" src="@/assets/ablestack-logo.png" class="logo-image" />
  </div>
  <div v-else class="logo">
    <img alt="ABLESTACK Works" src="@/assets/mini-ablestack-logo.png" class="mini-logo-image" />
  </div>
</template>

<script>
export default {
  name: "Logo",
  props: {
    collapsed: {
      type: Boolean,
      required: false,
      default: false,
    },
  },
};
</script>

<style type="less" scoped>
.logo {
  height: 64px;
  position: relative;
  line-height: 64px;
  -webkit-transition: all 0.3s;
  transition: all 0.3s;
  overflow: hidden;
  box-shadow: 1px 1px 0 0 #e8e8e8;
}

.logo-image {
  display: inline-block;
  vertical-align: middle;
  margin-left: 8px;
  margin-right: 8px;
  margin-bottom: -10px;
  height: 100%;
  width: 240px;
}

.mini-logo-image {
  display: inline-block;
  vertical-align: middle;
  margin-left: 8px;
  height: 80%;
  width: 80%;
}
</style>
