<template>
  <ASpace id="content-space" direction="horizontal">
    <ARow id="content-row" style="min-height: 400px">
      <ACol :span="8" style="background: #f0f2f5; padding-right: 8px">
        <!-- 왼쪽 detail 창 -->
        <ACard bordered style="min-height: 300px">
          <AccountInfoCard :user-data-info="userDataInfo" />
        </ACard>
      </ACol>

      <ACol :span="16" style="background: #f0f2f5; padding-left: 8px">
        <!-- 오른쪽 tab 창 -->
        <ACard bordered>
          <AccountTab :user-data-info="userDataInfo" />
        </ACard>
      </ACol>
    </ARow>
  </ASpace>
</template>

<script>
// import TabbedContent from "@/components/TabbedContent";
import AccountInfoCard from "./AccountInfoCard";
import AccountTab from "./AccountTab";
import { defineComponent, ref } from "vue";

export default defineComponent({
  components: {
    AccountInfoCard,
    AccountTab,
  },
  props: {
    userDataInfo: {
      type: Object,
      required: true,
      default: null,
    },
  },
  setup() {
    return {};
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
