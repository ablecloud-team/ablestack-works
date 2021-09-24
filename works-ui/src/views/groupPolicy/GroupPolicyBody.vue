<template>
  <ASpace direction="horizontal" id="content-space">
    <ARow style="min-height: 400px" id="content-row">
      <ACol :span="8" style="background: #f0f2f5; padding-right: 8px">
        <!-- 왼쪽 detail 창 -->
        <ACard bordered style="min-height: 300px">
          <GroupPolicyInfoCard
          :groupDataInfo="groupDataInfo" />
        </ACard>
      </ACol>

      <ACol :span="16" style="background: #f0f2f5; padding-left: 8px">
        <!-- 오른쪽 tab 창 -->
        <ACard bordered>
          <GroupPolicyTab
          :groupDataInfo="groupDataInfo"/>
        </ACard>
      </ACol>
    </ARow>
  </ASpace>
</template>

<script>
// import TabbedContent from "@/components/TabbedContent";
import GroupPolicyInfoCard from "./GroupPolicyInfoCard";
import GroupPolicyTab from "./GroupPolicyTab";
import { defineComponent, ref } from "vue";

export default defineComponent({
  components: {
    GroupPolicyInfoCard,
    GroupPolicyTab,
  },
  props: {
    groupDataInfo:{
      type: Object,
      required: true,
      default: null,
    },
 },
  setup() {
    return {
    };
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
