<template class="able-action">
  <a-space :size="size">
    <a-tooltip v-if="state.buttonBoolean.addButton" placement="top">
      <template #title>{{ title }}</template>
      <a-button
        type="primary"
        shape="round"
        size="size"
        @click="setAddModalValue(true)"
      >
        <template #icon>
          {{ title }}
          <PlusOutlined />
        </template>
      </a-button>
    </a-tooltip>
    <!--Start circle button start-->
    <a-tooltip v-if="state.buttonBoolean.start" placement="top">
      <template #title>{{ $t("tooltip.start") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('tooltip.start')">
        <CaretRightOutlined />
      </a-button>
    </a-tooltip>
    <!--Start circle button end-->
    <!--Stop circle button start-->
    <a-tooltip v-if="state.buttonBoolean.stop" placement="top">
      <template #title>{{ $t("tooltip.stop") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('tooltip.stop')">
        <PoweroffOutlined />
      </a-button>
    </a-tooltip>
    <!--Stop circle button end-->
    <!--reset circle button start-->
    <a-tooltip v-if="state.buttonBoolean.reset" placement="top">
      <template #title>{{ $t("tooltip.reset") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('tooltip.reset')">
        <ReloadOutlined />
      </a-button>
    </a-tooltip>
    <!--reinstall circle button start-->
    <a-tooltip v-if="state.buttonBoolean.reinstall" placement="top">
      <template #title>{{ $t("tooltip.reinstall") }}</template>
      <a-button
        shape="circle"
        @click="setCircleButtonModal('tooltip.reinstall')"
      >
        <SyncOutlined />
      </a-button>
    </a-tooltip>
    <!--reinstall circle button end-->
    <!--snapshot circle button start-->
    <a-tooltip v-if="state.buttonBoolean.snapshot" placement="top">
      <template #title>{{ $t("tooltip.snapshot") }}</template>
      <a-button
        shape="circle"
        @click="setCircleButtonModal('tooltip.snapshot')"
      >
        <CameraOutlined />
      </a-button>
    </a-tooltip>
    <!--snapshot circle button end-->
    <!--volsnapshot circle button start-->
    <a-tooltip v-if="state.buttonBoolean.volsnapshot" placement="top">
      <template #title>{{ $t("tooltip.volsnapshot") }}</template>
      <a-button
        shape="circle"
        @click="setCircleButtonModal('tooltip.volsnapshot')"
      >
        <i class="fas fa-camera-retro"></i>
        <VideoCameraAddOutlined />
      </a-button>
    </a-tooltip>
    <!--volsnapshot circle button end-->
    <!--isoattach circle button start-->
    <a-tooltip v-if="state.buttonBoolean.isoattach" placement="top">
      <template #title>{{ $t("tooltip.isoattach") }}</template>
      <a-button
        shape="circle"
        @click="setCircleButtonModal('tooltip.isoattach')"
      >
        <PaperClipOutlined />
      </a-button>
    </a-tooltip>
    <!--isoattach circle button end-->
    <!--edit circle button start-->
    <a-tooltip v-if="state.buttonBoolean.edit" placement="top">
      <template #title>{{ $t("tooltip.edit") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('tooltip.edit')">
        <EditOutlined />
      </a-button>
    </a-tooltip>
    <!--edit circle button end-->
    <!--pause circle button start-->
    <a-tooltip v-if="state.buttonBoolean.pause" placement="top">
      <template #title>{{ $t("tooltip.pause") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('tooltip.pause')">
        <PauseOutlined />
      </a-button>
    </a-tooltip>
    <!--pause circle button end-->
    <!--destroy circle button start-->
    <a-tooltip v-if="state.buttonBoolean.destroy" placement="top">
      <template #title>{{ $t("tooltip.destroy") }}</template>
      <a-button
        type="primary"
        shape="circle"
        danger
        @click="setCircleButtonModal('tooltip.destroy')"
      >
        <DeleteFilled />
      </a-button>
    </a-tooltip>
    <!--destroy circle button ens-->
    <!--AddModal start-->
    <!-- <AddModal
      :visible="showAddModal"
      :action-from="actionFrom"
      :title="title"
      @changeAddModal="setAddModalValue"
    /> -->
    <!--AddModal end-->
    <!--circle button check modal start-->
    <a-modal
      :title="$t(modalTitle)"
      :visible="confirmModalView"
      @cancel="handleCancel"
      @ok="handleCancel"
    >
      <p>{{ $t(modalTitle) }} 하시겠습니까?</p>
    </a-modal>
    <!--circle button check modal end-->
  </a-space>
</template>

<script>
import { defineComponent, onMounted, reactive, ref } from "vue";
import AddModal from "@/components/AddModal";

export default defineComponent({
  components: {
    AddModal,
  },
  props: {
    actionFrom: {
      type: String,
      requires: true,
      default: "",
    },
    title: {
      type: String,
      requires: false,
      default: "",
    },
  },
  setup(props) {
    const sizeValue = 8;
    console.log("==================== props.actionFrom ====================:::: "+props.actionFrom);
    const state = reactive({
      callComponent: ref(props.actionFrom),
      buttonBoolean: {
        showModal: ref(false),
        addButton: ref(false),
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
      if (state.callComponent === "Workspace") {
        state.buttonBoolean.addButton = true;
      } else if (state.callComponent === "WorkspaceList") {
        state.buttonBoolean.start = true;
        state.buttonBoolean.stop = true;
        state.buttonBoolean.edit = true;
        state.buttonBoolean.destroy = true;
        state.buttonBoolean.pause = true;
      } else if (state.callComponent === "WorkspaceDetail") {
        state.buttonBoolean.start = true;
        state.buttonBoolean.stop = true;
        state.buttonBoolean.edit = true;
        state.buttonBoolean.destroy = true;
        state.buttonBoolean.pause = true;
      } else if (state.callComponent === "VirtualMachine") {
        state.buttonBoolean.start = true;
        state.buttonBoolean.stop = true;
        state.buttonBoolean.reset = true;
        state.buttonBoolean.destroy = true;
      } else if (state.callComponent === "VirtualMachineList") {
        state.buttonBoolean.start = true;
        state.buttonBoolean.stop = true;
        state.buttonBoolean.edit = true;
        state.buttonBoolean.destroy = true;
      } else if (state.callComponent === "VirtualMachineDetail") {
        state.buttonBoolean.start = true;
        state.buttonBoolean.stop = true;
        state.buttonBoolean.edit = true;
        state.buttonBoolean.destroy = true;
      } else if (state.callComponent === "User") {
        state.buttonBoolean.addButton = true;
      } else if (state.callComponent === "UserDetail") {
        state.buttonBoolean.destroy = true;
        state.buttonBoolean.pause = true;
      }
    });
    return {
      size: ref(sizeValue),
      confirmModalView: ref(false),
      //TODO:2021.07.30 mobalTitle i18n 적용필요
      //fixm
      modalTitle: "",
      state,
    };
  },
  // emits: [
  //   'changeAddModal',
  // ],
  methods: {
    setCircleButtonModal: function (value) {
      this.confirmModalView = true;
      this.modalTitle = value;
    },
    handleCancel: function () {
      this.confirmModalView = false;
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
