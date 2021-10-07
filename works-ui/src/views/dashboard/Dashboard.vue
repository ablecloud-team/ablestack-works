<template>
  <div style="width:100%; padding: 10px;">
    <a-row>
      <a-col flex="100%">
        <a-row :gutter="12" type="flex">
          <a-col flex="50%" class="dashboard-a-col">
            <a-card :title="$t('label.workspace.count')" class="dashboard-a-card-cl" hoverable>
              <span style="font-size: 80px;">33</span>
            </a-card>
          </a-col>
          <a-col flex="50%" class="dashboard-a-col">
            <a-card :title="$t('label.desktop.count')" class="dashboard-a-card-cl" hoverable>
              <span style="font-size: 80px;">33</span>
            </a-card>
          </a-col>
          <!-- <a-col flex="25%" class="dashboard-a-col">
            <a-card :title="$t('label.allocated.cpu.count')" class="dashboard-a-card-cl" hoverable>
              <a-progress type="dashboard" :percent="33"/>
            </a-card>
          </a-col>
          <a-col flex="25%" class="dashboard-a-col">
            <a-card :title="$t('label.allocated.memory.count')" class="dashboard-a-card-cl" hoverable>
              <a-progress type="dashboard" :percent="22"/>
            </a-card>
          </a-col>
          <a-col flex="25%" class="dashboard-a-col">
            <a-card :title="$t('label.allocated.disk.count')" class="dashboard-a-card-cl" hoverable>
              <a-progress type="dashboard" :percent="70" />
            </a-card>
          </a-col>
          <a-col flex="25%" class="dashboard-a-col">
            <a-card :title="$t('label.allocated.IP.count')" class="dashboard-a-card-cl" hoverable>
              <a-progress type="dashboard" :percent="70" />
            </a-card>
          </a-col> -->
        </a-row>
        <a-row :gutter="12" type="flex">
          <a-col flex="50%" class="dashboard-a-col">
            <a-card :title="$t('label.desktop.connected.count')" class="dashboard-a-card-cl" hoverable>
              <span style="font-size: 80px;">33</span>
            </a-card>
          </a-col>
          <a-col flex="50%" class="dashboard-a-col">
            <a-card :title="$t('label.app.connected.count')" class="dashboard-a-card-cl" hoverable>
              <span style="font-size: 80px;">33</span>
            </a-card>
          </a-col>
        </a-row>
      </a-col>
      <!-- <a-col flex="30%" class="dashboard-a-col">
        <a-card :bordered="true" class="dashboard-right-card">
          <a-timeline style="text-align: left">
            <a-timeline-item color="green"
              >Create a services site 2015-09-01</a-timeline-item
            >
            <a-timeline-item color="green"
              >Create a services site 2015-09-01</a-timeline-item
            >
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
        </a-card>
      </a-col> -->
    </a-row>
  </div>
</template>
  
<script>
export default {
  name: "Dashboard",
};
</script>

<style scoped>
.dashboard-a-card-cl {
  width: 100%;
  height: 100%;
  text-align: center;
}
.dashboard-a-col {
  padding-bottom: 10px;
}
</style>