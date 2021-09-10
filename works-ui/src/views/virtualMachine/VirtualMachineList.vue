<template>
  <a-table
    size="middle"
    :loading="loading"
    class="ant-table-striped"
    :columns="VMListColumns"
    :data-source="vmDataList"
    :rowClassName="(record, index) => (index % 2 === 1 ? 'dark-row' : 'light-row')"
    :bordered="bordered ? bordered : false"
    style="overflow-y: auto; overflow: auto"
    :row-key="(record, index) => index"
    :row-selection="{selectedRowKeys: selectedRowKeys, onChange: onSelectChange}"
    :pagination="pagination"
  >

    <template #nameRender="{ record }">
      <router-link :to="{ path: '/virtualMachineDetail/'+record.uuid}">{{ record.name }}</router-link>
    </template>

    <template #actionRender>
      <a-Popover placement="topLeft">
        <template #content>
          <actions :action-from="actionFrom" />
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


export default defineComponent({
  props: {
  },
  components: {
    Actions,
  },
  setup() {
    // const rowSelection = {
    //   onChange: (selectedRowKeys, selectedRows) => {
    //     console.log(
    //       `selectedRowKeys: ${selectedRowKeys}`,
    //       "selectedRows: ",
    //       selectedRows
    //     );
    //   },
    //   onSelect: (record, selected, selectedRows) => {
    //     $emit('actionFromChange', 'VirtualMachineDetail');
    //     console.log(record, selected, selectedRows);
    //   },
    //   onSelectAll: (selected, selectedRows, changeRows) => {
    //     console.log(selected, selectedRows, changeRows);
    //   },
    // };
    return {
      // rowSelection,
      loading: ref(false),
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
      selectedRowKeys: [],
      VMListColumns : [
        {
          title: this.$t('label.name'),
          dataIndex: "name",
          key: "name",
          slots: { customRender: "nameRender" },
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: "",
          key: "action",
          dataIndex: "action",
          align: 'right',
          slots: { customRender: "actionRender" },
        },
        {
          title: this.$t('label.state'),
          dataIndex: "state",
          key: "state",
          sorter: (a, b) => (a.state < b.state ? -1 : a.state > b.state ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t('label.workspace'),
          dataIndex: "workspace",
          key: "workspace",
          sorter: (a, b) =>
            a.workspace < b.workspace ? -1 : a.workspace > b.workspace ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t('label.users'),
          dataIndex: "user",
          key: "user",
          sorter: (a, b) => (a.user < b.user ? -1 : a.user > b.user ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t('label.desktop.connect.boolean'),
          dataIndex: "conn",
          key: "conn",
          sorter: (a, b) => (a.conn < b.connoD ? -1 : a.conn > b.conn ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
      ],
      vmDataList : [
        {"uuid":"0101010101010101010101010101001", "name":"VM1","state":"Running","user":"user01","conn":"TRUE","workspace":"test1"},
        {"uuid":"20202020202020202020202020202020", "name":"VM2","state":"Stopped","user":"user03","conn":"FALSE","workspace":"test1"},
        {"uuid":"3030303030303030303030303030303030", "name":"VM3","state":"Running","user":"user02","conn":"TRUE","workspace":"test1"}
      ],
    }
  },
    created() {
    this.fetchData();
  },
  methods: {
    setSelection (selection) {
      this.selectedRowKeys = selection;
      if(this.selectedRowKeys.length == 0){
        this.$emit('actionFromChange', "VirtualMachine");
      }else{
        this.$emit('actionFromChange', this.actionFrom);
      }
    },
    resetSelection () {
      this.setSelection([])
    },
    onSelectChange (selectedRowKeys, selectedRows) {
      this.setSelection(selectedRowKeys)
    },
    fetchData() {
      this.loading = true;
      // worksApi
      //   .get("/api/v1/virtualMachines", { withCredentials: true })
      //   .then((response) => {
      //     if (response.data.result.status == 200) {
      //       this.vmDataList = response.data.result.list;
      //     } else {
      //       message.error(this.t('message.response.data.fail'));
      //       //console.log(response.message);
      //     }
      //   })
      //   .catch(function (error) {
      //     message.error(error.message);
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
</style>
