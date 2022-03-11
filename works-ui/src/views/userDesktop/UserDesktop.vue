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
                @click="fetchRefresh()"
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
            <a-col
              flex="auto"
              v-for="workspace in dataList"
              class="dashboard-a-col"
            >
              <a-card
                :title="workspace.name + '(' + workspace.description + ')'"
                bodyStyle="margin: 1px;"
                ><font-awesome-icon :icon="['fa-chrome']" />
                <a-row :gutter="4">
                  <a-col
                    v-for="vm in workspace.instanceList"
                    class="dashboard-a-col"
                  >
                    <a-card
                      hoverable
                      :title="vm.name"
                      style="width: 200px; text-align: center"
                    >
                      <DesktopOutlined
                        :style="{
                          fontSize: '40px',
                          color: vm.mold_status == 'Running' ? 'black' : 'pink',
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
                          @confirm="favorite(vm.uuid, vm.favorite)"
                        >
                          <a-tooltip placement="bottom">
                            <template #title>{{
                              $t("label.favorite.off")
                            }}</template>

                            <StarFilled
                              v-if="vm.favorite"
                              :id="vm.uuid + '-TRUE'"
                              :style="{ color: '#ffd700' }"
                            />
                          </a-tooltip>
                          <a-tooltip placement="bottom">
                            <template #title>{{
                              $t("label.favorite.on")
                            }}</template>
                            <StarOutlined
                              v-if="!vm.favorite"
                              :id="vm.uuid + '-FALSE'"
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
                          <template #title>{{
                            $t("label.rdp.connect")
                          }}</template>
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
                              ? $t("label.desktop.console.connect.ready")
                              : $t("label.desktop.console.connect.notready")
                          }}</template>

                          <CodeFilled
                            v-if="vm.handshake_status == 'Ready'"
                            :style="{ color: '#333' }"
                            @click="connectConsole(workspace.uuid, vm.uuid)"
                          />
                          <CodeFilled v-else :style="{ color: '#d9dbdf' }" />
                        </a-tooltip>
                        <!-- </a-popconfirm> -->

                        <a-Popover placement="topLeft" trigger="click">
                          <template #content>
                            <a-popconfirm
                              v-if="vm.mold_status == 'Running'"
                              :title="$t('modal.confirm.user.vmStop')"
                              :ok-text="$t('label.ok')"
                              :cancel-text="$t('label.cancel')"
                              @confirm="vmAction(vm.uuid, 'vmStop')"
                            >
                              <a-tooltip
                                placement="bottom"
                                style="padding-right: 40px"
                              >
                                <template #title>{{
                                  $t("tooltip.vmStop")
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
                              v-if="vm.mold_status == 'Running'"
                              :title="$t('modal.confirm.user.vmRestart')"
                              :ok-text="$t('label.ok')"
                              :cancel-text="$t('label.cancel')"
                              @confirm="vmAction(vm.uuid, 'vmRestart')"
                            >
                              <a-tooltip placement="bottom">
                                <template #title>{{
                                  $t("tooltip.vmRestart")
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
                              v-if="vm.mold_status == 'Stopped'"
                              :title="$t('modal.confirm.user.vmStart')"
                              :ok-text="$t('label.ok')"
                              :cancel-text="$t('label.cancel')"
                              @confirm="vmAction(vm.uuid, 'vmStart')"
                            >
                              <a-tooltip placement="bottom">
                                <template #title>{{
                                  $t("tooltip.vmStart")
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
                        />({{ vm.mold_status }})
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
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";
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
      cryptKey: "IgmTQVMISq9t4Bj7iRz7kZklqzfoXuq1",
      spinning: ref(false),
      favPopTitle: ref(""),
      actionFrom: "VirtualMachineList",
      loading: ref(true),
      pagination,
      dataList: ref([]),
      succCnt: ref(0),
      failCnt: ref(0),
    };
  },
  created() {
    this.fetchRefresh();
    this.timer = setInterval(() => {
      //60초 자동 갱신
      this.fetchData();
    }, 60000);
  },
  methods: {
    fetchRefresh() {
      this.spinning = true;
      this.fetchData();
    },
    async fetchData() {
      // console.log("fetchData!!");
      await worksApi
        .get("/api/v1/userdesktop/" + sessionStorage.getItem("userName"))
        .then((response) => {
          if (response.status == 200) {
            if (
              response.data.workspaceList !== null &&
              response.data.workspaceList !== undefined
            ) {
              this.dataList = response.data.workspaceList;
            } else {
              this.dataList = [];
            }
          } else {
            message.error(this.$t("message.response.data.fail"));
            //console.log(response.message);
          }
        })
        .catch((error) => {
          console.log(error);
          message.error(this.$t("message.response.data.fail"));
        })
        .finally(() => {
          this.spinning = false;
        });
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
      const liteParamArr = this.dataList
        .filter((dl) => dl.uuid === worksId)[0]
        .instanceList.filter((il) => il.uuid === vmId)[0];

      liteParamArr["hostname"] = liteParamArr.ipaddress;
      delete liteParamArr.ipaddress;
      delete liteParamArr.owner_account_id;

      liteParamArr["port"] = 3389;
      liteParamArr["username"] = sessionStorage.getItem("userName");
      liteParamArr["domain"] = sessionStorage.getItem("domainName");
      liteParamArr["enable-wallpaper"] = true;
      liteParamArr["enable-font-smoothing"] = true;
      liteParamArr["enable-theming"] = true;
      liteParamArr["enable-menu-animations"] = true;
      liteParamArr["resize-method"] = "display-update";
      liteParamArr["disable-upload"] = false;
      liteParamArr["create-drive-path"] = true;
      liteParamArr["drive-name"] = "Z";
      liteParamArr["drive-path"] = "/";
      liteParamArr["enable-drive"] = true;
      //liteParamArr["security"] = "rdp";

      console.log(liteParamArr);

      const encrypted = btoa(
        this.$CryptoJS.AES.encrypt(
          JSON.stringify(liteParamArr),
          this.cryptKey
        ).toString()
      );
      //console.log(encrypted);
      window.open("/client/?enc=" + encrypted, "_blank");
    },
    async vmAction(uuid, action) {
      if (action == "vmStart") {
        message.loading(this.$t("message.vm.status.starting"), 100);
        this.worksUrl = "/api/v1/instance/VMStart/";
        this.sucMessage = "message.vm.status.start.ok";
        this.failMessage = "message.vm.status.start.fail";
      }
      if (action == "vmStop") {
        message.loading(this.$t("message.vm.status.stopping"), 100);
        this.worksUrl = "/api/v1/instance/VMStop/";
        this.sucMessage = "message.vm.status.stop.ok";
        this.failMessage = "message.vm.status.stop.fail";
      }
      if (action == "vmRestart") {
        message.loading(this.$t("message.vm.status.restarting"), 100);
        this.worksUrl = "/api/v1/instance/VMReboot/";
        this.sucMessage = "message.vm.status.restart.ok";
        this.failMessage = "message.vm.status.restart.fail";
      }

      try {
        const res = await worksApi.patch(this.worksUrl + uuid);
        console.log(res.status);
        if (res.status == 200) {
          this.succCnt = this.succCnt + 1;
        }
      } catch (error) {
        console.log(error);
        this.failCnt = this.failCnt + 1;
      }

      await this.funcDelay(12000);
      this.funcEndMessage();
      this.fetchRefresh();
    },
    async funcDelay(delay) {
      return new Promise(function (resolve) {
        setTimeout(function () {
          resolve("delay call!");
        }, delay);
      });
    },
    funcEndMessage() {
      message.destroy();
      if (this.succCnt > 0) {
        message.success(
          this.$t(this.sucMessage, {
            count: this.succCnt,
          }),
          5
        );
      }
      if (this.failCnt > 0) {
        message.error(
          this.$t(this.failMessage, {
            count: this.failCnt,
          }),
          5
        );
      }
      this.failCnt = 0;
      this.succCnt = 0;
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
