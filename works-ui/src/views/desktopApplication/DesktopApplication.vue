<template>
  <a-card class="desktopapp-main-card" :bodyStyle="{ padding: '5px' }">
    <a-card class="desktopapp-row-card" :bodyStyle="{ padding: '5px' }">
      <a-card
        title="Workspace#1"
        class="desktopapp-left-card"
        :headStyle="{
          textAlign: 'center',
          fontSize: '25px',
          padding: '0px',
          height: '35px',
          verticalAlign: 'middle',
        }"
        :bodyStyle="{ marginBottom: '1px', fontSize: '18px' }"
      >
        <p>
          이 워크 스페이스의 설명을 적는 공간입니다.이 워크 스페이스의 설명을
          적는 공간입니다.
        </p>
        <p>사용가능한 가상머신 : 2EA</p>
        <p style="text-align: center">
          <a-button type="primary" :size="large" block>바로 실행</a-button>
        </p>
      </a-card>
      <a-card
        class="desktopapp-right-card"
        :bodyStyle="{ padding: '0px', fontSize: '18px' }"
        :bordered="false"
      >
        <a-card class="desktopapp-card-main">
          <span style="font-size: x-large; margin-bottom: 10px">
            Desktop #1
            <a-rate v-model:value="value" :tooltips="즐겨찾기" :count="1" />
          </span>
          <div style="text-align: center">
            <DesktopOutlined :style="{ fontSize: '85px' }" />
          </div>
          <a-card
            class="desktopapp-card-main-footer"
            :bodyStyle="{ padding: '5px', textAlign: 'center' }"
          >
            <a-button shape="circle" :size="small" style="margin: 0 2px 0 2px">
              <template #icon>
                <DownloadOutlined />
              </template>
            </a-button>
            <a-button shape="circle" :size="small" style="margin: 0 2px 0 2px">
              <template #icon>
                <Html5Outlined />
              </template>
            </a-button>
            <a-button
              type="text"
              shape="circle"
              :size="small"
              style="margin: 0 2px 0 2px"
            >
              <template #icon>
                <MoreOutlined />
              </template>
            </a-button>
          </a-card>
        </a-card>
        <a-card class="desktopapp-card-main">
          <span style="font-size: x-large; margin-bottom: 10px">
            Desktop #1
            <a-rate v-model:value="value" :tooltips="즐겨찾기" :count="1" />
          </span>
          <div style="text-align: center">
            <DesktopOutlined :style="{ fontSize: '85px' }" />
          </div>
          <a-card
            class="desktopapp-card-main-footer"
            :bodyStyle="{ padding: '5px', textAlign: 'center' }"
          >
            <a-button shape="circle" :size="small" style="margin: 0 2px 0 2px">
              <template #icon>
                <DownloadOutlined />
              </template>
            </a-button>
            <a-button shape="circle" :size="small" style="margin: 0 2px 0 2px">
              <template #icon>
                <Html5Outlined />
              </template>
            </a-button>
            <a-button
              type="text"
              shape="circle"
              :size="small"
              style="margin: 0 2px 0 2px"
            >
              <template #icon>
                <MoreOutlined />
              </template>
            </a-button>
          </a-card>
        </a-card>
        <a-card class="desktopapp-card-main">
          <span style="font-size: x-large; margin-bottom: 10px">
            Desktop #1
            <a-rate v-model:value="value" :tooltips="즐겨찾기" :count="1" />
          </span>
          <div style="text-align: center">
            <DesktopOutlined :style="{ fontSize: '85px' }" />
          </div>
          <a-card
            class="desktopapp-card-main-footer"
            :bodyStyle="{ padding: '5px', textAlign: 'center' }"
          >
            <a-button shape="circle" :size="small" style="margin: 0 2px 0 2px">
              <template #icon>
                <DownloadOutlined />
              </template>
            </a-button>
            <a-button shape="circle" :size="small" style="margin: 0 2px 0 2px">
              <template #icon>
                <Html5Outlined />
              </template>
            </a-button>
            <a-button
              type="text"
              shape="circle"
              :size="small"
              style="margin: 0 2px 0 2px"
            >
              <template #icon>
                <MoreOutlined />
              </template>
            </a-button>
          </a-card>
        </a-card>
        <a-card class="desktopapp-card-main">
          <span style="font-size: x-large; margin-bottom: 10px">
            Desktop #1
            <a-rate v-model:value="value" :tooltips="즐겨찾기" :count="1" />
          </span>
          <div style="text-align: center">
            <DesktopOutlined :style="{ fontSize: '85px' }" />
          </div>
          <a-card
            class="desktopapp-card-main-footer"
            :bodyStyle="{ padding: '5px', textAlign: 'center' }"
          >
            <a-button shape="circle" :size="small" style="margin: 0 2px 0 2px">
              <template #icon>
                <DownloadOutlined />
              </template>
            </a-button>
            <a-button shape="circle" :size="small" style="margin: 0 2px 0 2px">
              <template #icon>
                <Html5Outlined />
              </template>
            </a-button>
            <a-button
              type="text"
              shape="circle"
              :size="small"
              style="margin: 0 2px 0 2px"
            >
              <template #icon>
                <MoreOutlined />
              </template>
            </a-button>
          </a-card>
        </a-card>
        <a-card class="desktopapp-card-main">
          <span style="font-size: x-large; margin-bottom: 10px">
            Desktop #1
            <a-rate v-model:value="value" :tooltips="즐겨찾기" :count="1" />
          </span>
          <div style="text-align: center">
            <DesktopOutlined :style="{ fontSize: '85px' }" />
          </div>
          <a-card
            class="desktopapp-card-main-footer"
            :bodyStyle="{ padding: '5px', textAlign: 'center' }"
          >
            <a-button shape="circle" :size="small" style="margin: 0 2px 0 2px">
              <template #icon>
                <DownloadOutlined />
              </template>
            </a-button>
            <a-button shape="circle" :size="small" style="margin: 0 2px 0 2px">
              <template #icon>
                <Html5Outlined />
              </template>
            </a-button>
            <a-button
              type="text"
              shape="circle"
              :size="small"
              style="margin: 0 2px 0 2px"
            >
              <template #icon>
                <MoreOutlined />
              </template>
            </a-button>
          </a-card>
        </a-card>
      </a-card>
    </a-card>
  </a-card>
</template>

<script>
import {
  DesktopOutlined,
  DownloadOutlined,
  Html5Outlined,
  MoreOutlined,
} from "@ant-design/icons-vue";

export default {
  name: "DesktopApp",
  components: {
    DesktopOutlined,
    DownloadOutlined,
    Html5Outlined,
    MoreOutlined,
  },
};
</script>

<style scoped>
.desktopapp-main-card {
  height: 100%;
  margin: 12px;
  padding: 0;
  display: grid;
  background: #f0f2f5;
}
.desktopapp-row-card {
  display: grid;
}
.desktopapp-left-card {
  width: 450px;
  height: 250px;
  float: left;
  border-width: 3px;
}
.desktopapp-right-card {
  width: calc(100% - 510px);
  height: auto;
  float: right;
}
::v-deep .ant-card-head-title {
  padding: 5px;
}
.desktopapp-card-main {
  max-height: 270px;
  width: 270px;
  text-align: center;
  float: left;
  margin-right: 5px;
  margin-bottom: 5px;
  border-width: 3px;
}
.desktopapp-card-main-footer {
  border: 0px;
  border-top: 1px solid #e8e8e8;
  margin-top: 20px;
  padding: 0px;
}
</style>
