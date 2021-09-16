<template>
  <a-space direction="horizontal" id="content-space">
    <a-row style="min-height: 400px" id="content-row">
      <a-col :span="8" style="background: #f0f2f5; padding-right: 8px">
        <!-- 왼쪽 detail 창 -->
        <ACard bordered style="min-height: 300px">
          <WorkspaceInfoCard
          :workspaceInfo="workspaceInfo"
          :templateDataList="templateDataList"
          :offeringDataList="offeringDataList"/>
        </ACard>
      </a-col>

      <a-col :span="16" style="background: #f0f2f5; padding-left: 8px">
        <!-- 오른쪽 tab 창 -->
        <a-card bordered>
          <WorkSpaceTab 
          :workspaceInfo="workspaceInfo"/>
        </a-card>
      </a-col>
    </a-row>
  </a-space>
</template>

<script>
import { defineComponent, ref } from "vue";
import WorkSpaceTab from "./WorkSpaceTab";
import WorkspaceInfoCard from "./WorkspaceInfoCard.vue";

export default defineComponent({
  components: { WorkSpaceTab, WorkspaceInfoCard },
  props: {
    workspaceInfo: {
      type: Object,
      required: true,
      default: null,
    },
    templateDataList: {
      type: Object,
      required: true,
      default: null,
    },
    offeringDataList: {
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
