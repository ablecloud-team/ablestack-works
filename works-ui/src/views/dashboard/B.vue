<template class="able-action">
  <a-space :size="size">
    <a-tooltip placement="top" v-if="state.buttonBoolean.addButton">
      <template #title>{{ add }}</template>
      <a-button
        type="primary"
        shape="round"
        size="size"
        @click="setAddModalValue(true)"
      >
        <template #icon>
          {{ add }}
          <PlusOutlined />
        </template>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="state.buttonBoolean.start">
      <template #title>{{ $t("tooltip.start") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('tooltip.start')">
        <CaretRightOutlined />
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="state.buttonBoolean.stop">
      <template #title>{{ $t("tooltip.stop") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('tooltip.stop')">
        <PoweroffOutlined />
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="state.buttonBoolean.reset">
      <template #title>{{ $t("tooltip.reset") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('tooltip.reset')">
        <ReloadOutlined />
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="state.buttonBoolean.reinstall">
      <template #title>{{ $t("tooltip.reinstall") }}</template>
      <a-button
        shape="circle"
        @click="setCircleButtonModal('tooltip.reinstall')"
      >
        <SyncOutlined />
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="state.buttonBoolean.snapshot">
      <template #title>{{ $t("tooltip.snapshot") }}</template>
      <a-button
        shape="circle"
        @click="setCircleButtonModal('tooltip.snapshot')"
      >
        <CameraOutlined />
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="state.buttonBoolean.volsnapshot">
      <template #title>{{ $t("tooltip.volsnapshot") }}</template>
      <a-button
        shape="circle"
        @click="setCircleButtonModal('tooltip.volsnapshot')"
      >
        <i class="fas fa-camera-retro"></i>
        <VideoCameraAddOutlined />
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="state.buttonBoolean.isoattach">
      <template #title>{{ $t("tooltip.isoattach") }}</template>
      <a-button
        shape="circle"
        @click="setCircleButtonModal('tooltip.isoattach')"
      >
        <PaperClipOutlined />
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="state.buttonBoolean.start">
      <template #title>{{ $t("tooltip.edit") }}</template>
      <a-button shape="circle" @click="setCircleButtonModal('tooltip.edit')">
        <EditOutlined />
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="state.buttonBoolean.destroy">
      <template #title>{{ $t("tooltip.destroy") }}</template>
      <a-button
        type="primary"
        shape="circle"
        @click="setCircleButtonModal('tooltip.destroy')"
        danger
      >
        <DeleteFilled />
      </a-button>
    </a-tooltip>
    <AddModal
      :visible="showAddModal"
      :actionFrom="'Actions'"
      :add="add"
      @changeAddModal="setAddModalValue"
    />
    <a-modal
      :title="$t(modalTitle)"
      :visible="visible"
      @cancel="handleCancel"
      @ok="handleCancel"
    >
      <p>{{ $t(modalTitle) }} 하시겠습니까?</p>
    </a-modal>
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
      String,
    },
    add: {
      String,
    },
  },
  setup(props) {
    const sizeValue = 8;
    console.log("B-setup");
    console.log(props.actionFrom);
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
      },
    });
    onMounted(() => {
      if (state.callComponent === "Workspaces") {
        state.buttonBoolean.showModal = true;
        state.buttonBoolean.addButton = true;
        state.buttonBoolean.start = true;
        state.buttonBoolean.stop = true;
        state.buttonBoolean.reset = true;
        state.buttonBoolean.reinstall = true;
        state.buttonBoolean.snapshot = true;
        state.buttonBoolean.volsnapshot = true;
        state.buttonBoolean.iso = true;
        state.buttonBoolean.destroy = true;
        state.buttonBoolean.edit = true;
      }
    });
    return {
      size: ref(sizeValue),
      showAddModal: ref(false),
      visible: ref(false),
      modalTitle: ref(""),
      state,
    };
  },
  // emits: [
  //   'changeAddModal',
  // ],
  methods: {
    setAddModalValue: function (value) {
      this.showAddModal = value;
    },
    setCircleButtonModal: function (value) {
      this.visible = true;
      this.modalTitle = value;
    },
    handleCancel: function () {
      this.visible = false;
    },
  },
});
</script>

<style scoped>
#content-action .ant-btn {
  padding-top: 0;
  padding-bottom: 0;
  font-size: 14px;
}
</style>
