<template>
  <ASpace id="content-space" direction="horizontal">
    <ARow id="content-row" style="min-height: 400px">
      <ACol :span="8" style="background: #f0f2f5; padding-right: 8px">
        <!-- 왼쪽 detail 창 -->
        <ACard bordered style="min-height: 300px">
          <VmInfoCard :action-from="actionFrom" :vm-data-info="vmDataInfo" />
        </ACard>
      </ACol>

      <ACol :span="16" style="background: #f0f2f5; padding-left: 8px">
        <!-- 오른쪽 tab 창 -->
        <ACard bordered>
          <VirtualMachineTab :vm-data-info="vmDataInfo" />
        </ACard>
      </ACol>
    </ARow>
  </ASpace>
</template>

<script>
// import TabbedContent from "@/components/TabbedContent";
import VirtualMachineTab from "./VirtualMachineTab";
import { defineComponent, ref } from "vue";
import VmInfoCard from "./VirtualMachineInfoCard";

export default defineComponent({
  components: { VirtualMachineTab, VmInfoCard },
  props: {
    vmDataInfo: {
      type: Object,
      required: true,
      default: null,
    },
  },
  setup() {
    return {
      actionFrom: ref("VirtualMachineDetail"),
    };
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
