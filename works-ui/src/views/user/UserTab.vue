<template>
  <div id="ContentTab">
    <a-tabs v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane key="1" :tab="$t('label.detail')">
        <DetailContent
          ref="listRefleshCall1"
          :action-from="'UserDetail'"
        />
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script>
import { defineComponent, ref } from "vue";
import DetailContent from "@/components/DetailContent";

export default defineComponent({
  components: {
    DetailContent,
  },
  props: {},
  setup() {
    const tabPosition = ref("top");
    const activeKey = ref("1");
    return {
      tabPosition,
      activeKey,
    };
  },
  data() {
    return {
    };
  },
  methods: {
    reflesh() {
      this.$refs.listRefleshCall1.reflesh();
    },
  },
});
</script>

<style>
#ContentTab {
  text-align: left;
}
</style>
