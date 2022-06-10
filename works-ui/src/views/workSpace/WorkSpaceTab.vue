<template>
  <div id="ContentTab">
    <a-tabs v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane
        key="1"
        :tab="$t('label.desktop.vm.list')"
        :forceRender="forceRender"
      >
        <TableContent
          ref="listRefreshCall1"
          :tap-name="'desktop'"
          :action-from="'VirtualMachineList'"
          :workspace-info="workspaceInfo"
          :vm-list="vmList"
          @parentRefresh="parentRefresh"
        />
      </a-tab-pane>
      <a-tab-pane key="2" :tab="$t('label.users')" :forceRender="forceRender">
        <TableContent
          ref="listRefreshCall2"
          :tap-name="'user'"
          :action-from="'WorkspaceUserList'"
          :workspace-info="workspaceInfo"
          :group-info="groupInfo"
          @parentRefresh="parentRefresh"
        />
      </a-tab-pane>
      <a-tab-pane
        key="3"
        :tab="$t('label.policy.list')"
        :forceRender="forceRender"
      >
        <TableContent
          ref="listRefreshCall3"
          :tap-name="'policy'"
          :action-from="'policyList'"
          :workspace-info="workspaceInfo"
          :workspace-policy-list="workspacePolicyList"
        />
      </a-tab-pane>
      <a-tab-pane
          key="4"
          :tab="$t('label.network.list')"
          :forceRender="forceRender"
      >
        <TableContent
            ref="listRefreshCall4"
            :tap-name="'network'"
            :network-list="networkList"
            :workspace-info="workspaceInfo"
        />
      </a-tab-pane>
<!--      <a-tab-pane-->
<!--          key="5"-->
<!--          :tab="$t('label.ad.policy.list')"-->
<!--          :forceRender="forceRender"-->
<!--      >-->
<!--        <TableContent-->
<!--            ref="listRefreshCall5"-->
<!--            :tap-name="'policy'"-->
<!--            :network-list="networkList"-->
<!--            :workspace-info="workspaceInfo"-->
<!--        />-->
<!--      </a-tab-pane>-->
    </a-tabs>
  </div>
</template>
<script>
import { defineComponent, ref } from "vue";
import TableContent from "@/components/TableContent";

export default defineComponent({
  components: { TableContent },
  props: {
    workspaceInfo: {
      type: Object,
      required: false,
      default: null,
    },
    networkList: {
      type: Object,
      required: false,
      default: null,
    },
    policyList: {
      type: Object,
      required: false,
      default: null,
    },
    vmList: {
      type: Object,
      required: false,
      default: null,
    },
    groupInfo: {
      type: Object,
      required: false,
      default: null,
    },
    workspacePolicyList: {
      type: Object,
      required: false,
      default: null,
    },
  },
  emits: ["parentRefresh"],
  setup() {
    return {
      tabPosition: ref("top"),
      activeKey: ref("1"),
      forceRender: ref(true),
      vmActionFrom: ref("VirtualMachineList"),
      userActionFrom: ref("WorkspaceUserList"),
    };
  },
  created() {},
  methods: {
    fetchRefresh(refreshClick) {
      this.forceRender = true;
      // this.vmActionFrom = "VirtualMachineList";
      // this.userActionFrom = "WorkspaceUserList";
      // setTimeout(() => {
      this.$refs.listRefreshCall1.fetchRefresh(refreshClick);
      this.$refs.listRefreshCall2.fetchRefresh(refreshClick);
      this.$refs.listRefreshCall3.fetchRefresh(refreshClick);
      this.$refs.listRefreshCall4.fetchRefresh(refreshClick);
      this.$refs.listRefreshCall5.fetchRefresh(refreshClick);
      // }, 100);
    },
    parentRefresh() {
      // this.vmActionFrom = "";
      // this.userActionFrom = "";
      this.$emit("parentRefresh");
    },
  },
});
</script>

<style>
#ContentTab {
  text-align: left;
}
</style>
