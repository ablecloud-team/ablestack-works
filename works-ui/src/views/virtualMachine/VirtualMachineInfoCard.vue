<template>
  <div class="resource-details">
    <div class="resource-details__name" v-if="actionFrom==='VirtualMachineDetail'">
      <a-avatar shape="square" :size="60">
        <template #icon>
          <DesktopOutlined />
        </template>
      </a-avatar>
      <h4 style="margin-left:20px;">
        {{ vmDataInfo.Name }}
      </h4>
    </div>
  </div>
  <div id="Status" class="CardItem" >
    <div class="ItemName">{{ $t("label.state") }}</div>
    <div class="Item">
      <a-badge class="Status" :color="vmDataInfo.State == 'Disable' ?'grey' : vmDataInfo.State == 'Running' ? 'green' : 'red'" :text="vmDataInfo.State" />
    </div>
  </div>

  <div id="ID" class="CardItem">
    <div class="ItemName">ID</div>
    <div class="Item">
      <a-button shape="circle" type="dashed">
        <BarcodeOutlined />
      </a-button>
      {{ vmDataInfo.Uuid }}
    </div>
  </div>

 
</template>

<script>
import { DesktopOutlined } from '@ant-design/icons-vue';
import { defineComponent } from "vue";

// function guid() {
//   function _s4() {
//     return (((1 + Math.random()) * 0x10000) | 0).toString(16).substring(1);
//   }
//   return (
//     _s4() +
//     _s4() +
//     "-" +
//     _s4() +
//     "-" +
//     _s4() +
//     "-" +
//     _s4() +
//     "-" +
//     _s4() +
//     _s4() +
//     _s4()
//   );
// }
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
    //let infoID = guid();
    return { infoID:null };
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
