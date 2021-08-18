<template>
  <a-table
    size="middle"
    class="ant-table-striped"
    :columns="columns"
    :data-source="data"
    :rowClassName="
      (record, index) => (index % 2 === 1 ? 'dark-row' : 'light-row')
    "
    :bordered="bordered ? bordered : false"
    style="overflow-y: auto; overflow: auto"
    :row-key="(record, index) => index"
    :row-selection="rowSelection"
  >
    <template #nameRender="{ record }">
      <router-link
        :to="{
          name: 'VirtualMachineDetail',
          params: {
            name: record.Name,
            info: [record.IPAddress, record.Account, record.Zone],
          },
        }"
        >{{ record.Name }}</router-link
      >
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
import { defineComponent } from "vue";
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
  setup() {
    return {
      rowSelection,
    };
  },
  components: {
    Actions,
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
