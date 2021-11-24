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
    <a-tooltip
      v-if="state.buttonBoolean.workspaceAccountDestroy"
      placement="bottom"
    >
      <template #title>{{ $t("tooltip.destroy") }}</template>
      <a-button
        type="primary"
        shape="circle"
        danger
        @click="setCircleButtonModal('workspaceAccountDestroy')"
      >
        <DeleteFilled />
      </a-button>
    </a-tooltip>

    <a-modal
      v-model:visible="commonModalView"
      :title="$t('tooltip.' + modalTitle)"
      :ok-text="$t('label.ok')"
      :cancel-text="$t('label.cancel')"
      @cancel="handleCancel"
      @ok="handleSubmit(actionFrom)"
    >
      <a-alert :message="modalConfirm" :type="alertType" show-icon />
      <br />
      <a-table
        size="small"
        :columns="listColumns"
        :pagination="{ pageSize: 10 }"
        :data-source="eventList"
      >
      </a-table>
    </a-modal>

    <a-modal
      v-model:visible="userAllocateVmModalView"
      :title="$t('tooltip.desktop.allocate.user')"
      :ok-text="$t('label.ok')"
      :cancel-text="$t('label.cancel')"
      @cancel="handleCancel"
      @ok="handleSubmit(actionFrom)"
    >
      <a-alert :message="modalConfirm" type="info" show-icon />
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
      <br /><br />
      <a-table
        size="small"
        :columns="listColumns"
        :pagination="{ pageSize: 10 }"
        :data-source="eventList"
      >
      </a-table>
    </a-modal>
  </a-space>
</template>

