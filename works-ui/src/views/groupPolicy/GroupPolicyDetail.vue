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
                  { name: $t('label.group.policy'), component: 'GroupPolicy' },
                  { name: groupDataInfo.groupname, component: null },
                ]"
              />
            </a-col>

            <!-- 왼쪽 액션 -->
            <a-col id="content-action" :span="12">
              <Actions :actionFrom="actionFrom" />
            </a-col>
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
          <GroupPolicyBody :groupDataInfo="groupDataInfo" />
        </div>
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script>
import Actions from "../../components/Actions";
import Apath from "../../components/Apath";
import GroupPolicyBody from "./GroupPolicyBody.vue";
import { defineComponent, ref } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";
export default defineComponent({
  props: {},
  components: { GroupPolicyBody, Apath, Actions },
  setup() {
    return {
      actionFrom: ref("GroupPolicyDetail"),
    };
  },
  data() {
    return {
      groupDataInfo: [],
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      worksApi
        .get("/api/v1/group/" + this.$route.params.groupName)
        .then((response) => {
          if (response.status == 200) {
            this.groupDataInfo = response.data.result;
            //console.log(response.data.result.member);
          } else {
            message.error(this.$t("message.response.data.fail"));
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
