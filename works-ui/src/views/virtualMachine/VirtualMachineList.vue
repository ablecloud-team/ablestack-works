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
    :row-selection="{
      selectedRowKeys: state.selectedRowKeys,
      onChange: onSelectChange,
    }"
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
      <router-link
        :to="{
          path: '/virtualMachineDetail/' + record.uuid,
        }"
        >{{ record.name }}</router-link
      >
    </template>

    <template #actionRender="{ record }">
      <a-Popover placement="topLeft">
        <template #content>
          <Actions
            v-if="actionFrom === 'VirtualMachineList'"
            :action-from="actionFrom"
            :vm-info="record"
            @fetchData="fetchData"
          />
        </template>
        <MoreOutlined />
      </a-Popover>
    </template>
    <template #vmStateRender="{ record }">
      <a-badge
        class="head-example"
        :color="
          record.mold_status === 'Running'
            ? 'green'
            : record.mold_status === 'Stopping' ||
              record.mold_status === 'Starting'
            ? 'blue'
            : record.mold_status === 'Stopped'
            ? 'red'
            : ''
        "
        :text="
          record.mold_status === 'Running'
            ? $t('label.vm.status.running')
            : record.mold_status === 'Starting'
            ? $t('label.vm.status.starting')
            : record.mold_status === 'Stopping'
            ? $t('label.vm.status.stopping')
            : record.mold_status === 'Stopped'
            ? $t('label.vm.status.stopped')
            : ''
        "
      />
    </template>
    <template #vmReadyStateRender="{ record }">
      <a-tooltip placement="bottom">
        <template #title>{{ record.handshake_status }}</template>
        <a-badge
          class="head-example"
          :color="
            record.handshake_status === 'Not Ready' ||
            record.handshake_status === 'Pending'
              ? 'red'
              : record.handshake_status === 'Joining' ||
                record.handshake_status === 'Joined'
              ? 'yellow'
              : record.handshake_status === 'Ready'
              ? 'green'
              : 'red'
          "
          :text="
            record.handshake_status === 'Not Ready' ||
            record.handshake_status === 'Pending'
              ? $t('label.vm.status.initializing') +
                '(' +
                record.handshake_status +
                ')'
              : record.handshake_status === 'Joining' ||
                record.handshake_status === 'Joined'
              ? $t('label.vm.status.configuring') +
                '(' +
                record.handshake_status +
                ')'
              : record.handshake_status === 'Ready'
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
      selectedRowKeys: [],
      selectedRows: [],
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
      loading: ref(false),
      actionFrom: ref(""),
      searchInput,
      state,
      handleSearch,
      handleReset,
    };
  },
  data() {
    return {
      timer: ref(null),
      pagination: {
        pageSize: 10,
        showSizeChanger: true, // display can change the number of pages per page
        pageSizeOptions: ["10", "20", "50", "100", "200"], // number of pages per option
        showTotal: (total) =>
          this.$t("label.total") + ` ${total}` + this.$t("label.items"), // show total
        showSizeChange: (current, pageSize) => (this.pageSize = pageSize), // update display when changing the number of pages per page
      },
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
          width: "15%",
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
          width: "10%",
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
          width: "10%",
          sorter: (a, b) =>
            a.state < b.state ? -1 : a.status > b.status ? 1 : 0,
          sortDirections: ["descend", "ascend"],
          slots: { customRender: "vmStateRender" },
        },
        {
          title: this.$t("label.vm.ready.state"),
          dataIndex: "handshake_status",
          key: "handshake_status",
          width: "20%",
          sorter: (a, b) =>
            a.handshake_status < b.handshake_status
              ? -1
              : a.handshake_status > b.handshake_status
              ? 1
              : 0,
          sortDirections: ["descend", "ascend"],
          slots: { customRender: "vmReadyStateRender" },
        },
        {
          title: this.$t("label.vm.network.ip"),
          dataIndex: "ipaddress",
          key: "ipaddress",
          width: "15%",
          sorter: (a, b) =>
            a.ipaddress < b.ipaddress ? -1 : a.ipaddress > b.ipaddress ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.vm.session.count"),
          dataIndex: "connected",
          key: "connected",
          width: "10%",
          sorter: (a, b) =>
            a.connected < b.connected ? -1 : a.connected > b.connected ? 1 : 0,
          sortDirections: ["descend", "ascend"],
          slots: { customRender: "sessionRender" },
        },
      ],
    };
  },
  created() {
    this.fetchRefresh();
    this.timer = setInterval(() => {
      //60초 자동 갱신
      this.fetchData();
    }, 30000);
  },
  beforeUnmount() {
    clearInterval(this.timer);
  },
  methods: {
    fetchRefresh() {
      this.loading = true;
      this.state.selectedRowKeys = [];
      this.state.searchText = "";
      this.fetchData();
    },
    onSelectChange(selectedRowKeys, selectedRows) {
      this.state.selectedRowKeys = selectedRowKeys;
      this.state.selectedRows = selectedRows;
      if (this.state.selectedRows.length > 0) {
        this.$emit(
          "actionFromChange",
          "VirtualMachineList",
          this.state.selectedRows
        );
      } else {
        this.$emit("actionFromChange", "VirtualMachine", null);
      }
    },
    async fetchData() {
      this.actionFrom = "";
      await worksApi
        .get("/api/v1/instance/all")
        .then((response) => {
          if (response.status == 200) {
            if (response.data.result.instanceInfo !== null) {
              this.vmDataList = response.data.result.instanceInfo;
              this.vmDataList.forEach((value, index, array) => {
                this.vmDataList[index].key = index;
              });
            }
          } else {
            message.error(this.$t("message.response.data.fail"));
          }
        })
        .catch((error) => {
          message.error(this.$t("message.response.data.fail"));
          console.log(error);
        })
        .finally(() => {
          this.loading = false;
        });

      this.actionFrom = "VirtualMachineList";
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
