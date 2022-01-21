<template>
  <ASpace id="content-space" direction="horizontal">
    <ARow id="content-row" style="min-height: 400px">
      <ACol :span="8" style="background: #f0f2f5; padding-right: 8px">
        <!-- 왼쪽 detail 창 -->
        <ACard bordered style="min-height: 300px">
          <VirtualMachineInfoCard
            ref="listRefreshCall1"
            :vmDbDataInfo="vmDbDataInfo"
            :vmMoldDataInfo="vmMoldDataInfo"
            :vmNetworkInfo="vmNetworkInfo"
            :vmDiskInfo="vmDiskInfo"
            :cpuused="cpuused"
          />
        </ACard>
      </ACol>

      <ACol :span="16" style="background: #f0f2f5; padding-left: 8px">
        <!-- 오른쪽 tab 창 -->
        <ACard bordered>
          <VirtualMachineTab
            ref="listRefreshCall2"
            :vmDbDataInfo="vmDbDataInfo"
            :vmMoldDataInfo="vmMoldDataInfo"
            :vmNetworkInfo="vmNetworkInfo"
            :vmDiskInfo="vmDiskInfo"
          />
        </ACard>
      </ACol>
    </ARow>
  </ASpace>
</template>

<script>
import { defineComponent, ref } from "vue";
import VirtualMachineInfoCard from "./VirtualMachineInfoCard.vue";
import VirtualMachineTab from "./VirtualMachineTab.vue";

export default defineComponent({
  components: { VirtualMachineTab, VirtualMachineInfoCard },
  props: {
    vmDbDataInfo: {
      type: Object,
      required: false,
      default: null,
    },
    vmMoldDataInfo: {
      type: Object,
      required: false,
      default: null,
    },
    vmNetworkInfo: {
      type: Object,
      required: false,
      default: null,
    },
    vmDiskInfo: {
      type: Object,
      required: false,
      default: null,
    },
    cpuused: {
      type: Number,
      required: false,
      default: 0,
    },
  },
  setup() {
    return {
      actionFrom: ref("VirtualMachineDetail"),
    };
  },
  methods: {
    fetchRefresh() {
      this.$refs.listRefreshCall1.fetchRefresh();
      this.$refs.listRefreshCall2.fetchRefresh();
    },
  },
});
</script>
<style>
#content-space,
#content-row {
  width: 100%;
}
#content-space > .ant-space-item {
  width: 100%;
}
</style>
