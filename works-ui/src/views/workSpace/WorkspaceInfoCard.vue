<template>
  <div class="resource-details">
    <div class="resource-details__name">
      <a-avatar shape="square" :size="60">
        <template #icon>
          <ClusterOutlined />
        </template>
      </a-avatar>
      <h4 style="margin-left:20px;">
        {{ workspaceDataList.Name }}
      </h4>
    </div>
  </div>
  <div id="Status" class="CardItem" >
    <div class="ItemName">{{ $t("label.state") }}</div>
    <div class="Item">
      <a-badge class="Status" :color="workspaceDataList.State == 'Disable' ?'grey' : workspaceDataList.State == 'Running' ? 'green' : 'red'" :text="workspaceDataList.State" />
    </div>
  </div>

  <div id="ID" class="CardItem">
    <div class="ItemName">ID</div>
    <div class="Item">
      <a-button shape="circle" type="dashed">
        <BarcodeOutlined />
      </a-button>
      {{ workspaceDataList.Uuid }}
    </div>
  </div>

  <div id="Type" class="CardItem">
    <div class="ItemName">{{ $t("label.type") }}</div>
    <div class="Item">{{ workspaceDataList.Type }}</div>
  </div>

  <div v-if="workspaceDataList.Type =='desktop'" id="shared" class="CardItem">
    <div class="ItemName">{{ $t("label.dedicated.shared") }}</div>
    <div class="Item">{{  workspaceDataList.Shared == 'false' ? 'Dedicated' : 'Shared' }}</div>
  </div>

  <div id="" class="CardItem">
    <div class="ItemName">{{ $t("label.desktop.quantity") }}</div>
    <div class="Item">{{ workspaceDataList.Quantity }}</div>
  </div>

  <div id="" class="CardItem">
    <div class="ItemName">{{ $t("label.desktop.connection.quantity") }}</div>
    <!-- <div class="Item">{{ workspaceDataList.Type }}</div> -->
  </div>

  <div class="CardItem">
    <div class="ItemName">{{ $t("label.template") }}</div>
    <div class="Item">{{ templateDataList.displaytext }}</div>
  </div>

  <div class="CardItem">
    <div class="ItemName">{{ $t("label.compute.offering") }}</div>
    <div class="Item">{{ offeringDataList.displaytext }}</div>
  </div>

  <!-- <div class="CardItem">
    <div class="ItemName">{{ $t("label.network") }}</div>
    <div class="Item">{{ networkDataList.displaytext }}</div>
  </div> -->
</template>

<script>
import { ClusterOutlined } from '@ant-design/icons-vue';
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
    ClusterOutlined,
  },
  props: {
    workspaceDataList: {
      type: Object,
      required: false,
      default: null,
    },
    templateDataList: {
      type: Object,
      required: false,
      default: null,
    },
    offeringDataList: {
      type: Object,
      required: true,
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
