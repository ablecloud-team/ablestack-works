<template>
  <div id="ContentTab">
    <a-tabs v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane key="1" :tab="$t('label.detail')" :forceRender="forceRender">
        <DetailContent
          ref="refreshCall"
          :action-from="'UserDetail'"
          :account-info="accountInfo"
        />
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script>
import { defineComponent, ref } from "vue";
import DetailContent from "@/components/DetailContent";

export default defineComponent({
  components: {
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
  methods: {
    fetchRefresh() {
      this.$refs.refreshCall.fetchRefresh();
    },
  },
});
</script>

<style>
#ContentTab {
  text-align: left;
}
</style>
