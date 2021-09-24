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
    <a-tooltip v-if="state.buttonBoolean.reset" placement="top">
      <template #title>{{ $t("tooltip.reset") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('reset')">
        <ReloadOutlined />
      </a-button>
    </a-tooltip>
    <!--reinstall-->
    <a-tooltip v-if="state.buttonBoolean.reinstall" placement="top">
      <template #title>{{ $t("tooltip.reinstall") }}</template>
      <a-button
        shape="circle"
        @click="setCircleButtonModal('reinstall')"
      >
        <SyncOutlined />
      </a-button>
    </a-tooltip>
    <!--snapshot-->
    <a-tooltip v-if="state.buttonBoolean.snapshot" placement="top">
      <template #title>{{ $t("tooltip.snapshot") }}</template>
      <a-button
        shape="circle"
        @click="setCircleButtonModal('snapshot')"
      >
        <CameraOutlined />
      </a-button>
    </a-tooltip>
    <!--volsnapshot-->
    <a-tooltip v-if="state.buttonBoolean.volsnapshot" placement="top">
      <template #title>{{ $t("tooltip.volsnapshot") }}</template>
      <a-button
        shape="circle"
        @click="setCircleButtonModal('volsnapshot')"
      >
        <i class="fas fa-camera-retro"></i>
        <VideoCameraAddOutlined />
      </a-button>
    </a-tooltip>
    <!--isoattach-->
    <a-tooltip v-if="state.buttonBoolean.isoattach" placement="top">
      <template #title>{{ $t("tooltip.isoattach") }}</template>
      <a-button
        shape="circle"
        @click="setCircleButtonModal('isoattach')"
      >
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
      <a-button shape="circle" @click="setCircleButtonModal('pause')">
        <PauseOutlined />
      </a-button>
    </a-tooltip>
    <!--destroy-->
    <a-tooltip v-if="state.buttonBoolean.destroy" placement="top">
      <template #title>{{ $t("tooltip.destroy") }}</template>
      <a-button
        type="primary"
        shape="circle"
        danger
        @click="setCircleButtonModal('destroy')"
      >
        <DeleteFilled />
      </a-button>
    </a-tooltip>

    <!-- Confirm Modal -->
    <a-modal
      :title="$t('tooltip.'+modalTitle)"
      :visible="confirmModalView"
      @cancel="handleCancel"
      @ok="handleSubmit(actionFrom, name, uuid)">
      <p>{{ $t(modalConfirm) }}</p>
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
    name: {
      type: String,
      requires: false,
      default: "",
    },
    status: {
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
      status: ref(props.status),
      buttonBoolean: {
        showModal: ref(false),
        start: ref(false),
        stop: ref(false),
        reset: ref(false),
        reinstall: ref(false),
        snapshot: ref(false),
        volsnapshot: ref(false),
        iso: ref(false),
        destroy: ref(false),
        edit: ref(false),
        pause: ref(false),
      },
    });
    onMounted(() => {
      if (state.callComponent === "WorkspaceList" || state.callComponent === "WorkspaceDetail") {
        state.buttonBoolean.start = true;
        state.buttonBoolean.stop = true;
        state.buttonBoolean.destroy = true;
      } else if (state.callComponent === "VirtualMachine") {
      } else if (state.callComponent === "VirtualMachineList" || state.callComponent === "VirtualMachineDetail") {
        if (state.status === "Running") {
          state.buttonBoolean.stop = true;
        } else {
          state.buttonBoolean.start = true;
        }
        state.buttonBoolean.iso = true;
        state.buttonBoolean.destroy = true;
      } else if (state.callComponent === "AccountList" || state.callComponent === "AccountDetail") {
        state.buttonBoolean.destroy = true;
      } else if (state.callComponent === "GroupPolicyList" || state.callComponent === "GroupPolicyDetail") {
        state.buttonBoolean.destroy = true;
      }
    });
    return {
      size: ref(sizeValue),
      confirmModalView: ref(false),
      modalTitle: ref(""),
      modalConfirm: ref(""),
      popView: ref(true),
      state,
    };
  },
  // emits: [
  //   'changeAddModal',
  // ],
  methods: {
    setCircleButtonModal: function (value) {
      this.confirmModalView = true;
      this.popView = false;
      this.modalTitle = value;
      if(value == "start") this.modalConfirm = "modal.confirm.start";
      if(value == "stop") this.modalConfirm = "modal.confirm.stop";
      if(value == "destroy") this.modalConfirm = "modal.confirm.destroy";

    },
    handleCancel: function () {
      this.confirmModalView = false;
    },
    handleSubmit: function (actionFrom, name, uuid) { //actionFrom, name, uuid
      console.log(this.modalTitle + "  ::  " +actionFrom + " ::  " + name + " :: "+ uuid);
      if(actionFrom == "VirtualMachineList"){
        let worksUrl, resMessage = "";
        if(this.modalTitle.includes("start")){
          worksUrl = "/api/v1/instance/VMStart/" + uuid;
          resMessage = this.$t("message.vm.status.update");
        }
        if(this.modalTitle.includes("stop")){
          worksUrl = "/api/v1/instance/VMStop/" + uuid;
          resMessage = this.$t("message.vm.status.update");
        }
        if(this.modalTitle.includes("destroy")){
          worksUrl = "/api/v1/instance/VMDestroy/"+ uuid;
          resMessage = this.$t("message.vm.status.delete");
        }
        worksApi
          .patch( worksUrl, { withCredentials: true })
          .then((response) => {
            if (response.status == 200) {
              this.vmDataList = response.data.result.list;
              message.loading(resMessage);
            } else {
              message.error("message.vm.status.update.fail");
            }
            setTimeout(() => {
              this.$emit("fetchData");
              this.handleCancel();
            }, 1000);

          })
          .catch(function (error) {
            message.error(error);
          });

      }
      setTimeout(() => {
        this.$emit("fetchData");
        this.handleCancel();
      }, 1000);
 
    } 
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
