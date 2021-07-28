<template>
  <div id="ContentTab">
    <a-tabs
        v-model:activeKey="activeKey"
        :tab-position="tabPosition"
    >
      <a-tab-pane key="1" tab="상세">
        <DetailContent />
      </a-tab-pane>
      <a-tab-pane key="2" tab="데이터 디스크">
        <TableContent
            :data="VMDiskListData"
            :columns="VMDiskListColumns"
        />
      </a-tab-pane>
      <a-tab-pane key="3" tab="네트워크">
        <TableContent
            :data="NWListData"
            :columns="NWListColumns"
        />
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script>
import { defineComponent, ref } from 'vue';
import TableContent from "@/components/TableContent";
import DetailContent from "@/components/DetailContent";
import {VMListData, VMListColumns, VMDiskListData, VMDiskListColumns, NWListData, NWListColumns} from "@/data"
export default defineComponent({
  components: {
    TableContent,
    DetailContent
  },
  setup() {
    const tabPosition = ref('left');
    const activeKey = ref('1');
    return {
      tabPosition,
      activeKey,
      VMListData,
      VMListColumns,
      VMDiskListData,
      VMDiskListColumns,
      NWListData,
      NWListColumns

    };
  },
});
</script>

<style>
#ContentTab{
  text-align: left;
}
</style>
