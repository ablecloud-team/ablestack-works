<template>
  <div id="ContentTab">
    <a-tabs v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane key="1" :tab="$t('label.vm.list')">
        <TableContent
          :tap-name="'desktop'"
          :action-from="'VirtualMachineList'"
        />
      </a-tab-pane>
      <a-tab-pane key="2" :tab="$t('label.users')">
        <TableContent
          :tap-name="'user'"
          :action-from="'UserDetail'"
        />
      </a-tab-pane>
      <a-tab-pane key="3" :tab="$t('label.policy.list')">
        <TableContent :tap-name="'policy'" :action-from="'policyDetail'" />
      </a-tab-pane>
      <a-tab-pane key="4" :tab="$t('label.network.list')">
        <TableContent :tap-name="'network'" />
      </a-tab-pane>
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
      required: true,
      default: null,
    },
  },
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
  methods: {},
});
</script>

<style>
#ContentTab {
  text-align: left;
}
</style>
