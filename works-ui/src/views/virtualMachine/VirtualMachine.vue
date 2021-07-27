<template>
  <div id="content-layout">
    <a-layout>
      <a-layout-header id="content-header">
        <div id="content-header-body">
          <a-row id="content-header-row">
            <!-- 오른쪽 경로 -->
            <a-col id='content-path' :span="12">
              <Apath v-bind:paths="[$t('label.vm')]"
              />

            </a-col>

            <!-- 왼쪽 액션 -->
            <a-col id="content-action" :span="12">
              <actions :poweroff=true :destroy="true" :reset="true"/>
            </a-col>

          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
            <VirtualMachineList :data="VMListData" :columns="VMListColumns" :bordered="false"/>
        </div>
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script>
import Actions from "@/components/Actions";
import Apath from "@/components/Apath";
import {VMListColumns, VMListData} from "@/data"
import VirtualMachineList from "@/views/virtualMachine/VirtualMachineList";
import {defineComponent} from "vue";

export default defineComponent({
  props: {
    msg: String
  },
  components: {VirtualMachineList,  Apath, Actions},

  setup() {
    return {
      VMListColumns, VMListData
    };
  }
})
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
