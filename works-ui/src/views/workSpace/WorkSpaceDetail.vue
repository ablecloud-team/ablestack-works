<template>
  <div id="content-layout">
    <a-layout>
      <a-layout-header id="content-header">
        <div>
          <a-row>
            <!-- 오른쪽 경로 -->
            <a-col id="content-path" :span="12">
              <Apath
                :paths="[
                  { name: $t('label.workspace'), component: 'Workspace' },
                  { name: workspaceInfo.name, component: null },
                ]"
              />
              <a-button
                shape="round"
                style="margin-left: 20px"
                size="small"
                @click="refresh(true)"
              >
                <template #icon>
                  <ReloadOutlined /> {{ $t("label.refresh") }}
                </template>
              </a-button>
            </a-col>
            <!-- 우측 액션 -->
            <a-col id="content-action" :span="12">
              <Actions
                v-if="actionFrom === 'WorkspaceDetail'"
                :action-from="actionFrom"
                :workspace-info="workspaceInfo"
              />
            </a-col>
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
          <WorkSpaceBody
            v-if="actionFrom === 'WorkspaceDetail'"
            ref="listRefreshCall"
            :workspace-info="workspaceInfo"
            :offering-info="offeringDataList"
            :network-list="networkList"
            :vm-list="vmList"
            :group-info="groupInfo"
            :workspace-policy-list="workspacePolicyList"
            @parentRefresh="refresh"
          />
        </div>
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script>
import Actions from "@/components/Actions";
import Apath from "@/components/Apath";
import WorkSpaceBody from "./WorkSpaceBody.vue";
import { defineComponent, ref } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";

export default defineComponent({
  components: { Apath, Actions, WorkSpaceBody },
  props: {},
  setup() {
    return {
      actionFrom: ref(""),
    };
  },
  data() {
    return {
      workspaceInfo: ref([]),
      offeringDataList: ref([]),
      networkList: ref([]),
      vmList: ref([]),
      workspacePolicyList: ref([]),
      groupInfo: ref([]),
      timer: ref(null),
    };
  },
  created() {
    this.fetchData();
    this.timer = setInterval(() => {
      //90초 자동 갱신
      this.refresh(false);
    }, 90000);
  },
  beforeUnmount() {
    clearInterval(this.timer);
  },
  methods: {
    async refresh(refreshClick) {
      await this.fetchData();
      this.$refs.listRefreshCall.fetchRefresh(refreshClick);
    },
    async fetchData() {
      await worksApi
        .get("/api/v1/workspace/" + this.$route.params.workspaceUuid)
        .then((response) => {
          if (response.status == 200) {
            this.workspaceInfo = response.data.result.workspaceInfo;
            this.offeringDataList =
              response.data.result.serviceOfferingInfo.serviceoffering[0];
            this.networkList = response.data.result.networkInfo.network;

            if (response.data.result.instanceList !== null) {
              this.vmList = response.data.result.instanceList;
            } else {
              this.vmList = ref([]);
            }
            if (response.data.result.groupDetail !== null) {
              this.groupInfo = response.data.result.groupDetail;
            }
            if (response.data.result.workspacePolicy !== null) {
              this.workspacePolicyList = response.data.result.workspacePolicy;
            }
          } else {
            message.error(this.$t("message.response.data.fail"));
          }
        })
        .catch((error) => {
          console.log(error);
          message.error(this.$t("message.response.data.fail"));
        })
        .finally(() => {
          this.actionFrom = "WorkspaceDetail";
        });
    },
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#content-layout .ant-layout-content {
  background: #f0f2f5;
  padding: 10px;
}
#content-layout .ant-layout-header {
  text-align: left;
  background: white;
  border: 1px solid #e8e8e8;
  /*color: #fff;*/
  font-size: 14px;
  line-height: 1.5;
  padding: 20px;
  height: auto;
}

#content-path {
  text-align: left;
  align-items: center;
  display: flex;
  height: 32px;
}

#content-action {
  text-align: right;
}
</style>
