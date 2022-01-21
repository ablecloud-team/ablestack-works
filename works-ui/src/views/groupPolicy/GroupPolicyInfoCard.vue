<template>
  <div class="resource-details">
    <div class="resource-details__name">
      <a-avatar shape="square" :size="60">
        <template #icon>
          <UsergroupAddOutlined />
        </template>
      </a-avatar>
      <h4 style="margin-left:20px;">
        {{ groupDataInfo.groupname }}
      </h4>
    </div>
  </div>

  <div id="firstname" class="CardItem">
    <div class="ItemName">{{ $t('label.distinguishedName') }}</div>
    <div class="Item">{{ groupDataInfo.distinguishedName }}</div>
  </div>


</template>

<script>
import { defineComponent } from "vue";
export default defineComponent({
  components: {},
  props: {
    groupDataInfo: {
      type: Object,
      required: true,
      default: null,
    },
  },
  setup() {
    return {};
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
