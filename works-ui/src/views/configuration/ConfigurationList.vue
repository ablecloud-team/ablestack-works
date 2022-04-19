<template>
  <a-table
    size="middle"
    :loading="loading"
    class="ant-table-striped"
    :columns="columns"
    :data-source="dataList"
    :row-class-name="
      (record, index) => (index % 2 === 1 ? 'dark-row' : 'light-row')
    "
    :bordered="false"
    style="overflow-y: auto; overflow: auto"
    :pagination="pagination"
  >
    <!-- 검색 필터링 template-->
    <template
      #customFilterDropdown="{
        setSelectedKeys,
        selectedKeys,
        confirm,
        clearFilters,
        column,
      }"
    >
      <div style="padding: 8px">
        <a-input
          ref="searchInput"
          :placeholder="$t('search.' + column.dataIndex)"
          :value="selectedKeys[0]"
          style="width: 188px; margin-bottom: 8px; display: block"
          @change="
            (e) => setSelectedKeys(e.target.value ? [e.target.value] : [])
          "
          @pressEnter="handleSearch(selectedKeys, confirm, column.dataIndex)"
        />
        <a-button
          type="primary"
          size="small"
          style="width: 90px; margin-right: 8px"
          @click="handleSearch(selectedKeys, confirm, column.dataIndex)"
        >
          <template #icon><search-outlined /></template>
          {{ $t("label.search") }}
        </a-button>
        <a-button
          size="small"
          style="width: 90px"
          @click="handleReset(clearFilters)"
        >
          {{ $t("label.reset") }}
        </a-button>
      </div>
    </template>
    <template #customFilterIcon="{ filtered }">
      <search-outlined :style="{ color: filtered ? '#108ee9' : undefined }" />
    </template>
    <!-- 검색 필터링 template-->

    <template #bodyCell="{ column, text, record }">
      <template v-if="column.dataIndex === 'value'">
        <div>
          <a-input
            v-if="editableData[record.name]"
            v-model:value="editableData[record.name][column.dataIndex]"
            style="margin: -5px 0"
          />
          <template v-else>
            {{ text }}
          </template>
        </div>
      </template>
      <template v-if="column.dataIndex === 'action'">
        <div class="editable-row-operations">
          <span v-if="editableData[record.name]">
            <a-button shape="circle" @click="cancel(record.name)">
              <CloseCircleOutlined style="color: red" />
            </a-button>
            <a-button shape="circle" @click="save(record.name)">
              <CheckCircleOutlined style="color: #52c41a" />
            </a-button>
          </span>
          <span v-else>
            <a-button shape="circle" @click="edit(record.name)">
              <EditOutlined />
            </a-button>
          </span>
        </div>
      </template>
    </template>
  </a-table>
</template>

<script>
import { cloneDeep } from "lodash-es";
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
      //rowSelection,
      loading: ref(false),
      actionFrom: ref("AccountList"),
      searchInput,
      state,
      handleSearch,
      handleReset,
      editingKey: "",
      editableData: reactive({}),
    };
  },
  data() {
    return {
      pagination: {
        // pageSize: 10,
        showSizeChanger: true, // display can change the number of pages per page
        pageSizeOptions: ["10", "20", "50", "100", "200"], // number of pages per option
        showTotal: (total) =>
          this.$t("label.total") + ` ${total}` + this.$t("label.items"), // show total
        // showSizeChange: (current, pageSize) => (this.pageSize = pageSize), // update display when changing the number of pages per page
      },
      dataList: [],
      columns: [
        {
          title: this.$t("label.name"),
          dataIndex: "name",
          width: "25%",
          // slots: {
          //   customRender: "name",
          //   filterDropdown: "filterDropdown",
          //   filterIcon: "filterIcon",
          // },
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
          ellipsis: true,
          customFilterDropdown: true,
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
          title: this.$t("label.description"),
          dataIndex: "description",
          width: "30%",
          sorter: (a, b) =>
            a.description < b.description
              ? -1
              : a.description > b.description
              ? 1
              : 0,
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.value"),
          dataIndex: "value",
          width: "30%",
          // slots: { customRender: "value" },
          sorter: (a, b) =>
            a.value < b.value ? -1 : a.value > b.value ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.action"),
          dataIndex: "action",
          align: "center",
          width: "15%",
          // slots: { customRender: "actionRender" },
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
      this.loading = true;
      this.state.searchText = "";
      this.fetchData();
    },
    async fetchData() {
      await worksApi
        .get("/api/v1/configuration")
        .then((response) => {
          if (response.status == 200) {
            if (response.data.result !== null) {
              this.dataList = response.data.result;
            } else {
              this.wsDataList = [];
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
    edit(name) {
      this.editableData[name] = cloneDeep(
        this.dataList.filter((item) => name === item.name)[0]
      );
    },
    async save(name) {
      Object.assign(
        this.dataList.filter((item) => name === item.name)[0],
        this.editableData[name]
      );

      console.log(this.editableData[name].name, this.editableData[name].value);
      let params = new URLSearchParams();

      params.append("id", this.editableData[name].id);
      params.append("value", this.editableData[name].value);

      try {
        const res = await worksApi.patch(
          "/api/v1/configuration/" + this.editableData[name].id,
          params
        );
        if (res.status == 200) {
          message.success(
            this.$t("message.configuration.data.update.success", {
              name: this.editableData[name].name,
            })
          );
        } else {
          message.error(
            this.$t("message.configuration.data.update.fail", {
              name: this.editableData[name].name,
            })
          );
        }
      } catch (error) {
        message.error(
          this.$t("message.configuration.data.update.fail", {
            name: this.editableData[name].name,
          })
        );
        console.log(error);
      }
      delete this.editableData[name];

      this.fetchRefresh();
    },
    cancel(name) {
      delete this.editableData[name];
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

.editable-row-operations a {
  margin-right: 8px;
}
</style>
