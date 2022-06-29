<template>
  <div id="ContentTab">
    <a-tabs v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane key="1" :tab="$t('label.detail')" :forceRender="forceRender">
        <DetailContent
          ref="listRefreshCall1"
          :action-from="'VMDetail'"
          :resource="resource"
        />
      </a-tab-pane>
      <a-tab-pane
        key="2"
        :tab="$t('label.disk.list')"
        :forceRender="forceRender"
      >
        <TableContent
          ref="listRefreshCall2"
          :tap-name="'datadisk'"
          :action-from="'VMDetail'"
          :resource="resource"
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
  props: {
    resource: {
      type: Object,
      required: true,
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
      this.$refs.listRefreshCall2.fetchRefresh();
    },
  },
});
</script>

<style>
#ContentTab {
  text-align: left;
}
</style>
