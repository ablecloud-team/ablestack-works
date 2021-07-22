<template>
  <div id="ContentTab">
    <a-tabs type="card" v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane key="1" tab="Tab 1"><TableContent :data="data" :columns="columns" /></a-tab-pane>
<!--      <a-tab-pane key="2" tab="Tab 2"><ListView :data="data" :columns="columns"  /></a-tab-pane>-->
      <a-tab-pane key="3" tab="Tab 3">Content of Tab Pane 3</a-tab-pane>
    </a-tabs>
  </div>
</template>

<script>
import TableContent from "./TableContent";

import { defineComponent, ref } from 'vue';
import ListView from "@/components/ListView";
import {data, columns} from "@/data"
export default defineComponent({
  components: {ListView, TableContent},
  setup() {
    const tabPosition = ref('top');
    const activeKey = ref('1');
    return {
      tabPosition,
      activeKey,
      data,
      columns
    };
  },
});
</script>

<style>
#ContentTab{
  text-align: left;
}
</style>
