<template>
  <a-button
    type="dashed"
    block
    style="margin-bottom: 14px"
    @click="changeModal(true)"
    v-if="state.addButtonBoolean"
  >
    <PlusOutlined />
    {{ state.addButtontext }}
  </a-button>
  <a-table
    size="middle"
    class="ant-table-striped"
    :columns="columns"
    :data-source="data"
    :rowClassName="
      (record, index) => (index % 2 === 1 ? 'dark-row' : 'light-row')
    "
    :bordered="bordered ? bordered : false"
    style="overflow: auto"
    :row-selection="rowSelection"
    :pagination="pagination"
  >
    <template #nameRender="{ record }">
      <span v-if="comp !== undefined">
        <router-link
          :to="{
            name: comp,
            params: {
              name: record.name,
              info: [record.IPAddress, record.Account, record.Zone],
            },
          }"
        >
          {{ record.name }}
        </router-link>
      </span>
      <span v-else>
        {{ record.name }}
      </span>
    </template>

    <template #actionRender>
      <a-Popover placement="top">
        <template #content>
          <ASpace direction="horizontal">
            <actions :actionFrom="actionFrom" />
          </ASpace>
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

  <a-modal
    v-model:title="state.addButtontext"
    v-model:visible="state.modalBoolean"
    @ok="changeModal(false)"
    width="400px"
  >
    <a-input-number
      id="inputNumber"
      :title="'Desktop Quantity'"
      v-model:value="addDesktopQuantity"
      :min="1"
      :max="10"
      style="width: 100%; margin-top: 7px"
      v-if="state.callTapName === 'desktop'"
    />
    <a-select v-if="state.callTapName === 'user'" style="width: 100%">
    </a-select>
  </a-modal>
</template>

<script>
// import {SmileOutlined} from '@ant-design/icons-vue';
import { defineComponent, reactive, ref } from "vue";
import Actions from "../components/Actions";

const rowSelection = {
  onChange: (selectedRowKeys, selectedRows) => {
    // console.log(
    //   `selectedRowKeys: ${selectedRowKeys}`,
    //   "selectedRows: ",
    //   selectedRows
    // );
  },
  onSelect: (record, selected, selectedRows) => {
    //console.log(record, selected, selectedRows);
  },
  onSelectAll: (selected, selectedRows, changeRows) => {
    //console.log(selected, selectedRows, changeRows);
  },
};
export default defineComponent({
  props: {
    tapName: String,
    data: Object,
    columns: Object,
    bordered: Boolean,
    comp: String,
    actionFrom: String,
  },
  setup(props) {
    //console.log("TableContent.vue props.tapName");
    //console.log(props.tapName);
    const state = reactive({
      modalBoolean: ref(false),
      callTapName: ref(props.tapName),
      addButtonBoolean: ref(false),
      addButtontext: ref(""),
      callModal: ref("desktop"),
    });  
    const addDesktopQuantity = ref(1);
    return {
      rowSelection,
      state,
      addDesktopQuantity,
      pagination: {
        pageSize: 20,
        showSizeChanger: true, // display can change the number of pages per page
        pageSizeOptions: ["10", "20", "30", "40"], // number of pages per option
        showTotal: (total) => `Total ${total} items`, // show total
        showSizeChange: (current, pageSize) => (this.pageSize = pageSize), // update display when changing the number of pages per page
      },
      actionFrom: ref(props.actionFrom),
    };
  },
  created() {
    this.fetchData();
  },
  data() {
    return {
    };
  },
  components: {
    Actions,
  },
  methods: {
    changeModal(value) {
      this.state.modalBoolean = ref(value);
    },
    fetchData() {
      if (this.state.callTapName === "desktop") {
        this.state.addButtonBoolean = ref(true);
        this.state.addButtontext = this.$t("label.desktop.vm.add");
      } else if (this.state.callTapName === "user") {
        this.state.addButtonBoolean = ref(true);
        this.state.addButtontext = this.$t("label.desktop.user.add");
      } else if (this.state.callTapName === "datadisk") {
        this.state.addButtonBoolean = ref(false);
      } else if (this.state.callTapName === "network") {
        this.state.addButtonBoolean = ref(false);
      }
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
