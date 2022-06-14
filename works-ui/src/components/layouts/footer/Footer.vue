<template>
  <div class="footer">
    <div class="line">Copyright (c) 2021-2022, ABLECLOUD.Co.Ltd</div>
    <div class="line">ABLESTACK({{ buildVersion }})</div>
  </div>
</template>

<script>
export default {
  name: "Footer",
  data() {
    return {
      buildVersion: process.env.VUE_APP_VERSION,
    };
  },
  created() {
    //젠킨스를 이용한 빌드가 아닐 경우 버전 및 빌드 날짜 표시하기 위함.
    if (this.buildVersion === "") {
      let version = "Bronto-v2.0.0";
      const m = new Date();
      const date =
        m.getFullYear() +
        ("0" + (m.getMonth() + 1)).slice(-2) +
        ("0" + m.getDate()).slice(-2);
      this.buildVersion = version + "-" + date + "-dev";
    }
  },
};
</script>

<style scoped></style>
