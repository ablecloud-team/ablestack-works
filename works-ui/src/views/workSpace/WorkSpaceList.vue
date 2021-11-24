<template>
  <a-table
    :loading="loading"
    size="middle"
    class="ant-table-striped"
    :columns="listColumns"
    :data-source="wsDataList"
    :row-class-name="
      (record, index) => (index % 2 === 1 ? 'dark-row' : 'light-row')
    "
    :bordered="false"
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
      <router-link :to="{ path: '/workspaceDetail/' + record.uuid }">
        {{ record.name }}
      </router-link>
    </template>
    <template #actionRender="{ record }">
      <a-Popover placement="topLeft">
        <template #content>
          <Actions
            :action-from="actionFrom"
            :workspace-info="record"
            @fetchData="fetchRefresh"
          />
        </template>
        <MoreOutlined />
      </a-Popover>
    </template>
    <template #stateRender="{ record }">
      <a-tooltip placement="bottom">
        <template #title>{{ record.template_ok_check }}</template>
        <a-badge
          class="head-example"
          :color="
            record.template_ok_check === 'Not Ready' ||
            record.template_ok_check === 'Pending'
              ? 'red'
              : record.template_ok_check === 'Joining' ||
                record.template_ok_check === 'Joined'
              ? 'yellow'
              : record.template_ok_check === 'Ready'
              ? 'green'
              : 'red'
          "
          :text="
            record.template_ok_check === 'Not Ready' ||
            record.template_ok_check === 'Pending'
              ? $t('label.vm.status.initializing') +
                '(' +
                record.template_ok_check +
                ')'
              : record.template_ok_check === 'Joining' ||
                record.template_ok_check === 'Joined'
              ? $t('label.vm.status.configuring') +
                '(' +
                record.template_ok_check +
                ')'
              : record.template_ok_check === 'Ready'
              ? $t('label.vm.status.ready')
              : $t('label.vm.status.notready')
          "
        />
      </a-tooltip>
    </template>
    <template #typeRender="{ record }">
      {{ record.workspace_type.toUpperCase() }}
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
    };
    const handleReset = (clearFilters) => {
      clearFilters();
      state.searchText = "";
    };
    return {
      loading: ref(false),
      actionFrom: ref("WorkspaceList"),
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
      wsDataList: ref([]),
      listColumns: [
        {
          title: this.$t("label.name"),
          dataIndex: "name",
          key: "name",
          width: "37%",
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
          title: this.$t("label.state"),
          dataIndex: "state",
          key: "state",
          width: "20%",
          sorter: (a, b) =>
            a.state < b.state ? -1 : a.state > b.state ? 1 : 0,
          sortDirections: ["descend", "ascend"],
          slots: { customRender: "stateRender" },
        },
        {
          title: this.$t("label.type"),
          dataIndex: "workspace_type",
          key: "workspace_type",
          width: "20%",
          slots: {
            customRender: "typeRender",
            filterDropdown: "filterDropdown",
            filterIcon: "filterIcon",
          },
          sorter: (a, b) =>
            a.workspace_type < b.workspace_type
              ? -1
              : a.workspace_type > b.workspace_type
              ? 1
              : 0,
          sortDirections: ["descend", "ascend"],
          //onFilter: (value, record) => record.Type.toUpperCase().indexOf(value) === 0,
          //filterMultiple: false,
          // filters: [
          //   {
          //     text: 'DESKTOP',
          //     value: 'DESKTOP',
          //   },
          //   {
          //     text: 'APP',
          //     value: 'APP',
          //   },
          // ]
          onFilter: (value, record) =>
            record.workspace_type
              .toString()
              .toLowerCase()
              .includes(value.toLowerCase()),
          onFilterDropdownVisibleChange: (visible) => {
            if (visible) {
              setTimeout(() => {
                this.$refs.searchInput.focus();
              }, 100);
            }
          },
        },
        {
          title: this.$t("label.desktop.quantity"),
          dataIndex: "quantity",
          key: "quantity",
          width: "20%",
          sorter: (a, b) =>
            a.quantity < b.quantity ? -1 : a.quantity > b.quantity ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
        // {
        //   title: this.$t("label.desktop.connection.quantity"),
        //   dataIndex: "NoC",
        //   key: "NoC",
        //   sorter: (a, b) => (a.NoC < b.NoC ? -1 : a.NoC > b.NoC ? 1 : 0),
        //   sortDirections: ["descend", "ascend"],
        // },
        // {
        //   title: this.$t("label.network.type"),
        //   dataIndex: "NetType",
        //   key: "NetType",
        //   sorter: (a, b) =>
        //     a.NetType < b.NetType ? -1 : a.NetType > b.NetType ? 1 : 0,
        //   sortDirections: ["descend", "ascend"],
        // },
        // {
        //   title: this.$t("label.restrict"),
        //   dataIndex: "Restrict",
        //   key: "Restrict",
        //   sorter: (a, b) =>
        //     a.Restrict < b.Restrict ? -1 : a.Restrict > b.Restrict ? 1 : 0,
        //   sortDirections: ["descend", "ascend"],
        // },
        // ,
        // {
        //     title: 'Tags',
        //     key: 'tags',
        //     dataIndex: 'tag',
        //     slots: {customRender: 'tags'},
        // }
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
          "WorkspaceList",
          this.state.selectedRows
        );
      } else {
        this.$emit("actionFromChange", "Workspace", null);
      }
    },
    async fetchData() {
      await worksApi
        .get("/api/v1/workspace")
        .then((response) => {
          if (response.status == 200) {
            if (response.data.result.list !== null) {
              this.wsDataList = response.data.result.list;

              this.wsDataList.forEach((value, index, array) => {
                this.wsDataList[index].key = index;
              });
            } else {
              this.wsDataList = ref([]);
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
    },
  },
});
</script>

<style scoped>
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
</style>

<style scoped lang="scss">
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
