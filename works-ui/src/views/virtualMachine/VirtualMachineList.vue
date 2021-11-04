<template>
  <a-table
    size="middle"
    :loading="loading"
    class="ant-table-striped"
    :columns="vmListColumns"
    :data-source="vmDataList"
    :row-class-name="
      (record, index) => (index % 2 === 1 ? 'dark-row' : 'light-row')
    "
    :bordered="bordered ? bordered : false"
    style="overflow-y: auto; overflow: auto"
    :pagination="pagination"
  >
    <!-- 검색 필터링 template-->
    <template
      #filterDropdown="{ setSelectedKeys, selectedKeys, confirm, column }"
    >
      <div style="padding: 8px">
        <a-input-search
          ref="searchInput"
          :placeholder="$t('search.' + column.dataIndex)"
          :value="selectedKeys[0]"
          @change="
            (e) => setSelectedKeys(e.target.value ? [e.target.value] : [])
          "
          @search="handleSearch(selectedKeys, confirm, column.dataIndex)"
        />
        <!-- <a-button
          type="primary"
          size="small"
          style="width: 90px; margin-right: 8px"
          @click="handleSearch(selectedKeys, confirm, column.dataIndex)"
        >
          <template #icon><SearchOutlined /></template>
          {{ $t("label.search")}}
        </a-button>
        <a-button size="small" style="width: 90px" @click="handleReset(clearFilters)">
          {{ $t("label.reset")}}
        </a-button> -->
      </div>
    </template>
    <template #filterIcon="filtered">
      <SearchOutlined :style="{ color: filtered ? '#108ee9' : undefined }" />
    </template>
    <!-- 검색 필터링 template-->

    <template #nameRender="{ record }">
      <router-link :to="{ path: '/virtualMachineDetail/' + record.uuid +'/' + record.name}">{{
        record.name
      }}</router-link>
    </template>

    <template #actionRender="{ record }">
      <a-Popover placement="topLeft">
        <template #content>
          <Actions
            :action-from="actionFrom"
            :vm-uuid="record.uuid"
            @fetchData="fetchData"
          />
        </template>
        <MoreOutlined />
      </a-Popover>
    </template>
    <template #vmStateRender="{ record }">
      <a-badge
        class="head-example"
        :color="record.mold_status === 'Running' ? 'green' : 'red'"
        :text="
          record.mold_status === 'Running'
            ? $t('label.vm.status.running')
            : $t('label.vm.status.stopped')
        "
      />
    </template>
    <template #vmReadyStateRender="{ record }">
      <a-tooltip placement="bottom">
        <template #title>{{ record.handshake_status }}</template>
          <a-badge
            class="head-example"
            :color="'red'"
            :text="
              (record.mold_status == 'Running' && record.handshake_status === 'Ready') 
              ? $t('label.vm.status.ready')
              : $t('label.vm.status.notready')
            "
          />
      </a-tooltip>
    </template>
    <template #userRender="{ record }">
      <!-- {{
        record.owner_account_id === ""
          ? $t("label.owner.account.false")
          : record.owner_account_id
      }} -->
      {{ record.owner_account_id }}
    </template>
    <template #sessionRender="{ record }">
      {{ record.connected }}
    </template>
  </a-table>
</template>

