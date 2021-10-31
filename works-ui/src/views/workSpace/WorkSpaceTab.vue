<template>
  <div id="ContentTab">
    <a-tabs v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane key="1" :tab="$t('label.vm.list')">
        <TableContent
          ref="listRefleshCall1"
          :tap-name="'desktop'"
          :action-from="'VirtualMachineList'"
        />
      </a-tab-pane>
      <a-tab-pane key="2" :tab="$t('label.users')">
        <TableContent
          ref="listRefleshCall2"
          :tap-name="'user'"
          :action-from="'UserDetail'"
        />
      </a-tab-pane>
      <a-tab-pane key="3" :tab="$t('label.policy.list')">
        <TableContent
          ref="listRefleshCall3"
          :tap-name="'policy'"
          :action-from="'policyDetail'"
        />
      </a-tab-pane>
      <a-tab-pane key="4" :tab="$t('label.network.list')">
        <TableContent ref="listRefleshCall4" :tap-name="'network'" />
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
  setup(props) {
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
    reflesh() {
      this.$refs.listRefleshCall1.fetchData();
      this.$refs.listRefleshCall2.fetchData();
      this.$refs.listRefleshCall3.fetchData();
      this.$refs.listRefleshCall4.fetchData();
    },
  },
});
</script>

<style>
#ContentTab {
  text-align: left;
}
</style>
