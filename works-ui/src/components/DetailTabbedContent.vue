<template>
  <div id="content-layout">
    <a-layout>
      <a-layout-header id="content-header">
        <div>
          <a-row>
            <!-- 오른쪽 경로 -->
            <a-col id='content-path' :span="12">
              <Apath v-bind:paths="[{name:$t('label.Workspace'), component: 'Workspaces'}, {name:$t('label.WorkspaceDetail'), component: null}]"
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
          <Bodycontent :name="name" :info="info"/>
        </div>
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script>
import Bodycontent from "@/components/BodyContent";
import Actions from "@/components/actions";
import Apath from "@/components/Apath";
export default {
  name: 'DetailTabbedContent',
  props: {
    name: String,
    info: Object
  },
  components: {Apath, Actions, Bodycontent},
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#content-layout, #content-body{
  width:calc(100%);
}
#content-layout .ant-layout-header {
  text-align: left;
  background: #e8eaed;
  color: #fff;
  font-size: 14px;
  line-height: 1.5;
  padding: 24px;
}

#content-layout .ant-layout-header #content-path {
  text-align: left;
  /*background: #ea7da3;*/
}


#content-layout .ant-layout-header #content-action {
  text-align: right;
  /*background: #b0ea7d;*/
  margin-bottom: -6px;
}
#content-layout .ant-layout-content {
  /*background: rgb(233, 211, 16);*/
  color: #fff;
  min-height: 120px;
  line-height: 120px;
}

#content-layout > .ant-layout {
  margin-bottom: 48px;
}

#content-layout > .ant-layout:last-child {
  margin: 0;
}

</style>
