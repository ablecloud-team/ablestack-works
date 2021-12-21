<template>
  <div id="content-layout">
    <a-layout>
      <a-layout-header id="content-header">
        <div>
          <a-row>
            <!-- 오른쪽 경로 -->
            <a-col id="content-path" :span="12">
              <Apath :paths="[$t('label.desktop')]" />
              <a-button
                shape="round"
                style="margin-left: 18px"
                size="small"
                @click="refresh()"
              >
                <template #icon>
                  <ReloadOutlined /> {{ $t("label.refresh") }}
                </template>
              </a-button>
            </a-col>
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <a-spin :spinning="spinning" size="large">
          <a-card bodyStyle="">
            <a-list
              :grid="{ gutter: 16, xs: 1, sm: 2, md: 3, lg: 3, xl: 5, xxl: 10 }"
              :pagination="pagination"
              :data-source="instanceList"
            >
              <template #renderItem="{ item }">
                <a-list-item>
                  <a-card
                    hoverable
                    style="width: 100%; text-align: center"
                    :title="item.name"
                  >
                    <DesktopOutlined
                      :style="{
                        fontSize: '40px',
                        color:
                          item.handshake_status === 'Ready' ? 'black' : 'pink',
                      }"
                    />

                    <template class="ant-card-actions" #actions>
                      <a-popconfirm
                        :title="
                          item.favorite == true
                            ? '즐겨찾기 해제하시겠습니까?'
                            : '즐겨찾기에 추가하시겠습니까?'
                        "
                        :ok-text="$t('label.ok')"
                        :cancel-text="$t('label.cancel')"
                        @confirm="favorite(item.id, item.favorite)"
                      >
                        <a-tooltip v-if="item.favorite" placement="bottom">
                          <template #title>{{ $t("즐겨찾기 해제") }}</template>
                          <StarFilled
                            :id="item.id + '-TRUE'"
                            :style="{ color: '#ffd700', fontSize: '18px' }"
                          />
                        </a-tooltip>
                        <a-tooltip v-else placement="bottom">
                          <template #title>{{ $t("즐겨찾기 추가") }}</template>
                          <StarOutlined
                            :id="item.id + '-FALSE'"
                            :style="{ color: '#d9dbdf', fontSize: '18px' }"
                          />
                        </a-tooltip>
                      </a-popconfirm>

                      <a-popconfirm
                        :title="'RDP 파일을 다운로드 하시겠습니까?'"
                        :ok-text="$t('label.ok')"
                        :cancel-text="$t('label.cancel')"
                        @confirm="downloadRDP(item.name)"
                      >
                        <a-tooltip placement="bottom">
                          <template #title>{{ $t("RDP 다운로드") }}</template>
                          <CloudDownloadOutlined
                            :style="{ color: '#292929', fontSize: '18px' }"
                          />
                        </a-tooltip>
                      </a-popconfirm>

                      <a-popconfirm
                        v-if="item.handshake_status == 'Ready'"
                        :title="'데스크톱에 접속하시겠습니까?'"
                        :ok-text="$t('label.ok')"
                        :cancel-text="$t('label.cancel')"
                        @confirm="connectConsole(item.id)"
                      >
                        <a-tooltip placement="bottom">
                          <template #title>{{ $t("데스크톱 접속") }}</template>
                          <Html5Outlined
                            :style="{
                              color: '#333',
                              fontSize: '18px',
                            }"
                          />
                        </a-tooltip>
                      </a-popconfirm>
                      <a-Popover placement="topLeft">
                        <template #content>
                          <a-popconfirm
                            v-if="item.state == 'Running'"
                            :title="'가상머신을 종료하시겠습니까?'"
                            :ok-text="$t('label.ok')"
                            :cancel-text="$t('label.cancel')"
                            @confirm="vmState(item.id, item.state)"
                          >
                            <a-tooltip
                              placement="bottom"
                              style="padding-right: 40px"
                            >
                              <template #title>{{
                                $t("가상머신 종료")
                              }}</template>
                              <a-button shape="circle">
                                <template #icon>
                                  <PoweroffOutlined
                                    :style="{
                                      color: '#333',
                                      fontSize: '18px',
                                      marginTop: '3px',
                                    }"
                                  />
                                </template>
                              </a-button>
                            </a-tooltip>
                          </a-popconfirm>
                          <a-popconfirm
                            v-if="item.state == 'Running'"
                            :title="'가상머신을 재시작하시겠습니까?'"
                            :ok-text="$t('label.ok')"
                            :cancel-text="$t('label.cancel')"
                            @confirm="vmState(item.id, item.state)"
                          >
                            <a-tooltip placement="bottom">
                              <template #title>{{
                                $t("가상머신 재시작")
                              }}</template>

                              <a-button shape="circle">
                                <template #icon>
                                  <ReloadOutlined
                                    :style="{
                                      color: '#333',
                                      fontSize: '18px',
                                      marginTop: '3px',
                                    }"
                                  />
                                </template>
                              </a-button>
                            </a-tooltip>
                          </a-popconfirm>
                          <a-popconfirm
                            v-if="item.state == 'Stopped'"
                            :title="'가상머신을 시작하시겠습니까?'"
                            :ok-text="$t('label.ok')"
                            :cancel-text="$t('label.cancel')"
                            @confirm="vmState(item.id, item.state)"
                          >
                            <a-tooltip placement="bottom">
                              <template #title>{{
                                $t("가상머신 시작")
                              }}</template>
                              <a-button shape="circle">
                                <template #icon>
                                  <CaretRightOutlined
                                    :style="{
                                      color: '#333',
                                      fontSize: '18px',
                                      marginTop: '3px',
                                    }"
                                  />
                                </template>
                              </a-button>
                            </a-tooltip>
                          </a-popconfirm>
                        </template>
                        <MoreOutlined
                          :style="{ color: '#333', fontSize: '18px' }"
                        />
                      </a-Popover>
                    </template>

                    <br /><br />
                    <a-tooltip placement="bottom">
                      <template #title>{{ item.handshake_status }}</template>
                      <a-badge
                        class="head-example"
                        :color="
                          item.handshake_status === 'Not Ready' ||
                          item.handshake_status === 'Pending'
                            ? 'red'
                            : item.handshake_status === 'Joining' ||
                              item.handshake_status === 'Joined'
                            ? 'yellow'
                            : item.handshake_status === 'Ready'
                            ? 'green'
                            : 'red'
                        "
                        :text="
                          item.handshake_status === 'Not Ready' ||
                          item.handshake_status === 'Pending'
                            ? $t('label.vm.status.initializing')
                            : item.handshake_status === 'Joining' ||
                              item.handshake_status === 'Joined'
                            ? $t('label.vm.status.configuring')
                            : item.handshake_status === 'Ready'
                            ? $t('label.vm.status.ready')
                            : $t('label.vm.status.notready')
                        "
                      />
                      ({{ item.state }})
                    </a-tooltip>

                    <br />
                    <WindowsFilled /> {{ item.ostype }}
                  </a-card>
                </a-list-item>
              </template>
            </a-list>
          </a-card>
        </a-spin>
      </a-layout-content>
    </a-layout>
  </div>
