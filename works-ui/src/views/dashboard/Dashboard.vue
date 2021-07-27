<template>
  <div class="dashboard-main-grid">
    <div class="dashboard-left-grid">
      <a-card class="dashboard-card" hoverable>
        <span style="font-size: x-large; margin-bottom: 10px;">
          {{ $t('label.workspace') }}
        </span>
        <a-progress
            type="dashboard"
            :percent="70"
            style="margin-top: 10px;"
        />
        <a-card
            class="dashboard-card-footer"
            :bodyStyle="{padding: '5px'}"
        >
          <div class="dashboard-card-field">20 / 30</div>
        </a-card>
      </a-card>
      <a-card class="dashboard-card" hoverable>
        <span style="font-size: x-large; margin-bottom: 10px;">
          {{ $t('label.workspace') }}
        </span>
        <a-progress
            type="dashboard"
            :percent="70"
            style="margin-top: 10px;"
        />
        <a-card
            class="dashboard-card-footer"
            :bodyStyle="{padding: '5px'}"
        >
          <span class="dashboard-card-field">20 / 30</span>
        </a-card>
      </a-card>
      <a-card class="dashboard-card" hoverable>
      <span style="font-size: x-large; margin-bottom: 10px;">
        {{ $t('label.workspace') }}
      </span>
        <a-progress
            type="dashboard"
            :percent="70"
            style="margin-top: 10px;"
        />
        <a-card
            class="dashboard-card-footer"
            :bodyStyle="{padding: '5px'}"
        >
          <span class="dashboard-card-field">20 / 30</span>
        </a-card>
      </a-card>
      <a-card class="dashboard-card" hoverable>
        <span style="font-size: x-large; margin-bottom: 10px;">
          {{ $t('label.workspace') }}
        </span>
        <a-progress
            type="dashboard"
            :percent="70"
            style="margin-top: 10px;"
            />
        <a-card
            class="dashboard-card-footer"
            :bodyStyle="{padding: '5px'}"
        >
          <span class="dashboard-card-field">20 / 30</span>
        </a-card>
      </a-card>
      <a-card class="dashboard-card" hoverable>
        <span style="font-size: x-large; margin-bottom: 10px;">
          {{ $t('label.workspace') }}
        </span>
        <a-progress
            type="dashboard"
            :percent="70"
            style="margin-top: 10px;"
            />
        <a-card
            class="dashboard-card-footer"
            :bodyStyle="{padding: '5px'}"
        >
          <span class="dashboard-card-field">20 / 30</span>
        </a-card>
      </a-card>
      <a-card class="dashboard-card" hoverable>
        <span style="font-size: x-large; margin-bottom: 10px;">
          {{ $t('label.workspace') }}
        </span>
        <a-progress
            type="dashboard"
            :percent="70"
            style="margin-top: 10px;"
            />
        <a-card
            class="dashboard-card-footer"
            :bodyStyle="{padding: '5px'}"
        >
          <span class="dashboard-card-field">20 / 30</span>
        </a-card>
      </a-card>
      <a-card class="dashboard-card" hoverable>
        <span style="font-size: x-large; margin-bottom: 10px;">
          {{ $t('label.workspace') }}
        </span>
        <a-progress
            type="dashboard"
            :percent="70"
            style="margin-top: 10px;"
            />
        <a-card
            class="dashboard-card-footer"
            :bodyStyle="{padding: '5px'}"
        >
          <span class="dashboard-card-field">20 / 30</span>
        </a-card>
      </a-card>
      <a-card class="dashboard-card" hoverable>
        <span style="font-size: x-large; margin-bottom: 10px;">
          {{ $t('label.workspace') }}
        </span>
        <a-progress
            type="dashboard"
            :percent="70"
            style="margin-top: 10px;"
            />
        <a-card
            class="dashboard-card-footer"
            :bodyStyle="{padding: '5px'}"
        >
          <span class="dashboard-card-field">20 / 30</span>
        </a-card>
      </a-card>
    </div>
    <div class="dashboard-right-grid">
      <a-card :bordered="true" class="dashboard-right-card">
        <a-timeline style="text-align: left;">
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
      </a-card>
    </div>
  </div>
</template>

<script>
export default {
  name: "Dashboard"
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

}

.dashboard-card{
  max-height: 291px;
}
.dashboard-right-card{
  width: 100%;
  height: auto;
  margin: 14px
}

.dashboard-card-footer{
  border: 0px;
  border-top: 1px solid #e8e8e8;
  margin-top: 20px;
  padding: 0px;
}

.dashboard-card-field{
  position: relative;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  margin: 0;
}::after{
 }
</style>
