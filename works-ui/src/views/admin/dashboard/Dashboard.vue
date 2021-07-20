<template>
  <a-layout>
    <a-layout-content class="dashboard-left-content">
      <div class="dashboard-div-grid">
        <a-row :gutter="[4,4]" style="padding-top: 6px; padding-bottom: 3px">
          <a-col :span="8">
            <a-card
                :loading="loading"
            >
              <h3>가상머신 수</h3>
              <a-progress
                  type="dashboard"
                  :percent="70"
                  :size="large" />
              <p>20 / 30</p>
            </a-card>
          </a-col>
          <a-col :span="8">
            <a-card
                :loading="loading"
            >
              <h3>가상머신 수</h3>
              <a-progress
                  type="dashboard"
                  :percent="70"
                  :size="large" />
              <p>20 / 30</p>
            </a-card>
          </a-col>
          <a-col :span="8">
            <a-card
                :loading="loading"
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
        <a-row :gutter="[4,4]" style="padding-top: 3px; padding-bottom: 6px">
          <a-col :span="8">
            <a-card
                :loading="loading"
            >
              <h3>가상머신 수</h3>
              <a-progress
                  type="dashboard"
                  :percent="70"
                  :size="large" />
              <p>20 / 30</p>
            </a-card>
          </a-col>
          <a-col :span="8">
            <a-card
                :loading="loading"
            >
              <h3>가상머신 수</h3>
              <a-progress
                  type="dashboard"
                  :percent="70"
                  :size="large" />
              <p>20 / 30</p>
            </a-card>
          </a-col>
          <a-col :span="8">
            <a-card
                :loading="loading"
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
      </div>
    </a-layout-content>
    <a-layout-sider class="dashboard-right-content">
      <a-timeline>
        <a-timeline-item color="green">Create a services site 2015-09-01</a-timeline-item>
        <a-timeline-item color="green">Create a services site 2015-09-01</a-timeline-item>
        <a-timeline-item color="red">
          <p>Solve initial network problems 1</p>
          <p>Solve initial network problems 2</p>
          <p>Solve initial network problems 3 2015-09-01</p>
        </a-timeline-item>
        <a-timeline-item>
          <p>Technical testing 1</p>
          <p>Technical testing 2</p>
          <p>Technical testing 3 2015-09-01</p>
        </a-timeline-item>
        <a-timeline-item color="gray">
          <p>Technical testing 1</p>
          <p>Technical testing 2</p>
          <p>Technical testing 3 2015-09-01</p>
        </a-timeline-item>
        <a-timeline-item color="gray">
          <p>Technical testing 1</p>
          <p>Technical testing 2</p>
          <p>Technical testing 3 2015-09-01</p>
        </a-timeline-item>
      </a-timeline>
    </a-layout-sider>
  </a-layout>
</template>

<script>
import { defineComponent } from 'vue';
export default defineComponent({})
</script>

<style scoped>
.dashboard-right-content{
  background: white;
  flex: 0 0 30% !important;
  max-width: 30% !important;
  min-width: 30% !important;
  width: 30% !important;
  margin-top: 5px;
  padding: 3px 3px 3px 3px;
}

.dashboard-div-grid{
  text-align: center;
  max-width: 90rem;
}
.dashboard-left-content{
  width: 70%;
  padding-left: 6px;
  padding-right: 6px;
}
</style>
