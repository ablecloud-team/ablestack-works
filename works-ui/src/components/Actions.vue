<template class="able-action">
  <a-space :size="size">
    <a-tooltip placement="top" v-if="addButton === true">
      <template #title>{{ add }}</template>
      <a-button type="primary" shape="round" size="size" @click="changeAddModalValue">
        <template #icon>
          {{ add }}
          <PlusOutlined />
        </template>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="power === true">
      <template #title>{{ $t('tooltip.poweron') }}</template>
      <a-button shape="circle">
        <PoweroffOutlined/>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="poweroff === true" >
      <template #title>{{ $t('tooltip.poweroff') }}</template>
      <a-button shape="circle">
        <PoweroffOutlined/>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="reset === true">
      <template #title>{{ $t('tooltip.reset') }}</template>
      <a-button shape="circle" >
        <ReloadOutlined/>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="reinstall === true">
      <template #title>{{ $t('tooltip.reinstall') }}</template>
      <a-button shape="circle" >
        <SyncOutlined/>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="snapshot === true">
      <template #title>{{ $t('tooltip.snapshot') }}</template>
      <a-button shape="circle" >
        <CameraOutlined/>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="volsnapshot === true">
      <template #title>{{ $t('tooltip.volsnapshot') }}</template>
      <a-button shape="circle" >
        <i class="fas fa-camera-retro"></i>
        <font-awesome-icon icon="camera-retro"/>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="iso === true">
      <template #title>{{ $t('tooltip.isoattach') }}</template>
      <a-button shape="circle" >
        <PaperClipOutlined/>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="edit === true">
      <template #title>{{ $t('tooltip.edit') }}</template>
      <a-button shape="circle" >
        <EditOutlined/>
      </a-button>
    </a-tooltip>

    <a-tooltip placement="top" v-if="destroy === true">
      <template #title>{{ $t('tooltip.destroy') }}</template>
      <a-button type="primary" shape="circle" danger>
        <DeleteFilled/>
      </a-button>
    </a-tooltip>
  </a-space>
</template>

<script>
import {defineComponent, ref} from 'vue';
export default defineComponent({
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
    power: {
      Boolean,
      default: false,
    },
    poweroff: {
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
    };
  },
  emits: [
    'changeAddModal',
  ],
  methods: {
    changeAddModalValue() {
      this.$emit('changeAddModal', true);
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
