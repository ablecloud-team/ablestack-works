<template>
  <div id="ContentTab">
    <a-tabs v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane key="1" :tab="$t('label.detail')">
        <DetailContent ref="listRefreshCall1" :action-from="'AccountDetail'" />
      </a-tab-pane>
      <a-tab-pane key="2" :tab="$t('label.vm.list')">
        <!-- <TableContent
          ref="listRefreshCall2"
          :tap-name="'desktop'"
          :action-from="'VirtualMachineList'"
        /> -->
      </a-tab-pane>
      <a-tab-pane key="3" :tab="$t('label.app.list')">
        <!-- <TableContent
          ref="listRefreshCall3"
          :tap-name="'app'"
          :action-from="'VirtualMachineList'"
        /> -->
      </a-tab-pane>
      <a-tab-pane key="4" :tab="$t('label.policy.list')">
        <!-- <TableContent
          ref="listRefreshCall4"
          :tap-name="'policy'"
          :action-from="'PolicyGroupList'"
        /> -->
      </a-tab-pane>
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
  methods: {
    refresh() {
      this.$refs.listRefreshCall1.refresh();
      // this.$refs.listRefreshCall2.fetchData();
      // this.$refs.listRefreshCall3.fetchData();
      // this.$refs.listRefreshCall4.fetchData();
    },
  },
});
</script>

<style>
#ContentTab {
  text-align: left;
}
</style>
