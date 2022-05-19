<template>
  <div class="footer">
    <div class="line">Copyright (c) 2021-2022, ABLECLOUD.Co.Ltd</div>
    <div class="line">
      ABLESTACK Bronto(v{{
        buildVersion + "-" + buildDate + "-" + buildRelease + buildDev
      }})
    </div>
  </div>
</template>

<script>
export default {
  name: "Footer",
  data() {
    return {
      buildDate: process.env.VUE_APP_BUILD_DATE,
      buildVersion: process.env.VUE_APP_VERSION,
      buildRelease: process.env.VUE_APP_RELEASE,
      buildDev: process.env.VUE_APP_API_URL == "" ? "" : "-dev",
    };
  },
};
</script>

<style scoped></style>
