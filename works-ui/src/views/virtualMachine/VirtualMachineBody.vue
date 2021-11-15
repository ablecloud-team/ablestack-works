<template>
  <ASpace id="content-space" direction="horizontal">
    <ARow id="content-row" style="min-height: 400px">
      <ACol :span="8" style="background: #f0f2f5; padding-right: 8px">
        <!-- 왼쪽 detail 창 -->
        <ACard bordered style="min-height: 300px">
          <VirtualMachineInfoCard 
            ref="listRefreshCall1"
          />
        </ACard>
      </ACol>

      <ACol :span="16" style="background: #f0f2f5; padding-left: 8px">
        <!-- 오른쪽 tab 창 -->
        <ACard bordered>
          <VirtualMachineTab 
            ref="listRefreshCall2"
          />
        </ACard>
      </ACol>
    </ARow>
  </ASpace>
</template>

<script>
// import TabbedContent from "@/components/TabbedContent";
import VirtualMachineTab from "./VirtualMachineTab.vue";
import { defineComponent, ref } from "vue";
import VirtualMachineInfoCard from "./VirtualMachineInfoCard.vue";

export default defineComponent({
  components: { VirtualMachineTab, VirtualMachineInfoCard },
  props: {
  },
  setup() {
    return {
      actionFrom: ref("VirtualMachineDetail"),
    };
  },
  methods: {
    refresh() {
      this.$refs.listRefreshCall1.refresh();
      this.$refs.listRefreshCall2.refresh();
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
