<template>
  <ASpace direction="horizontal" id="content-space">
    <ARow style="min-height: 400px" id="content-row">
      <ACol :span="8" style="background: #ececec; padding: 8px">
        <!-- 왼쪽 detail 창 -->
        <ACard bordered style=" min-height: 300px">
          <InfoCard :name="name" :tags="['test', 'tag', 'list']" :info="info" />
        </ACard>
      </ACol>

      <ACol :span="16" style="background: #ececec; padding: 8px">
        <!-- 오른쪽 tab 창 -->
        <ACard bordered>
          <TabbedContent />
        </ACard>
      </ACol>
    </ARow>
  </ASpace>
</template>

<script>
import TabbedContent from "./TabbedContent";
import InfoCard from "@/components/InfoCard";

export default {
  name: 'Bodycontent',
  components: {InfoCard, TabbedContent,},
  props: {
    name: String,
    info: Object
  }

};
</script>
<style>
#content-space, #content-row{
  width: 100%;
}
#content-space > .ant-space-item{
  width: 100%;
}

</style>
