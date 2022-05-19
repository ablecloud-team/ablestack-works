


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
                <template #icon><ReloadOutlined /></template>
                {{ $t("label.refresh") }}
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
                :bodyStyle="{ margin: '1px' }"
              >
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
                      <!-- <a-popconfirm
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
                        </a-popconfirm> -->
                      <template #actions>
                        <a-tooltip placement="bottom">
                          <template #title>{{
                            $t("label.rdp.connect")
                          }}</template>
                          <Icon>
                            <template #component>
                              <img
                                src="@/assets/icons8-remote-desktop-97.png"
                                width="30"
                                height="30"
                                @click="
                                  workspace.policy.rdp_access_allow == '1'
                                    ? connectRdpClient(workspace.uuid, vm.uuid)
                                    : false
                                "
                              />
                            </template>
                          </Icon>
                        </a-tooltip>

                        <a-tooltip placement="bottom">
                          <template #title>{{
                            vm.handshake_status === "Ready"
                              ? $t("label.desktop.console.connect.ready")
                              : $t("label.desktop.console.connect.notready")
                          }}</template>
                          <Icon>
                            <template #component>
                              <img
                                src="@/assets/icons8-internet-64.png"
                                width="30"
                                height="30"
                                @click="
                                  workspace.policy.rdp_access_allow == '1'
                                    ? connectConsole(workspace.uuid, vm.uuid)
                                    : false
                                "
                              />
                            </template>
                          </Icon>
                        </a-tooltip>

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
                                      }"
                                    />
                                  </template>
                                </a-button>
                              </a-tooltip>
                            </a-popconfirm>
                          </template>
                          <MoreOutlined
                            :style="{
                              color: '#333',
                              fontSize: '24px',
                              marginTop: '5px',
                            }"
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
import { defineComponent, ref, h } from "vue";
import customProtocolCheck from "custom-protocol-check";
import Icon from "@ant-design/icons-vue";
import Apath from "@/components/Apath";
import Actions from "@/components/Actions";
import { worksApi } from "@/api/index";
import { notification, message, Button } from "ant-design-vue";
const hostname =
  process.env.VUE_APP_API_URL == ""
    ? window.location.hostname
    : process.env.VUE_APP_API_URL;
const apiPort =
  process.env.VUE_APP_API_PORT == "" ? "8082" : process.env.VUE_APP_API_PORT;
