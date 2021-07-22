<template>

  <div class="resource-details">
    <div class="resource-details__name">
      <a-avatar id='oslogo' :size="100" style="background-color: rgba(0,0,0,0); font-size: 20px" >
        <font-awesome-icon :icon="['fab', 'centos']" :size="'4x'" style="color: #666;" />
      </a-avatar>
      <h4>
        {{ name }}
      </h4>
    </div>
    <div id="div_tag" v-if="tags !== undefined">
      <a-tag v-for="tag in tags" v-bind:key="tags.indexOf(tag)">
        {{ tag }}
      </a-tag>
    </div>
    <!--      <a-tooltip placement="right">-->
    <!--        <template #title>connect portal</template>-->
    <!--        <a-button shape="circle" type="dashed">-->
    <!--          <PaperClipOutlined/>-->
    <!--        </a-button>-->
    <!--      </a-tooltip>-->
  </div>
  <ADivider />

  <div id="Status" class="CardItem">
    <div class="ItemName">Status</div>
    <div class="Item">
      <a-badge class='Status' :color="'green'" :text="'Implemented'" />
    </div>
  </div>

  <div id="ID" class="CardItem">
    <div class="ItemName">ID</div>
    <div class="Item">
      <a-button shape="circle" type="dashed" >
        <BarcodeOutlined />
      </a-button>{{infoID}}
    </div>
  </div>


  <div id="Type" class="CardItem">
    <div class="ItemName">Type</div>
    <div class="Item">
      Desktop
    </div>
  </div>


  <div id="NoD" class="CardItem">
    <div class="ItemName">Number Of Desktop</div>
    <div class="Item">
      83
    </div>
  </div>

  <div id="NoC" class="CardItem">
    <div class="ItemName">Number Of Connection</div>
    <div class="Item">
      42
    </div>
  </div>

  <div id="Network" class="CardItem">
    <div class="ItemName">Network</div>
    <div class="Item">
      DaaS-ISO-NW
    </div>
  </div>
<!--  <div v-if="info !== undefined">-->
<!--    <a-tag v-for="tag in info" v-bind:key="info.indexOf(tag)">-->
<!--      {{tag}}-->
<!--    </a-tag>-->
<!--  </div>-->
</template>

<script>
function guid() {
  function _s4() {
    return ((1 + Math.random()) * 0x10000 | 0).toString(16).substring(1);
  }
  return _s4() + _s4() + '-' + _s4() + '-' + _s4() + '-' + _s4() + '-' + _s4() + _s4() + _s4();
}
export default {
  props:{
    name: String,
    tags: {
      type:Array,
      required: false,
    },
    info: {
      type:Object,
      required: false,
    }
  },
  setup(){
    let infoID = guid()
    return {infoID}
  }
}
</script>

<style  lang="scss" scoped>
.CardItem{
  margin-bottom: 20px;
}
::v-deep(.ant-badge-status-dot) {
  width: 12px ;
  height: 12px ;
}
.ItemName{
  font-size: 14px;
  font-weight: bold;
}
.resource-details {
  text-align: center;
  margin-bottom: 20px;

  &__name {
    display: flex;
    align-items: center;
    align-content: center;
    text-align: center;

    .ant-avata {
      margin-right: 20px;
      overflow: hidden;
      min-width: 50px;
      cursor: pointer;
      img {
        height: 100%;
        width: 100%;
      }
    }
    h4{
      font-size: 18px;
    }
  }
}
</style>
