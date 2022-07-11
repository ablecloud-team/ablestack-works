<template>
  <div id="ContentTab">
    <a-tabs v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane
        key="1"
        :tab="$t('label.desktop.vm.list')"
        :forceRender="forceRender"
      >
        <TableContent
          ref="listRefreshCall1"
          :tap-name="'desktop'"
          :action-from="'VMList'"
          :resource="resource"
          @parentRefresh="parentRefresh"
        />
      </a-tab-pane>
      <a-tab-pane key="2" :tab="$t('label.users')" :forceRender="forceRender">
        <TableContent
          ref="listRefreshCall2"
          :tap-name="'user'"
          :action-from="'WSUserList'"
          :resource="resource"
          @parentRefresh="parentRefresh"
        />
      </a-tab-pane>
      <a-tab-pane
        key="3"
        :tab="$t('label.policy.list')"
        :forceRender="forceRender"
      >
        <a-tabs v-model:activeKey="policyActiveKey" type="card">
          <a-tab-pane
            key="1"
            :tab="$t('label.group.policy')"
            :forceRender="forceRender"
          >
            <TableContent
              ref="listRefreshCall31"
              :tap-name="'gpolicy'"
              :action-from="'WSPolicyList'"
              :resource="resource"
            />
          </a-tab-pane>
          <a-tab-pane
            key="2"
            :tab="$t('label.webclient.policy')"
            :forceRender="forceRender"
          >
            <TableContent
              ref="listRefreshCall32"
              :tap-name="'wpolicy'"
              :action-from="'WSPolicyList'"
              :resource="resource"
            />
          </a-tab-pane>
        </a-tabs>

        <!-- <TableContent
          ref="listRefreshCall3"
          :tap-name="'policy'"
          :action-from="'WSPolicyList'"
          :resource="resource"
        /> -->
      </a-tab-pane>
      <a-tab-pane
        key="4"
        :tab="$t('label.network.list')"
        :forceRender="forceRender"
      >
        <TableContent
          ref="listRefreshCall4"
          :tap-name="'network'"
          :resource="resource"
        />
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
    resource: {
      type: Object,
      required: true,
      default: null,
    },
  },
  emits: ["parentRefresh"],
  setup() {
    return {
      tabPosition: ref("top"),
      activeKey: ref("1"),
      policyActiveKey: ref("1"),
      forceRender: ref(true),
    };
  },
  created() {},
  methods: {
    fetchRefresh(refreshClick) {
      this.forceRender = true;
      this.$refs.listRefreshCall1.fetchRefresh(refreshClick);
      this.$refs.listRefreshCall2.fetchRefresh(refreshClick);
      this.$refs.listRefreshCall31.fetchRefresh(refreshClick);
      this.$refs.listRefreshCall32.fetchRefresh(refreshClick);
      this.$refs.listRefreshCall4.fetchRefresh(refreshClick);
    },
    parentRefresh() {
      this.$emit("parentRefresh");
    },
  },
});
</script>

<style>
#ContentTab {
  text-align: left;
}
</style>
