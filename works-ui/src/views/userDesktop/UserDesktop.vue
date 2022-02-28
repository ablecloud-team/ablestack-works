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
          <a-row class="dashboard-a-row" :gutter="4">
            <a-col v-for="workspace in dataList" class="dashboard-a-col">
              <a-card
                :title="
                  'WORKSPACE : ' +
                  workspace.workspace +
                  '(' +
                  workspace.desc +
                  ')'
                "
                bodyStyle="margin: 1px; fontSize: 100px;"
              >
                <a-row :gutter="4">
                  <a-col
                    v-for="vm in workspace.instanceList"
                    class="dashboard-a-col"
                  >
                    <a-card
                      hoverable
                      style="width: 200px; text-align: center"
                      :title="vm.name"
                    >
                      <DesktopOutlined
                        :style="{
                          fontSize: '40px',
                          color: vm.state == 'Running' ? 'black' : 'pink',
                        }"
                      />

                      <template class="ant-card-actions" #actions>
                        <a-popconfirm
                          :title="
                            vm.favorite == true
                              ? '즐겨찾기 해제하시겠습니까?'
                              : '즐겨찾기에 추가하시겠습니까?'
                          "
                          :ok-text="$t('label.ok')"
                          :cancel-text="$t('label.cancel')"
                          @confirm="favorite(vm.id, vm.favorite)"
                        >
                          <a-tooltip placement="bottom">
                            <template #title>{{
                              $t("즐겨찾기 해제")
                            }}</template>

                            <StarFilled
                              v-if="vm.favorite"
                              :id="vm.id + '-TRUE'"
                              :style="{ color: '#ffd700' }"
                            />
                          </a-tooltip>
                          <a-tooltip placement="bottom">
                            <template #title>{{
                              $t("즐겨찾기 추가")
                            }}</template>
                            <StarOutlined
                              v-if="!vm.favorite"
                              :id="vm.id + '-FALSE'"
                              :style="{ color: '#d9dbdf' }"
                            />
                          </a-tooltip>
                        </a-popconfirm>
                        <!-- 
                        <a-popconfirm
                          :title="'RDP 파일을 다운로드 하시겠습니까?'"
                          :ok-text="$t('label.ok')"
                          :cancel-text="$t('label.cancel')"
                          @confirm="downloadRDP(vm.name)"
                        > -->
                        <a-tooltip placement="bottom">
                          <template #title>{{ $t("RDP 접속") }}</template>
                          <CloudDownloadOutlined
                            :style="{ color: '#292929' }"
                            @click="downloadRDP(vm.name)"
                          />
                        </a-tooltip>
                        <!-- </a-popconfirm> -->

                        <!-- <a-popconfirm
                          :title="'데스크톱에 접속하시겠습니까?'"
                          :ok-text="$t('label.ok')"
                          :cancel-text="$t('label.cancel')"
                          @confirm="connectConsole(vm.id)"
                          :disabled="
                            vm.handshake_status == 'Ready' ? false : true
                          "
                        > -->
                        <a-tooltip placement="bottom">
                          <template #title>{{
                            vm.handshake_status === "Ready"
                              ? $t("데스크톱 접속")
                              : $t("데스크톱 접속 불가")
                          }}</template>
                          <CodeFilled
                            v-if="vm.handshake_status == 'Ready'"
                            :style="{ color: '#333' }"
                            @click="connectConsole(workspace.id, vm.id)"
                          />
                          <CodeFilled v-else :style="{ color: '#d9dbdf' }" />
                        </a-tooltip>
                        <!-- </a-popconfirm> -->

                        <a-Popover placement="topLeft" trigger="click">
                          <template #content>
                            <a-popconfirm
                              v-if="vm.state == 'Running'"
                              :title="'가상머신을 종료하시겠습니까?'"
                              :ok-text="$t('label.ok')"
                              :cancel-text="$t('label.cancel')"
                              @confirm="vmState(vm.id, vm.state)"
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
                              v-if="vm.state == 'Running'"
                              :title="'가상머신을 재시작하시겠습니까?'"
                              :ok-text="$t('label.ok')"
                              :cancel-text="$t('label.cancel')"
                              @confirm="vmState(vm.id, vm.state)"
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
                              v-if="vm.state == 'Stopped'"
                              :title="'가상머신을 시작하시겠습니까?'"
                              :ok-text="$t('label.ok')"
                              :cancel-text="$t('label.cancel')"
                              @confirm="vmState(vm.id, vm.state)"
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
                        <template #title>{{ vm.handshake_status }}</template>
                        <a-badge
                          class="head-example"
                          :color="
                            vm.handshake_status == 'Ready' ? 'green' : 'red'
                          "
                          :text="
                            vm.handshake_status == 'Ready'
                              ? $t('label.vm.status.ready')
                              : $t('label.vm.status.notready')
                          "
                        />({{ vm.state }})
                      </a-tooltip>

                      <br />
                      <WindowsFilled /> {{ vm.ostype }}
                    </a-card>
                  </a-col>
                </a-row>
              </a-card>
            </a-col>
          </a-row>
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
      cryptKey: "ikAkd39aszkdEghj",
      spinning: ref(false),
      favPopTitle: ref(""),
      actionFrom: "VirtualMachineList",
      loading: ref(true),
      pagination,
      // instanceList: [
      //   {
      //     id: "11111111-111111-1111111-1-123123",
      //     name: "ws1-014",
      //     state: "Running",
      //     handshake_status: "Ready",
      //     ostype: "Windows 10 (64bit)",
      //     favorite: false,
      //     hostname: "10.1.1.82",
      //     port: 3389,
      //     username: "user1",
      //     password: "~!fkal1228",
      //     domain: "able",
      //     "enable-wallpaper": true,
      //     "enable-font-smoothing": true,
      //     "enable-theming": true,
      //     "enable-menu-animations": true,
      //     "resize-method": "display-update",
      //   },
      //   {
      //     id: "11111111-111111-1111111-1-234234",
      //     name: "ws1-015",
      //     state: "Running",
      //     handshake_status: "Ready",
      //     ostype: "Windows 10 (64bit)",
      //     favorite: false,
      //     hostname: "10.1.1.111",
      //     port: 3389,
      //     username: "user1",
      //     password: "~!fkal1228",
      //     domain: "able",
      //     "enable-wallpaper": false,
      //     "enable-font-smoothing": false,
      //     "enable-theming": false,
      //     "enable-menu-animations": false,
      //     "resize-method": "reconnect",
      //   },
      //   {
      //     id: "11111111-111111-1111111-1-345345",
      //     name: "ws1-016",
      //     state: "Running",
      //     handshake_status: "Ready",
      //     ostype: "Windows 10 (64bit)",
      //     favorite: false,
      //     hostname: "10.1.1.73",
      //     port: 3389,
      //     username: "user1",
      //     password: "~!fkal1228",
      //     domain: "able",
      //     "enable-wallpaper": true,
      //     "enable-font-smoothing": true,
      //     "enable-theming": true,
      //     "enable-menu-animations": true,
      //     "resize-method": "display-update",
      //   },
      //   {
      //     id: "11111111-111111-1111111-1-8888",
      //     name: "vm822",
      //     state: "Running",
      //     handshake_status: "Ready",
      //     ostype: "Windows 10 (64bit)",
      //     favorite: false,
      //   },
      //   {
      //     id: "11111111-111111-1111111-1-9999",
      //     name: "vm9",
      //     state: "Stopped",
      //     handshake_status: "Ready",
      //     ostype: "Windows 10 (64bit)",
      //     favorite: false,
      //   },
      //   {
      //     id: "11111111-111111-1111111-1-1231",
      //     name: "vm10",
      //     state: "Stopped",
      //     handshake_status: "Ready",
      //     ostype: "Windows 10 (64bit)",
      //     favorite: false,
      //   },
      //   {
      //     id: "11111111-111111-1111111-1-1231",
      //     name: "vm10",
      //     state: "Stopped",
      //     handshake_status: "Ready",
      //     ostype: "Windows 10 (64bit)",
      //     favorite: false,
      //   },
      // ],
      dataList: [
        {
          id: "1111111-1111111111",
          workspace: "workspace 1",
          desc: "개인업무용1",
          instanceList: [
            {
              id: "111111111111",
              name: "ws1-014",
              state: "Running",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
              hostname: "10.1.1.82",
              port: 3389,
              username: "user1",
              password: "~!fkal1228",
              domain: "able",
              "enable-wallpaper": true,
              "enable-font-smoothing": true,
              "enable-theming": true,
              "enable-menu-animations": true,
              "resize-method": "display-update",
            },
            {
              id: "222222222222222",
              name: "ws1-015",
              state: "Running",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
              hostname: "10.1.1.111",
              port: 3389,
              username: "user1",
              password: "~!fkal1228",
              domain: "able",
              "enable-wallpaper": false,
              "enable-font-smoothing": false,
              "enable-theming": false,
              "enable-menu-animations": false,
              "resize-method": "reconnect",
            },
            {
              id: "3333333333333333333333333",
              name: "ws1-016",
              state: "Running",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
              hostname: "10.1.1.73",
              port: 3389,
              username: "user1",
              password: "~!fkal1228",
              domain: "able",
              "enable-wallpaper": true,
              "enable-font-smoothing": true,
              "enable-theming": true,
              "enable-menu-animations": true,
              "resize-method": "display-update",
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
              id: "11111111-111111-1111111-1-1231",
              name: "vm10",
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
              id: "11111111-111111-1111111-1-1231",
              name: "vm10",
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
              id: "11111111-111111-1111111-1-1231",
              name: "vm10",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
          ],
        },
        {
          id: "22222222-22222222",
          workspace: "workspace 2",
          desc: "개인업무용2",
          instanceList: [
            {
              id: "22222222-11111",
              name: "vm10",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
            {
              id: "22222222-2222",
              name: "vm10",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
            {
              id: "22222222-33333",
              name: "vm10",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
          ],
        },
        {
          id: "3333333-3333333",
          workspace: "workspace 3",
          desc: "개인업무용3",
          instanceList: [
            {
              id: "3333333-11111",
              name: "vm10",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
            {
              id: "3333333-22222",
              name: "vm10",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
          ],
        },
        {
          id: "44444-44444",
          workspace: "workspace 4",
          desc: "개인업무용4",
          instanceList: [
            {
              id: "44444-1111111",
              name: "vm10",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
          ],
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
      // console.log("fetchData!!");
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
      let rdpConfig =
        "targetisaadjoined:i:0\n" +
        "hubdiscoverygeourl:s:\n" +
        "redirected video capture encoding quality:i:0\n" +
        "camerastoredirect:s:\n" +
        "gatewaybrokeringtype:i:0\n" +
        "use redirection server name:i:0\n" +
        "alternate shell:s:\n" +
        "disable themes:i:0\n" +
        "disable cursor setting:i:1\n" +
        "remoteapplicationname:s:\n" +
        "resourceprovider:s:\n" +
        "disable menu anims:i:1\n" +
        "remoteapplicationcmdline:s:\n" +
        "promptcredentialonce:i:0\n" +
        "gatewaycertificatelogonauthority:s:\n" +
        "audiocapturemode:i:0\n" +
        "prompt for credentials on client:i:0\n" +
        "gatewayhostname:s:\n" +
        "remoteapplicationprogram:s:\n" +
        "gatewayusagemethod:i:2\n" +
        "screen mode id:i:2\n" +
        "use multimon:i:0\n" +
        "authentication level:i:3\n" +
        "desktopwidth:i:0\n" +
        "desktopheight:i:0\n" +
        "redirectsmartcards:i:0\n" +
        "redirectclipboard:i:1\n" +
        "forcehidpioptimizations:i:0\n" +
        "full address:s:10.10.1.110:3389\n" +
        "username:s:user1\n" +
        "password 51:01000000d08c9ddf0115d1118c7a00c04fc297eb01000000f7d63b57853e464ab1d317a8417beae70000000002000000000003660000c000000010000000bb23d3b56c695c0006389b38891c5b1a0000000004800000a00000001000000031421b3dbd931cda62251dd2120d53e4180000008240be3347e4a17740eadf42040e72b26f63d948260d626814000000be3e14b3083cb81535647a6a5d7755daf9f06dc1\n" +
        "domain:s:able\n" +
        "drivestoredirect:s:\n" +
        "loadbalanceinfo:s:\n" +
        "networkautodetect:i:1\n" +
        "enablecredsspsupport:i:2\n" +
        "redirectprinters:i:0\n" +
        "autoreconnection enabled:i:1\n" +
        "session bpp:i:32\n" +
        "administrative session:i:0\n" +
        "audiomode:i:0\n" +
        "bandwidthautodetect:i:1\n" +
        "authoring tool:s:\n" +
        "connection type:i:7\n" +
        "remoteapplicationmode:i:0\n" +
        "disable full window drag:i:0\n" +
        "gatewayusername:s:\n" +
        "shell working directory:s:\n" +
        "wvd endpoint pool:s:\n" +
        "remoteapplicationappid:s:\n" +
        "allow font smoothing:i:1\n" +
        "connect to console:i:0\n" +
        "disable wallpaper:i:0\n" +
        "gatewayaccesstoken:s:\n" +
        "promptcredentialonce:i:0\n";

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
    connectConsole(worksId, vmId) {
      // console.log(worksId, vmId);
      console.log(
        this.dataList
          .filter((dl) => dl.id === worksId)[0]
          .instanceList.filter((il) => il.id === vmId)
      );
      const liteParamArr = this.dataList
        .filter((dl) => dl.id === worksId)[0]
        .instanceList.filter((il) => il.id === vmId)[0];
      console.log(new URLSearchParams(liteParamArr).toString());

      const encrypted = btoa(
        this.$CryptoJS.AES.encrypt(
          JSON.stringify(liteParamArr),
          this.cryptKey
        ).toString()
      );
      console.log(encrypted);
      window.open("/client/?enc=" + encrypted, "_blank");
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
.dashboard-a-row {
  width: 100%;
  padding-top: 5px;
  padding-right: 5px;
  padding-left: 5px;
}
.dashboard-a-col {
  padding-bottom: 5px;
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
