<template>
  <div id="ContentTab">
    <a-tabs v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane key="1" :tab="$t('label.detail')">
        <DetailContent :data-info="vmDataInfo" :action-from="'vmDetail'" />
      </a-tab-pane>
      <a-tab-pane key="2" :tab="$t('label.disk.list')">
        <TableContent :data="vmDiskList" :columns="vmDiskListColumns" />
      </a-tab-pane>
      <a-tab-pane key="3" :tab="$t('label.network.list')">
        <TableContent :data="vmNetworkList" :columns="vmNetworkListColumns" />
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script>
import { defineComponent, ref } from "vue";
import TableContent from "@/components/TableContent";
import DetailContent from "@/components/DetailContent";
export default defineComponent({
  components: {
    TableContent,
    DetailContent,
  },
  props: {
    vmDataInfo: {
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
      vmDiskList: JSON.parse(
        '[{"name":"Datadisk1","state":"Allocated","size":"50GB","conn":"VM1","action":""},' +
          '{"name":"Datadisk2","state":"Allocated","size":"100GB","conn":"VM2","action":""},' +
          '{"name":"Datadisk3","state":"Allocated","size":"200GB","conn":"VM3","action":""}]'
      ),
      vmNetworkList: JSON.parse(
        '[{"name":"Datadisk1","state":"Allocated","size":"50GB","conn":"VM1","action":""},' +
          '{"name":"Datadisk2","state":"Allocated","size":"100GB","conn":"VM2","action":""},' +
          '{"name":"Datadisk3","state":"Allocated","size":"200GB","conn":"VM3","action":""}]'
      ),
      vmDiskListColumns: [
        {
          dataIndex: "name",
          key: "name",
          slots: { customRender: "nameRender" },
          title: this.$t("label.name"),
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: "",
          key: "action",
          dataIndex: "action",
          align: "right",
          width: "5px",
          slots: { customRender: "actionRender" },
        },
        {
          title: this.$t("label.state"),
          dataIndex: "state",
          key: "state",
          sorter: (a, b) =>
            a.state < b.state ? -1 : a.state > b.state ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.size"),
          dataIndex: "size",
          key: "size",
          sorter: (a, b) => (a.size < b.size ? -1 : a.size > b.size ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.connected.desktop"),
          dataIndex: "conn",
          key: "conn",
          sorter: (a, b) => (a.conn < b.conn ? -1 : a.conn > b.conn ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
      ],

      vmNetworkListColumns: [
        {
          dataIndex: "name",
          key: "name",
          slots: { customRender: "nameRender" },
          title: this.$t("label.name"),
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: "",
          key: "action",
          dataIndex: "action",
          align: "right",
          width: "5px",
          slots: { customRender: "actionRender" },
        },
        {
          title: this.$t("label.state"),
          dataIndex: "state",
          key: "state",
          sorter: (a, b) =>
            a.state < b.state ? -1 : a.state > b.state ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
      ],
    };
  },
});
</script>

<style>
#ContentTab {
  text-align: left;
}
</style>
