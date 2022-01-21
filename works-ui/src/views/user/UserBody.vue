<template>
  <ASpace id="content-space" direction="horizontal">
    <ARow id="content-row" style="min-height: 400px">
      <ACol :span="8" style="background: #f0f2f5; padding-right: 8px">
        <!-- 왼쪽 detail 창 -->
        <ACard bordered style="min-height: 300px">
          <UserInfoCard ref="refreshCall1" :account-info="accountInfo" />
        </ACard>
      </ACol>

      <ACol :span="16" style="background: #f0f2f5; padding-left: 8px">
        <!-- 오른쪽 tab 창 -->
        <ACard bordered>
          <UserTab ref="refreshCall2" :account-info="accountInfo" />
        </ACard>
      </ACol>
    </ARow>
  </ASpace>
</template>

<script>
// import TabbedContent from "@/components/TabbedContent";
import UserInfoCard from "./UserInfoCard";
import UserTab from "./UserTab";
import { defineComponent } from "vue";

export default defineComponent({
  components: {
    UserInfoCard,
    UserTab,
  },
  props: {
    accountInfo: {
      type: Object,
      required: false,
      default: null,
    },
  },
  setup() {
    return {};
  },
  methods: {
    fetchRefresh() {
      this.$refs.refreshCall1.fetchRefresh();
      this.$refs.refreshCall2.fetchRefresh();
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
