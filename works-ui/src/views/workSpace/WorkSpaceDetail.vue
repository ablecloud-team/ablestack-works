<template>
  <div id="content-layout">
    <a-layout>
      <a-layout-header id="content-header">
        <div>
          <a-row>
            <!-- 오른쪽 경로 -->
            <a-col id="content-path" :span="12">
              <Apath
                v-if="resource.workspaceInfo"
                :paths="[
                  {
                    name: $t('label.workspace'),
                    component: 'Workspace',
                  },
                  { name: resource.workspaceInfo.name, component: null },
                ]"
              />
              <a-button
                shape="round"
                style="margin-left: 20px"
                size="small"
                @click="refresh(true)"
              >
                <template #icon><ReloadOutlined /></template>
                {{ $t("label.refresh") }}
              </a-button>
            </a-col>
            <!-- 우측 액션 -->
            <a-col id="content-action" :span="12">
              <Actions
                v-if="actionFrom"
                :action-from="actionFrom"
                :ws-info="resource.workspaceInfo"
              />
            </a-col>
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
          <WorkSpaceBody
            v-if="resource.workspaceInfo"
            ref="listRefreshCall"
            :resource="resource"
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
export default defineComponent({
  components: { Apath, Actions, WorkSpaceBody },
  props: {},
  setup() {
    return {};
  },
  data() {
    return {
      actionFrom: ref(""),
      resource: ref([]),
      timer: ref(null),
    };
  },
  created() {
    this.fetchData();
    this.timer = setInterval(() => {
      //60초 자동 갱신
      this.refresh(false);
    }, 60000);
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
      await this.$worksApi
        .get("/api/v1/workspace/" + this.$route.params.workspaceUuid)
        .then((response) => {
          if (response.status == 200) {
            this.resource = response.data.result;
          } else {
            this.$message.error(this.$t("message.response.data.fail"));
          }
        })
        .catch((error) => {
          console.log(error);
          this.$message.error(this.$t("message.response.data.fail"));
        })
        .finally(() => {
          this.actionFrom = "WSDetail";
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
