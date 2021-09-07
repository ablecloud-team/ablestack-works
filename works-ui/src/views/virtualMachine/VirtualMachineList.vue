<template>
  <a-table
    size="middle"
    class="ant-table-striped"
    :columns="VMListColumns"
    :data-source="VMListData"
    :rowClassName="
      (record, index) => (index % 2 === 1 ? 'dark-row' : 'light-row')
    "
    :bordered="bordered ? bordered : false"
    style="overflow-y: auto; overflow: auto"
    :row-key="(record, index) => index"
    :row-selection="rowSelection"
    :pagination="pagination"
  >
    <template #nameRender="{ record }">
      <!-- <router-link
        :to="{
          name: 'VirtualMachineDetail',
          params: {
            name: record.Name,
            info: [record.IPAddress, record.Account, record.Zone],
          },
        }"
        >{{ record.Name }}</router-link
      > -->
      <router-link :to="{ path: '/virtualMachineDetail/'+record.Uuid}">{{ record.Name }}</router-link>
    </template>

    <template #actionRender="{ record }">
      <a-Popover placement="topLeft">
        <template #content>
          <ASpace direction="horizontal">
            <Actions
              :power="record.State === 'Running'"
              :destroy="true"
              :reset="true"
              :iso="true"
            />
          </ASpace>
        </template>
        :
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
  props: {
    data: Object,
    columns: Object,
    bordered: Boolean,
  },
  components: {
    Actions,
  },
  setup() {
    return {
      rowSelection,
      actionFrom: ref("VirtualMachineList"),
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
      VMListColumns : [
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
          title: this.$t('label.workspace'),
          dataIndex: "Workspace",
          key: "Workspace",
          sorter: (a, b) =>
            a.Workspace < b.Workspace ? -1 : a.Workspace > b.Workspace ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t('label.users'),
          dataIndex: "User",
          key: "User",
          sorter: (a, b) => (a.Type < b.Type ? -1 : a.Type > b.Type ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t('label.desktop.connect.boolean'),
          dataIndex: "Conn",
          key: "Conn",
          sorter: (a, b) => (a.NoD < b.NoD ? -1 : a.NoD > b.NoD ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
      ],
      VMListData : [
        {"Uuid":"sdfasdfasdfasdf", "Name":"VM1","State":"Running","User":"user01","Conn":"TRUE","Workspace":"test1"},
        {"Uuid":"sdfasdfasdfasdf", "Name":"VM2","State":"Stopped","User":"user03","Conn":"FALSE","Workspace":"test1"},
        {"Uuid":"sdfasdfasdfasdf", "Name":"VM3","State":"Running","User":"user02","Conn":"TRUE","Workspace":"test1"}
      ],
    }
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
