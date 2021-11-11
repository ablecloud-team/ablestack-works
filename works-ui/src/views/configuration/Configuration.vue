<template>
  <div id="content-layout">
    <a-layout>
      <a-layout-header id="content-header">
        <div id="content-header-body">
          <a-row id="content-header-row">
            <!-- 오른쪽 경로 -->
            <a-col id="content-path" :span="12">
              <Apath :paths="[$t('label.configuration')]" />
              <a-button
                shape="round"
                style="margin-left: 20px"
                size="small"
                @click="refresh()"
              >
                <template #icon>
                  <ReloadOutlined /> {{ $t("label.refresh") }}
                </template>
              </a-button>
            </a-col>

            <!-- 왼쪽 액션 -->
            <a-col id="content-action" :span="12">
              <div></div>
            </a-col>
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
          <ConfigurationList
            ref="listRefreshCall"
            @actionFromChange="actionFromChange"
          />
        </div>
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script>
import Actions from "@/components/Actions";
import Apath from "@/components/Apath";
import ConfigurationList from "@/views/configuration/ConfigurationList";
import { defineComponent, ref, reactive } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";

export default defineComponent({
  components: {
    ConfigurationList,
    Apath,
    Actions,
  },
  props: {},
  setup() {
    return {};
  },
  data() {
    return {};
  },
  created() {},
  methods: {
    refresh() {
      this.$refs.listRefreshCall.fetchData();
    },
    actionFromChange(val) {
      //console.log(val);
      this.actionFrom = ref(val);
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
