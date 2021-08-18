<template>
  <div id="content-layout">
    <a-layout>
      <a-layout-header id="content-header">
        <div>
          <a-row>
            <!-- 오른쪽 경로 -->
            <a-col id="content-path" :span="12">
              <Apath
                v-bind:paths="[
                  { name: $t('label.group.policy'), component: 'GroupPolicy' },
                  { name: name, component: null },
                ]"
              />
            </a-col>

            <!-- 왼쪽 액션 -->
            <a-col id="content-action" :span="12">
              <actions :addButton="true" :add="$t('label.group.policy')" />
            </a-col>
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
          <GroupPolicyList
            :data="GroupPolicyData"
            :columns="GroupPolicyColumns"
            :bordered="false"
          />
        </div>
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import Actions from "../../components/Actions";
import Apath from "../../components/Apath";
import { GroupPolicyColumns, GroupPolicyData } from "../../data";
import GroupPolicyList from "./GroupPolicyList";
export default defineComponent({
  props: {
    name: String,
    info: Object,
  },
  components: { GroupPolicyList, Apath, Actions },
  setup() {
    return {
      GroupPolicyColumns,
      GroupPolicyData,
    };
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#content-layout .ant-layout-content {
  background: #f0f2f5;
  padding: 10px;
}
#content-layout .ant-layout-header {
  text-align: left;
  background: white;
  border: 1px solid #e8e8e8;
  /*color: #fff;*/
  font-size: 14px;
  line-height: 1.5;
  padding: 24px;
  height: auto;
}

#content-path {
  text-align: left;
  align-items: center;
  display: flex;
}

#content-action {
  text-align: right;
}
</style>
