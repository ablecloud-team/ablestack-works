<template>
  <div class="exception">
    <div class="img" v-if="type == '403'">
      <img
        src="@/assets/404.png"
        :style="{
          width: '256px',
          height: '256px',
        }"
      />
    </div>
    <div class="img" v-if="type == '404'">
      <img
        src="@/assets/404.png"
        :style="{
          width: '256px',
          height: '256px',
        }"
      />
    </div>
    <div class="img" v-if="type == '500'">
      <img
        src="@/assets/404.png"
        :style="{
          width: '256px',
          height: '256px',
        }"
      />
    </div>
    <div class="content">
      <h1>{{ config[type].title }}</h1>
      <div class="desc">{{ config[type].desc }}</div>
      <div class="action">
        <a-button type="primary" @click="handleToHome">{{
          $t("label.go.back")
        }}</a-button>
      </div>
    </div>
  </div>
</template>

<script>
const types = {
  403: {
    title: "403",
    desc: "Access Forbidden.",
  },
  404: {
    title: "404",
    desc: "Not Found.",
  },
  500: {
    title: "500",
    desc: "Internal Server Error.",
  },
};

export default {
  name: "Exception",
  props: {
    type: {
      type: String,
      default: "404",
    },
  },
  data() {
    return {
      config: types,
    };
  },
  methods: {
    handleToHome() {
      this.$router.go(-1);
    },
  },
};
</script>

<style lang="less" scoped>
.exception {
  height: 100%;
  align-items: center;
  text-align: center;
  .img {
    top: 50%;
    position: relative;
    transform: translateY(-50%);
    margin-top: -50px;
    display: inline-block;
    padding-right: 52px;
    zoom: 1;
  }
  .content {
    top: 50%;
    position: relative;
    transform: translateY(-50%);
    margin-top: -50px;
    display: inline-block;
    flex: auto;
    h1 {
      color: #434e59;
      font-size: 72px;
      font-weight: 600;
      line-height: 72px;
      margin-bottom: 24px;
    }
    .desc {
      color: rgba(0, 0, 0, 0.45);
      font-size: 20px;
      line-height: 28px;
      margin-bottom: 16px;
    }
  }
}

.mobile {
  .exception {
    margin-top: 30px;
    .img {
      padding-right: unset;

      img {
        max-width: 80%;
      }
    }
  }
}
</style>
