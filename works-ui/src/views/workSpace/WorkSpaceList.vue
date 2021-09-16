<template>
  <a-table
    :loading="loading"
    size="middle"
    class="ant-table-striped"
    :columns="listColumns"
    :data-source="dataList"
    :row-class-name="
      (record, index) => (index % 2 === 1 ? 'dark-row' : 'light-row')
    "
    :bordered="false"
    style="overflow-y: auto; overflow: auto"
    :row-key="(record, index) => index"
    :pagination="pagination"
  >
    <!-- 검색 필터링 template-->
    <template #filterDropdown="{ setSelectedKeys, selectedKeys, confirm, clearFilters, column }">
      <div style="padding: 8px">
        <a-input-search
          ref="searchInput"
          :placeholder="$t('search.'+column.dataIndex)"
          :value="selectedKeys[0]"
          @change="e => setSelectedKeys(e.target.value ? [e.target.value] : [])"
          @pressEnter="handleSearch(selectedKeys, confirm, column.dataIndex)"
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
      <router-link :to="{ path: '/workspaceDetail/' + record.uuid + '/' + record.name }">{{ record.name }}</router-link>
    </template>
    <template #actionRender>
      <a-Popover placement="topLeft">
        <template #content>
          <actions :actionFrom="actionFrom" />
        </template>
        <MoreOutlined />
      </a-Popover>
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
import { SearchOutlined } from "@ant-design/icons-vue";
// const rowSelection = {
//   onChange: (selectedRowKeys, selectedRows) => {
//     // console.log(
//     //   `selectedRowKeys: ${selectedRowKeys}`,
//     //   "selectedRows: ",
//     //   selectedRows
//     // );
//   },
//   onSelect: (record, selected, selectedRows) => {
//     // console.log(record, selected, selectedRows);
//   },
//   onSelectAll: (selected, selectedRows, changeRows) => {
//     // console.log(selected, selectedRows, changeRows);
//   },
// };

export default defineComponent({
  name: "WorkspaceList",
  components: {
    Actions,
    SearchOutlined,
  },
  props: {},
  setup() {
    const state = reactive({
      searchText: '',
      searchedColumn: '',
    });
    const searchInput = ref();
    const handleSearch = (selectedKeys, confirm, dataIndex) => {
      confirm();
      state.searchText = selectedKeys[0];
      state.searchedColumn = dataIndex;
    };
    const handleReset = clearFilters => {
      clearFilters();
      state.searchText = '';
    };
    return {
      //rowSelection,
      loading: ref(false),
      actionFrom: ref("WorkspaceList"),
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
      selectedRowKeys: [],
      dataList: [],
      listColumns: [
        {
          title: this.$t("label.name"),
          dataIndex: "name",
          key: "name",
          width: "39%",
          slots: { 
            customRender: "nameRender",
            filterDropdown: 'filterDropdown',
            filterIcon: 'filterIcon',
          },
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
          onFilter: (value, record) => record.Name.toString().toLowerCase().includes(value.toLowerCase()),
          onFilterDropdownVisibleChange: visible => {
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
          width: "1%",
          slots: { customRender: "actionRender" },
        },
        {
          title: this.$t("label.state"),
          dataIndex: "state",
          key: "state",
          width: "20%",
          sorter: (a, b) => (a.state < b.state ? -1 : a.state > b.state ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.type"),
          dataIndex: "workspace_type",
          key: "workspace_type",
          width: "20%",
          slots: { 
            customRender: "typeRender",
            filterDropdown: 'filterDropdown',
            filterIcon: 'filterIcon',
          },
          sorter: (a, b) => (a.workspace_type < b.workspace_type ? -1 : a.workspace_type > b.workspace_type ? 1 : 0),
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
          onFilter: (value, record) => record.workspace_type.toString().toLowerCase().includes(value.toLowerCase()),
          onFilterDropdownVisibleChange: visible => {
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
          sorter: (a, b) => (a.quantity < b.quantity ? -1 : a.quantity > b.quantity ? 1 : 0),
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
    this.fetchData();
  },
  methods: {
    setSelection(selection) {
      this.selectedRowKeys = selection;
      if (this.selectedRowKeys.length == 0) {
        this.$emit("actionFromChange", "Workspace");
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
    fetchData() {
      this.loading = true;
      worksApi
        .get("/api/v1/workspace", { withCredentials: true })
        .then((response) => {
          if (response.status == 200) {
            this.dataList = response.data.result.list;
          } else {
            message.error(this.$t("message.response.data.fail"));
            //console.log(response.message);
          }
        })
        .catch(function (error) {
          message.error(error.message);
          //console.log(error);
        });
      setTimeout(() => {
        this.loading = false;
      }, 500);
    },
  },
});
</script>

<style scoped>
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
