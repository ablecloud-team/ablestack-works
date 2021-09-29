<template class="able-action">
  <a-space :size="size">
    <!--Start-->
    <a-tooltip v-if="state.buttonBoolean.start" placement="top">
      <template #title>{{ $t("tooltip.start") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('start')">
        <CaretRightOutlined />
      </a-button>
    </a-tooltip>
    <!--Stop-->
    <a-tooltip v-if="state.buttonBoolean.stop" placement="top">
      <template #title>{{ $t("tooltip.stop") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('stop')">
        <PoweroffOutlined />
      </a-button>
    </a-tooltip>
    <!--reset -->
    <a-tooltip v-if="state.buttonBoolean.userAllocate" placement="top">
      <template #title>{{ $t("tooltip.userAllocate") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('userAllocate')">
        <UserAddOutlined />
      </a-button>
    </a-tooltip>
    <a-tooltip v-if="state.buttonBoolean.userUnlock" placement="top">
      <template #title>{{ $t("tooltip.userUnlock") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('userUnlock')">
        <UserDeleteOutlined />
      </a-button>
    </a-tooltip>
    <!--reinstall-->
    <a-tooltip v-if="state.buttonBoolean.reinstall" placement="top">
      <template #title>{{ $t("tooltip.reinstall") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('reinstall')">
        <SyncOutlined />
      </a-button>
    </a-tooltip>
    <!--snapshot-->
    <a-tooltip v-if="state.buttonBoolean.snapshot" placement="top">
      <template #title>{{ $t("tooltip.snapshot") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('snapshot')">
        <CameraOutlined />
      </a-button>
    </a-tooltip>
    <!--volsnapshot-->
    <a-tooltip v-if="state.buttonBoolean.volsnapshot" placement="top">
      <template #title>{{ $t("tooltip.volsnapshot") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('volsnapshot')">
        <i class="fas fa-camera-retro"></i>
        <VideoCameraAddOutlined />
      </a-button>
    </a-tooltip>
    <!--isoattach-->
    <a-tooltip v-if="state.buttonBoolean.isoattach" placement="top">
      <template #title>{{ $t("tooltip.isoattach") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('isoattach')">
        <PaperClipOutlined />
      </a-button>
    </a-tooltip>
    <!--edit-->
    <a-tooltip v-if="state.buttonBoolean.edit" placement="top">
      <template #title>{{ $t("tooltip.edit") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('edit')">
        <EditOutlined />
      </a-button>
    </a-tooltip>
    <!--pause-->
    <a-tooltip v-if="state.buttonBoolean.pause" placement="top">
      <template #title>{{ $t("tooltip.pause") }}</template>
      <a-button shape="circle" @click="asdf()">
        <PauseOutlined />
      </a-button>
    </a-tooltip>
    <!--workspaceDestroy-->
    <a-tooltip v-if="state.buttonBoolean.workspaceDestroy" placement="top">
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
    <a-tooltip v-if="state.buttonBoolean.vmDestroy" placement="top">
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
    <a-tooltip v-if="state.buttonBoolean.accountDestroy" placement="top">
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
      :title="$t('tooltip.' + modalTitle)"
      :visible="confirmModalView"
      :ok-text="$t('label.ok')"
      :cancel-text="$t('label.cancel')"
      @cancel="handleCancel"
      @ok="handleSubmit(actionFrom, workspace, uuid)"
    >
      <p>{{ $t(modalConfirm) }}</p>
    </a-modal>

    <a-modal
      :title="$t('tooltip.desktop.allocate.user')"
      v-model:visible="userAllocateVmModalBoolean"
      width="400px"
      :ok-text="$t('label.ok')"
      :cancel-text="$t('label.cancel')"
      @ok="putUserAllocateVm(workspace, uuid)"
    >
      <a-select
        v-model:value="selectedUser"
        :placeholder="$t('tooltip.user')"
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

