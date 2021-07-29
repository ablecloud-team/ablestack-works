<template class="able-action">
  <a-space :size="size">
    <a-tooltip placement="top" v-if="addButton === true">
      <template #title>{{ add }}</template>
      <a-button type="primary" shape="round" size="size" @click="setAddModalValue(true)">
        <template #icon>
          {{ add }}
          <PlusOutlined />
        </template>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="start">
      <template #title>{{ $t('tooltip.start') }}</template>
      <a-button
          shape="circle"
          @click="setCircleButtonModal('tooltip.start')"
      >
        <CaretRightOutlined />
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="stop" >
      <template #title>{{ $t('tooltip.stop') }}</template>
      <a-button
          shape="circle"
          @click="setCircleButtonModal('tooltip.stop')"
      >
        <PoweroffOutlined/>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="reset">
      <template #title>{{ $t('tooltip.reset') }}</template>
      <a-button
          shape="circle"
          @click="setCircleButtonModal('tooltip.reset')"
      >
        <ReloadOutlined/>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="reinstall">
      <template #title>{{ $t('tooltip.reinstall') }}</template>
      <a-button
          shape="circle"
          @click="setCircleButtonModal('tooltip.reinstall')"
      >
        <SyncOutlined/>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="snapshot">
      <template #title>{{ $t('tooltip.snapshot') }}</template>
      <a-button
          shape="circle"
          @click="setCircleButtonModal('tooltip.snapshot')"
      >
        <CameraOutlined/>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="volsnapshot">
      <template #title>{{ $t('tooltip.volsnapshot') }}</template>
      <a-button
          shape="circle"
          @click="setCircleButtonModal('tooltip.volsnapshot')"
      >
        <i class="fas fa-camera-retro"></i>
        <font-awesome-icon icon="camera-retro"/>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="iso">
      <template #title>{{ $t('tooltip.isoattach') }}</template>
      <a-button
          shape="circle"
          @click="setCircleButtonModal('tooltip.isoattach')"
      >
        <PaperClipOutlined/>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="edit">
      <template #title>{{ $t('tooltip.edit') }}</template>
      <a-button
          shape="circle"
          @click="setCircleButtonModal('tooltip.edit')"
      >
        <EditOutlined/>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="destroy">
      <template #title>{{ $t('tooltip.destroy') }}</template>
      <a-button
          type="primary"
          shape="circle"
          @click="setCircleButtonModal('tooltip.destroy')"
          danger
      >
        <DeleteFilled/>
      </a-button>
    </a-tooltip>
  <AddModal
      :visible="showAddModal"
      :actionFrom="'Actions'"
      @changeAddModal="setAddModalValue"
  />
    <a-modal
        :title="$t(modalTitle)"
        :visible="visible"
        :confirm-loading="confirmLoading"
        @cancel="handleCancel"
        @ok="handleCancel"
    >
      <p>{{ $t(modalTitle) }} 하시겠습니까?</p>
    </a-modal>
  </a-space>
</template>

<script>
import {defineComponent, ref} from 'vue';
import AddModal from "@/components/AddModal";
export default defineComponent({
  components: {
    AddModal
  },
  props:{
    actionFrom:  {
      String,
    },
    add: {
      String,
    },
    showModal: {
      Boolean,
      default: false,
    },
    addButton: {
      Boolean,
      default: false,
    },
    start: {
      Boolean,
      default: false,
    },
    stop: {
      Boolean,
      default: false,
    },
    reset: {
      Boolean,
      default: false,
    },
    reinstall: {
      Boolean,
      default: false,
    },
    snapshot: {
      Boolean,
      default: false,
    },
    volsnapshot: {
      Boolean,
      default: false,
    },
    iso: {
      Boolean,
      default: false,
    },
    destroy: {
      Boolean,
      default: false,
    },
    edit: {
      Boolean,
      default: false,
    },
  },
  setup() {
    const sizeValue = 8;
    return {
      size: ref(sizeValue),
      showAddModal: ref(false),
      visible: ref(false),
      modalTitle: ref('')
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
    }
  },
})
</script>

<style scoped>
#content-action .ant-btn {
  padding-top:0px ;
  padding-bottom: 0px;
  font-size: 14px;
}
</style>
