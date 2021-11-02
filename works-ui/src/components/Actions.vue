<template class="able-action">
  <a-space :size="8">
    <!--Start-->
    <a-tooltip v-if="state.buttonBoolean.vmStart" placement="bottom">
      <template #title>{{ $t("tooltip.vmStart") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('vmStart')">
        <CaretRightOutlined />
      </a-button>
    </a-tooltip>
    <!--Stop-->
    <a-tooltip v-if="state.buttonBoolean.vmStop" placement="bottom">
      <template #title>{{ $t("tooltip.vmStop") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('vmStop')">
        <PoweroffOutlined />
      </a-button>
    </a-tooltip>
    <!--reset -->
    <a-tooltip v-if="state.buttonBoolean.userAllocate" placement="bottom">
      <template #title>{{ $t("tooltip.userAllocate") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('userAllocate')">
        <UserAddOutlined />
      </a-button>
    </a-tooltip>
    <a-tooltip v-if="state.buttonBoolean.userUnlock" placement="bottom">
      <template #title>{{ $t("tooltip.userUnlock") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('userUnlock')">
        <UserDeleteOutlined />
      </a-button>
    </a-tooltip>
    <!--workspaceDestroy-->
    <a-tooltip v-if="state.buttonBoolean.workspaceDestroy" placement="bottom">
      <template #title>{{ $t("tooltip.destroy") }}</template>
      <a-button
        type="primary"
        shape="circle"
        danger
        @click="setCircleButtonModal('workspaceDestroy')"
      >
        <DeleteFilled />
      </a-button>
    </a-tooltip>
    <!--vmDestroy-->
    <a-tooltip v-if="state.buttonBoolean.vmDestroy" placement="bottom">
      <template #title>{{ $t("tooltip.destroy") }}</template>
      <a-button
        type="primary"
        shape="circle"
        danger
        @click="setCircleButtonModal('vmDestroy')"
      >
        <DeleteFilled />
      </a-button>
    </a-tooltip>
    <!--accountDestroy-->
    <a-tooltip v-if="state.buttonBoolean.accountDestroy" placement="bottom">
      <template #title>{{ $t("tooltip.destroy") }}</template>
      <a-button
        type="primary"
        shape="circle"
        danger
        @click="setCircleButtonModal('accountDestroy')"
      >
        <DeleteFilled />
      </a-button>
    </a-tooltip>

    <!-- Confirm Modal -->
    <a-modal
      v-model:visible="confirmModalView"
      :title="$t('tooltip.' + modalTitle)"
      :ok-text="$t('label.ok')"
      :cancel-text="$t('label.cancel')"
      @cancel="handleCancel"
      @ok="handleSubmit(actionFrom)"
    >
      <p>{{ $t(modalConfirm) }}</p>
    </a-modal>

    <a-modal
      v-model:visible="userAllocateVmModalBoolean"
      :title="$t('tooltip.desktop.allocate.user')"
      width="400px"
      :ok-text="$t('label.ok')"
      :cancel-text="$t('label.cancel')"
      @ok="putUserAllocateVm()"
    >
      <a-select
        v-model:value="selectedUser"
        show-search
        style="width: 100%; margin-top: 7px"
        option-filter-prop="label"
        class="addmodal-aform-item-div"
      >
        <a-select-option
          v-for="option in workspaceUserDataList"
          :key="option.name"
          :value="option.name"
          :label="option.name"
        >
          {{ option.name }}
        </a-select-option>
      </a-select>
    </a-modal>
  </a-space>
</template>

<script>
import { defineComponent, onMounted, reactive, ref } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";
import router from "@/router";

