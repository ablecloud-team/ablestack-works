<template>

  <a-tooltip placement="top" v-if="power === true">
    <template #title>{{ $t('tooltip.poweron') }}</template>
    <a-button type="primary" shape="circle">
      <PoweroffOutlined/>
    </a-button>
  </a-tooltip>

  <a-tooltip placement="top" v-if="poweroff === true" >
    <template #title>{{ $t('tooltip.poweroff') }}</template>
    <a-button type="primary" shape="circle" danger>
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

</template>

<script>
export default {
  name: "actions",
  props:{
    power: Boolean,
    poweroff: Boolean,
    reset: Boolean,
    reinstall: Boolean,
    snapshot: Boolean,
    volsnapshot: Boolean,
    iso: Boolean,
    destroy: Boolean,
    edit: Boolean,
  }
}
</script>

<style scoped>

</style>