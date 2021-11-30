<template>
  <div id="content-layout">
    <a-layout>
      <a-layout-header id="content-header">
        <div id="content-header-body">
          <a-row id="content-header-row">
            <!-- 좌측 경로 -->
            <a-col id="content-path" :span="12">
              <Apath :paths="[$t('label.vm')]" />
              <a-button
                shape="round"
                style="margin-left: 20px"
                size="small"
                @click="refresh()"
              >
                <template #icon>
                  <ReloadOutlined /> {{ $t("label.refresh") }}
                </template>
              </a-button>
            </a-col>

            <!-- 우측 액션 -->
            <a-col id="content-action" :span="12">
              <Actions
                v-if="actionFrom === 'VirtualMachineList'"
                :action-from="actionFrom"
                :multi-select-list="multiSelectList"
                @fetchData="refresh"
              />
            </a-col>
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
          <VirtualMachineList
            ref="listRefreshCall"
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
import VirtualMachineList from "./VirtualMachineList";

export default defineComponent({
  name: "VirtualMachine",
  components: {
    VirtualMachineList,
    Apath,
    Actions,
  },
  props: {},
  data() {
    return {
      actionFrom: ref(""),
      multiSelectList: ref([]),
    };
  },
  methods: {
    refresh() {
      this.$refs.listRefreshCall.fetchRefresh();
    },
    actionFromChange(val, obj) {
      this.actionFrom = "";
      setTimeout(() => {
        this.actionFrom = val;
        this.multiSelectList = obj;
      }, 100);
    },
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
