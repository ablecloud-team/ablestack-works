<template>
  <a-breadcrumb>

    <a-breadcrumb-item href="">
      <HomeOutlined style="font-size: 16px"/>
    </a-breadcrumb-item>


    <a-breadcrumb-item href="" v-for="path in paths" v-bind:key="paths.indexOf(path)">
      <span v-if="path.component !== undefined"><router-link :to="{ name: path.component}">{{ path.name }}</router-link></span>
      <span v-else-if="path.name !== undefined">{{ path.name }}</span>
      <span v-else>{{ path }}</span>
    </a-breadcrumb-item>


  </a-breadcrumb>
</template>

<script>
export default {
  name: "Apath",
  props:{
    paths: Array,
  },
  setup(){
    return
  }
}
</script>

<style scoped>

</style>