<template>
  <div id="ContentTab">
    <a-tabs
      v-model:activeKey="activeKey"
      type="card"
      :tab-position="tabPosition"
      @change="changeTap"
    >
      <a-tab-pane key="1" :tab="$t('label.vm.list')">
        <TableContent
          v-model:tapName="state.callTap"
          :data="VMListData"
          :columns="VMListColumns"
          :actionFrom="'VirtualMachineList'"
          comp="VirtualMachineDetail"
        />
      </a-tab-pane>
      <a-tab-pane key="2" :tab="$t('label.users')">
        <TableContent
          v-model:tapName="state.callTap"
          :data="UserListData"
          :columns="UserListColumns"
        />
      </a-tab-pane>
      <a-tab-pane key="3" :tab="$t('label.disk.list')">
        <TableContent
          v-model:tapName="state.callTap"
          :data="VMDiskListData"
          :columns="VMDiskListColumns"
        />
      </a-tab-pane>
      <a-tab-pane key="4" :tab="$t('label.network.list')">
        <TableContent
          v-model:tapName="state.callTap"
          :data="networkDataList"
          :columns="NWListColumns"
        />
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script>
import { defineComponent, reactive, ref } from "vue";
import TableContent from "@/components/TableContent";

export default defineComponent({
  components: { TableContent },
  props: {
    workspaceDataList: {
      type: Object,
      required: true,
      default: null,
    },
    networkDataList: {
      type: Object,
      required: true,
      default: null,
    },
  },
  setup() {
    const tabPosition = ref("top");
    const activeKey = ref("1");
    const state = reactive({
      callTap: ref("desktop"),
    });
    const changeTap = (value) => {
      console.log(`${value}`);
      if (`${value}` === "1") {
        state.callTap = ref("desktop");
      } else if (`${value}` === "2") {
        state.callTap = ref("user");
      } else if (`${value}` === "3") {
        state.callTap = ref("datadisk");
      } else if (`${value}` === "4") {
        state.callTap = ref("network");
      }
      console.log(state.callTap);
    };
    return {
      state,
      changeTap,
      tabPosition,
      activeKey,
    };
  },
  data() {
    return {
      VMDiskListColumns : [
        {
          dataIndex: "name",
          key: "name",
          slots: { customRender: "nameRender" },
          title: this.$t('label.name'),
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: "",
          key: "action",
          dataIndex: "action",
          slots: { customRender: "actionRender" },
        },
        {
          title: this.$t('label.state'),
          dataIndex: "state",
          key: "state",
          sorter: (a, b) => (a.state < b.state ? -1 : a.state > b.state ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t('label.size'),
          dataIndex: "size",
          key: "size",
          sorter: (a, b) => (a.size < b.size ? -1 : a.size > b.size ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t('label.connected.desktop'),
          dataIndex: "conn",
          key: "conn",
          sorter: (a, b) => (a.conn < b.conn ? -1 : a.conn > b.conn ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
      ],
      VMDiskListData : JSON.parse(
        '[{"name":"Datadisk1","state":"Allocated","size":"50GB","conn":"VM1","action":""},' +
        '{"name":"Datadisk2","state":"Allocated","size":"100GB","conn":"VM2","action":""},' +
        '{"name":"Datadisk3","state":"Allocated","size":"200GB","conn":"VM3","action":""}]'
      ),
      VMListColumns : [
        {
          dataIndex: "name",
          key: "name",
          slots: { customRender: "nameRender" },
          title: this.$t('label.name'),
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: "",
          key: "action",
          dataIndex: "action",
          slots: { customRender: "actionRender" },
        },
        {
          title: this.$t('label.state'),
          dataIndex: "state",
          key: "state",
          sorter: (a, b) => (a.state < b.state ? -1 : a.state > b.state ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t('label.workspace'),
          dataIndex: "workspace",
          key: "workspace",
          sorter: (a, b) =>
            a.workspace < b.workspace ? -1 : a.workspace > b.workspace ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t('label.users'),
          dataIndex: "user",
          key: "user",
          sorter: (a, b) => (a.user < b.user ? -1 : a.user > b.user ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t('label.desktop.connect.boolean'),
          dataIndex: "conn",
          key: "conn",
          sorter: (a, b) => (a.conn < b.conn ? -1 : a.conn > b.conn ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
      ],
      VMListData : [
        {"name":"VM1","state":"Running","user":"user01","conn":"TRUE","workspace":"work11"},
        {"name":"VM2","state":"Stopped","user":"user03","conn":"FALSE","workspace":"work112"},
        {"name":"VM3","state":"Running","user":"user02","conn":"TRUE","workspace":"work113"}
      ],
      NWListColumns : [
        {
          dataIndex: "name",
          key: "name",
          slots: { customRender: "nameRender" },
          title: this.$t('label.name'),
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: "",
          key: "action",
          dataIndex: "action",
          slots: { customRender: "actionRender" },
        },
        {
          title: this.$t('label.state'),
          dataIndex: "state",
          key: "state",
          sorter: (a, b) => (a.state < b.state ? -1 : a.state > b.state ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
      ],
      UserListData : [
        {"name":"user01","state":"Allocated","desktop":"Desktop1"}
      ],
      UserListColumns : [
        {
          dataIndex: "name",
          key: "name",
          slots: { customRender: "nameRender" },
          title: this.$t('label.name'),
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        // {
        //     title: '',
        //     key: 'action',
        //     dataIndex: 'action',
        //     slots: {customRender: 'actionRender'}
        // },
        {
          title: this.$t('label.state'),
          dataIndex: "state",
          key: "state",
          sorter: (a, b) => (a.state < b.state ? -1 : a.state > b.state ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t('label.allocated.desktop'),
          dataIndex: "Deskdesktoptop",
          key: "desktop",
          sorter: (a, b) => (a.desktop < b.desktop ? -1 : a.desktop > b.desktop ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
      ],
    }
  },
});
</script>

<style>
#ContentTab {
  text-align: left;
}
</style>
