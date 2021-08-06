<template>
  <div class="pager">
    <a-pagination
      size="small"
      :total="50"
      show-size-changer
      show-quick-jumper
    />
  </div>
</template>

<script>
export default {
  name: "Pager",
};
</script>
