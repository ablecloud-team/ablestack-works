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
                  { name: workspaceName, component: null },
                ]"
              />
              <a-button
                shape="round"
                style="margin-left: 20px; height: 30px"
                @click="reflesh()"
              >
                <template #icon>
                  <ReloadOutlined /> {{ $t("label.reflesh") }}
                </template>
              </a-button>
            </a-col>
            <!-- 우측 액션 -->
            <a-col id="content-action" :span="12">
              <Actions
                :action-from="actionFrom"
                :workspace-uuid="workspaceUuid"
              />
            </a-col>
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
          <WorkSpaceBody
            ref="listRefleshCall"
          />
        </div>
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script>
import Actions from "@/components/Actions";
import Apath from "@/components/Apath";
import WorkSpaceBody from "@/views/workSpace/WorkSpaceBody";
import { defineComponent, ref } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";

export default defineComponent({
  components: { Apath, Actions, WorkSpaceBody },
  props: {},
  setup(props) {
    return {
      actionFrom: ref("WorkspaceDetail"),
    };
  },
  data() {
    return {
      workspaceUuid: ref(this.$route.params.workspaceUuid),
      workspaceName: ref(this.$route.params.workspaceName),
    };
  },
  created() {},
  methods: {
    reflesh() {
      this.$refs.listRefleshCall.reflesh();
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
  padding: 24px;
  height: auto;
}

#content-path {
  text-align: left;
  align-items: center;
  display: flex;
}

#content-action {
  text-align: right;
}
</style>
