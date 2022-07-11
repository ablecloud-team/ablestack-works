<template>
  <a-space id="content-space" direction="horizontal">
    <a-row id="content-row" style="min-height: 400px">
      <a-col :span="6" style="background: #f0f2f5; padding-right: 8px">
        <!-- 왼쪽 detail 창 -->
        <a-card bordered style="min-height: 300px">
          <VirtualMachineInfoCard ref="listRefreshCall1" :resource="resource" />
        </a-card>
      </a-col>

      <a-col :span="18" style="background: #f0f2f5; padding-left: 8px">
        <!-- 오른쪽 tab 창 -->
        <a-card bordered>
          <VirtualMachineTab ref="listRefreshCall2" :resource="resource" />
        </a-card>
      </a-col>
    </a-row>
  </a-space>
</template>

<script>
import { defineComponent, ref } from "vue";
import VirtualMachineInfoCard from "./VirtualMachineInfoCard.vue";
import VirtualMachineTab from "./VirtualMachineTab.vue";

export default defineComponent({
  components: { VirtualMachineTab, VirtualMachineInfoCard },
  props: {
    resource: {
      type: Object,
      required: true,
      default: null,
    },
  },
  setup() {},
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
