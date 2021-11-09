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
                  { name: $t('label.users'), component: 'Account' },
                  { name: userName, component: null },
                ]"
              />
              <a-button
                shape="round"
                style="margin-left: 20px;"
                size="small"
                @click="reflesh()"
              >
                <template #icon>
                  <ReloadOutlined /> {{ $t("label.reflesh") }}
                </template>
              </a-button>
            </a-col>

            <!-- 왼쪽 액션 -->
            <a-col id="content-action" :span="12">
              <Actions :action-from="actionFrom" />
            </a-col>
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
          <AccountBody 
            ref="listRefleshCall"
          />
        </div>
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script>
import Actions from "../../components/Actions";
import Apath from "../../components/Apath";
import AccountBody from "./AccountBody";
import { defineComponent, ref } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";
export default defineComponent({
  components: {
    AccountBody,
    Apath,
    Actions,
  },
  props: {},
  setup() {
    return {
      actionFrom: ref("AccountDetail"),
    };
  },
  data() {
    return {
      userName: ref(this.$route.params.userName),
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    reflesh() { 
      this.$refs.listRefleshCall.reflesh();
    },
    fetchData() {
      worksApi
        .get("/api/v1/user/" + this.$route.params.userName)
        .then((response) => {
          if (response.status == 200) {
            this.userDataInfo = response.data.result;
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