export default defineComponent({
  components: {},
  props: {
    actionFrom: {
      type: String,
      requires: true,
      default: "",
    },
    workspaceUuid: {
      type: String,
      requires: false,
      default: "",
    },
    vmUuid: {
      type: String,
      requires: false,
      default: "",
    },
    allocateStatus: {
      type: String,
      requires: false,
      default: "",
    },
  },
  emits: ["fetchData"],
  setup(props) {
    //console.log("==================== props.actionFrom ====================:::: "+props.actionFrom);
    const state = reactive({
      callComponent: ref(props.actionFrom),
      workspaceUuid: ref(props.workspaceUuid),
      allocateStatus: ref(props.allocateStatus),
      workspaceName: ref(""),
      vmUuid: ref(props.vmUuid),
      vmStatus: ref(""),
      buttonBoolean: {
        showModal: ref(false),
        vmStart: ref(false),
        vmStop: ref(false),
        userAllocate: ref(false),
        userUnlock: ref(false),
        reinstall: ref(false),
        snapshot: ref(false),
        volsnapshot: ref(false),
        iso: ref(false),
        workspaceDestroy: ref(false),
        vmDestroy: ref(false),
        accountDestroy: ref(false),
        edit: ref(false),
        pause: ref(false),
      },
    });
    return {
      confirmModalView: ref(false),
      userAllocateVmModalBoolean: ref(false),
      workspaceUserDataList: ref([]),
      modalTitle: ref(""),
      modalConfirm: ref(""),
      state,
    };
  },
  data() {
    return {
      selectedUser: ref(""),
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      if (this.state.callComponent.includes("Workspace")) {
        this.state.buttonBoolean.workspaceDestroy = true;
      } else if (this.state.callComponent.includes("Account")) {
        this.state.buttonBoolean.accountDestroy = true;
      } else if (this.state.callComponent.includes("GroupPolicy")) {
        this.state.buttonBoolean.destroy = true;
      }

      if (this.state.callComponent.includes("VirtualMachine")) {
        let _uuid = this.$route.params.vmUuid === undefined ? this.state.vmUuid : this.$route.params.vmUuid;
        //console.log("this.$route.params.vmUuid :: "+this.$route.params.vmUuid+ " :: this.state.vmUuid :: " +this.state.vmUuid);
        worksApi
          .get("/api/v1/instance/detail/" + _uuid)
          .then((response) => {
            if (response.status == 200) {
              //this.vmDbDataInfo = response.data.result.instanceDBInfo;
              this.state.vmUuid = response.data.result.instanceDBInfo.uuid;
              this.state.vmStatus = response.data.result.instanceDBInfo.mold_status;
              this.state.workspaceName = response.data.result.instanceDBInfo.workspace_name;
              this.state.allocateStatus = response.data.result.instanceDBInfo.owner_account_id;

              if (this.state.vmStatus === "Running") {
                this.state.buttonBoolean.vmStop = true;
                this.state.buttonBoolean.vmStart = false;
              } else {
                this.state.buttonBoolean.vmStop = false;
                this.state.buttonBoolean.vmStart = true;
              }
              if (this.state.allocateStatus == "") {
                this.state.buttonBoolean.userAllocate = true;
                this.state.buttonBoolean.userUnlock = false;
              } else {
                this.state.buttonBoolean.userAllocate = false;
                this.state.buttonBoolean.userUnlock = true;
              }
              this.state.buttonBoolean.vmDestroy = true;

              //해당 워크스페이스에 추가 된 사용자 목록 조회
              worksApi
                .get("/api/v1/group/" + this.state.workspaceName)
                .then((response) => {
                  if (response.status == 200) {
                    const temp =
                      response.data.result.member == undefined
                        ? ""
                        : response.data.result.member;
                    for (let str of temp) {
                      this.workspaceUserDataList.push({ name: str.split(",")[0].split("CN=")[1] });
                    }
                  } else {
                    //message.error(this.t("message.response.data.fail"));
                  }
                })
                .catch(function (error) {
                  //message.error(error);
                });
            } else {
              //console.log("데이터를 정상적으로 가져오지 못했습니다.");
            }
          })
          .catch(function (error) {
            console.log(error);
          });
      }
    },
    putUserAllocateVm() {
      let params = new URLSearchParams();
      params.append("instanceUuid", this.state.vmUuid);
      params.append("username", this.selectedUser);
      worksApi
        .post("/api/v1/instance", params)
        .then((response) => {
          if (response.status === 200) {
            message.success(this.$t("message.user.vm.allocated.completed"), 3);
            setTimeout(() => {
              this.$emit("fetchData");
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
    setCircleButtonModal(value) {
      if (value == "userAllocate") {
        this.userAllocateVmModalBoolean = true;
      } else {
        this.confirmModalView = true;
        this.modalTitle = value;
      }
      if (value == "workspaceStart") this.modalConfirm = "modal.confirm.workspaceStart";
      if (value == "workspaceStop") this.modalConfirm = "modal.confirm.workspaceStop";
      if (value == "workspaceDestroy") this.modalConfirm = "modal.confirm.workspaceDestroy";

      if (value == "vmStart") this.modalConfirm = "modal.confirm.vmStart";
      if (value == "vmStop") this.modalConfirm = "modal.confirm.vmStop";
      if (value == "vmDestroy") this.modalConfirm = "modal.confirm.vmDestroy";

      if (value == "accountDestroy") this.modalConfirm = "modal.confirm.accountDestroy";
      if (value == "userUnlock") this.modalConfirm = "modal.confirm.userUnlock";
    },
    handleCancel() {
      this.confirmModalView = false;
      this.userAllocateVmModalBoolean = false;
    },
    handleSubmit(actionFrom) {
      //console.log(this.modalTitle + "  ::  " + this.state.vmUuid);
      if (actionFrom.includes("VirtualMachine")) {
        let worksUrl, resMessage = "";
        if (this.modalTitle.includes("vmStart")) {
          message.loading(this.$t("message.vm.status.starting"), 12);
          worksUrl = "/api/v1/instance/VMStart/" + this.state.vmUuid;
          resMessage = this.$t("message.vm.status.update");
        }
        if (this.modalTitle.includes("vmStop")) {
          message.loading(this.$t("message.vm.status.stopping"), 12);
          worksUrl = "/api/v1/instance/VMStop/" + this.state.vmUuid;
          resMessage = this.$t("message.vm.status.update");
        }
        if (this.modalTitle.includes("vmDestroy")) {
          message.loading(this.$t("message.vm.status.destroying"), 12);
          worksUrl = "/api/v1/instance/VMDestroy/" + this.state.vmUuid;
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
                if (actionFrom =="VirtualMachineDetail" && this.modalTitle.includes("vmDestroy")){
                  router.push({ name: "VirtualMachine" });
                } else {
                  this.$emit("fetchData");
                  this.fetchData();
                }
              }, 12000);
            } else {
              message.error(this.$t("message.vm.status.update.fail"));
            }
          })
          .catch(function (error) {
            message.error(error);
          });
      }

      if (actionFrom.includes("Workspace")) {
        worksApi
          .get("/api/v1/workspace/" + this.state.workspaceUuid)
          .then((response) => {
            if (response.status == 200) {
              if(response.data.result.workspaceInfo.quantity == 0){
                message.loading(this.$t("message.workspace.status.destroying"), 6);
                worksApi
                  .delete("/api/v1/workspace/" + this.state.workspaceUuid)
                  .then((response) => {
                    if (response.status == 200) {
                      this.vmDataList = response.data.result.list;
                      this.handleCancel();
                      setTimeout(() => {
                        message.destroy();
                        message.success(this.$t("message.workspace.status.delete"));
                        if(actionFrom =="WorkspaceDetail") router.push({ name: "Workspace" });
                        if(actionFrom =="WorkspaceList") this.$emit("fetchData");
                      }, 3000);
                    } else {
                      message.error(this.$t("message.workspace.delete.fail"), 5);
                    }
                  })
                  .catch(function (error) {
                    message.error(error);
                  });
              } else {
                message.error(this.$t("message.workspace.delete.existvm"));
              }
            }
          })
          .catch(function (error) {
            message.error(error);
          });
      }
    },
  },
});
</script>

<style scoped>
#content-action .ant-btn {
  padding-top: 0px;
  padding-bottom: 0px;
  font-size: 14px;
}
</style>
