<template>
  <div id="ContentTab">
    <a-tabs v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane key="1" :tab="$t('label.detail')">
        <DetailContent
          ref="listRefleshCall1"
          :action-from="'UserDetail'"
        />
      </a-tab-pane>
      <a-tab-pane key="2" :tab="$t('label.vm.list')">
        <!-- <TableContent
          ref="listRefleshCall2"
          :tap-name="'desktop'"
          :action-from="'VirtualMachineList'"
        /> -->
      </a-tab-pane>
      <a-tab-pane key="3" :tab="$t('label.app.list')">
        <!-- <TableContent
          ref="listRefleshCall3"
          :tap-name="'app'"
          :action-from="'VirtualMachineList'"
        /> -->
      </a-tab-pane>
      <a-tab-pane key="4" :tab="$t('label.policy.list')">
        <!-- <TableContent
          ref="listRefleshCall4"
          :tap-name="'policy'"
          :action-from="'PolicyGroupList'"
        /> -->
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script>
import { defineComponent, ref } from "vue";
import TableContent from "@/components/TableContent";
import DetailContent from "@/components/DetailContent";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";

export default defineComponent({
  components: {
    TableContent,
    DetailContent,
  },
  props: {
    userDataInfo: {
      type: Object,
      required: true,
      default: null,
    },
    // vmDiskList:{
    //   type: Object,
    //   required: false,
    //   default: null,
    // },
    // vmNetworkList:{
    //   type: Object,
    //   required: false,
    //   default: null,
    // },
  },
  setup() {
    const tabPosition = ref("top");
    const activeKey = ref("1");
    return {
      tabPosition,
      activeKey,
    };
  },
  data() {
    return {
    };
  },
  methods: {
    reflesh() {
      this.$refs.listRefleshCall1.reflesh();
      // this.$refs.listRefleshCall2.fetchData();
      // this.$refs.listRefleshCall3.fetchData();
      // this.$refs.listRefleshCall4.fetchData();
    },
  },
});
</script>

<style>
#ContentTab {
  text-align: left;
}
</style>
