<template>
  <div id="ContentTab">
    <a-tabs type="card" v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane key="1" :tab="$t('label.vm.list')">
        <TableContent :data="VMListData" :columns="VMListColumns" comp="VirtualMachineDetail" />
      </a-tab-pane>
      <a-tab-pane key="2" :tab="$t('label.users')">
        <TableContent :data="UserListData" :columns="UserListColumns" />
      </a-tab-pane>
      <a-tab-pane key="3" :tab="$t('label.disk.list')">
        <TableContent :data="VMDiskListData" :columns="VMDiskListColumns" />
      </a-tab-pane>
      <a-tab-pane key="4" :tab="$t('label.network.list')">
        <TableContent :data="NWListData" :columns="NWListColumns" />
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script>
import TableContent from "@/components/TableContent";

import { defineComponent, ref } from 'vue';
// import ListView from "@/components/ListView";

import {VMListData,
  VMListColumns,
  VMDiskListData,
  VMDiskListColumns,
  NWListData,
  NWListColumns,
  UserListData,
  UserListColumns,
} from "@/data"
export default defineComponent({
  components: { TableContent },
  setup() {
    const tabPosition = ref('top');
    const activeKey = ref('1');
    return {
      tabPosition,
      activeKey,
      VMListData,
      VMListColumns,
      VMDiskListData,
      VMDiskListColumns,
      NWListData,
      NWListColumns,
      UserListData,
      UserListColumns,
    };
  },
});
</script>

<style>
#ContentTab{
  text-align: left;
}
</style>
