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
    <template #nameRender="{ record }">
        <router-link :to="{ path: '/accountDetail/'+record.name}">{{ record.name }}</router-link>
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
import { defineComponent, ref } from "vue";
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
  setup() {
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
          width: "39%",
          slots: { customRender: "nameRender" },
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: "",
          key: "action",
          dataIndex: "action",
          align: "right",
          width: "1%",
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
          width: "30%",
          sorter: (a, b) => (a.mail < b.mail ? -1 : a.mail > b.mail ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.allocateddesktop"),
          dataIndex: "desktop",
          key: "desktop",
          width: "30%",
          sorter: (a, b) => (a.desktop < b.desktop ? -1 : a.desktop > b.desktop ? 1 : 0),
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
        .get("/api/v1/user", { withCredentials: true })
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
