<template>
  <div class="user-menu">
    <FontSizeOutlined class="header-notice-icon" />
    <BellOutlined class="header-notice-icon" />
    <UserOutlined class="header-notice-icon" />
  </div>
</template>

<script>
export default {
  name: "AdminHeader"
}
</script>

<style scoped>
.header-notice-icon{
  font-size: 22px;
  padding: 6px;
}

.user-menu{
  float: right;
  height: 100%;
  padding: 0 2px;
}
</style>
