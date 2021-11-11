<template>
  <a-table
    size="middle"
    :loading="loading"
    class="ant-table-striped"
    :columns="groupListColumns"
    :data-source="groupDataList"
    :row-class-name="
      (record, index) => (index % 2 === 1 ? 'dark-row' : 'light-row')
    "
    :bordered="false"
    style="overflow-y: auto; overflow: auto"
    :row-key="(record, index) => index"
    :pagination="pagination"
  >
    <template #nameRender="{ record }">
      <router-link :to="{ path: '/groupPolicyDetail/' + record.name }">{{
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

    <!-- <template #tags="{ text: tags }">
      <span>
        <a-tag
          v-for="tag in tags"
          :key="tag"
          :color="
            tag === 'loser' ? 'volcano' : tag.length > 5 ? 'geekblue' : 'green'
          "
        >
          {{ tag.toUpperCase() }}
        </a-tag>
      </span>
    </template> -->
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
  props: {},
  setup() {
    return {
      //rowSelection,
      loading: ref(false),
      actionFrom: ref("GroupPolicyList"),
      pagination: {
        pageSize: 10,
        showSizeChanger: true, // display can change the number of pages per page
        pageSizeOptions: ["10", "20", "50", "100"], // number of pages per option
        showTotal: (total) => `Total ${total} items`, // show total
        showSizeChange: (current, pageSize) => (this.pageSize = pageSize), // update display when changing the number of pages per page
      },
    };
  },
  components: {
    Actions,
  },
  data() {
    return {
      selectedRowKeys: [],
      groupDataList: [],
      groupListColumns: [
        {
          title: this.$t("label.name"),
          dataIndex: "name",
          key: "name",
          width: "39%",
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
          slots: { customRender: "nameRender" },
        },
        {
          title: "",
          key: "action",
          dataIndex: "action",
          width: "1%",
          align: "right",
          slots: { customRender: "actionRender" },
        },
        {
          title: this.$t("label.description"),
          dataIndex: "description",
          key: "description",
          width: "60%",
          sorter: (a, b) =>
            a.description < b.description
              ? -1
              : a.description > b.description
              ? 1
              : 0,
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
        this.$emit("actionFromChange", "GroupPolicy");
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
        .get("/api/v1/group")
        .then((response) => {
          if (response.status == 200) {
            this.groupDataList = response.data.result;
          } else {
            message.error(this.$t("message.response.data.fail"));
            //console.log(response.message);
          }
        })
        .catch((error) => {
          message.error(this.$t("message.response.data.fail"));
          console.log(error);
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
