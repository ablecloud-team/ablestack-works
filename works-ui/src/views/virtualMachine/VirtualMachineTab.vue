<template>
  <div id="ContentTab">
    <a-tabs v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane key="1" :tab="$t('label.detail')">
        <DetailContent
          ref="listRefreshCall1"
          :action-from="'VirtualMachineDetail'"
        />
      </a-tab-pane>
      <a-tab-pane key="2" :tab="$t('label.disk.list')">
        <TableContent
          ref="listRefreshCall2"
          :tap-name="'datadisk'"
          :action-from="'VirtualMachineDetail'"
        />
      </a-tab-pane>
      <!-- <a-tab-pane key="3" :tab="$t('label.network.list')">
        <TableContent :data="vmNetworkList" :columns="vmNetworkListColumns" />
      </a-tab-pane> -->
    </a-tabs>
  </div>
</template>

<script>
import { defineComponent, ref } from "vue";
import TableContent from "@/components/TableContent";
import DetailContent from "@/components/DetailContent";
export default defineComponent({
  components: {
    TableContent,
    DetailContent,
  },
  props: {},
  setup() {
    const tabPosition = ref("top");
    const activeKey = ref("1");
    return {
      tabPosition,
      activeKey,
    };
  },
  data() {
    return {};
  },
  created() {
  },
  methods: {
    refresh() {
      this.$refs.listRefreshCall1.refresh();
      this.$refs.listRefreshCall2.fetchData();
    },
  },
});
</script>

<style>
#ContentTab {
  text-align: left;
}
</style>
