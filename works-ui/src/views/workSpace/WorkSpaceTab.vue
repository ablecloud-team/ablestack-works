<template>
  <div id="ContentTab">
    <a-tabs
        type="card"
        v-model:activeKey="activeKey"
        :tab-position="tabPosition"
        @change="changeTap"
    >
      <a-tab-pane key="1" :tab="$t('label.vm.list')">
        <TableContent
            v-model:tapName="status.callTap"
            :data="VMListData"
            :columns="VMListColumns"
            comp="VirtualMachineDetail"
        />
      </a-tab-pane>
      <a-tab-pane key="2" :tab="$t('label.users')">
        <TableContent
            v-model:tapName="status.callTap"
            :data="UserListData"
            :columns="UserListColumns"
        />
      </a-tab-pane>
      <a-tab-pane key="3" :tab="$t('label.disk.list')">
        <TableContent
            v-model:tapName="status.callTap"
            :data="VMDiskListData"
            :columns="VMDiskListColumns"
        />
      </a-tab-pane>
      <a-tab-pane key="4" :tab="$t('label.network.list')">
        <TableContent
            v-model:tapName="status.callTap"
            :data="NWListData"
            :columns="NWListColumns"
        />
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script>
import TableContent from "@/components/TableContent";

import {defineComponent, reactive, ref} from 'vue';
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
    const status = reactive({
      callTap: ref('desktop'),
    });
    const changeTap = value =>{
      console.log(`${value}`);
      if(`${value}` === '1'){
        status.callTap = ref('desktop');
      }else if(`${value}` === '2'){
        status.callTap = ref('user');
      }else if(`${value}` === '3'){
        status.callTap = ref('datadisk');
      }else if(`${value}` === '4'){
        status.callTap = ref('network');
      }
      console.log('sdfsdfsdf');
      console.log(status.callTap);
    }
    return {
      status,
      changeTap,
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
