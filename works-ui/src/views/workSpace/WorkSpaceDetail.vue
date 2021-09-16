<template>
  <div id="content-layout">
    <a-layout>
      <a-layout-header id="content-header">
        <div>
          <a-row>
            <!-- 오른쪽 경로 -->
            <a-col id="content-path" :span="12">
              <Apath
                v-bind:paths="[
                  { name: $t('label.workspace'), component: 'Workspace' },
                  { name: workspaceInfo.name, component: null }, 
                ]"
              />
            </a-col>
            <!-- 우측 액션 -->
            <a-col id="content-action" :span="12">
              <actions :actionFrom="actionFrom" />
            </a-col>
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
          <WorkSpaceBody
          :workspaceInfo="workspaceInfo"
          :templateDataList="templateDataList"
          :offeringDataList="offeringDataList"/>
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
  props: {
  },
  components: { Apath, Actions, WorkSpaceBody },
  setup(props) {
    return {
      actionFrom: ref("WorkspaceDetail"),
    };
  },
  data() {
    return {
      workspaceInfo: [],
      templateDataList: [],
      offeringDataList: [],
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      //alert(this.$route.params.uuid);
      worksApi
        .get("/api/v1/workspace/"+this.$route.params.uuid, { withCredentials: true })
        .then((response) => {
          if (response.status == 200) {
            this.workspaceInfo = response.data.result.workspaceInfo;
            this.templateDataList = response.data.result.templateInfo.listtemplatesresponse.template[0];
            this.offeringDataList = response.data.result.serviceOfferingInfo.listserviceofferingsresponse.serviceoffering[0];            
          } else {
            message.error(this.$t('message.response.data.fail'));
            //console.log("데이터를 정상적으로 가져오지 못했습니다.");
          }

        })
        .catch(function (error) {
          console.log(error);
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
