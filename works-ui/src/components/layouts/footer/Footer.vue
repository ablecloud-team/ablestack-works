<template>
  <div class="footer">
    <div class="line">Copyright (c) 2021-2022, ABLECLOUD.Co.Ltd</div>
    <div class="line">
      ABLESTACK Works Cerato(3.0.0- {{ buildDate + "-" + buildSeq + "-dev" }} )
    </div>
  </div>
</template>

<script>
export default {
  name: "Footer",
  data() {
    return {
      buildDate: process.env.VUE_APP_RELEASE_DATE,
      buildSeq: process.env.VUE_APP_VERSION_CODE,
    };
  },
};
</script>

<style scoped></style>
