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
                style="margin-left: 20px"
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
          <div
            id="content-body"
            v-for="workspace in dataList"
            style="width: 100%; padding: 5px"
          >
            <a-card
              :title="'WORKSPACE : ' + workspace.workspace"
              bodyStyle="margin: 1px; fontSize: 100px;"
            >
              <a-row>
                <a-col flex="100%">
                  <a-row :gutter="12" type="flex">
                    <a-col
                      v-for="vm in workspace.instanceList"
                      flex="10%"
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

                          <a-popconfirm
                            :title="'RDP 파일을 다운로드 하시겠습니까?'"
                            :ok-text="$t('label.ok')"
                            :cancel-text="$t('label.cancel')"
                            @confirm="downloadRDP(vm.name)"
                          >
                            <a-tooltip placement="bottom">
                              <template #title>{{
                                $t("RDP 다운로드")
                              }}</template>
                              <CloudDownloadOutlined
                                :style="{ color: '#292929' }"
                              />
                            </a-tooltip>
                          </a-popconfirm>

                          <a-popconfirm
                            :title="'데스크톱에 접속하시겠습니까?'"
                            :ok-text="$t('label.ok')"
                            :cancel-text="$t('label.cancel')"
                            @confirm="connectConsole(vm.id)"
                            :disabled="
                              vm.handshake_status == 'Ready' ? false : true
                            "
                          >
                            <a-tooltip placement="bottom">
                              <template #title>{{
                                vm.handshake_status === "Ready"
                                  ? $t("데스크톱 접속")
                                  : $t("데스크톱 접속 불가")
                              }}</template>
                              <CodeFilled
                                :style="{
                                  color:
                                    vm.handshake_status == 'Ready'
                                      ? '#333'
                                      : '#d9dbdf',
                                }"
                              />
                            </a-tooltip>
                          </a-popconfirm>

                          <a-popconfirm
                            :title="
                              vm.state == 'Running'
                                ? '가상머신을 종료하시겠습니까?'
                                : '가상머신을 시작하시겠습니까?'
                            "
                            :ok-text="$t('label.ok')"
                            :cancel-text="$t('label.cancel')"
                            @confirm="vmState(vm.id, vm.state)"
                          >
                            <a-tooltip placement="bottom">
                              <template #title>{{
                                vm.state === "Running"
                                  ? $t("가상머신 종료")
                                  : $t("가상머신 시작")
                              }}</template>
                              <PoweroffOutlined
                                v-if="vm.state == 'Running'"
                                :style="{ color: '#333' }"
                              />
                              <CaretRightOutlined
                                v-else
                                :style="{ color: '#333' }"
                              />
                            </a-tooltip>
                          </a-popconfirm>
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
                </a-col>
              </a-row>
            </a-card>
          </div>

          <!-- <a-card bodyStyle="margin: 1px; fontSize: 100px;">
            <a-list
              :grid="{ gutter: 16, xs: 1, sm: 2, md: 3, lg: 3, xl: 6, xxl: 10 }"
              :pagination="pagination"
              :data-source="instanceList"
            >
              <template #renderItem="{ item }">
                <a-list-item>
                  <a-card
                    hoverable
                    style="width: 200px; text-align: center"
                    :title="item.name"
                  >
                    <DesktopOutlined
                      :style="{
                        fontSize: '40px',
                        color: item.state == 'Running' ? 'black' : 'pink',
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
                        <a-tooltip placement="bottom">
                          <template #title>{{ $t("즐겨찾기 해제") }}</template>

                          <StarFilled
                            v-if="item.favorite"
                            :id="item.id + '-TRUE'"
                            :style="{ color: '#ffd700' }"
                          />
                        </a-tooltip>
                        <a-tooltip placement="bottom">
                          <template #title>{{ $t("즐겨찾기 추가") }}</template>
                          <StarOutlined
                            v-if="!item.favorite"
                            :id="item.id + '-FALSE'"
                            :style="{ color: '#d9dbdf' }"
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
                            :style="{ color: '#292929' }"
                          />
                        </a-tooltip>
                      </a-popconfirm>

                      <a-popconfirm
                        :title="'데스크톱에 접속하시겠습니까?'"
                        :ok-text="$t('label.ok')"
                        :cancel-text="$t('label.cancel')"
                        @confirm="connectConsole(item.id)"
                        :disabled="
                          item.handshake_status == 'Ready' ? false : true
                        "
                      >
                        <a-tooltip placement="bottom">
                          <template #title>{{
                            item.handshake_status === "Ready"
                              ? $t("데스크톱 접속")
                              : $t("데스크톱 접속 불가")
                          }}</template>
                          <CodeFilled
                            :style="{
                              color:
                                item.handshake_status == 'Ready'
                                  ? '#333'
                                  : '#d9dbdf',
                            }"
                          />
                        </a-tooltip>
                      </a-popconfirm>

                      <a-popconfirm
                        :title="
                          item.state == 'Running'
                            ? '가상머신을 종료하시겠습니까?'
                            : '가상머신을 시작하시겠습니까?'
                        "
                        :ok-text="$t('label.ok')"
                        :cancel-text="$t('label.cancel')"
                        @confirm="vmState(item.id, item.state)"
                      >
                        <a-tooltip placement="bottom">
                          <template #title>{{
                            item.state === "Running"
                              ? $t("가상머신 종료")
                              : $t("가상머신 시작")
                          }}</template>
                          <PoweroffOutlined
                            v-if="item.state == 'Running'"
                            :style="{ color: '#333' }"
                          />
                          <CaretRightOutlined
                            v-else
                            :style="{ color: '#333' }"
                          />
                        </a-tooltip>
                      </a-popconfirm>
                    </template>
                    <br /><br />

                    <a-tooltip placement="bottom">
                      <template #title>{{ item.handshake_status }}</template>
                      <a-badge
                        class="head-example"
                        :color="
                          item.handshake_status == 'Ready' ? 'green' : 'red'
                        "
                        :text="
                          item.handshake_status == 'Ready'
                            ? $t('label.vm.status.ready')
                            : $t('label.vm.status.notready')
                        "
                      />({{ item.state }})
                    </a-tooltip>

                    <br />
                    <WindowsFilled /> {{ item.ostype }}
                  </a-card>
                </a-list-item>
              </template>
            </a-list>
          </a-card> -->
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
      pageSize: 10,
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
          handshake_status: "Ready",
          ostype: "Windows 10 (64bit)",
          favorite: true,
        },
        {
          id: "11111111-111111-1111111-1-2222",
          name: "vm62",
          state: "Stopped",
          handshake_status: "Not Ready",
          ostype: "Windows 10 (64bit)",
          favorite: false,
        },
        {
          id: "11111111-111111-1111111-1-3333",
          name: "vm36",
          state: "Stopped",
          handshake_status: "Not Ready",
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
      ],
      dataList: [
        {
          workspace: "workspace 1",
          instanceList: [
            {
              id: "11111111-111111-1111111-1-1111",
              name: "vm18898",
              state: "Running",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: true,
            },
            {
              id: "11111111-111111-1111111-1-2222",
              name: "vm62",
              state: "Stopped",
              handshake_status: "Not Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
            {
              id: "11111111-111111-1111111-1-3333",
              name: "vm36",
              state: "Stopped",
              handshake_status: "Not Ready",
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
          ],
        },
        {
          workspace: "workspace 2",
          instanceList: [
            {
              id: "11111111-111441-1111111-1-1231",
              name: "vm10",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
            {
              id: "11111111-666666-1111111-1-1231",
              name: "vm10",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
            {
              id: "777777-111111-1111111-1-1231",
              name: "vm10",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
          ],
        },
        {
          workspace: "workspace 3",
          instanceList: [
            {
              id: "11111111-111441-1111111-1-1231",
              name: "vm10",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
            {
              id: "11111111-666666-1111111-1-1231",
              name: "vm10",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
            {
              id: "777777-111111-1111111-1-1231",
              name: "vm10",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
          ],
        },
        {
          workspace: "workspace 4",
          instanceList: [
            {
              id: "11111111-111441-1111111-1-1231",
              name: "vm10",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
            {
              id: "11111111-666666-1111111-1-1231",
              name: "vm10",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
            {
              id: "777777-111111-1111111-1-1231",
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
      console.log(uuid + " :: " + bool);
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
  padding: 20px;
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
</style>
