<template>
  <div id="content-layout">
    <a-layout>
      <a-layout-header id="content-header">
        <div id="content-header-body">
          <a-row id="content-header-row">
            <!-- 오른쪽 경로 -->
            <a-col id='content-path' :span="12">
              <Apath v-bind:paths="[$t('label.Workspace')]"
              />

            </a-col>

            <!-- 왼쪽 액션 -->
            <a-col id="content-action" :span="12">
              <actions :add="$t('label.addworkspace')" :poweroff=true :destroy="true" :reset="true"/>
            </a-col>

          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
            <WorkSpaceList :data="WorkspaceListData" :columns="WorkspaceListColumns" :bordered="false"/>
        </div>
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script>
import Actions from "@/components/actions";
import Apath from "@/components/Apath";
import {WorkspaceListData, WorkspaceListColumns} from "@/data"
import WorkSpaceList from "./WorkSpaceList";
import {defineComponent} from "vue";

export default defineComponent({
  props: {
    msg: String
  },
  components: {WorkSpaceList, Apath, Actions},

  setup() {
    return {
      WorkspaceListData, WorkspaceListColumns
    };
  }
})
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
