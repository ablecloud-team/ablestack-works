<template>
  <div class="logo">
    <img
      alt="ABLESTACK Works"
      src="@/assets/ablestack-logo.png"
      class="logo-image"
    />
  </div>
</template>

<script>
export default {
  name: "Logo",
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
  height: 43px;
}
</style>
