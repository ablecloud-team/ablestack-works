<template class="able-action">
  <a-space :size="8">
    <!-- 단일 Select 일때 버튼 표시 -->
    <a-tooltip v-if="state.buttonBoolean.vmStart" placement="bottom">
      <template #title>{{ $t("tooltip.vmStart") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('vmStart')">
        <CaretRightOutlined />
      </a-button>
    </a-tooltip>
    <a-tooltip v-if="state.buttonBoolean.vmStop" placement="bottom">
      <template #title>{{ $t("tooltip.vmStop") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('vmStop')">
        <PoweroffOutlined />
      </a-button>
    </a-tooltip>
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
    <!-- 단일 Select 일때 버튼 표시 -->

    <!-- Multi Select 일때 일괄 버튼 표시 -->
    <a-tooltip v-if="state.buttonBoolean.multiVmStart" placement="bottom">
      <template #title>{{ $t("tooltip.multiVmStart") }}</template>
      <a-button shape="round" @click="setCircleButtonModal('vmStart')">
        <CaretRightOutlined /> {{ $t("tooltip.multiVmStart") }}
      </a-button>
    </a-tooltip>
    <a-tooltip v-if="state.buttonBoolean.multiVmStop" placement="bottom">
      <template #title>{{ $t("tooltip.multiVmStop") }}</template>
      <a-button shape="round" @click="setCircleButtonModal('vmStop')">
        <PoweroffOutlined /> {{ $t("tooltip.multiVmStop") }}
      </a-button>
    </a-tooltip>

    <a-tooltip v-if="state.buttonBoolean.multiUserAllocate" placement="bottom">
      <template #title>{{ $t("tooltip.multiUserAllocate") }}</template>
      <a-button shape="round" @click="setCircleButtonModal('userAllocate')">
        <UserAddOutlined /> {{ $t("tooltip.multiUserAllocate") }}
      </a-button>
    </a-tooltip>
    <a-tooltip v-if="state.buttonBoolean.multiUserUnlock" placement="bottom">
      <template #title>{{ $t("tooltip.multiUserUnlock") }}</template>
      <a-button shape="round" @click="setCircleButtonModal('userUnlock')">
        <UserDeleteOutlined /> {{ $t("tooltip.multiUserUnlock") }}
      </a-button>
    </a-tooltip>
    <a-tooltip v-if="state.buttonBoolean.multiVmDestroy" placement="bottom">
      <template #title>{{ $t("tooltip.multiDestroy") }}</template>
      <a-button shape="round" danger @click="setCircleButtonModal('vmDestroy')">
        <DeleteFilled /> {{ $t("tooltip.multiDestroy") }}
      </a-button>
    </a-tooltip>
    <a-tooltip
      v-if="state.buttonBoolean.multiWorkspaceDestroy"
      placement="bottom"
    >
      <template #title>{{ $t("tooltip.multiDestroy") }}</template>
      <a-button
        shape="round"
        danger
        @click="setCircleButtonModal('workspaceDestroy')"
      >
        <DeleteFilled />{{ $t("tooltip.multiDestroy") }}
      </a-button>
    </a-tooltip>
    <a-tooltip
      v-if="state.buttonBoolean.multiAccountDestroy"
      placement="bottom"
    >
      <template #title>{{ $t("tooltip.multiDestroy") }}</template>
      <a-button
        shape="round"
        danger
        @click="setCircleButtonModal('accountDestroy')"
      >
        <DeleteFilled />{{ $t("tooltip.multiDestroy") }}
      </a-button>
    </a-tooltip>
    <a-tooltip
      v-if="state.buttonBoolean.multiWorkspaceAccountDestroy"
      placement="bottom"
    >
      <template #title>{{ $t("tooltip.multiDestroy") }}</template>
      <a-button
        shape="round"
        danger
        @click="setCircleButtonModal('workspaceAccountDestroy')"
      >
        <DeleteFilled />{{ $t("tooltip.multiDestroy") }}
      </a-button>
    </a-tooltip>

    <!-- Multi Select 일때 일괄 버튼 표시 -->
    <a-modal
      v-model:visible="commonModalView"
      :title="$t('tooltip.' + modalTitle)"
      :ok-text="$t('label.ok')"
      :cancel-text="$t('label.cancel')"
      @cancel="handleCancel"
      @ok="handleSubmit()"
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
      @ok="handleSubmit()"
    >
      <a-alert :message="modalConfirm" type="info" show-icon />
      <a-select
        v-model:value="selectedUser"
        show-search
        style="width: 100%; margin-top: 7px"
        option-filter-prop="label"
        class="addmodal-aform-item-div"
        :placeholder="$t('tooltip.vm.account.select')"
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
        vmDestroy: ref(false),
        userAllocate: ref(false),
        userUnlock: ref(false),
        workspaceDestroy: ref(false),
        workspaceAccountDestroy: ref(false),
        accountDestroy: ref(false),
        multiVmStart: ref(false),
        multiVmStop: ref(false),
        multiVmDestroy: ref(false),
        multiUserAllocate: ref(false),
        multiUserUnlock: ref(false),
        multiWorkspaceDestroy: ref(false),
        multiWorkspaceAccountDestroy: ref(false),
        multiAccountDestroy: ref(false),
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
      selectedUser: ref(undefined),
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
      worksUrl: ref(""),
      sucMessage: ref(""),
      failMessage: ref(""),
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
        if (this.workspaceInfo) {
          this.eventList = [this.workspaceInfo];
          if (this.callComponent === "WorkspaceUserList")
            this.state.buttonBoolean.workspaceAccountDestroy = true;
          else this.state.buttonBoolean.workspaceDestroy = true;
        } else {
          this.eventList = this.multiSelectList;
          if (this.callComponent === "WorkspaceUserList")
            this.state.buttonBoolean.multiWorkspaceAccountDestroy = true;
          else this.state.buttonBoolean.multiWorkspaceDestroy = true;
        }
      } else if (this.callComponent.includes("Account")) {
        if (this.accountInfo) {
          this.eventList = [this.accountInfo];
          this.state.buttonBoolean.accountDestroy = true;
        } else {
          this.eventList = this.multiSelectList;
          this.state.buttonBoolean.multiAccountDestroy = true;
        }
      } else if (this.callComponent.includes("GroupPolicy")) {
        this.state.buttonBoolean.destroy = true;
      }
      if (this.callComponent.includes("VirtualMachine")) {
        if (this.vmInfo) this.eventList = [this.vmInfo];
        else this.eventList = this.multiSelectList;

        // console.log(this.eventList);

        //시작 버튼 체크
        let res = null;
        res = this.eventList.filter((it) => it.mold_status === "Running");
        if (this.eventList.length === res.length) {
          if (this.vmInfo) this.state.buttonBoolean.vmStop = true;
          else this.state.buttonBoolean.multiVmStop = true;
        }

        //정지 버튼 체크
        res = this.eventList.filter((it) => it.mold_status === "Stopped");
        if (this.eventList.length === res.length) {
          if (this.vmInfo) this.state.buttonBoolean.vmStart = true;
          else this.state.buttonBoolean.multiVmStart = true;
        }
        //사용자 할당 해제 버튼 체크
        res = this.eventList.filter((it) => it.owner_account_id !== "");
        if (res.length > 0) {
          if (this.vmInfo) this.state.buttonBoolean.userUnlock = true;
          else this.state.buttonBoolean.multiUserUnlock = true;
        }

        //가상머신 삭제버튼 세팅
        if (this.vmInfo) this.state.buttonBoolean.vmDestroy = true;
        else this.state.buttonBoolean.multiVmDestroy = true;

        //사용자 할당 버튼 체크
        res = this.eventList.filter((it) => it.owner_account_id === "");
        if (this.eventList.length === res.length) {
          //같은 워크스페이스면 할당 버튼 활성화, 아니면 비활성화
          res = this.eventList.filter(function (obj, i, s) {
            return (
              i ===
              s.findIndex(function (t) {
                return t.workspace_name === obj.workspace_name;
              })
            );
          });
          //사용자 할당 버튼
          if (res.length === 1) {
            //사용자 할당 버튼
            if (this.vmInfo) this.state.buttonBoolean.userAllocate = true;
            else this.state.buttonBoolean.multiUserAllocate = true;

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
    handleSubmit() {
      //console.log(this.modalTitle + "  ::  " + this.callComponent);
      if (this.callComponent.includes("VirtualMachine")) {
        if (this.modalTitle.includes("vm")) this.vmAction();

        if (this.modalTitle.includes("userUnlock")) this.vmUserUnlockAction();

        if (this.modalTitle.includes("userAllocate"))
          this.vmUserAllocateAction();
      }

      if (this.callComponent.includes("Workspace")) {
        if (this.modalTitle.includes("workspaceDestroy"))
          this.workspaceDestroyAction();

        if (this.modalTitle.includes("workspaceAccountDestroy"))
          this.workspaceAccountDestroyAction();
      }

      if (this.callComponent.includes("Account")) {
        if (this.modalTitle.includes("accountDestroy"))
          this.accountDestroyAction();
      }
    },
    async funcDelay(delay) {
      return new Promise(function (resolve) {
        setTimeout(function () {
          resolve("delay call!");
        }, delay);
      });
    },
    funcEndMessage() {
      message.destroy();
      if (this.succCnt > 0) {
        message.success(
          this.$t(this.sucMessage, {
            count: this.succCnt,
          }),
          5
        );
      }
      if (this.failCnt > 0) {
        message.error(
          this.$t(this.failMessage, {
            count: this.failCnt,
          }),
          5
        );
      }
      this.failCnt = 0;
      this.succCnt = 0;
    },
    async workspaceAccountDestroyAction() {
      this.sucMessage = "message.workspace.user.delete.ok";
      this.failMessage = "message.workspace.user.delete.fail";

      message.loading(this.$t("message.workspace.vm.user.deleting"), 100);

      //console.log(this.eventList);
      for (let val of this.eventList) {
        try {
          const response = await worksApi.delete(
            "/api/v1/group/" + this.wsName + "/" + val.name
          );
          if (response.status == 200) {
            this.succCnt = this.succCnt + 1;
          }
        } catch (error) {
          console.log(error);
          this.failCnt = this.failCnt + 1;
        }
      }

      this.$emit("fetchData");
      await this.funcDelay(2000);
      this.handleCancel();
      this.funcEndMessage();
    },
    async vmUserAllocateAction() {
      if (this.selectedUser.length == 0) return false;

      this.sucMessage = "message.user.vm.allocated.ok";
      this.failMessage = "message.user.vm.allocated.fail";
      message.loading(this.$t("message.user.vm.allocating"), 100);

      for (let val of this.eventList) {
        try {
          const response = await worksApi.put(
            "/api/v1/connection/" + val.uuid + "/" + this.selectedUser
          );
          if (response.status == 200) {
            this.succCnt = this.succCnt + 1;
          }
        } catch (error) {
          console.log(error);
          this.failCnt = this.failCnt + 1;
        }
      }

      this.$emit("fetchData");
      await this.funcDelay(2000);
      this.handleCancel();
      this.funcEndMessage();
    },
    async vmUserUnlockAction() {
      this.sucMessage = "message.user.vm.unlock.ok";
      this.failMessage = "message.user.vm.unlock.fail";
      message.loading(this.$t("message.user.vm.unlocking"), 100);

      for (let val of this.eventList) {
        try {
          const response = await worksApi.delete(
            "/api/v1/connection/" + val.uuid
          );
          if (response.status == 204) {
            this.succCnt = this.succCnt + 1;
          }
        } catch (error) {
          console.log(error);
          this.failCnt = this.failCnt + 1;
        }
      }
      this.$emit("fetchData");
      await this.funcDelay(2000);
      this.handleCancel();
      this.funcEndMessage();
    },
    async vmAction() {
      if (this.modalTitle.includes("vmStart")) {
        message.loading(this.$t("message.vm.status.starting"), 100);
        this.worksUrl = "/api/v1/instance/VMStart/";
        this.sucMessage = "message.vm.status.start.ok";
        this.failMessage = "message.vm.status.start.fail";
      }
      if (this.modalTitle.includes("vmStop")) {
        message.loading(this.$t("message.vm.status.stopping"), 100);
        this.worksUrl = "/api/v1/instance/VMStop/";
        this.sucMessage = "message.vm.status.stop.ok";
        this.failMessage = "message.vm.status.stop.fail";
      }
      if (this.modalTitle.includes("vmDestroy")) {
        message.loading(this.$t("message.vm.status.destroying"), 100);
        this.worksUrl = "/api/v1/instance/VMDestroy/";
        this.sucMessage = "message.vm.status.delete.ok";
        this.failMessage = "message.vm.status.delete.fail";
      }
      for (let val of this.eventList) {
        try {
          const res = await worksApi.patch(this.worksUrl + val.uuid);
          if (res.status == 200) {
            this.succCnt = this.succCnt + 1;
          }
        } catch (error) {
          console.log(error);
          this.failCnt = this.failCnt + 1;
        }
      }

      await this.funcDelay(12000);
      if (
        this.callComponent == "VirtualMachineDetail" &&
        this.modalTitle.includes("vmDestroy")
      ) {
        router.push({ name: "VirtualMachine" });
      } else {
        this.$emit("fetchData");
      }
      
      //workspace 상세화면의 데스크톱 VM탭인지 여부를 확인하기 위함.
      if (this.wsName.length > 0) {
        await this.funcDelay(2000);
      }

      this.handleCancel();
      this.funcEndMessage();
    },
    async workspaceDestroyAction() {
      this.sucMessage = "message.workspace.status.delete";
      this.failMessage = "message.workspace.delete.existvm";
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

      await this.funcDelay(1000);
      this.funcEndMessage();
      this.handleCancel();
      if (this.callComponent == "WorkspaceDetail") {
        router.push({ name: "Workspace" });
      } else {
        this.$emit("fetchData");
      }
    },
    async accountDestroyAction() {
      this.sucMessage = "message.account.destroy.ok";
      this.failMessage = "message.account.destroy.fail";
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
      await this.funcDelay(1000);
      this.funcEndMessage();
      this.handleCancel();
      if (this.callComponent == "AccountDetail") {
        router.push({ name: "Account" });
      } else {
        this.$emit("fetchData");
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
