<template>
  <a-list
      size="small"
      :dataSource="data"
      :pagination="false"
      style="overflow-y: auto"
  >
    <template #renderItem="{ item }">
      <a-list-item>
        <template #actions>
          <actions :power="item.State!=='Running'" :edit="item.State!=='Running'" :poweroff="item.State === 'Running'" :destroy="true" :reset="item.State === 'Running'"/>
        </template>
        <a-list-item-meta :description="item.State"  >

        <template #title>
          {{ item.Name }}
        </template>
        </a-list-item-meta>
      </a-list-item>
    </template>

    <template #header>
      <div>Header</div>
    </template>
    <template #footer>
      <div>Footer</div>
    </template>
  </a-list>
</template>

<script>

import {columns, data, listdata} from "../data";
import Actions from "@/components/actions";
export default {
  name: 'ListView',
  components: {
    Actions
  },
  props: {
  },
  setup(){
    return {
      data,
      columns,
      listdata,
    };
  }
}
</script>

<style scoped>
:deep(.ant-table-thead) {
  background-color: #f9f9f9;
}

:deep(.ant-table-small) > .ant-table-content > .ant-table-body {
  margin: 0;
}

:deep(.light-row) {
  background-color: #fff;
}

:deep(.dark-row) {
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
