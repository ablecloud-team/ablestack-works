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
            </a-col>
            <!-- 우측 액션 -->
            <a-col id="content-action" :span="12">
              <Actions
                :action-from="actionFrom"
                :vm-uuid="uuid"
                @fetchData="fetchData"
              />
            </a-col>
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
          <VirtualMachineBody
            :vm-db-data-info="vmDbDataInfo"
            :vm-mold-data-info="vmMoldDataInfo"
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
      vmDbDataInfo: ref([]),
      vmMoldDataInfo: ref([]),
      actionFrom: ref("VirtualMachineDetail"),
      modalConfirm: ref(""),
      modalTitle: ref(""),
      confirmModalView: ref(false),
      userAllocateVmModalBoolean: ref(false),
      workspaceUserDataList: ref([]),
      selectedUser: ref(""),
    };
  },
  data() {
    return {};
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      // 가상머신 상세조회
      worksApi
        .get("/api/v1/instance/detail/" + this.$route.params.vmUuid)
        .then((response) => {
          if (response.status === 200) {
            this.vmDbDataInfo = response.data.result.instanceDBInfo;
            this.vmMoldDataInfo =
              response.data.result.instanceMoldInfo.virtualmachine[0];

            //해당 워크스페이스에 추가 된 사용자 목록 조회
            worksApi
              .get("/api/v1/group/" + this.vmDbDataInfo.workspace_name)
              .then((response) => {
                if (response.status == 200) {
                  //console.log(response.data.result.member);
                  const temp =
                    response.data.result.member == undefined
                      ? ""
                      : response.data.result.member;
                  for (let str of temp) {
                    this.workspaceUserDataList.push({
                      name: str.split(",")[0].split("CN=")[1],
                    });
                  }
                } else {
                }
              })
              .catch(function (error) {
                message.error(error);
              });
          } else {
            message.error(this.$t("message.response.data.fail"));
            //console.log("데이터를 정상적으로 가져오지 못했습니다.");
          }
        })
        .catch(function (error) {
          console.log(error);
        });
    },
    setCircleButtonModal: function (value) {
      if (value == "userAllocate") {
        this.userAllocateVmModalBoolean = true;
      } else {
        this.confirmModalView = true;
        this.modalTitle = value;
        if (value == "vmStart") this.modalConfirm = "modal.confirm.start";
        if (value == "vmStop") this.modalConfirm = "modal.confirm.stop";
        if (value == "vmDestroy") this.modalConfirm = "modal.confirm.vmDestroy";
        if (value == "userUnlock") this.userUnlock = "modal.confirm.userUnlock";
      }
    },
    handleCancel: function () {
      this.confirmModalView = false;
      this.userAllocateVmModalBoolean = false;
    },
    handleSubmit: function () {
      //console.log(this.modalTitle + "  ::  " + this.$route.params.vmUuid);
      let worksUrl,
        resMessage = "";
      if (this.modalTitle.includes("vmStart")) {
        message.loading(this.$t("message.vm.status.starting"), 6);
        worksUrl = "/api/v1/instance/VMStart/" + this.$route.params.vmUuid;
        resMessage = this.$t("message.vm.status.update");
      }
      if (this.modalTitle.includes("vmStop")) {
        message.loading(this.$t("message.vm.status.stopping"), 6);
        worksUrl = "/api/v1/instance/VMStop/" + this.$route.params.vmUuid;
        resMessage = this.$t("message.vm.status.update");
      }
      if (this.modalTitle.includes("vmDestroy")) {
        message.loading(this.$t("message.vm.status.destroying"), 6);
        worksUrl = "/api/v1/instance/VMDestroy/" + this.$route.params.vmUuid;
        resMessage = this.$t("message.vm.status.delete");
      }
      worksApi
        .patch(worksUrl)
        .then((response) => {
          if (response.status == 200) {
            this.vmDataList = response.data.result.list;
            this.handleCancel();
            setTimeout(() => {
              message.destroy();
              message.success(resMessage);
              if (this.modalTitle.includes("vmDestroy")) {
                this.$router.push({ name: "VirtualMachine" });
              } else {
                this.fetchData();
              }
            }, 5000);
          } else {
            message.error(this.$t("message.vm.status.update.fail"));
          }
        })
        .catch(function (error) {
          message.error(error);
        });
    },
    putUserAllocateVm: function () {
      //console.log(this.selectedUser + "  ::  " + uuid + "  :: " + workspace);
      let params = new URLSearchParams();
      params.append("instanceUuid", this.$route.params.vmUuid);
      params.append("username", this.selectedUser);
      worksApi
        .post("/api/v1/instance", params)
        .then((response) => {
          if (response.status === 200) {
            message.success(this.$t("message.user.vm.allocated.completed"), 3);
            setTimeout(() => {
              this.fetchData();
            }, 1000);
            this.handleCancel();
          } else {
            message.error("message.user.vm.allocated.fail");
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