<script>
import { defineComponent, ref, reactive } from "vue";
import Actions from "@/components/Actions";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";
export default defineComponent({
  components: {
    Actions,
  },
  props: {},
  emits: ["actionFromChange"],
  setup() {
    const state = reactive({
      searchText: "",
      searchedColumn: "",
    });
    const searchInput = ref();
    const handleSearch = (selectedKeys, confirm, dataIndex) => {
      confirm();
      state.searchText = selectedKeys[0];
      state.searchedColumn = dataIndex;
      //console.log(selectedKeys + "  ::  " + confirm + "  ::  " +dataIndex);
    };
    const handleReset = (clearFilters) => {
      clearFilters();
      state.searchText = "";
    };
    return {
      // rowSelection,
      loading: ref(false),
      actionFrom: ref("VirtualMachineList"),
      pagination: {
        pageSize: 10,
        showSizeChanger: true, // display can change the number of pages per page
        pageSizeOptions: ["10", "20", "50", "100"], // number of pages per option
        showTotal: (total) => `Total ${total} items`, // show total
        showSizeChange: (current, pageSize) => (this.pageSize = pageSize), // update display when changing the number of pages per page
      },
      searchInput,
      state,
      handleSearch,
      handleReset,
    };
  },
  data() {
    return {
      timer: ref(null),
      selectedRowKeys: [],
      vmDataList: [],
      vmListColumns: [
        {
          title: this.$t("label.name"),
          dataIndex: "name",
          key: "name",
          width: "22%",
          slots: {
            customRender: "nameRender",
            filterDropdown: "filterDropdown",
            filterIcon: "filterIcon",
          },
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
          ellipsis: true,
          onFilter: (value, record) =>
            record.name.toString().toLowerCase().includes(value.toLowerCase()),
          onFilterDropdownVisibleChange: (visible) => {
            if (visible) {
              setTimeout(() => {
                this.$refs.searchInput.focus();
              }, 100);
            }
          },
        },
        {
          title: "",
          key: "action",
          dataIndex: "action",
          align: "right",
          width: "3%",
          slots: { customRender: "actionRender" },
        },
        {
          title: this.$t("label.workspace"),
          dataIndex: "workspace_name",
          key: "workspace_name",
          width: "20%",
          sorter: (a, b) =>
            a.workspace_name < b.workspace_name
              ? -1
              : a.workspace_name > b.workspace_name
              ? 1
              : 0,
          sortDirections: ["descend", "ascend"],
          ellipsis: true,
        },
        {
          title: this.$t("label.users"),
          dataIndex: "owner_account_id",
          key: "owner_account_id",
          width: "11%",
          sorter: (a, b) =>
            a.owner_account_id < b.owner_account_id
              ? -1
              : a.owner_account_id > b.owner_account_id
              ? 1
              : 0,
          sortDirections: ["descend", "ascend"],
          slots: { customRender: "userRender" },
        },
        {
          title: this.$t("label.vm.state"),
          dataIndex: "status",
          key: "status",
          width: "11%",
          sorter: (a, b) =>
            a.state < b.state ? -1 : a.status > b.status ? 1 : 0,
          sortDirections: ["descend", "ascend"],
          slots: { customRender: "vmStateRender" },
        },
        {
          title: this.$t("label.vm.ready.state"),
          dataIndex: "status",
          key: "status",
          width: "11%",
          sorter: (a, b) =>
            a.status < b.status ? -1 : a.status > b.status ? 1 : 0,
          sortDirections: ["descend", "ascend"],
          slots: { customRender: "vmReadyStateRender" },
        },
        {
          title: this.$t("label.vm.network.ip"),
          dataIndex: "ipaddress",
          key: "ipaddress",
          width: "11%",
          sorter: (a, b) =>
            a.ipaddress < b.ipaddress ? -1 : a.ipaddress > b.ipaddress ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.vm.session.count"),
          dataIndex: "connected",
          key: "connected",
          width: "11%",
          sorter: (a, b) =>
            a.connected < b.connected ? -1 : a.connected > b.connected ? 1 : 0,
          sortDirections: ["descend", "ascend"],
          slots: { customRender: "sessionRender" },
        },
      ],
    };
  },
  created() {
    this.fetchData();
    this.timer = setInterval(() => {
      //10초 자동 갱신
      this.fetchData();
    }, 15000);
  },
  beforeUnmount() {
    clearInterval(this.timer);
  },
  methods: {
    setSelection(selection) {
      this.selectedRowKeys = selection;
      if (this.selectedRowKeys.length == 0) {
        this.$emit("actionFromChange", "VirtualMachine");
      } else {
        this.$emit("actionFromChange", this.actionFrom);
      }
    },
    resetSelection() {
      this.setSelection([]);
    },
    onSelectChange(selectedRowKeys, selectedRows) {
      this.setSelection(selectedRowKeys);
    },
    fetchData(val) {
      this.loading = true;
      worksApi
        .get("/api/v1/instance/all")
        .then((response) => {
          if (response.status == 200) {
            this.vmDataList = response.data.result.instanceInfo;
          } else {
            message.error(this.$t("message.response.data.fail"));
          }
        })
        .catch(function (error) {
          message.error(error);
          //console.log(error);
        });
      setTimeout(() => {
        this.loading = false;
      }, 500);
    },
  },
});
</script>
<style scoped lang="scss">
::v-deep(.ant-badge-status-dot) {
  width: 12px;
  height: 12px;
}
::v-deep(.ant-table-thead) {
  background-color: #f9f9f9;
}

::v-deep(.ant-table-small) > .ant-table-content > .ant-table-body {
  margin: 0;
}

::v-deep(.light-row) {
  background-color: #fff;
}

::v-deep(.dark-row) {
  background-color: #f9f9f9;
}

.shift-btns {
  display: flex;
}
.shift-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 20px;
  height: 20px;
  font-size: 12px;

  &:not(:last-child) {
    margin-right: 5px;
  }

  &--rotated {
    font-size: 10px;
    transform: rotate(90deg);
  }
}

.alert-notification-threshold {
  background-color: rgba(255, 231, 175, 0.75);
  color: #e87900;
  padding: 10%;
}

.alert-disable-threshold {
  background-color: rgba(255, 190, 190, 0.75);
  color: #f50000;
  padding: 10%;
}
</style>
