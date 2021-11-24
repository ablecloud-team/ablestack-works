<template>
  <div id="ContentTab">
    <a-tabs v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane key="1" :tab="$t('label.detail')" :forceRender="forceRender">
        <DetailContent
          ref="listRefreshCall1"
          :action-from="'AccountDetail'"
          :account-info="accountInfo"
        />
      </a-tab-pane>
      <!-- <a-tab-pane key="2" :tab="$t('label.vm.list')" :forceRender="forceRender">
        <TableContent
          ref="listRefreshCall2"
          :tap-name="'desktop'"
          :action-from="'VirtualMachineList'"
        />
      </a-tab-pane>
      <a-tab-pane
        key="3"
        :tab="$t('label.app.list')"
        :forceRender="forceRender"
      >
        <TableContent
          ref="listRefreshCall3"
          :tap-name="'app'"
          :action-from="'VirtualMachineList'"
        />
      </a-tab-pane>
      <a-tab-pane
        key="4"
        :tab="$t('label.policy.list')"
        :forceRender="forceRender"
      >
        <TableContent
          ref="listRefreshCall4"
          :tap-name="'policy'"
          :action-from="'PolicyGroupList'"
        />
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
  props: {
    accountInfo: {
      type: Object,
      required: false,
      default: null,
    },
  },
  setup() {
    return {
      tabPosition: ref("top"),
      activeKey: ref("1"),
      forceRender: ref(false),
    };
  },
  data() {
    return {};
  },
  created() {
    setTimeout(() => {
      this.forceRender = true;
    }, 1000);
  },
  methods: {
    fetchRefresh() {
      this.$refs.listRefreshCall1.fetchRefresh();
      // this.$refs.listRefreshCall2.fetchRefresh();
      // this.$refs.listRefreshCall3.fetchRefresh();
      // this.$refs.listRefreshCall4.fetchRefresh();
    },
  },
});
</script>

<style>
#ContentTab {
  text-align: left;
}
</style>