export default defineComponent({
  name: "UserDesktop",
  components: {
    Icon,
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
      osPass: ref(false),
    };
  },
  created() {
    this.fetchRefresh();
    this.timer = setInterval(() => {
      //60초 자동 갱신
      this.fetchData();
    }, 60000);
  },
  mounted() {
    this.funcOsCheck();
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
    async connectRdpClient(workspaceUuid, vmUuid) {
      if (!this.osPass) {
        message.warning(this.$t("message.userdesktop.rdp.connect.os.limit"));
        return false;
      }
      const paramArr = this.dataList
        .filter((dl) => dl.uuid === workspaceUuid)[0]
        .instanceList.filter((il) => il.uuid === vmUuid)[0];

      await worksApi
        .get(
          "/api/v1/connection/rdp/" +
            vmUuid +
            "/" +
            sessionStorage.getItem("userName")
        )
        .then((res) => {
          if (res.status == 200) {
            const url =
              "worksapp://" +
              hostname +
              ":" +
              apiPort +
              "/?full address=" +
              hostname +
              "&server port=" +
              res.data.instance.public_port +
              "&domain=" +
              sessionStorage.getItem("domainName") +
              "&username=" +
              sessionStorage.getItem("userName") +
              "&password=" +
              res.data.instance.password +
              "&instanceUuid=" +
              vmUuid +
              "&hash=" +
              res.data.instance.hash;

            console.log(url);
            customProtocolCheck(
              url,
              () => {
                this.downloadRdpClient("works-client.zip");
                message.warning(
                  this.$t("message.userdesktop.rdp.connect.client.notinstall")
                );
                //worksapp 프로토콜이 설치안됨
                const key = `open${Date.now()}`;
                notification.info({
                  message: "RDP 클라이언트 설치 안내",
                  description: h("div", [
                    this.$t("message.userdesktop.client.install.notice1"),
                    h("br"),
                    this.$t("message.userdesktop.client.install.notice2"),
                    h("br"),
                    this.$t("message.userdesktop.client.install.notice3"),
                  ]),
                  duration: 0,
                  style: {
                    width: "500px",
                  },
                  btn: () =>
                    h(
                      Button,
                      {
                        type: "primary",
                        size: "small",
                        onClick: () => notification.close(key),
                      },
                      { default: () => this.$t("label.ok") }
                    ),
                  key,
                  onClose: close,
                });
              },
              () => {},
              100
            );
          } else {
            message.error(this.$t("message.response.data.fail"));
          }
        })
        .catch((error) => {
          console.log(error);
          message.error(this.$t("fail"));
        })
        .finally(() => {});
    },
    downloadRdpClient(filename) {
      const downlink = document.createElement("a");
      downlink.setAttribute("href", "/client/" + filename);
      downlink.setAttribute("download", filename);
      downlink.style.display = "none";
      document.body.appendChild(downlink);
      downlink.click();
      document.body.removeChild(downlink);
    },
    connectConsole(workspaceUuid, vmUuid) {
      // console.log(worksId, vmId);
      const liteParamArr = this.dataList
        .filter((dl) => dl.uuid === workspaceUuid)[0]
        .instanceList.filter((il) => il.uuid === vmUuid)[0];

      liteParamArr["hostname"] = liteParamArr.ipaddress;
      delete liteParamArr.ipaddress;
      delete liteParamArr.owner_account_id;

      liteParamArr["port"] = 3389;
      liteParamArr["username"] = sessionStorage.getItem("userName");
      liteParamArr["domain"] = sessionStorage.getItem("domainName");

      liteParamArr["enable-wallpaper"] = true;
      liteParamArr["enable-font-smoothing"] = true;
      liteParamArr["enable-theming"] = false;
      liteParamArr["enable-menu-animations"] = false;
      liteParamArr["resize-method"] = "display-update";

      //liteParamArr["create-drive-path"] = true;
      liteParamArr["drive-name"] = "GUACD";
      liteParamArr["drive-path"] = "/share";
      liteParamArr["enable-drive"] = true;
      liteParamArr["disable-upload"] = false;
      liteParamArr["disable-download"] = false;
      liteParamArr["enable-printing"] = true;
      liteParamArr["printer-name"] = "VDI-PRINTER";

      liteParamArr["enable-touch"] = true;

      liteParamArr["timestamp"] = Math.floor(Date.now() / 1000);
      //liteParamArr["security"] = "rdp";

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
    funcOsCheck() {
      var os,
        ua = navigator.userAgent;

      if (ua.indexOf("Windows") != -1) {
        os = "Windows";
        this.osPass = true;
      } else if (
        ua.indexOf("iPad") != -1 ||
        ua.indexOf("iPhone") != -1 ||
        ua.indexOf("iPod") != -1
      ) {
        os = "Apple iOS";
        this.osPass = false;
      } else if (ua.indexOf("Android" != -1)) {
        os = "Android OS";
        this.osPass = false;
      } else if (ua.match(/Mac|PPC/)) {
        os = "Mac";
        this.osPass = false;
      } else if (ua.match(/Linux/)) {
        os = "Linux";
        this.osPass = false;
      } else if (ua.match(/(Free|Net|Open)BSD/)) {
        os = RegExp.$1 + "BSD";
        this.osPass = false;
      } else if (ua.match(/SunOS/)) {
        os = "Solaris";
        this.osPass = false;
      }
      //자세한 os 분류가 필요할때 사용
      // if (ua.match(/Win(dows )?NT 6\.0/)) {
      //   os = "Windows Vista";
      //   this.osPass = true;
      // } else if (ua.match(/Win(dows )?(NT 5\.1|XP)/)) {
      //   os = "Windows XP";
      //   this.osPass = true;
      // } else {
      //   if (
      //     ua.indexOf("Windows NT 5.1") != -1 ||
      //     ua.indexOf("Windows XP") != -1
      //   ) {
      //     os = "Windows XP";
      //     this.osPass = true;
      //   } else if (
      //     ua.indexOf("Windows NT 7.0") != -1 ||
      //     ua.indexOf("Windows NT 6.1") != -1
      //   ) {
      //     os = "Windows 7";
      //     this.osPass = true;
      //   } else if (
      //     ua.indexOf("Windows NT 8.0") != -1 ||
      //     ua.indexOf("Windows NT 6.2") != -1
      //   ) {
      //     os = "Windows 8";
      //     this.osPass = true;
      //   } else if (
      //     ua.indexOf("Windows NT 8.1") != -1 ||
      //     ua.indexOf("Windows NT 6.3") != -1
      //   ) {
      //     os = "Windows 8.1";
      //     this.osPass = true;
      //   } else if (
      //     ua.indexOf("Windows NT 10.0") != -1 ||
      //     ua.indexOf("Windows NT 6.4") != -1
      //   ) {
      //     os = "Windows 10";
      //     this.osPass = true;
      //   } else if (ua.match(/Win(dows )?NT( 4\.0)?/)) {
      //     os = "Windows NT";
      //     this.osPass = true;
      //   } else if (
      //     ua.indexOf("iPad") != -1 ||
      //     ua.indexOf("iPhone") != -1 ||
      //     ua.indexOf("iPod") != -1
      //   ) {
      //     os = "Apple iOS";
      //     this.osPass = false;
      //   } else if (ua.indexOf("Android" != -1)) {
      //     os = "Android OS";
      //     this.osPass = false;
      //   } else if (ua.match(/Mac|PPC/)) {
      //     os = "Mac";
      //     this.osPass = false;
      //   } else if (ua.match(/Linux/)) {
      //     os = "Linux";
      //     this.osPass = false;
      //   } else if (ua.match(/(Free|Net|Open)BSD/)) {
      //     os = RegExp.$1 + "BSD";
      //     this.osPass = false;
      //   } else if (ua.match(/SunOS/)) {
      //     os = "Solaris";
      //     this.osPass = false;
      //   }
      // }
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

