<template>
  <a-space id="content-space" direction="horizontal">
    <a-row id="content-row" style="min-height: 400px">
      <a-col :span="6" style="background: #f0f2f5; padding-right: 8px">
        <!-- 왼쪽 detail 창 -->
        <a-card bordered style="min-height: 300px">
          <WorkspaceInfoCard ref="listRefreshCall1" :resource="resource" />
        </a-card>
      </a-col>

      <a-col :span="18" style="background: #f0f2f5; padding-left: 8px">
        <!-- 오른쪽 tab 창 -->
        <a-card bordered>
          <WorkSpaceTab
            ref="listRefreshCall2"
            :resource="resource"
            @parentRefresh="parentRefresh"
          />
        </a-card>
      </a-col>
    </a-row>
  </a-space>
</template>

<script>
import { defineComponent } from "vue";
import WorkSpaceTab from "./WorkSpaceTab";
import WorkspaceInfoCard from "./WorkspaceInfoCard.vue";

export default defineComponent({
  components: { WorkSpaceTab, WorkspaceInfoCard },
  props: {
    resource: {
      type: Object,
      required: true,
      default: null,
    },
  },
  emits: ["parentRefresh"],
  setup() {
    return {};
  },
  created() {},
  methods: {
    fetchRefresh(refreshClick) {
      this.$refs.listRefreshCall1.fetchRefresh(refreshClick);
      this.$refs.listRefreshCall2.fetchRefresh(refreshClick);
    },
    parentRefresh() {
      this.$emit("parentRefresh");
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
