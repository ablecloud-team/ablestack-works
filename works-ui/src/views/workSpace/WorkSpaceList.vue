<template>
  <a-table
    :loading="loading"
    size="middle"
    :columns="listColumns"
    :data-source="wsDataList"
    :bordered="false"
    style="overflow-y: auto; overflow: auto"
    :pagination="pagination"
    :row-selection="{
      selectedRowKeys: state.selectedRowKeys,
      onChange: onSelectChange,
    }"
  >
    <!-- 검색 필터링 template-->
    <template #customFilterDropdown="{ setSelectedKeys, selectedKeys, confirm, clearFilters, column }">
      <div style="padding: 8px">
        <a-input
          ref="searchInput"
          :placeholder="$t('search.' + column.dataIndex)"
          :value="selectedKeys[0]"
          style="width: 188px; margin-bottom: 8px; display: block"
          @change="(e) => setSelectedKeys(e.target.value ? [e.target.value] : [])"
          @pressEnter="handleSearch(selectedKeys, confirm, column.dataIndex)"
        />
        <a-button type="primary" size="small" style="width: 90px; margin-right: 8px" @click="handleSearch(selectedKeys, confirm, column.dataIndex)">
          <template #icon><search-outlined /></template>
          {{ $t("label.search") }}
        </a-button>
        <a-button size="small" style="width: 90px" @click="handleReset(clearFilters)">
          {{ $t("label.reset") }}
        </a-button>
      </div>
    </template>
    <template #customFilterIcon="filtered">
      <search-outlined :style="{ color: filtered ? '#108ee9' : undefined }" />
    </template>
    <!-- 검색 필터링 template-->

    <template #bodyCell="{ column, text, record }">
      <template v-if="column.dataIndex === 'name'">
        <router-link :to="{ path: '/workspaceDetail/' + record.uuid }">
          {{ text }}
        </router-link>
      </template>
      <template v-if="column.dataIndex === 'action'">
        <a-Popover placement="topLeft">
          <template #content>
            <Actions v-if="actionFrom == 'WSList'" :action-from="actionFrom" :ws-info="record" @fetchData="fetchRefresh" />
          </template>
          <MoreOutlined />
        </a-Popover>
      </template>
      <template v-if="column.dataIndex === 'state'">
        <a-tooltip placement="bottom">
          <template #title>{{ record.template_ok_check }}</template>
          <a-badge
            class="head-example"
            :color="
              record.template_ok_check === 'Not Ready' || record.template_ok_check === 'Pending'
                ? 'red'
                : record.template_ok_check === 'Joining' || record.template_ok_check === 'Joined'
                ? 'yellow'
                : record.template_ok_check === 'Ready'
                ? 'green'
                : 'red'
            "
            :text="
              record.template_ok_check === 'Not Ready' || record.template_ok_check === 'Pending'
                ? $t('label.vm.status.initializing') + '(' + record.template_ok_check + ')'
                : record.template_ok_check === 'Joining' || record.template_ok_check === 'Joined'
                ? $t('label.vm.status.configuring') + '(' + record.template_ok_check + ')'
                : record.template_ok_check === 'Ready'
                ? $t('label.vm.status.ready')
                : $t('label.vm.status.notready')
            "
          />
        </a-tooltip>
      </template>

      <template v-if="column.dataIndex === 'workspace_type'">
        {{ record.workspace_type === "desktop" ? $t("label.desktop") : record.workspace_type === "application" ? $t("label.application") : "" }}
      </template>
    </template>
  </a-table>
</template>

