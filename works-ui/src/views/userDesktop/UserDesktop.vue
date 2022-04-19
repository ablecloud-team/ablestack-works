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
                bodyStyle="margin: 1px;"
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

                      <template #actions>
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

                        <a-tooltip placement="bottom">
                          <template #title>{{
                            $t("label.rdp.connect")
                          }}</template>
                          <Icon>
                            <template #component>
                              <svg
                                id="remote-desktop"
                                width="24"
                                height="24"
                                :fill="
                                  workspace.policy.rdp_access_allow == '1'
                                    ? ''
                                    : '#d9dbdf'
                                "
                                viewBox="0 0 24 24"
                                @click="
                                  workspace.policy.rdp_access_allow == '1'
                                    ? connectRdp(workspace.uuid, vm.uuid)
                                    : false
                                "
                              >
                                <path
                                  d="M3,2A2,2 0 0,0 1,4V16C1,17.11 1.9,18 3,18H10V20H8V22H16V20H14V18H21A2,2 0 0,0 23,16V4A2,2 0 0,0 21,2M3,4H21V16H3M15,5L11.5,8.5L15,12L16.4,10.6L14.3,8.5L16.4,6.4M9,8L7.6,9.4L9.7,11.5L7.6,13.6L9,15L12.5,11.5"
                                />
                              </svg>
                              <!-- <svg
                                fill="none"
                                height="24"
                                viewBox="0 0 24 24"
                                width="24"
                                xmlns="http://www.w3.org/2000/svg"
                                @click="connectRdp(workspace.uuid, vm.uuid)"
                              >
                                <path
                                  d="M17.0514 4.32178L18.4656 5.73599L14.223 9.97863L18.4656 14.2213L17.0514 15.6355L11.3946 9.97863L17.0514 4.32178Z"
                                  fill="currentColor"
                                />
                                <path
                                  d="M6.94864 19.6785L5.53442 18.2643L9.77706 14.0216L5.53442 9.77897L6.94864 8.36476L12.6055 14.0216L6.94864 19.6785Z"
                                  fill="currentColor"
                                />
                              </svg> -->
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
                              <svg
                                id="mdi-console"
                                width="24"
                                height="24"
                                :fill="
                                  vm.handshake_status == 'Ready'
                                    ? ''
                                    : '#d9dbdf'
                                "
                                viewBox="0 0 24 24"
                                @click="connectConsole(workspace.uuid, vm.uuid)"
                              >
                                <path
                                  d="M20,19V7H4V19H20M20,3A2,2 0 0,1 22,5V19A2,2 0 0,1 20,21H4A2,2 0 0,1 2,19V5C2,3.89 2.9,3 4,3H20M13,17V15H18V17H13M9.58,13L5.57,9H8.4L11.7,12.3C12.09,12.69 12.09,13.33 11.7,13.72L8.42,17H5.59L9.58,13Z"
                                />
                              </svg>
                            </template>
                          </Icon>
                          <!-- <CodeFilled
                            v-if="vm.handshake_status == 'Ready'"
                            :style="{ color: '#333' }"
                            @click="connectConsole(workspace.uuid, vm.uuid)"
                          /> -->
                          <!-- <CodeFilled v-else :style="{ color: '#d9dbdf' }" /> -->
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
import { defineComponent, ref, TrackOpTypes } from "vue";
import Icon from "@ant-design/icons-vue";
import Apath from "@/components/Apath";
import Actions from "@/components/Actions";
import axios from "axios";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";
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
    async connectRdp(workspaceUuid, vmUuid) {
      await worksApi
        .get(
          "/api/v1/connection/rdp/" +
            vmUuid +
            "/" +
            sessionStorage.getItem("userName")
        )
        .then((response) => {
          if (response.status == 200) {
            console.log(response.data.instance.public_port);
          } else {
            message.error(this.$t("message.response.data.fail"));
          }
        })
        .catch((error) => {
          console.log(error);
          message.error(this.$t("message.response.data.fail"));
        })
        .finally(() => {});

      const paramArr = this.dataList
        .filter((dl) => dl.uuid === workspaceUuid)[0]
        .instanceList.filter((il) => il.uuid === vmUuid)[0];
      console.log(paramArr);

      // const agent = navigator.userAgent.toLowerCase();
      // // console.log(navigator);
      // // console.log(window);
      // if (
      //   (navigator.appName == "Netscape" &&
      //     navigator.userAgent.search("Trident") != -1) ||
      //   agent.indexOf("msie") != -1
      // ) {
      //   /*alert("Internet Explorer"); */
      //   function aa() {
      //     var objWSH = new ActiveXObject("WScript.Shell");
      //     var retval = objWSH.Run("C:/Windows/SysWOW64/notepad.exe", 1, true);
      //   }
      // } else if (agent.indexOf("chrome") != -1) {
      //   function aa() {
      //     /*alert("HAVE TO INSTALL."); */
      //     var objWSH = new ActiveXObject("WScript.Shell");
      //     var retval = objWSH.Run("C:/Windows/SysWOW64/notepad.exe", 1, true);
      //   }
      // }

      // let rdpParam = new URLSearchParams();
      // rdpParam.append("full address", paramArr.ipaddress);
      // rdpParam.append("port", paramArr.port || 3389);
      // rdpParam.append("domain", sessionStorage.getItem("domainName"));
      // rdpParam.append("username", sessionStorage.getItem("userName"));
      // rdpParam.append("password 51", paramArr.password);

      // await axios
      //   .get("worksapp://works/", rdpParam)
      //   .then((response) => {})
      //   .catch((error) => {})
      //   .finally(() => {});

      // try {
      //   const link = document.createElement("a");
      //   link.setAttribute("href", "worksapp://aa.aaa.com");
      //   // link.setAttribute("download", "");
      //   link.style.display = "none";
      //   document.body.appendChild(link);
      //   link.click();
      //   document.body.removeChild(link);
      // } catch (error) {
      //   console.log(error);
      // }
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

      liteParamArr["enable-wallpaper"] = false;
      liteParamArr["enable-font-smoothing"] = false;
      liteParamArr["enable-theming"] = false;
      liteParamArr["enable-menu-animations"] = false;
      liteParamArr["resize-method"] = "display-update";

      liteParamArr["create-drive-path"] = true;
      liteParamArr["drive-name"] = "GUACD";
      liteParamArr["drive-path"] = "/share";
      liteParamArr["enable-drive"] = true;
      liteParamArr["disable-upload"] = false;
      liteParamArr["disable-download"] = false;

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
