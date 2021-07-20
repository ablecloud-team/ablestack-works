<template>
  <a-row
      :gutter="12"
  >
    <a-col
        :xl="18"
        class="login-arow-acol1"
    >
      <a-row :gutter="12" >
        <a-col
            :xs="12"
            :sm="8"
            :md="6"
          >
          <a-card
              class="capacity-dashboard-chart-card-inner"
              :loading="loading"
              :width="200"
          >
            <h3>가상머신 수</h3>
            <a-progress
                type="dashboard"
                :percent="70"
                :size="large" />
            <p>20 / 30</p>
          </a-card>
          <a-card
              class="capacity-dashboard-chart-card-inner"
              :loading="loading"
              :width="200"
          >
            <h3>가상머신 수</h3>
            <a-progress
                type="dashboard"
                :percent="70"
                :size="large" />
            <p>20 / 30</p>
          </a-card>
          <a-card
              class="capacity-dashboard-chart-card-inner"
              :loading="loading"
              :width="200"
          >
            <h3>가상머신 수</h3>
            <a-progress
                type="dashboard"
                :percent="70"
                :size="large" />
            <p>20 / 30</p>
          </a-card>
        </a-col>
      </a-row>
    </a-col>
    <a-col
        :xl="6"
        class="login-arow-acol2"
    >
      col-6
    </a-col>
  </a-row>
</template>

<script>
import { defineComponent } from 'vue';
export default defineComponent({})
</script>

<style scoped>
.login-arow-acol1{
  background: #1890ff;
  display: flex;
  justify-content: center;
}

.login-arow-acol2{
  background: darkorchid;
}

.capacity-dashboard-chart-card-inner{
  text-align: center;
  padding: 20px;
  margin: 5px;
  width: 300px;
}
</style>
