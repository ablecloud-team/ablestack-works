<template>
  <a-table
    size="middle"
    :loading="loading"
    class="ant-table-striped"
    :columns="columns"
    :data-source="dataSource"
    :row-class-name="
      (record, index) => (index % 2 === 1 ? 'dark-row' : 'light-row')
    "
    :bordered="false"
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
      </div>
    </template>
    <template #filterIcon="filtered">
      <SearchOutlined :style="{ color: filtered ? '#108ee9' : undefined }" />
    </template>
    <!-- 검색 필터링 template-->
    <template v-for="col in ['value']" #[col]="{ text, record }" :key="col">
      <div>
        <a-input
          v-if="editableData[record.name]"
          v-model:value="editableData[record.name][col]"
          style="margin: -5px 0"
        />
        <template v-else>
          {{ text }}
        </template>
      </div>
    </template>
    <template #actionRender="{ record }">
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
    const dataSource = ref([
      {
        id: 1,
        name: "mold.default.domain.account",
        category: "Advanced",
        value: "works",
        default_value: null,
        description: "Mold 의 domain 계정",
        update_date: null,
      },
      {
        id: 2,
        name: "mold.default.domain.id",
        category: "Advanced",
        value: "d294bfc7-f86e-4f67-ac83-80220110b23f",
        default_value: null,
        description: "Mold 의 domain ID",
        update_date: null,
      },
      {
        id: 3,
        name: "mold.default.api.key",
        category: "Advanced",
        value:
          "WuiJ_rWUllvp5a9Wkj1ZzwE_VqShM-3eJr1O2Jvi4RmaIfsWgbJh-GKxMpJ78V47wVX8BWACi9KilsaRusjZ7w",
        default_value: null,
        description: "Mold 의 도메인 관리자의 API Key 를 설정하는 항목입니다.",
        update_date: null,
      },
      {
        id: 4,
        name: "mold.default.secret.key",
        category: "Advanced",
        value:
          "Y3sJNJuK6fHuMigdOYqaJHAGgVaIb9RDUTfTfAAeQzR7UAUPUfcOjhekIHStirtLIfgb6Nbre_x1Lz9S7c7HHQ",
        default_value: null,
        description:
          "Mold 의 도메인 관리자의 Secret Key 를 설정하는 항목입니다.",
        update_date: null,
      },
      {
        id: 5,
        name: "mold.default.protocol",
        category: "Advanced",
        value: "http://",
        default_value: null,
        description: "Mold 의 프로토콜 타입",
        update_date: null,
      },
      {
        id: 6,
        name: "mold.default.url",
        category: "Advanced",
        value: "10.10.1.10",
        default_value: null,
        description: "Mold 의 URL",
        update_date: null,
      },
      {
        id: 7,
        name: "mold.default.port",
        category: "Advanced",
        value: "8080",
        default_value: null,
        description: "Mold 의 Port",
        update_date: null,
      },
      {
        id: 8,
        name: "mold.default.url.postfix",
        category: "Advanced",
        value: "/client/api",
        default_value: null,
        description: "Mold 의 URL 의 Postfix",
        update_date: null,
      },
      {
        id: 9,
        name: "dc.default.protocol",
        category: "Advanced",
        value: "http://",
        default_value: null,
        description: "Domain Controller 의 프로토콜 타입",
        update_date: null,
      },
      {
        id: 10,
        name: "dc.default.url",
        category: "Advanced",
        value: "10.1.1.172",
        default_value: null,
        description: "Domain Controller 의 URL",
        update_date: null,
      },
      {
        id: 11,
        name: "dc.default.port",
        category: "Advanced",
        value: "8083",
        default_value: null,
        description: "Domain Controller 의 Port",
        update_date: null,
      },
      {
        id: 12,
        name: "dc.default.url.postfix",
        category: "Advanced",
        value: "/api",
        default_value: null,
        description: "Domain Controller 의 URL 의 Postfix",
        update_date: null,
      },
      {
        id: 13,
        name: "mold.network.uuid",
        category: "Advanced",
        value: "205",
        default_value: null,
        description: "Works 에서 사용할 Mold network uuid",
        update_date: null,
      },
      {
        id: 14,
        name: "mold.zone.id",
        category: "Advanced",
        value: "82bfaf86-ea6b-4cbc-88c5-0a951fd57aa0",
        default_value: null,
        description: "Works 가 생성된 Zone ID",
        update_date: null,
      },
      {
        id: 15,
        name: "works.default.url",
        category: "Advanced",
        value: "10.1.1.171",
        default_value: null,
        description: "Works 의 URL",
        update_date: null,
      },
      {
        id: 16,
        name: "works.default.port",
        category: "Advanced",
        value: "8082",
        default_value: null,
        description: "Works 의 Port",
        update_date: null,
      },
      {
        id: 17,
        name: "samba.default.url",
        category: "Advanced",
        value: "10.1.1.171",
        default_value: null,
        description: "Samba 의 URL",
        update_date: null,
      },
      {
        id: 18,
        name: "samba.default.domain",
        category: "Advanced",
        value: "cl",
        default_value: null,
        description: "Samba 의 Domain",
        update_date: null,
      },
      {
        id: 19,
        name: "guacamole.default.protocol",
        category: "Advanced",
        value: "http://",
        default_value: null,
        description: "guacamole 의 protocol",
        update_date: null,
      },
      {
        id: 20,
        name: "guacamole.default.url",
        category: "Advanced",
        value: "10.10.1.100",
        default_value: null,
        description: "guacamole 의 URL",
        update_date: null,
      },
      {
        id: 21,
        name: "guacamole.default.port",
        category: "Advanced",
        value: "8080",
        default_value: null,
        description: "guacamole 의 Port",
        update_date: null,
      },
      {
        id: 22,
        name: "guacamole.default.url.postfix",
        category: "Advanced",
        value: "/guacamole/api",
        default_value: null,
        description: "guacamole 의 URL 의 Postfix",
        update_date: null,
      },
      {
        id: 23,
        name: "guacamole.default.username",
        category: "Advanced",
        value: "guacadmin",
        default_value: null,
        description: "guacamole 의 기본 도메인",
        update_date: null,
      },
      {
        id: 24,
        name: "cluster.default.name",
        category: "Advanced",
        value: "cl",
        default_value: null,
        description: "Cluster 의 이름",
        update_date: null,
      },
      {
        id: 25,
        name: "log.default.level",
        category: "Advanced",
        value: "logrus.DebugLevel",
        default_value: null,
        description: "Works-API 의 Log Level",
        update_date: null,
      },
    ]);
    const editableData = reactive({});
    const edit = (name) => {
      editableData[name] = cloneDeep(
        dataSource.value.filter((item) => name === item.name)[0]
      );
    };

    const save = (name) => {
      Object.assign(
        dataSource.value.filter((item) => name === item.name)[0],
        editableData[name]
      );
      delete editableData[name];
    };

    const cancel = (name) => {
      delete editableData[name];
    };

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
      dataSource,
      editingKey: "",
      editableData,
      edit,
      save,
      cancel,
    };
  },
  data() {
    return {
      dataList: [],
      columns: [
        {
          title: this.$t("label.account"),
          dataIndex: "name",
          width: "25%",
          slots: {
            customRender: "name",
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
          title: this.$t("label.description"),
          dataIndex: "description",
          width: "30%",
          slots: { customRender: "description" },
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
          slots: { customRender: "value" },
          sorter: (a, b) =>
            a.value < b.value ? -1 : a.value > b.value ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.action"),
          dataIndex: "action",
          align: "center",
          width: "15%",
          slots: { customRender: "actionRender" },
        },
      ],
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      this.loading = true;
      // worksApi
      //   .get("/api/v1/user")
      //   .then((response) => {
      //     if (response.data.result.status == 200) {
      //       this.userDataList = response.data.result.result;
      //     } else {
      //       message.error(this.$t("message.response.data.fail"));
      //       //console.log(response.message);
      //     }
      //   })
      //   .catch((error) => {
      //     message.error(this.$t("message.response.data.fail"));
      //     //console.log(error);
      //   });
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

.editable-row-operations a {
  margin-right: 8px;
}
</style>
