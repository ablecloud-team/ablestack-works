<template>
  <div id="content-layout">
    <a-layout>
      <a-layout-header id="content-header">
        <div>
          <a-row>
            <!-- 왼쪽 경로 -->
            <a-col id="content-path" :span="12">
              <Apath
                :paths="[
                  { name: $t('label.vm'), component: 'VirtualMachine' },
                  { name: vmName, component: null },
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
            <!-- 우측 액션 -->
            <a-col id="content-action" :span="12">
              <Actions
                :action-from="actionFrom"
                @fetchData="reflesh"
              />
            </a-col>
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
          <VirtualMachineBody ref="listRefleshCall"/>
        </div>
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script>
import Actions from "@/components/Actions";
import Apath from "@/components/Apath";
import VirtualMachineBody from "./VirtualMachineBody.vue";
import { defineComponent, ref } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";

export default defineComponent({
  components: { VirtualMachineBody, Apath, Actions },
  props: {},
  setup() {
    return {
      vmDbDataInfo: ref([]),
      vmMoldDataInfo: ref([]),
      actionFrom: ref("VirtualMachineDetail"),
    };
  },
  data() {
    return {
      vmUuid: ref(this.$route.params.vmUuid),
      vmName: ref(this.$route.params.vmName),
    };
  },
  created() {
  },
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