<script>
import { defineComponent, reactive, ref } from "vue";
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
    vmInfo: {
      type: Object,
      requires: false,
      default: null,
    },
    workspaceInfo: {
      type: Object,
      requires: false,
      default: null,
    },
    accountInfo: {
      type: Object,
      requires: false,
      default: null,
    },
    multiSelectList: {
      type: Object,
      requires: false,
      default: null,
    },
    wsName: {
      type: String,
      requires: false,
      default: "",
    },
  },
  emits: ["fetchData"],
  setup() {
    const state = reactive({
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
        workspaceAccountDestroy: ref(false),
        vmDestroy: ref(false),
        accountDestroy: ref(false),
        edit: ref(false),
        pause: ref(false),
      },
    });
    return {
      userAllocateVmModalView: ref(false),
      commonModalView: ref(false),
      workspaceUserDataList: ref([]),
      modalTitle: ref(""),
      modalConfirm: ref(""),
      state,
    };
  },
  data(props) {
    return {
      selectedUser: ref(""),
      callComponent: ref(props.actionFrom),
      multiSelectList: props.multiSelectList,
      vmInfo: ref(props.vmInfo),
      accountInfo: ref(props.accountInfo),
      workspaceInfo: ref(props.workspaceInfo),
      wsName: ref(props.wsName),
      eventList: [],
      alertType: ref("info"),
      succCnt: ref(0),
      failCnt: ref(0),
      listColumns: [
        {
          title: this.$t("label.name"),
          dataIndex: "name",
          key: "name",
          width: "100%",
        },
      ],
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      if (this.callComponent.includes("Workspace")) {
        if (this.callComponent === "WorkspaceUserList") {
          this.state.buttonBoolean.workspaceAccountDestroy = true;
        } else {
          this.state.buttonBoolean.workspaceDestroy = true;
        }
        if (this.workspaceInfo) {
          this.eventList = [this.workspaceInfo];
        } else {
          this.eventList = this.multiSelectList;
        }
      } else if (this.callComponent.includes("Account")) {
        this.state.buttonBoolean.accountDestroy = true;

        if (this.accountInfo) {
          this.eventList = [this.accountInfo];
        } else {
          this.eventList = this.multiSelectList;
        }
      } else if (this.callComponent.includes("GroupPolicy")) {
        this.state.buttonBoolean.destroy = true;
      }
      if (this.callComponent.includes("VirtualMachine")) {
        this.state.buttonBoolean.vmDestroy = true;

        if (this.vmInfo) {
          this.eventList = [this.vmInfo];
        } else {
          this.eventList = this.multiSelectList;
        }

        // console.log(this.eventList);

        let res = null;
        res = this.eventList.filter(
          (it) => it.mold_status === "Running" || it.mold_status === ""
        ); //시작 버튼 체크
        // console.log(res.length);
        if (res.length === 0) this.state.buttonBoolean.vmStart = true;

        res = this.eventList.filter(
          (it) => it.mold_status === "Stopped" || it.mold_status === ""
        ); //정지 버튼 체크
        // console.log(res.length);
        if (res.length === 0) this.state.buttonBoolean.vmStop = true;

        res = this.eventList.filter((it) => it.owner_account_id === ""); //사용자 할당 해지 버튼 체크
        // console.log(res.length);
        if (res.length === 0) this.state.buttonBoolean.userUnlock = true;

        res = this.eventList.filter((it) => it.owner_account_id !== ""); //사용자 할당 버튼 체크
        // console.log(res.length);
        if (res.length === 0) {
          //같은 워크스페이스면 할당 버튼 활성화, 아니면 비활성화
          res = this.eventList.filter(function (obj, i, s) {
            return (
              i ===
              s.findIndex(function (t) {
                return t.workspace_name === obj.workspace_name;
              })
            );
          });

          if (res.length === 1) {
            this.state.buttonBoolean.userAllocate = true; //사용자 할당 버튼

            worksApi
              .get("/api/v1/group/" + res[0].workspace_name)
              .then((response) => {
                if (response.status == 200) {
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
                  message.error(this.$t("message.response.data.fail"));
                }
              })
              .catch((error) => {
                message.destroy();
                message.error(this.$t("message.response.data.fail"));
                console.log(error);
              });
          }
        }
      }
    },
    setCircleButtonModal(value) {
      this.modalTitle = value;
      if (value == "userAllocate") {
        this.userAllocateVmModalView = true;
        this.modalConfirm = this.$t(
          "modal.confirm.workspace.allocate.vm.user",
          {
            count: this.eventList.length,
          }
        );
      } else {
        this.commonModalView = true;
      }
      // if (value == "workspaceStart")
      //   this.modalConfirm =
      //     "[" +
      //     this.workspaceName +
      //     "] " +
      //     this.$t("modal.confirm.workspaceStart");
      // if (value == "workspaceStop")
      //   this.modalConfirm =
      //     "[" +
      //     this.workspaceName +
      //     "] " +
      //     this.$t("modal.confirm.workspaceStop");
      if (value == "workspaceDestroy") {
        this.modalConfirm = this.$t("modal.confirm.workspaceDestroy", {
          count: this.eventList.length,
        });
        this.alertType = "warning";
      }
      if (value == "workspaceAccountDestroy") {
        this.modalConfirm = this.$t("modal.confirm.workspaceAccountDestroy", {
          count: this.eventList.length,
        });
        this.alertType = "warning";
      }

      if (value == "vmStart") {
        this.modalConfirm = this.$t("modal.confirm.vmStart", {
          count: this.eventList.length,
        });
      }
      if (value == "vmStop") {
        this.modalConfirm = this.$t("modal.confirm.vmStop", {
          count: this.eventList.length,
        });
      }
      if (value == "vmDestroy") {
        this.modalConfirm = this.$t("modal.confirm.vmDestroy", {
          count: this.eventList.length,
        });
        this.alertType = "warning";
      }
      if (value == "userUnlock") {
        this.modalConfirm = this.$t("modal.confirm.userUnlock", {
          count: this.eventList.length,
        });
      }
      if (value == "accountDestroy") {
        this.modalConfirm = this.$t("modal.confirm.accountDestroy", {
          count: this.eventList.length,
        });
        this.alertType = "warning";
      }
    },
    handleCancel() {
      this.commonModalView = false;
      this.userAllocateVmModalView = false;
    },
    handleSubmit(actionFrom) {
      //console.log(this.modalTitle + "  ::  " + actionFrom);
      if (actionFrom.includes("VirtualMachine")) {
        if (this.modalTitle.includes("vm")) this.vmAction(actionFrom);

        if (this.modalTitle.includes("userUnlock")) this.vmUserUnlockAction();

        if (this.modalTitle.includes("userAllocate"))
          this.vmUserAllocateAction(actionFrom);
      }

      if (actionFrom.includes("Workspace")) {
        if (this.modalTitle.includes("workspaceDestroy"))
          this.workspaceDestroyAction(actionFrom);

        if (this.modalTitle.includes("workspaceAccountDestroy"))
          this.workspaceAccountDestroyAction(actionFrom);
      }

      if (actionFrom.includes("Account")) {
        if (this.modalTitle.includes("accountDestroy"))
          this.accountDestroyAction(actionFrom);
      }
    },
    async workspaceAccountDestroyAction(actionFrom) {
      let sucMessage = "message.workspace.user.delete.ok";
      let failMessage = "message.workspace.user.delete.fail";

      message.loading(this.$t("message.workspace.vm.user.deleting"), 100);

      //console.log(this.eventList);
      for (let val of this.eventList) {
        try {
          const response = await worksApi.delete(
            "/api/v1/group/" + this.wsName + "/" + val.name
          );
          console.log(response.status);
          if (response.status == 200) {
            this.succCnt = this.succCnt + 1;
          }
        } catch (error) {
          console.log(error);
          this.failCnt = this.failCnt + 1;
        }
      }
      console.log(this.succCnt, this.failCnt);
      this.handleCancel();
      setTimeout(() => {
        this.$emit("fetchData");

        message.destroy();

        if (this.succCnt > 0) {
          message.success(
            this.$t(sucMessage, {
              count: this.succCnt,
            }),
            5
          );
        }
        if (this.failCnt > 0) {
          message.error(
            this.$t(failMessage, {
              count: this.failCnt,
            }),
            5
          );
        }
        this.failCnt = 0;
        this.succCnt = 0;
      }, 2000);
    },
    async vmUserAllocateAction(actionFrom) {
      let sucMessage = "message.user.vm.allocated.ok";
      let failMessage = "message.user.vm.allocated.fail";
      message.loading(this.$t("message.user.vm.allocating"), 100);

      for (let val of this.eventList) {
        try {
          const res = await worksApi.put(
            "/api/v1/connection/" + val.uuid + "/" + this.selectedUser
          );
          if (res.status == 200) {
            this.succCnt = this.succCnt + 1;
          }
        } catch (error) {
          console.log(error);
          this.failCnt = this.failCnt + 1;
        }
      }
      this.handleCancel();
      setTimeout(() => {
        this.$emit("fetchData");

        message.destroy();
        if (this.succCnt > 0) {
          message.success(
            this.$t(sucMessage, {
              count: this.succCnt,
            }),
            5
          );
        }
        if (this.failCnt > 0) {
          message.error(
            this.$t(failMessage, {
              count: this.failCnt,
            }),
            5
          );
        }

        this.failCnt = 0;
        this.succCnt = 0;
      }, 3000);
    },
    async vmUserUnlockAction() {
      let sucMessage = "message.user.vm.unlock.ok";
      let failMessage = "message.user.vm.unlock.fail";
      message.loading(this.$t("message.user.vm.unlocking"), 100);

      for (let val of this.eventList) {
        try {
          const res = await worksApi.delete("/api/v1/connection/" + val.uuid);
          if (res.status == 200) {
            this.succCnt = this.succCnt + 1;
          }
        } catch (error) {
          console.log(error);
          this.failCnt = this.failCnt + 1;
        }
      }

      this.handleCancel();
      this.$emit("fetchData");

      setTimeout(() => {
        message.destroy();
        if (this.succCnt > 0) {
          message.success(
            this.$t(sucMessage, {
              count: this.succCnt,
            }),
            5
          );
        }
        if (this.failCnt > 0) {
          message.error(
            this.$t(failMessage, {
              count: this.failCnt,
            }),
            5
          );
        }

        this.failCnt = 0;
        this.succCnt = 0;
      }, 3000);
    },
    async vmAction(actionFrom) {
      let worksUrl,
        sucMessage,
        failMessage = "";
      if (this.modalTitle.includes("vmStart")) {
        message.loading(this.$t("message.vm.status.starting"), 100);
        worksUrl = "/api/v1/instance/VMStart/";
        sucMessage = "message.vm.status.start.ok";
        failMessage = "message.vm.status.start.fail";
      }
      if (this.modalTitle.includes("vmStop")) {
        message.loading(this.$t("message.vm.status.stopping"), 100);
        worksUrl = "/api/v1/instance/VMStop/";
        sucMessage = "message.vm.status.stop.ok";
        failMessage = "message.vm.status.stop.fail";
      }
      if (this.modalTitle.includes("vmDestroy")) {
        message.loading(this.$t("message.vm.status.destroying"), 100);
        worksUrl = "/api/v1/instance/VMDestroy/";
        sucMessage = "message.vm.status.delete.ok";
        failMessage = "message.vm.status.delete.fail";
      }
      for (let val of this.eventList) {
        try {
          const res = await worksApi.patch(worksUrl + val.uuid);
          if (res.status == 200) {
            this.succCnt = this.succCnt + 1;
          }
        } catch (error) {
          console.log(error);
          this.failCnt = this.failCnt + 1;
        }
      }

      this.handleCancel();
      setTimeout(() => {
        if (
          actionFrom == "VirtualMachineDetail" &&
          this.modalTitle.includes("vmDestroy")
        ) {
          router.push({ name: "VirtualMachine" });
        } else {
          this.$emit("fetchData");
        }

        message.destroy();
        if (this.succCnt > 0) {
          message.success(
            this.$t(sucMessage, {
              count: this.succCnt,
            })
          );
        }
        if (this.failCnt > 0) {
          message.error(
            this.$t(failMessage, {
              count: this.failCnt,
            }),
            5
          );
        }
        this.failCnt = 0;
        this.succCnt = 0;
      }, 10000);
    },
    async workspaceDestroyAction(actionFrom) {
      let sucMessage = "message.workspace.status.delete";
      let failMessage = "message.workspace.delete.existvm";
      message.loading(this.$t("message.workspace.status.destroying"));

      for (let val of this.eventList) {
        if (val.quantity === 0) {
          const res = await worksApi.delete("/api/v1/workspace/" + val.uuid);
          if (res.status == 200) {
            this.succCnt = this.succCnt + 1;
          }
        } else {
          this.failCnt = this.failCnt + 1;
        }
      }

      this.handleCancel();
      if (actionFrom == "WorkspaceDetail") {
        router.push({ name: "Workspace" });
      } else {
        this.$emit("fetchData");
      }
      setTimeout(() => {
        message.destroy();
        if (this.succCnt > 0) {
          message.success(
            this.$t(sucMessage, {
              count: this.succCnt,
            }),
            5
          );
        }
        if (this.failCnt > 0) {
          message.error(
            this.$t(failMessage, {
              count: this.failCnt,
            }),
            5
          );
        }
        this.failCnt = 0;
        this.succCnt = 0;
      }, 2000);
    },
    async accountDestroyAction(actionFrom) {
      let sucMessage = "message.account.destroy.ok";
      let failMessage = "message.account.destroy.fail";
      message.loading(this.$t("message.account.destroying"));

      for (let val of this.eventList) {
        try {
          const res = await worksApi.delete("/api/v1/user/" + val.name);
          if (res.status == 204) {
            this.succCnt = this.succCnt + 1;
          }
        } catch (error) {
          console.log(error);
          this.failCnt = this.failCnt + 1;
        }
      }

      this.handleCancel();
      if (actionFrom == "AccountDetail") {
        router.push({ name: "Account" });
      } else {
        this.$emit("fetchData");
      }
      setTimeout(() => {
        message.destroy();
        if (this.succCnt > 0) {
          message.success(
            this.$t(sucMessage, {
              count: this.succCnt,
            }),
            5
          );
        }
        if (this.failCnt > 0) {
          message.error(
            this.$t(failMessage, {
              count: this.failCnt,
            }),
            5
          );
        }
        this.failCnt = 0;
        this.succCnt = 0;
      }, 1000);
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
