<template>
  <a-table
    size="middle"
    class="ant-table-striped"
    :columns="listColumns"
    :data-source="data"
    :row-class-name="
      (record, index) => (index % 2 === 1 ? 'dark-row' : 'light-row')
    "
    :bordered="bordered ? bordered : false"
    style="overflow-y: auto; overflow: auto"
    :row-key="(record, index) => index"
    :row-selection="rowSelection"
    :pagination="pagination"
  >
    <template #nameRender="{ record }">
      <router-link :to="{ path: '/workspaceDetail/'+record.Uuid}">{{ record.Name }}</router-link>
    </template>

    <template #actionRender>
      <a-Popover placement="topLeft">
        <template #content>
          <actions :action-from="actionFrom" />
        </template>
        <MoreOutlined />
      </a-Popover>
    </template>

    <template #tags="{ text: tags }">
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
    </template>
  </a-table>
</template>

<script>
import { defineComponent, ref } from "vue";
import Actions from "@/components/Actions";

const rowSelection = {
  onChange: (selectedRowKeys, selectedRows) => {
    console.log(
      `selectedRowKeys: ${selectedRowKeys}`,
      "selectedRows: ",
      selectedRows
    );
  },
  onSelect: (record, selected, selectedRows) => {
    console.log(record, selected, selectedRows);
  },
  onSelectAll: (selected, selectedRows, changeRows) => {
    console.log(selected, selectedRows, changeRows);
  },
};
export default defineComponent({
  components: {
    Actions,
  },
  props: {
    data: {
      type: Object,
      requires: true,
      default: null,
    },
    bordered: {
      type: Boolean,
      requires: true,
      default: null,
    },
  },
  setup() {
    return {
      rowSelection,
      actionFrom: ref("WorkspaceList"),
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
      listColumns: [
        {
          dataIndex: "Name",
          key: "Name",
          slots: { customRender: "nameRender" },
          title: this.$t('label.name'),
          sorter: (a, b) => (a.Name < b.Name ? -1 : a.Name > b.Name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: "",
          key: "action",
          dataIndex: "action",
          slots: { customRender: "actionRender" },
        },
        {
          title: this.$t('label.state'),
          dataIndex: "State",
          key: "State",
          sorter: (a, b) => (a.State < b.State ? -1 : a.State > b.State ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t('label.type'),
          dataIndex: "Type",
          key: "Type",
          sorter: (a, b) => (a.Type < b.Type ? -1 : a.Type > b.Type ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t('label.desktop.quantity'),
          dataIndex: "Quantity",
          key: "Quantity",
          sorter: (a, b) => (a.NoD < b.NoD ? -1 : a.NoD > b.NoD ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t('label.desktop.connection.quantity'),
          dataIndex: "NoC",
          key: "NoC",
          sorter: (a, b) => (a.NoC < b.NoC ? -1 : a.NoC > b.NoC ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t('label.network.type'),
          dataIndex: "NetType",
          key: "NetType",
          sorter: (a, b) =>
            a.NetType < b.NetType ? -1 : a.NetType > b.NetType ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t('label.restrict'),
          dataIndex: "Restrict",
          key: "Restrict",
          sorter: (a, b) =>
            a.Restrict < b.Restrict ? -1 : a.Restrict > b.Restrict ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
        // ,
        // {
        //     title: 'Tags',
        //     key: 'tags',
        //     dataIndex: 'tag',
        //     slots: {customRender: 'tags'},
        // }
      ],
    }
  }
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
