<template>
  <div id="ContentTab">
    <a-tabs v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane key="1" :tab="$t('label.vm.list')">
        <TableContent
          :tapName="'desktop'"
          :data="VMListData"
          :columns="VMListColumns"
          :actionFrom="'VirtualMachineList'"
          :comp="'VirtualMachineDetail'"
          @tabReflesh="tabReflesh"
        />
      </a-tab-pane>
      <a-tab-pane key="2" :tab="$t('label.users')">
        <TableContent
          :tapName="'user'"
          :data="workspaceUserList"
          :columns="UserListColumns"
          :actionFrom="'UserDetail'"
          @tabReflesh="tabReflesh"
        />
      </a-tab-pane>
      <!-- <a-tab-pane key="3" :tab="$t('label.disk.list')">
        <TableContent
          :tapName="'disk'"
          :data="VMDiskListData"
          :columns="VMDiskListColumns"
          @tabReflesh="tabReflesh"
        />
      </a-tab-pane> -->
      <a-tab-pane key="4" :tab="$t('label.network.list')">
        <TableContent
          :tapName="'network'"
          :data="networkDataList"
          :columns="NWListColumns"
          @tabReflesh="tabReflesh"
        />
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script>
import { defineComponent, reactive, ref } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";
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
    return {
      tabPosition,
      activeKey,
    };
  },
  data() {
    return {
      VMDiskListColumns: [
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
          sorter: (a, b) => (a.state < b.state ? -1 : a.state > b.state ? 1 : 0),
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
      VMDiskListData: JSON.parse(
        '[{"name":"Datadisk1","state":"Allocated","size":"50GB","conn":"VM1","action":""},' +
        '{"name":"Datadisk2","state":"Allocated","size":"100GB","conn":"VM2","action":""},' +
        '{"name":"Datadisk3","state":"Allocated","size":"200GB","conn":"VM3","action":""}]'
      ),
      VMListColumns: [
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
          sorter: (a, b) => (a.state < b.state ? -1 : a.state > b.state ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.workspace"),
          dataIndex: "workspace",
          key: "workspace",
          sorter: (a, b) =>
            a.workspace < b.workspace ? -1 : a.workspace > b.workspace ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.users"),
          dataIndex: "user",
          key: "user",
          sorter: (a, b) => (a.user < b.user ? -1 : a.user > b.user ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.desktop.connect.boolean"),
          dataIndex: "conn",
          key: "conn",
          sorter: (a, b) => (a.conn < b.conn ? -1 : a.conn > b.conn ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
      ],
      VMListData: [
        {"name":"VM1","state":"Running","user":"user01","conn":"TRUE","workspace":"work11"},
        {"name":"VM2","state":"Stopped","user":"user03","conn":"FALSE","workspace":"work112"},
        {"name":"VM3","state":"Running","user":"user02","conn":"TRUE","workspace":"work113"}
      ],
      NWListColumns: [
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
          sorter: (a, b) => (a.state < b.state ? -1 : a.state > b.state ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
      ],
      workspaceUserList: [],
      UserListColumns: [
        {
          dataIndex: "name",
          key: "name",
          slots: { customRender: "nameRender" },
          title: this.$t("label.name"),
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
            title: '',
            key: 'action',
            dataIndex: 'action',
            align: 'right',
            slots: {customRender: 'actionRender'}
        },
        // {
        //   title: this.$t("label.state"),
        //   dataIndex: "state",
        //   key: "state",
        //   sorter: (a, b) => (a.state < b.state ? -1 : a.state > b.state ? 1 : 0),
        //   sortDirections: ["descend", "ascend"],
        // },
        {
          title: this.$t("label.allocateddesktop"),
          dataIndex: "desktop",
          key: "desktop",
          sorter: (a, b) => (a.desktop < b.desktop ? -1 : a.desktop > b.desktop ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
      ],
    }
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      //해당 워크스페이스에 추가된 사용자 목록 조회
      worksApi
        .get("/api/v1/user", { withCredentials: true })
        .then((response) => {
          if (response.status == 200) {
            this.workspaceUserList = response.data.result.result;
          } else {
            message.error(this.t('message.response.data.fail'));
            //console.log(response.message);
          }
        })
        .catch(function (error) {
          message.error(error.message);
          //console.log(error);
        });

      this.loading = ref(false);
    },
    tabReflesh() {
      this.fetchData();
    }
  },
});
</script>

<style>
#ContentTab {
  text-align: left;
}
</style>