export default defineComponent({
  components: {},
  props: {
    actionFrom: {
      type: String,
      requires: true,
      default: "",
    },
    uuid: {
      type: String,
      requires: false,
      default: "",
    },
    workspace: {
      type: String,
      requires: false,
      default: "",
    },
    status: {
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
  setup(props) {
    const sizeValue = 8;
    //console.log("==================== props.actionFrom ====================:::: "+props.actionFrom);
    const state = reactive({
      callComponent: ref(props.actionFrom),
      workspace: ref(props.workspace),
      uuid: ref(props.uuid),
      allocateStatus: ref(props.allocateStatus),
      status: ref(props.status),
      buttonBoolean: {
        showModal: ref(false),
        start: ref(false),
        stop: ref(false),
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
    onMounted(() => {
      if (state.callComponent === "WorkspaceList" || state.callComponent === "WorkspaceDetail") {
        state.buttonBoolean.start = true;
        state.buttonBoolean.stop = true;
        state.buttonBoolean.workspaceDestroy = true;
      } else if (state.callComponent === "VirtualMachineList" || state.callComponent === "VirtualMachineDetail") {
        if (state.status === "Running") {
          state.buttonBoolean.stop = true;
        } else {
          state.buttonBoolean.start = true;
        }
        if (state.allocateStatus == "") {
          state.buttonBoolean.userAllocate = true;
        } else {
          state.buttonBoolean.userUnlock = true;
        }
        state.buttonBoolean.vmDestroy = true;
      } else if (state.callComponent === "AccountList" || state.callComponent === "AccountDetail") {
        state.buttonBoolean.accountDestroy = true;
      } else if (state.callComponent === "GroupPolicyList" || state.callComponent === "GroupPolicyDetail") {
        state.buttonBoolean.destroy = true;
      }
    });
    return {
      size: ref(sizeValue),
      confirmModalView: ref(false),
      userAllocateVmModalBoolean: ref(false),
      workspaceUserDataList: ref([]),
      modalTitle: ref(""),
      modalConfirm: ref(""),
      popView: ref(true),
      state,
    };
  },
  data() {
    return {
      selectedUser: ref(""),
    }
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData: function () {
      if (this.state.callComponent === "VirtualMachineList" ) {
        //해당 워크스페이스에 추가 된 사용자 목록 조회
        worksApi
          .get("/api/v1/group/" + this.state.workspace)
          .then((response) => {
            if (response.status == 200) {
              //console.log(response.data.result.member);
              const temp =
                response.data.result.member == undefined
                  ? ""
                  : response.data.result.member;
              for (let str of temp) {
                this.workspaceUserDataList.push({ name: str.split(",")[0].split("CN=")[1] });
              }
            } else {
              //message.error(this.t("message.response.data.fail"));
              //console.log(response.message);
            }
          })
          .catch(function (error) {
            message.error(error);
            //console.log(error);
          });
      }
      
    },
    putUserAllocateVm: function (workspace, uuid) {
      //console.log(this.selectedUser + "  ::  " + uuid + "  :: " + workspace);
      let params = new URLSearchParams();
      params.append("uuid", uuid);
      params.append("username", this.selectedUser);
      worksApi
        .post("/api/v1/instance", params)
        .then((response) => {
          if (response.status === 200) {
            message.success(this.$t("message.user.vm.allocated.completed"), 3);
            setTimeout(() => {
              this.$emit("fetchData", "desktop");
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
    setCircleButtonModal: function (value) {
      if (value == "userAllocate") {
        this.userAllocateVmModalBoolean = true;
      } else {
        this.confirmModalView = true;
        this.popView = false;
        this.modalTitle = value;
      }
      if (value == "start") this.modalConfirm = "modal.confirm.start";
      if (value == "stop") this.modalConfirm = "modal.confirm.stop";
      if (value == "workspaceDestroy") this.modalConfirm = "modal.confirm.workspaceDestroy";
      if (value == "vmDestroy") this.modalConfirm = "modal.confirm.vmDestroy";
      if (value == "accountDestroy") this.modalConfirm = "modal.confirm.accountDestroy";
      if (value == "userUnlock") this.userUnlock = "modal.confirm.userUnlock";
    },
    handleCancel: function () {
      this.confirmModalView = false;
      this.userAllocateVmModalBoolean = false;
    },
    handleSubmit: function (actionFrom, workspace, uuid) {
      //console.log(this.modalTitle + "  ::  " + actionFrom + " ::  " + workspace + " :: " + uuid);
      
      if (actionFrom == "VirtualMachineList") {
        let worksUrl, resMessage = "";
        if (this.modalTitle.includes("start")) {
          message.loading(this.$t("message.vm.status.starting"), 6);
          worksUrl = "/api/v1/instance/VMStart/" + uuid;
          resMessage = this.$t("message.vm.status.update");
        }
        if (this.modalTitle.includes("stop")) {
          message.loading(this.$t("message.vm.status.stopping"), 6);
          worksUrl = "/api/v1/instance/VMStop/" + uuid;
          resMessage = this.$t("message.vm.status.update");
        }
        if (this.modalTitle.includes("vmDestroy")) {
          message.loading(this.$t("message.vm.status.destroying"), 6);
          worksUrl = "/api/v1/instance/VMDestroy/" + uuid;
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
                this.$emit("fetchData", "desktop");
              }, 6000);
            } else {
              message.error(this.$t("message.vm.status.update.fail"));
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
