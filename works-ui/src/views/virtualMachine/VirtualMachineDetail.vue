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
                  { name: vmDbDataInfo.name, component: null },
                ]"
              />
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
            <!-- 우측 액션 -->
            <a-col id="content-action" :span="12">
              <Actions
                v-if="actionFrom === 'VirtualMachineDetail'"
                :action-from="actionFrom"
                :vm-info="vmDbDataInfo"
                @fetchData="refresh"
              />
            </a-col>
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
          <VirtualMachineBody
            ref="listRefreshCall"
            :vmDbDataInfo="vmDbDataInfo"
            :vmMoldDataInfo="vmMoldDataInfo"
            :vmNetworkInfo="vmNetworkInfo"
            :vmDiskInfo="vmDiskInfo"
            :cpuused="cpuused"
          />
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
      actionFrom: ref(""),
    };
  },
  data() {
    return {
      vmDbDataInfo: ref([]),
      vmMoldDataInfo: ref([]),
      vmNetworkInfo: ref([]),
      vmDiskInfo: ref([]),
      cpuused: ref(0),
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    refresh() {
      this.fetchData();
      this.$refs.listRefreshCall.fetchRefresh();
    },
    async fetchData() {
      try {
        const response = await worksApi.get(
          "/api/v1/instance/detail/" + this.$route.params.vmUuid
        );

        if (response.status === 200) {
          //console.log(response.data.result.instanceDBInfo);
          this.vmDbDataInfo = response.data.result.instanceDBInfo;
          this.vmMoldDataInfo =
            response.data.result.instanceMoldInfo.virtualmachine[0];
          this.vmNetworkInfo = this.vmMoldDataInfo.nic[0];
          this.vmDiskInfo =
            response.data.result.instanceInstanceVolumeInfo.volume[0];
          this.cpuused =
            response.data.result.instanceMoldInfo.virtualmachine[0].cpuused.split(
              "%"
            )[0];
        } else {
          message.error(this.$t("message.response.data.fail"));
        }
      } catch (error) {
        console.log(error);
        message.error(this.$t("message.response.data.fail"));
      }
      this.actionFrom = "VirtualMachineDetail";
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
