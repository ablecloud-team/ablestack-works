<template>
  <a-breadcrumb>

    <a-breadcrumb-item href="">
      <HomeFilled/>
    </a-breadcrumb-item>


    <a-breadcrumb-item href="" v-for="path in paths">
      <span>{{ path }}</span>
    </a-breadcrumb-item>


  </a-breadcrumb>
</template>

<script>
export default {
  name: "Apath",
  props:{
    paths: Array,
  }
}
</script>

<style scoped>

</style>