<template>
  <div id="content-layout">
    <a-layout>
      <a-layout-header id="content-header">
        <div id="content-header-body">
          <a-row id="content-header-row">
            <!-- 좌측 경로 -->
            <a-col id="content-path" :span="12">
              <Apath v-bind:paths="[$t('label.vm')]"/>
              <a-button shape="round" style="margin-left: 20px; height:30px;" @click="reflesh()">
                <template #icon>
                  <ReloadOutlined /> {{$t('label.reflesh')}}
                </template>
              </a-button>
            </a-col>

            <!-- 우측 액션 -->
            <a-col id="content-action" :span="12" >
              <actions :actionFrom="actionFrom" v-if="actionFrom === 'VirtualMachineList'"/>
            </a-col>
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
          <VirtualMachineList
            ref="listRefleshCall"
            @actionFromChange="actionFromChange"
          />
        </div>
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script>
import Actions from "@/components/Actions";
import Apath from "@/components/Apath";
import { defineComponent, ref } from "vue";
import VirtualMachineList from "./VirtualMachineList.vue";

export default defineComponent ({
  name: "VirtualMachine",
  props: {},
  components: {
    VirtualMachineList,
    Apath,
    Actions,
  },
  data() {
    return {
      actionFrom: ref("VirtualMachine"),
    };
  },
  methods: {
    reflesh(){
      this.$refs.listRefleshCall.fetchData();
    },
    actionFromChange(val) {
      //console.log(val);
      this.actionFrom = ref(val);
    },
  }
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