<script>
import { defineComponent, ref, reactive } from "vue";
import Actions from "@/components/Actions";
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
        // pageSize: 10,
        showSizeChanger: true, // display can change the number of pages per page
        pageSizeOptions: ["10", "20", "50", "100", "200"], // number of pages per option
        showTotal: (total) => this.$t("label.total") + ` ${total}` + this.$t("label.items"), // show total
        // showSizeChange: (current, pageSize) => (this.pageSize = pageSize), // update display when changing the number of pages per page
      },
      wsDataList: ref([]),
      listColumns: [
        {
          title: this.$t("label.name"),
          dataIndex: "name",
          key: "name",
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
          ellipsis: true,
          customFilterDropdown: true,
          onFilter: (value, record) => record.name.toString().toLowerCase().includes(value.toLowerCase()),
          onFilterDropdownVisibleChange: (visible) => {
            if (visible) {
              setTimeout(() => {
                this.$refs.searchInput.focus();
              }, 100);
            }
          },
        },
        {
          title: this.$t("label.action"),
          key: "action",
          dataIndex: "action",
          width: "50px",
        },
        {
          title: this.$t("label.state"),
          dataIndex: "state",
          key: "state",
          width: "10%",
          sorter: (a, b) => (a.state < b.state ? -1 : a.state > b.state ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.type"),
          dataIndex: "workspace_type",
          key: "workspace_type",
          width: "10%",
          sorter: (a, b) => (a.workspace_type < b.workspace_type ? -1 : a.workspace_type > b.workspace_type ? 1 : 0),
          sortDirections: ["descend", "ascend"],
          // filters: [
          //   {
          //     text: "DESKTOP",
          //     value: "DESKTOP",
          //   },
          //   {
          //     text: "APP",
          //     value: "APP",
          //   },
          // ],
          customFilterDropdown: true,
          onFilter: (value, record) => record.workspace_type.toString().toUpperCase().includes(value.toUpperCase()),
          onFilterDropdownVisibleChange: (visible) => {
            if (visible) {
              setTimeout(() => {
                this.$refs.searchInput.focus();
              }, 100);
            }
          },
        },
        {
          title: this.$t("label.mastertemplate"),
          dataIndex: "master_template_name",
          key: "master_template_name",
          width: "20%",
          sorter: (a, b) => (a.master_template_name < b.master_template_name ? -1 : a.master_template_name > b.master_template_name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.desktop.quantity"),
          dataIndex: "quantity",
          key: "quantity",
          width: "10%",
          sorter: (a, b) => (a.quantity < b.quantity ? -1 : a.quantity > b.quantity ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.description"),
          dataIndex: "description",
          key: "description",
          width: "15%",
          sorter: (a, b) => (a.description < b.description ? -1 : a.description > b.description ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
      ],
    };
  },
  created() {
    this.fetchRefresh();
    this.timer = setInterval(() => {
      //60초 자동 갱신
      this.fetchData();
    }, 60000);
  },
  beforeUnmount() {
    clearInterval(this.timer);
  },
  methods: {
    fetchRefresh() {
      this.$emit("actionFromChange", "WS", null);
      this.loading = true;
      this.actionFrom = "";
      this.state.selectedRowKeys = [];
      this.state.searchText = "";
      this.fetchData();
    },
    onSelectChange(selectedRowKeys, selectedRows) {
      this.state.selectedRowKeys = selectedRowKeys;
      this.state.selectedRows = selectedRows;
      if (this.state.selectedRows.length > 0) {
        this.$emit("actionFromChange", "WSList", this.state.selectedRows);
      } else {
        this.$emit("actionFromChange", "WS", null);
      }
    },
    async fetchData() {
      await this.$worksApi
        .get("/api/v1/workspace")
        .then((response) => {
          if (response.status == 200) {
            if (response.data.result.list !== null) {
              this.wsDataList = response.data.result.list;

              this.wsDataList.forEach((value, index, array) => {
                this.wsDataList[index].key = index;
              });
            } else {
              this.wsDataList = [];
            }
          } else {
            this.$message.error(this.$t("message.response.data.fail"));
          }
        })
        .catch((error) => {
          this.$message.error(this.$t("message.response.data.fail"));
          console.log(error);
        })
        .finally(() => {
          this.loading = false;
        });
      this.actionFrom = "WSList";
    },
  },
});
</script>

<style scoped></style>
