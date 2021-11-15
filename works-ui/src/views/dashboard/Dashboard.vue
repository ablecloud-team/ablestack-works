<template>
  <div style="width: 100%">
    <a-row id="content-header-row">
      <a-col flex="100%" class="dashboard-a-col-one">
        <a-button
          type="primary"
          ghost
          shape="round"
          size="medium"
          @click="refresh()"
        >
          <template #icon>
            <ReloadOutlined /> {{ $t("label.fetch.latest") }}
          </template>
        </a-button>
      </a-col>
    </a-row>
  </div>
  <div style="width: 100%; padding: 10px">
    <a-spin :spinning="spinning" size="large">
      <a-row>
        <a-col flex="100%">
          <a-row :gutter="12" type="flex">
            <a-col flex="50%" class="dashboard-a-col">
              <a-card
                :title="$t('label.workspace.count')"
                class="dashboard-a-card-cl"
                hoverable
                @click="$router.push({ name: 'Workspace' })"
              >
                <span style="font-size: 80px">{{ workspaceCount }}</span>
              </a-card>
            </a-col>
            <a-col flex="50%" class="dashboard-a-col">
              <a-card
                :title="$t('label.desktop.count')"
                class="dashboard-a-card-cl"
                hoverable
                @click="$router.push({ name: 'VirtualMachine' })"
              >
                <span style="font-size: 80px">{{ instanceCount }}</span>
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
              <a-card
                :title="$t('label.desktop.connected.count')"
                class="dashboard-a-card-cl"
                hoverable
              >
                <span style="font-size: 80px">{{ desktopConCount }}</span>
              </a-card>
            </a-col>
            <a-col flex="50%" class="dashboard-a-col">
              <a-card
                :title="$t('label.app.connected.count')"
                class="dashboard-a-card-cl"
                hoverable
              >
                <span style="font-size: 80px">{{ appConCount }}</span>
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
    </a-spin>
  </div>
</template>
<script>
import { defineComponent, reactive, ref } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";
export default defineComponent({
  name: "Dashboard",
  components: {},
  props: {},
  setup() {
    const state = reactive({});
    return {
      state,
    };
  },
  data() {
    return {
      spinning: ref(false),
      workspaceCount: ref("0"),
      instanceCount: ref("0"),
      desktopConCount: ref("0"),
      appConCount: ref("0"),
    };
  },
  created() {
    this.fetchData();
    this.timer = setInterval(() => {
      //60초 자동 갱신
      this.fetchData();
    }, 60000);
  },
  unmounted() {
    clearInterval(this.timer);
  },
  methods: {
    refresh() {
      this.spinning = true;
      setTimeout(() => {
        this.fetchData();
        this.spinning = false;
      }, 1000);
    },
    fetchData() {
      worksApi
        .get("/api/v1/dashboard")
        .then((response) => {
          if (response.status == 200) {
            this.workspaceCount = response.data.result.workspaceCount;
            this.instanceCount = response.data.result.instanceCount;
            this.desktopConCount = "0";
            this.appConCount = "0";
          }
        })
        .catch((error) => {
          message.error(this.$t("message.response.data.fail"));
          console.log(error.message);
        });
    },
  },
});
</script>

<style scoped>
.dashboard-a-card-cl {
  width: 100%;
  height: 100%;
  text-align: center;
}
.dashboard-a-col-one {
  text-align: right;
  padding: 10px;
  padding-bottom: 0px;
}

.dashboard-a-col {
  padding-bottom: 10px;
}
#content-action {
  text-align: right;
}
</style>
