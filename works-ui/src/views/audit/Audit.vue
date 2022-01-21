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
                  { name: $t('label.audit'), component: 'Audit' },
                  { name: name, component: null },
                ]"
              />
            </a-col>

            <!-- 왼쪽 액션 -->
            <a-col id="content-action" :span="12">
              <actions :addButton="true" :add="$t('label.audit')" />
            </a-col>
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
          <AuditList
            :data="AuditData"
            :columns="AuditColumns"
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
import { AuditColumns, AuditData } from "../../data.js";
import AuditList from "./AuditList";

export default defineComponent({
  props: {
    name: String,
    info: Object,
  },
  components: { AuditList, Apath, Actions },
  setup() {
    return {
      AuditColumns,
      AuditData,
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
  padding: 20px;
  height: auto;
}

#content-path {
  text-align: left;
  align-items: center;
  display: flex;
  height: 32px;
}

#content-action {
  text-align: right;
}
</style>
