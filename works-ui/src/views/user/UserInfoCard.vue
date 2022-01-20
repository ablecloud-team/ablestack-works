<template>
  <a-spin :spinning="spinning">
    <div class="resource-details">
      <div class="resource-details__name">
        <a-avatar shape="square" :size="60">
          <template #icon>
            <UserOutlined />
          </template>
        </a-avatar>
        <h4 style="margin-left: 20px">
          {{ accountInfo.username }}
        </h4>
      </div>
    </div>
    <div id="firstname" class="CardItem">
      <div class="ItemName">{{ $t("label.lastname") }}</div>
      <div class="Item">{{ accountInfo.givenName }}</div>
    </div>
    <div id="firstname" class="CardItem">
      <div class="ItemName">{{ $t("label.firstname") }}</div>
      <div class="Item">{{ accountInfo.sn }}</div>
    </div>
    <div id="email" class="CardItem">
      <div class="ItemName">{{ $t("label.email") }}</div>
      <div class="Item">{{ accountInfo.mail }}</div>
    </div>
    <div id="phone" class="CardItem">
      <div class="ItemName">{{ $t("label.phone") }}</div>
      <div class="Item">{{ accountInfo.telephoneNumber }}</div>
    </div>
  </a-spin>
</template>

<script>
import { defineComponent, ref } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";

export default defineComponent({
  components: {},
  props: {
    accountInfo: {
      type: Object,
      required: false,
      default: null,
    },
  },
  data() {
    return {
      spinning: ref(true),
    };
  },
  created() {
    this.fetchRefresh();
  },
  methods: {
    fetchRefresh() {
      this.spinning = true;
      setTimeout(() => {
        this.spinning = false;
      }, 500);
    },
  },
});
</script>

<style lang="scss" scoped>
.CardItem {
  margin-bottom: 20px;
}
::v-deep(.ant-badge-status-dot) {
  width: 12px;
  height: 12px;
}
.ItemName {
  font-size: 14px;
  font-weight: bold;
}
.resource-details {
  text-align: center;
  margin-bottom: 20px;

  &__name {
    display: flex;
    align-items: center;
    align-content: center;
    text-align: center;

    .ant-avata {
      margin-right: 20px;
      overflow: hidden;
      min-width: 50px;
      cursor: pointer;
      img {
        height: 100%;
        width: 100%;
      }
    }
    h4 {
      font-size: 18px;
    }
  }
}
</style>
