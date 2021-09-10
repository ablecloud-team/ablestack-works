<template>
  <div class="resource-details">
    <div class="resource-details__name" v-if="actionFrom==='VirtualMachineDetail'">
      <a-avatar shape="square" :size="60">
        <template #icon>
          <DesktopOutlined />
        </template>
      </a-avatar>
      <h4 style="margin-left:20px;">
        {{ vmDataInfo.name }}
      </h4>
    </div>
  </div>
  <div id="Status" class="CardItem" >
    <div class="ItemName">{{ $t("label.state") }}</div>
    <div class="Item">
      <a-badge class="Status" :color="vmDataInfo.state == 'Disable' ?'grey' : vmDataInfo.state == 'Running' ? 'green' : 'red'" :text="vmDataInfo.state" />
    </div>
  </div>

  <div id="ID" class="CardItem">
    <div class="ItemName">{{ $t("label.uuid") }}</div>
    <div class="Item">
      <a-button shape="circle" type="dashed">
        <BarcodeOutlined />
      </a-button>
      {{ vmDataInfo.uuid }}
    </div>
  </div>

 
</template>

<script>
import { DesktopOutlined } from '@ant-design/icons-vue';
import { defineComponent } from "vue";

export default defineComponent({
  components: {
    DesktopOutlined,
  },
  props: {
    actionFrom: {
      type: String,
      required: true,
      default: '',
    },
    vmDataInfo: {
      type: Object,
      required: false,
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
