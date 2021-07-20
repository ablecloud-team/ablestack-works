<template>
  <div style="background-color: #ffffff; padding: 20px;" class="layout-content">
    <div style="display: inline-block;">
    <a-row :gutter="12">
      <a-col
          :xs="12"
          :sm="8"
          :md="6"
          :style="{ marginBottom: '12px' }"
      >
        <div class="ant-card-body" style="padding: 20px 24px 8px;">
          <chart-card>
            <div class="capacity-dashboard-chart-card-inner">
              <h3>워크스페이스 수량</h3>
              <a-progress
                  type="dashboard"
                  :percent="75"
                  :width="100" />
            </div>
            <template>Footer</template>
          </chart-card>
          <div class="chart-card-footer">
            <div class="chart-card-field">
              224.00 GB / 750.91 GB
            </div>
          </div>
        </div>
      </a-col>
      <a-col
          :xs="12"
          :sm="8"
          :md="6"
          :style="{ marginBottom: '12px' }"
      >
        <div class="ant-card-body" style="padding: 20px 24px 8px;">
          <chart-card>
            <div class="capacity-dashboard-chart-card-inner">
              <h3>가상머신 수</h3>
              <a-progress
                  type="dashboard"
                  :percent="75"
                  :width="100" />
            </div>
            <template>Footer</template>
          </chart-card>
          <div class="chart-card-footer">
            <div class="chart-card-field">
              224.00 GB / 750.91 GB
            </div>
          </div>
        </div>
      </a-col>
      <a-col
          :xs="12"
          :sm="8"
          :md="6"
          :style="{ marginBottom: '12px' }"
      >
        <div class="ant-card-body" style="padding: 20px 24px 8px;">
          <chart-card>
            <div class="capacity-dashboard-chart-card-inner">
              <h3>가상데스크톱 연결 갯수</h3>
              <a-progress
                  type="dashboard"
                  :percent="75"
                  :width="100" />
            </div>
            <template>Footer</template>
          </chart-card>
          <div class="chart-card-footer">
            <div class="chart-card-field">
              224.00 GB / 750.91 GB
            </div>
          </div>
        </div>
      </a-col>
      <a-col
          :xs="12"
          :sm="8"
          :md="6"
          :style="{ marginBottom: '12px' }"
      >
        <div class="ant-card-body" style="padding: 20px 24px 8px;">
          <chart-card>
            <div class="capacity-dashboard-chart-card-inner">
              <h3>앱 가상화 연결</h3>
              <a-progress
                  type="dashboard"
                  :percent="75"
                  :width="100" />
            </div>
            <template>Footer</template>
          </chart-card>
          <div class="chart-card-footer">
            <div class="chart-card-field">
              224.00 GB / 750.91 GB
            </div>
          </div>
        </div>
      </a-col>
    </a-row>
    <a-row :gutter="12">
      <a-col
          :xs="12"
          :sm="8"
          :md="6"
          :style="{ marginBottom: '12px' }"
      >
        <div class="ant-card-body" style="padding: 20px 24px 8px;">
          <chart-card>
            <div class="capacity-dashboard-chart-card-inner">
              <h3>전체 CPU 할당량</h3>
              <a-progress
                  type="dashboard"
                  :percent="75"
                  :width="100" />
            </div>
            <template>Footer</template>
          </chart-card>
          <div class="chart-card-footer">
            <div class="chart-card-field">
              224.00 GB / 750.91 GB
            </div>
          </div>
        </div>
      </a-col>
      <a-col
          :xs="12"
          :sm="8"
          :md="6"
          :style="{ marginBottom: '12px' }"
      >
        <div class="ant-card-body" style="padding: 20px 24px 8px;">
          <chart-card>
            <div class="capacity-dashboard-chart-card-inner">
              <h3>전체 메모리 할당량</h3>
              <a-progress
                  type="dashboard"
                  :percent="75"
                  :width="100" />
            </div>
            <template>Footer</template>
          </chart-card>
          <div class="chart-card-footer">
            <div class="chart-card-field">
              224.00 GB / 750.91 GB
            </div>
          </div>
        </div>
      </a-col>
      <a-col
          :xs="12"
          :sm="8"
          :md="6"
          :style="{ marginBottom: '12px' }"
      >
        <div class="ant-card-body" style="padding: 20px 24px 8px;">
          <chart-card>
            <div class="capacity-dashboard-chart-card-inner">
              <h3>전체 Disk 할당량</h3>
              <a-progress
                  type="dashboard"
                  :percent="75"
                  :width="100" />
            </div>
            <template>Footer</template>
          </chart-card>
          <div class="chart-card-footer">
            <div class="chart-card-field">
              224.00 GB / 750.91 GB
            </div>
          </div>
        </div>
      </a-col>
      <a-col
          :xs="12"
          :sm="8"
          :md="6"
          :style="{ marginBottom: '12px' }"
      >
        <div class="ant-card-body" style="padding: 20px 24px 8px;">
          <chart-card>
            <div class="capacity-dashboard-chart-card-inner">
              <h3>IP 사용 수량</h3>
              <a-progress
                  type="dashboard"
                  :percent="75"
                  :width="100" />
            </div>
            <template>Footer</template>
          </chart-card>
          <div class="chart-card-footer">
            <div class="chart-card-field">
              224.00 GB / 750.91 GB
            </div>
          </div>
        </div>
      </a-col>
    </a-row>
    </div>
    <div class="ant-col ant-col-xl-6" style="padding-left: 6px; padding-right: 6px; display: inline-block;">
      <div class="ant-card">
        <div class="ant-card-body" style="padding: 20px 24px 8px;">
          <div class="chart-card-header">
            <div  class="chart-card-meta">
              <span  class="chart-card-title">
              </span>
              <span class="chart-card-action"></span>
            </div>
            <div class="chart-card-total">
              <span ></span>
            </div>
          </div>
          <div class="chart-card-footer">
            <div  class="chart-card-field">
              <div class="capacity-dashboard-footer">
                <ul class="ant-timeline">
                  <li class="ant-timeline-item">
                    <div class="ant-timeline-item-tail"></div>
                    <div class="ant-timeline-item-head ant-timeline-item-head-green"></div>
                    <div class="ant-timeline-item-content">\
                      <span  style="color: rgb(153, 153, 153);">
                        <small >20 Jul 2021 05:51:11</small></span>
                      <br >
                      <span style="color: rgb(102, 102, 102);">
                        <small >
                          <a href="#/event/8781bc55-4f9a-4032-bf71-2958a2a21918" class="">USER.LOGIN</a>
                        </small>
                      </span>
                      <br >
                      <span style="color: rgb(170, 170, 170);">(admin) user has logged in from IP Address /211.230.14.134</span>
                    </div>
                  </li>
                  <li class="ant-timeline-item">
                    <div class="ant-timeline-item-tail"></div>
                    <div class="ant-timeline-item-head ant-timeline-item-head-green"></div>
                    <div class="ant-timeline-item-content">
                      <span style="color: rgb(153, 153, 153);">
                        <small >20 Jul 2021 05:23:10</small>
                      </span>
                      <br >
                      <span style="color: rgb(102, 102, 102);">
                        <small >
                          <a href="#/event/1f01450e-77d5-4f9b-9f94-696011415486" class="">USER.LOGIN</a>
                        </small>
                      </span>
                      <br >
                      <span style="color: rgb(170, 170, 170);">(admin) user has logged in from IP Address /172.22.0.6</span>
                    </div>
                  </li>
                  <li class="ant-timeline-item">
                    <div class="ant-timeline-item-tail"></div>
                    <div class="ant-timeline-item-head ant-timeline-item-head-green"></div>
                    <div class="ant-timeline-item-content">
                      <span style="color: rgb(153, 153, 153);">
                        <small>20 Jul 2021 04:06:14</small>
                      </span>
                      <br >
                      <span style="color: rgb(102, 102, 102);">
                        <small >
                          <a href="#/event/f924f194-a4e6-433b-b8e2-ddc158e131b3" class="">USER.LOGIN</a>
                        </small>
                      </span>
                      <br >
                      <span style="color: rgb(170, 170, 170);">(ycyun) user has logged in from IP Address /211.230.14.134</span>
                    </div>
                  </li>
                  <li class="ant-timeline-item">
                    <div class="ant-timeline-item-tail"></div>
                    <div class="ant-timeline-item-head ant-timeline-item-head-green"></div>
                    <div class="ant-timeline-item-content">
                      <span style="color: rgb(153, 153, 153);">
                        <small>20 Jul 2021 03:56:55</small>
                      </span>
                      <br >
                      <span style="color: rgb(102, 102, 102);"><small >
                        <a href="#/event/debc2c90-5545-4ea1-89e2-6fbcf2f24367" class="">USER.LOGIN</a>
                      </small>
                      </span>
                      <br >
                      <span style="color: rgb(170, 170, 170);">(admin) user has logged in from IP Address /172.22.0.6</span>
                    </div>
                  </li>
                  <li class="ant-timeline-item">
                    <div class="ant-timeline-item-tail"></div>
                    <div class="ant-timeline-item-head ant-timeline-item-head-green"></div>
                    <div class="ant-timeline-item-content">
                      <span style="color: rgb(153, 153, 153);">
                        <small>20 Jul 2021 03:30:41</small>
                      </span>
                      <br >
                      <span style="color: rgb(102, 102, 102);">
                        <small>
                        <a href="#/event/42f2556b-4d22-4f04-8aaf-d039fe0a09ea" class="">USER.LOGIN</a>
                        </small>
                      </span>
                      <br >
                      <span style="color: rgb(170, 170, 170);">(admin) user has logged in from IP Address /211.230.14.134</span>
                    </div>
                  </li>
                  <li class="ant-timeline-item ant-timeline-item-last">
                    <div class="ant-timeline-item-tail"></div>
                    <div class="ant-timeline-item-head ant-timeline-item-head-green"></div>
                    <div class="ant-timeline-item-content">
                      <span style="color: rgb(153, 153, 153);">
                        <small>20 Jul 2021 03:24:46</small>
                      </span>
                      <br >
                      <span style="color: rgb(102, 102, 102);">
                        <small>
                          <a href="#/event/c87c3bec-f69e-4e0c-9e5a-102203ae4ebe" class="">VM.START</a>
                        </small>
                      </span>
                      <br >
                      <span style="color: rgb(170, 170, 170);">(admin) Successfully completed starting Vm. Vm Id: 264e05fd-9e89-43a0-a8c4-c8e898c34f3b</span>
                    </div>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Dashboard"
}
</script>

<style lang="less" scoped>
.capacity-dashboard {
  &-wrapper {
    display: flex;
    margin-bottom: 12px;
  }

  &-chart-card-inner {
    text-align: center;
    white-space: nowrap;
    overflow: hidden;
  }

  &-select {
    width: 100%; // for flexbox causes

    .ant-select {
      width: 100%; // to fill flex item width
    }
  }

  &-button-wrapper {
    margin-left: 12px;
  }

  &-button {
    width: auto;
    padding-left: 12px;
  }

  &-button-icon {
    font-size: 16px;
    padding: 2px;
  }

  &-footer {
    padding-top: 12px;
    padding-left: 3px;
    white-space: normal;

  }
}
.layout-content{
  flex: auto;
}
.chart-card-footer{
  text-align: center;
}

.chart-card-field{
  display: inline-block;
  display: block;
}
</style>
