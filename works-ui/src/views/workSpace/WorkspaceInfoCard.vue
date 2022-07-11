<template>
  <a-spin :spinning="spinning">
    <div class="resource-details">
      <div class="resource-details__name">
        <a-avatar shape="square" :size="60">
          <template #icon>
            <ClusterOutlined />
          </template>
        </a-avatar>
        <h4 style="margin-left: 20px">
          {{ resource.workspaceInfo.name }}
        </h4>
      </div>
    </div>
    <div id="Description" class="CardItem">
      <div class="ItemName">{{ $t("label.description") }}</div>
      <div class="Item">{{ resource.workspaceInfo.description }}</div>
    </div>

    <div id="Status" class="CardItem">
      <div class="ItemName">{{ $t("label.state") }}</div>
      <div class="Item">
        <a-tooltip placement="bottom">
          <template #title>{{
            resource.workspaceInfo.template_ok_check
          }}</template>
          <a-badge
            class="head-example"
            :color="
              resource.workspaceInfo.template_ok_check === 'Not Ready' ||
              resource.workspaceInfo.template_ok_check === 'Pending'
                ? 'red'
                : resource.workspaceInfo.template_ok_check === 'Joining' ||
                  resource.workspaceInfo.template_ok_check === 'Joined'
                ? 'yellow'
                : resource.workspaceInfo.template_ok_check === 'Ready'
                ? 'green'
                : 'red'
            "
            :text="
              resource.workspaceInfo.template_ok_check === 'Not Ready' ||
              resource.workspaceInfo.template_ok_check === 'Pending'
                ? $t('label.vm.status.initializing')
                : resource.workspaceInfo.template_ok_check === 'Joining' ||
                  resource.workspaceInfo.template_ok_check === 'Joined'
                ? $t('label.vm.status.configuring')
                : resource.workspaceInfo.template_ok_check === 'Ready'
                ? $t('label.vm.status.ready')
                : $t('label.vm.status.notready')
            "
          />
          : [{{ resource.workspaceInfo.template_ok_check }}]
        </a-tooltip>
      </div>
    </div>

    <div id="ID" class="CardItem">
      <div class="ItemName">{{ $t("label.uuid") }}</div>
      <div class="Item">
        {{ resource.workspaceInfo.uuid }}
      </div>
    </div>

    <div id="createdate" class="CardItem">
      <div class="ItemName">{{ $t("label.createdate") }}</div>
      <div class="Item">{{ resource.workspaceInfo.create_date }}</div>
    </div>

    <div id="Type" class="CardItem">
      <div class="ItemName">{{ $t("label.type") }}</div>
      <div class="Item">
        {{
          resource.workspaceInfo.workspace_type === "desktop"
            ? $t("label.desktop")
            : resource.workspaceInfo.workspace_type === "application"
            ? $t("label.application")
            : ""
        }}
      </div>
    </div>

    <div
      v-if="resource.workspaceInfo.workspace_type === 'desktop'"
      id="shared"
      class="CardItem"
    >
      <div class="ItemName">{{ $t("label.dedicated.shared") }}</div>
      <div class="Item">
        {{
          resource.workspaceInfo.shared === false
            ? $t("label.dedicated")
            : $t("label.shared")
        }}
      </div>
    </div>

    <div id="" class="CardItem">
      <div class="ItemName">{{ $t("label.desktop.quantity") }}</div>
      <div class="Item">{{ resource.workspaceInfo.quantity }}</div>
    </div>

    <!-- <div id="" class="CardItem">
      <div class="ItemName">{{ $t("label.desktop.connection.quantity") }}</div>
      <div class="Item">{{ workspaceInfo.Type }}</div>
    </div> -->

    <div class="CardItem">
      <div class="ItemName">{{ $t("label.mastertemplate") }}</div>
      <div class="Item">{{ resource.workspaceInfo.master_template_name }}</div>
    </div>

    <div class="CardItem">
      <div class="ItemName">{{ $t("label.compute.offering") }}</div>
      <div class="Item">
        {{ resource.serviceOfferingInfo.serviceoffering[0].displaytext }}
      </div>
    </div>
  </a-spin>
</template>

<script>
import { defineComponent, ref } from "vue";
export default defineComponent({
  components: {},
  props: {
    resource: {
      type: Object,
      required: true,
      default: null,
    },
  },
  data() {
    return {
      spinning: ref(true),
    };
  },
  created() {
    this.fetchRefresh(true);
  },
  methods: {
    fetchRefresh(refreshClick) {
      if (refreshClick) this.spinning = true;
      else this.spinning = false;
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