</template>

<script>
import { defineComponent, ref } from "vue";
import Apath from "@/components/Apath";
import Actions from "@/components/Actions";

export default defineComponent({
  name: "UserDesktop",
  components: {
    Apath,
    Actions,
  },
  setup() {
    const pagination = {
      onChange: (page) => {
        console.log(page);
      },
      pageSize: 20,
    };
    return {
      //desktopIconStyle: ref([ {"fontSize": '40px', "color": '#cc0000'} ]),
      spinning: ref(false),
      favPopTitle: ref(""),
      actionFrom: "VirtualMachineList",
      loading: ref(true),
      pagination,
      instanceList: [
        {
          id: "11111111-111111-1111111-1-1111",
          name: "vm18898",
          state: "Running",
          handshake_status: "Pending",
          ostype: "Windows 10 (64bit)",
          favorite: true,
          hostname: "10.1.1.150",
          port: 3389,
          username: "user1",
          password: "Ablecloud1!",
          domain: "cjs",
        },
        {
          id: "11111111-111111-1111111-1-2222",
          name: "vm62",
          state: "Stopped",
          handshake_status: "Joining",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
        {
          id: "11111111-111111-1111111-1-3333",
          name: "vm36",
          state: "Stopped",
          handshake_status: "Joining",
          ostype: "Windows 10 (64bit)",
          favorite: true,
        },
        {
          id: "11111111-111111-1111111-1-4444",
          name: "vm46",
          state: "Running",
          handshake_status: "Not Ready",
          ostype: "Windows 10 (64bit)",
          favorite: true,
        },
        {
          id: "11111111-111111-1111111-1-5555",
          name: "vm55",
          state: "Stopped",
          handshake_status: "Not Ready",
          ostype: "Windows 10 (64bit)",
          favorite: true,
        },
        {
          id: "11111111-111111-1111111-1-6666",
          name: "vm645",
          state: "Running",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
          hostname: "10.1.1.150",
          port: 3389,
          username: "user1",
          password: "Ablecloud1!",
          domain: "cjs",
        },
        {
          id: "11111111-111111-1111111-1-7777",
          name: "vm27",
          state: "Stopped",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: true,
        },
        {
          id: "11111111-111111-1111111-1-8888",
          name: "vm822",
          state: "Running",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
        {
          id: "11111111-111111-1111111-1-9999",
          name: "vm9",
          state: "Stopped",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
        {
          id: "11111111-111111-1111111-1-1231",
          name: "vm10",
          state: "Stopped",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
        {
          id: "11111111-111111-1111111-1-8888",
          name: "vm822",
          state: "Running",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
        {
          id: "11111111-111111-1111111-1-9999",
          name: "vm9",
          state: "Stopped",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
        {
          id: "11111111-111111-1111111-1-1231",
          name: "vm10",
          state: "Stopped",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
        {
          id: "11111111-111111-1111111-1-8888",
          name: "vm822",
          state: "Running",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
        {
          id: "11111111-111111-1111111-1-9999",
          name: "vm9",
          state: "Stopped",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
        {
          id: "11111111-111111-1111111-1-1231",
          name: "vm10",
          state: "Stopped",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
        {
          id: "11111111-111111-1111111-1-8888",
          name: "vm822",
          state: "Running",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
        {
          id: "11111111-111111-1111111-1-9999",
          name: "vm9",
          state: "Stopped",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
        {
          id: "11111111-111111-1111111-1-1231",
          name: "vm10",
          state: "Stopped",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
        {
          id: "11111111-111111-1111111-1-8888",
          name: "vm822",
          state: "Running",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
        {
          id: "11111111-111111-1111111-1-9999",
          name: "vm9",
          state: "Stopped",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
        {
          id: "11111111-111111-1111111-1-1231",
          name: "vm10",
          state: "Stopped",
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
      ],
    };
  },
  created() {
    this.refresh();
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
      console.log("fetchData!!");
    },
    favorite(uuid, bool) {
      console.log(uuid + " :: " + bool);
      // //const elId = event.currentTarget.id;

      // document.getElementById(elId).style.display = "none";
      // if (elId.includes("-TRUE")) {
      //   document.getElementById(val + "-FALSE").style.display = "block";
      // } else if (elId.includes("-FALSE")) {
      //   document.getElementById(val + "-TRUE").style.display = "block";
      // }

      // document.getElementById(val).style.display = "none";
      // document.getElementById(val).style.display = "none";
    },
    downloadRDP(vmName) {
      let rdpConfig = ",fd,f,f,f,\ny777m7m\njjmjmjmjm\ndg34g3g3g3g3g3";

      var element = document.createElement("a");
      element.setAttribute(
        "href",
        "data:text/plain;charset=utf-8," + encodeURIComponent(rdpConfig)
      );

      element.setAttribute("download", vmName + ".rdp");
      element.style.display = "none";
      document.body.appendChild(element);
      element.click();

      document.body.removeChild(element);
    },
    connectConsole(uuid, bool) {
      const selArr = this.instanceList.filter((it) => it.id === uuid);
      console.log(new URLSearchParams(selArr[0]).toString());

      window.open(
        // window.location.hostname +
        "http://10.10.1.24:8088/#/?" +
          new URLSearchParams(selArr[0]).toString(),
        "_blank"
      );
    },
    vmState(uuid, bool) {
      console.log(uuid + " :: " + bool);
    },
  },
});
</script>

<style scoped>
::v-deep(.ant-badge-status-dot) {
  width: 12px;
  height: 12px;
}
.dashboard-a-col {
  padding-bottom: 10px;
}
.ant-card-actions {
  height: 10px;
}

#content-layout .ant-layout-header {
  text-align: left;
  background: white;
  border: 1px solid #e8e8e8;
  /*color: #fff;*/
  font-size: 14px;
  line-height: 1.5;
  padding: 18px;
  height: auto;
}

#content-path {
  text-align: left;
  align-items: center;
  display: flex;
  height: 32px;
}

#content-action {
  text-align: right;
}

.ant-btn {
  margin-right: 8px;
}
</style>
