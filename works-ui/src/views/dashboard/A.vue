<template style="text-align: center;">
  <Actions
      :actionFrom="status.actionFrom"
      :add="$t('label.add.workspace')"
      :showModal="status.showModal"
      :addButton="status.addButton"
      :start="status.power"
      :stop="status.poweroff"
      :reset="status.reset"
      :reinstall="status.reinstall"
      :snapshot="status.snapshot"
      :volsnapshot="status.volsnapshot"
      :iso="status.iso"
      :destroy="status.destroy"
      :edit="status.edit"
  />
</template>
<script>
import Actions from "../../components/Actions";
import {reactive, ref} from "vue";

export default {
  name: "A",
  components:{
    Actions,
  },
  setup(){
    const status = reactive({
      actionFrom: ref('Workspaces'),
      showModal: ref(true),
      addButton: ref(true),
      power: ref(true),
      poweroff: ref(true),
      reset: ref(true),
      reinstall: ref(true),
      snapshot: ref(true),
      volsnapshot: ref(true),
      iso: ref(true),
      destroy: ref(true),
      edit: ref(true),
    })
    return{
      status
    }
  }
}
</script>

<style scoped>
.dashboard-main-grid{
  display: grid;
  grid-template-columns: 3fr 1fr;
  grid-template-rows: 1fr;
  height: 100%;
  text-align: center;
  padding: 5px;
}
.dashboard-left-grid{
  display: grid;
  grid-template-columns: 1fr 1fr 1fr 1fr;
  grid-template-rows: 291px 291px;
  grid-gap: 14px;
  padding: 14px;
  text-align: center;
}

.dashboard-right-grid{
  display: grid;
  grid-template-columns: 1fr;
  grid-template-rows: 1fr 1fr 1fr;
  grid-gap: 14px;
  padding: 14px;
}

.dashboard-card{
  max-height: 291px;
}
.dashboard-right-card{
  width: 100%;
  text-align: left;
  margin-bottom: 10px;
  box-shadow: 5px 5px;
}
.dashboard-right-card-p{
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  line-height: 16px;
  max-height: 50px;
  -webkit-line-clamp: 3; /* 표시하고자 하는 라인 수 */
  -webkit-box-orient: vertical;
}
</style>
