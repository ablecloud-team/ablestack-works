<template>
  <div class="resource-details">
    <div class="resource-details__name">
      <a-avatar shape="square" :size="60">
        <template #icon>
          <ClusterOutlined />
        </template>
      </a-avatar>
      <h4 style="margin-left:20px;">
        {{ workspaceInfo.name }}
      </h4>
    </div>
  </div>
  <div id="Description" class="CardItem">
    <div class="ItemName">{{ $t("label.description") }}</div>
    <div class="Item">{{ workspaceInfo.description }}</div>
  </div>

  <div id="Status" class="CardItem" >
    <div class="ItemName">{{ $t("label.state") }}</div>
    <div class="Item">
      <a-badge class="Status" :color="workspaceInfo.state == 'Disable' ?'grey' : workspaceInfo.state == 'Running' ? 'green' : 'red'" :text="workspaceInfo.state" />
    </div>
  </div>

  <div id="ID" class="CardItem">
    <div class="ItemName">ID</div>
    <div class="Item">
      <a-button shape="circle" type="dashed">
        <BarcodeOutlined />
      </a-button>
      {{ workspaceInfo.uuid }}
    </div>
  </div>

  <div id="createdate" class="CardItem">
    <div class="ItemName">{{ $t("label.createdate") }}</div>
    <div class="Item">{{ workspaceInfo.create_date }}</div>
  </div>

  <div id="Type" class="CardItem">
    <div class="ItemName">{{ $t("label.type") }}</div>
    <div class="Item">{{ workspaceInfo.workspace_type }}</div>
  </div>

  <div v-if="workspaceInfo.workspace_type =='desktop'" id="shared" class="CardItem">
    <div class="ItemName">{{ $t("label.dedicated.shared") }}</div>
    <div class="Item">{{  workspaceInfo.shared == false ? 'Dedicated' : 'Shared' }}</div>
  </div>

  <div id="" class="CardItem">
    <div class="ItemName">{{ $t("label.desktop.quantity") }}</div>
    <div class="Item">{{ workspaceInfo.quantity }}</div>
  </div>

  <!-- <div id="" class="CardItem">
    <div class="ItemName">{{ $t("label.desktop.connection.quantity") }}</div>
    <div class="Item">{{ workspaceInfo.Type }}</div>
  </div> -->

  <div class="CardItem">
    <div class="ItemName">{{ $t("label.template") }}</div>
    <div class="Item">{{ templateDataList.displaytext }}</div>
  </div>

  <div class="CardItem">
    <div class="ItemName">{{ $t("label.compute.offering") }}</div>
    <div class="Item">{{ offeringDataList.displaytext }}</div>
  </div>
</template>

<script>
import { ClusterOutlined } from '@ant-design/icons-vue';
import { defineComponent } from "vue";
export default defineComponent({
  components: {
    ClusterOutlined,
  },
  props: {
    workspaceInfo: {
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
