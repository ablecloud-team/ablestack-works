<template>
  <a-layout>
    <a-layout-content class="dashboard-left-content">
      <div class="dashboard-div-grid">
        <a-row :gutter="[4,4]" style="padding-top: 6px; padding-bottom: 3px">
          <a-col :span="8">
            <a-card
                :loading="loading"
            >
              <h3><a-badge status="success"/>Desktop name1</h3>
              <DesktopOutlined style="font-size: 90px"/>
              <br>
              <a-button shape="round">
                <template #icon>
                  <DownloadOutlined />
                </template>
              </a-button>
              <a-button shape="round">
                <template #icon>
                  <Html5Outlined />
                </template>
              </a-button>
              <a-popover placement="right">
                <template #content>
                  <a-button shape="circle" :size="small" style="margin: 0 5px">
                    <template #icon>
                      <PoweroffOutlined />
                    </template>
                  </a-button>
                  <a-button shape="circle" :size="small" style="margin: 0 5px">
                    <template #icon>
                      <PlayCircleOutlined />
                    </template>
                  </a-button>
                  <a-button shape="circle" :size="small" style="margin: 0 5px">
                    <template #icon>
                      <ReloadOutlined />
                    </template>
                  </a-button>
                  <a-button shape="circle" :size="small" style="margin: 0 5px">
                    <template #icon>
                      <SwapOutlined />
                    </template>
                  </a-button>
                  <a-button shape="circle" :size="small" style="margin: 0 5px" danger>
                    <template #icon>
                      <DeleteOutlined />
                    </template>
                  </a-button>
                </template>
                <a-button type="text" class="ant-btn-icon-only"><MoreOutlined style="font-size: large"/></a-button>
              </a-popover>

            </a-card>
          </a-col>
          <a-col :span="8">
            <a-card
                :loading="loading"
            >
              <h3><a-badge status="error"/>Application name1</h3>
              <AppstoreOutlined style="font-size: 90px"/>
              <br>
              <a-button type="primary" shape="round" :size="size">
                <template #icon>
                  <DownloadOutlined />
                </template>
              </a-button>
              <a-button type="primary" shape="round" :size="size">
                <template #icon>
                  <Html5Outlined />
                </template>
              </a-button>
            </a-card>
          </a-col>
          <a-col :span="8">
            <a-card
                :loading="loading"
            >
              <h3><a-badge status="success"/>Application name2</h3>
              <AppstoreOutlined style="font-size: 90px"/>
              <br>
              <a-button type="primary" shape="round" :size="size">
                <template #icon>
                  <DownloadOutlined />
                </template>
              </a-button>
              <a-button type="primary" shape="round" :size="size">
                <template #icon>
                  <Html5Outlined />
                </template>
              </a-button>
            </a-card>
          </a-col>
        </a-row>
        <a-row :gutter="[4,4]" style="padding-top: 3px; padding-bottom: 6px">
          <a-col :span="8">
            <a-card
                :loading="loading"
            >
              <h3><a-badge status="error"/>Desktop name2</h3>
              <DesktopOutlined style="font-size: 90px"/>
              <br>
              <a-button type="primary" shape="round" :size="size">
                <template #icon>
                  <DownloadOutlined />
                </template>
              </a-button>
              <a-button type="primary" shape="round" :size="size">
                <template #icon>
                  <Html5Outlined />
                </template>
              </a-button>
              <a-popover placement="topRight">
                <template #content>
                  <p>시작</p>
                  <p>정지</p>
                </template>
                <a-button class="ant-btn-icon-only"><MoreOutlined style="font-size: large"/></a-button>
              </a-popover>

            </a-card>
          </a-col>
          <a-col :span="8">

          </a-col>
          <a-col :span="8">

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
import {
  AppstoreOutlined, DeleteOutlined,
  DesktopOutlined,
  DownloadOutlined,
  Html5Outlined,
  MoreOutlined, PlayCircleOutlined,
  PoweroffOutlined, ReloadOutlined, SwapOutlined
} from '@ant-design/icons-vue'
export default defineComponent({
  components:{
    DesktopOutlined,
    DownloadOutlined,
    Html5Outlined,
    MoreOutlined,
    AppstoreOutlined,
    PoweroffOutlined,
    PlayCircleOutlined,
    DeleteOutlined,
    ReloadOutlined,
    SwapOutlined,
  }
})
</script>

<style scoped>
.dashboard-right-content{
  background: white;
  flex: 0 0 30% !important;
  max-width: 30% !important;
  min-width: 30% !important;
  width: 30% !important;
  margin-top: 5px;
  padding: 5px 5px 5px 5px;
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
