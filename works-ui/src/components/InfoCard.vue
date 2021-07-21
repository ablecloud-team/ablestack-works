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
      <a-tag v-for="tag in tags">
        {{ tag }}
      </a-tag>
      <a-tooltip placement="right">
        <template #title>connect portal</template>
        <a-button shape="circle" type="dashed">
          <PaperClipOutlined/>
        </a-button>
      </a-tooltip>
    </div>

    <ADivider />
    <AEmpty />
</template>

<script>
export default {
  name: "InfoCard",
  props:{
    name: String,
    tags: Array,
  }
}
</script>

<style  lang="scss" scoped>

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
