<template>
  <a-card
      title="이름"
      :bordered="false"
      :headStyle="{minHeight: '0px',padding: '7px 24px 7px 24px', fontWeight:'700'}"
      :bodyStyle="{padding: '7px 24px 7px 24px'}"
      class="detail-content-card"
  >
    VM1
  </a-card>
  <a-card
      title="상태"
      :bordered="false"
      :headStyle="{minHeight: '0px',padding: '7px 24px 7px 24px'}"
      :bodyStyle="{padding: '7px 24px 7px 24px'}"
      class="detail-content-card"
  >
    Running
  </a-card>
  <a-card
      title="IP 주소"
      :bordered="false"
      :headStyle="{minHeight: '0px',padding: '7px 24px 7px 24px'}"
      :bodyStyle="{padding: '7px 24px 7px 24px'}"
      class="detail-content-card"
  >
    10.1.1.16
  </a-card>
  <a-card
      title="템플릿"
      :bordered="false"
      :headStyle="{minHeight: '0px',padding: '7px 24px 7px 24px'}"
      :bodyStyle="{padding: '7px 24px 7px 24px'}"
      class="detail-content-card"
  >
    Windows10-v1.0
  </a-card>
  <a-card
      title="컴퓨트 오퍼링"
      :bordered="false"
      :headStyle="{minHeight: '0px',padding: '7px 24px 7px 24px'}"
      :bodyStyle="{padding: '7px 24px 7px 24px'}"
      class="detail-content-card"
  >
    2 vCPU 4GB Mem, 1GbE NIC, 50GB Root (GlueFS)
  </a-card>
</template>

<script>
export default {
  name: "DetailContent"
}
</script>

<style scoped>
.detail-content-card{
  margin-bottom: 10px;
}::v-deep(.ant-card-head-title) {
   padding: 0px;
 }
</style>
