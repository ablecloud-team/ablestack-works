<template>
  <div id="ContentTab">
    <a-tabs v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane key="1" :tab="$t('label.vm.list')">
        <TableContent
          ref="listRefreshCall1"
          :tap-name="'desktop'"
          :action-from="'VirtualMachineList'"
        />
      </a-tab-pane>
      <a-tab-pane key="2" :tab="$t('label.users')">
        <TableContent
          ref="listRefreshCall2"
          :tap-name="'user'"
          :action-from="'UserDetail'"
        />
      </a-tab-pane>
      <a-tab-pane key="3" :tab="$t('label.policy.list')">
        <TableContent
          ref="listRefreshCall3"
          :tap-name="'policy'"
          :action-from="'policyDetail'"
        />
      </a-tab-pane>
      <a-tab-pane key="4" :tab="$t('label.network.list')">
        <TableContent ref="listRefreshCall4" :tap-name="'network'" />
      </a-tab-pane>
    </a-tabs>
  </div>
</template>
<script>
import { defineComponent, ref } from "vue";
import TableContent from "@/components/TableContent";

export default defineComponent({
  components: { TableContent },
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
  created() {},
  methods: {
    refresh() {
      this.$refs.listRefreshCall1.fetchData();
      this.$refs.listRefreshCall2.fetchData();
      this.$refs.listRefreshCall3.fetchData();
      this.$refs.listRefreshCall4.fetchData();
    },
  },
});
</script>

<style>
#ContentTab {
  text-align: left;
}
</style>
