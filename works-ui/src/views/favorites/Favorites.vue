<template>
  <div class="dashboard-main-grid">
    <div class="dashboard-left-grid">
      <a-card class="dashboard-card">
        <span style="font-size: x-large; margin-bottom: 10px">
          Desktop #1
          <a-rate v-model:value="value" :tooltips="즐겨찾기" :count="1" />
        </span>
        <div>
          <DesktopOutlined :style="{ fontSize: '100px' }" />
        </div>
        <a-card class="dashboard-card-footer">
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
      <a-card class="dashboard-card">
        <span style="font-size: x-large; margin-bottom: 10px">
          Desktop #2
          <a-rate v-model:value="value" :tooltips="즐겨찾기" :count="1" />
        </span>
        <div>
          <DesktopOutlined :style="{ fontSize: '100px' }" />
        </div>
        <a-card class="dashboard-card-footer">
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
      <a-card class="dashboard-card">
        <span style="font-size: x-large; margin-bottom: 10px">
          Desktop #3
          <a-rate v-model:value="value" :tooltips="즐겨찾기" :count="1" />
        </span>
        <div>
          <DesktopOutlined :style="{ fontSize: '100px' }" />
        </div>
        <a-card class="dashboard-card-footer">
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
      <a-card class="dashboard-card">
        <span style="font-size: x-large; margin-bottom: 10px">
          Desktop #4
          <a-rate v-model:value="value" :tooltips="즐겨찾기" :count="1" />
        </span>
        <div>
          <DesktopOutlined :style="{ fontSize: '100px' }" />
        </div>
        <a-card class="dashboard-card-footer">
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
      <a-card class="dashboard-card">
        <span style="font-size: x-large; margin-bottom: 10px">
          EXEL #1
          <a-rate v-model:value="value" :tooltips="즐겨찾기" :count="1" />
        </span>
        <div>
          <AppstoreAddOutlined :style="{ fontSize: '100px' }" />
        </div>
        <a-card class="dashboard-card-footer">
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
        </a-card>
      </a-card>
      <a-card class="dashboard-card">
        <span style="font-size: x-large; margin-bottom: 10px">
          WORD #1
          <a-rate v-model:value="value" :tooltips="즐겨찾기" :count="1" />
        </span>
        <div>
          <AppstoreAddOutlined :style="{ fontSize: '100px' }" />
        </div>
        <a-card class="dashboard-card-footer">
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
        </a-card>
      </a-card>
      <a-card class="dashboard-card">
        <span style="font-size: x-large; margin-bottom: 10px">
          POWERPOINT #1
          <a-rate v-model:value="value" :tooltips="즐겨찾기" :count="1" />
        </span>
        <div>
          <AppstoreAddOutlined :style="{ fontSize: '100px' }" />
        </div>
        <a-card class="dashboard-card-footer">
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
        </a-card>
      </a-card>
      <a-card class="dashboard-card">
        <span style="font-size: x-large; margin-bottom: 10px">
          세무서프로그램 #1
          <a-rate v-model:value="value" :tooltips="즐겨찾기" :count="1" />
        </span>
        <div>
          <AppstoreAddOutlined :style="{ fontSize: '100px' }" />
        </div>
        <a-card class="dashboard-card-footer">
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
        </a-card>
      </a-card>
    </div>
    <div class="dashboard-right-grid">
      <a-card
        title="공지사항 제목"
        :bordered="true"
        class="dashboard-right-card"
        hoverable
      >
        <p style="font-weight: bold">작성자 : ABLECLOUD</p>
        <p class="dashboard-right-card-p">
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다.
        </p>
      </a-card>
      <a-card
        title="공지사항 제목"
        :bordered="true"
        class="dashboard-right-card"
        hoverable
      >
        <p style="font-weight: bold">작성자 : ABLECLOUD</p>
        <p class="dashboard-right-card-p">
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다.
        </p>
      </a-card>
      <a-card
        title="공지사항 제목"
        :bordered="true"
        class="dashboard-right-card"
        hoverable
      >
        <p style="font-weight: bold">작성자 : ABLECLOUD</p>
        <p class="dashboard-right-card-p">
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다. 공지사항 내용입니다. 공지사항 내용입니다.
          공지사항 내용입니다.
        </p>
      </a-card>
    </div>
  </div>
</template>

<script>
import {
  AppstoreAddOutlined,
  DesktopOutlined,
  DownloadOutlined,
  Html5Outlined,
  MoreOutlined,
} from "@ant-design/icons-vue";
import { defineComponent, ref } from "vue";

export default defineComponent({
  name: "UserDashboard",
  components: {
    DesktopOutlined,
    DownloadOutlined,
    Html5Outlined,
    MoreOutlined,
    AppstoreAddOutlined,
  },
  setup() {
    const value = ref(3);
    return {
      value,
    };
  },
});
</script>

<style scoped>
.dashboard-main-grid {
  display: grid;
  grid-template-columns: 3fr 1fr;
  grid-template-rows: 1fr;
  height: 100%;
  text-align: center;
  padding: 5px;
}
.dashboard-left-grid {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr 1fr;
  grid-template-rows: 291px 291px;
  grid-gap: 14px;
  padding: 14px;
  text-align: center;
}

.dashboard-right-grid {
  display: grid;
  grid-template-columns: 1fr;
  grid-template-rows: 1fr 1fr 1fr;
  grid-gap: 14px;
  padding: 14px;
  margin: 14px;
  background: white;
}

.dashboard-card {
  max-height: 291px;
}
.dashboard-right-card {
  width: 100%;
  text-align: left;
  margin-bottom: 10px;
  border-width: 8px;
  box-shadow: 5px 5px;
}
.dashboard-right-card-p {
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  line-height: 16px;
  max-height: 50px;
  -webkit-line-clamp: 3; /* 표시하고자 하는 라인 수 */
  -webkit-box-orient: vertical;
}
.dashboard-card-footer {
  border: 0px;
  border-top: 1px solid #e8e8e8;
  margin-top: 20px;
  padding: 0px;
}
.dashboard-card-field {
  position: relative;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  margin: 0;
}
</style>
