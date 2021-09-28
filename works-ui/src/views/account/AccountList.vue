<template>
  <a-table
    size="middle"
    :loading="loading"
    class="ant-table-striped"
    :columns="UserListColumns"
    :data-source="userDataList"
    :row-class-name="
      (record, index) => (index % 2 === 1 ? 'dark-row' : 'light-row')
    "
    :bordered="false"
    style="overflow-y: auto; overflow: auto"
    :row-key="(record, index) => index"
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
      <router-link :to="{ path: '/accountDetail/' + record.name }">{{
        record.name
      }}</router-link>
    </template>
    <template #actionRender>
      <a-Popover placement="topLeft">
        <template #content>
          <Actions :action-from="actionFrom" />
        </template>
        <MoreOutlined />
      </a-Popover>
    </template>
  </a-table>
</template>

<script>
import { defineComponent, ref, reactive } from "vue";
import Actions from "@/components/Actions";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";

// const rowSelection = {
//   onChange: (selectedRowKeys, selectedRows) => {
//     console.log(
//       `selectedRowKeys: ${selectedRowKeys}`,
//       "selectedRows: ",
//       selectedRows
//     );
//   },
//   onSelect: (record, selected, selectedRows) => {
//     console.log(record, selected, selectedRows);
//   },
//   onSelectAll: (selected, selectedRows, changeRows) => {
//     console.log(selected, selectedRows, changeRows);
//   },
// };
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
      //rowSelection,
      loading: ref(false),
      actionFrom: ref("AccountList"),
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
      userDataList: [],
      UserListColumns: [
        {
          title: this.$t("label.account"),
          dataIndex: "name",
          key: "name",
          width: "47%",
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
        // {
        //   title: this.$t('label.state'),
        //   dataIndex: "state",
        //   key: "state",
        //   sorter: (a, b) => (a.state < b.state ? -1 : a.state > b.state ? 1 : 0),
        //   sortDirections: ["descend", "ascend"],
        // },
        {
          title: this.$t("label.email"),
          dataIndex: "mail",
          key: "mail",
          width: "50%",
          sorter: (a, b) => (a.mail < b.mail ? -1 : a.mail > b.mail ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
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
        this.$emit("actionFromChange", "User");
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
        .get("/api/v1/user")
        .then((response) => {
          if (response.data.result.status == 200) {
            this.userDataList = response.data.result.result;
          } else {
            message.error(this.t("message.response.data.fail"));
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
