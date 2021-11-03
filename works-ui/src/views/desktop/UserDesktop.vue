<template>
  <div v-for="workspace in dataList" style="width: 100%; padding: 5px">
    <a-card
      :title="'WORKSPACE : ' + workspace.workspace"
      bodyStyle="margin: 1px; fontSize: 100px;"
      :loading="loading"
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
                    @confirm="favoriteHandle(vm.id, !vm.favorite)"
                  >
                    <a-tooltip placement="bottom">
                      <template #title>{{ $t("즐겨찾기") }}</template>

                      <StarFilled
                        v-if="vm.favorite"
                        :id="vm.id + '-TRUE'"
                        :style="{ color: '#fcf00a' }"
                        @click="favoriteHandle(vm.favorite, vm.id)"
                      />
                      <StarOutlined
                        v-else
                        :id="vm.id + '-FALSE'"
                        @click="favoriteHandle(vm.favorite, vm.id)"
                      />
                    </a-tooltip>
                  </a-popconfirm>


                  <a-popconfirm
                    :title="'RDP 파일을 다울로드 하시겠습니까?'"
                    :ok-text="$t('label.ok')"
                    :cancel-text="$t('label.cancel')"
                    @confirm="favoriteHandle(vm.id, !vm.favorite)"
                  >
                    <a-tooltip placement="bottom">
                      <template #title>{{ $t("RDP 다운로드") }}</template>
                      <CloudDownloadOutlined />
                    </a-tooltip>
                  </a-popconfirm>
                  
                  <a-popconfirm
                    :title="'데스크톱에 접속하시겠습니까?'"
                    :ok-text="$t('label.ok')"
                    :cancel-text="$t('label.cancel')"
                    @confirm="favoriteHandle(vm.id, !vm.favorite)"
                    :disabled="vm.handshake_status == 'Ready' ? false : true"
                  >
                    <a-tooltip placement="bottom">
                      <template #title>{{  $t("데스크톱 접속") }}</template>
                      <CodeFilled />
                    </a-tooltip>
                  </a-popconfirm>

                  <a-popconfirm
                    :title="vm.state == 'Running' ? '가상머신을 종료하시겠습니까?' : '가상머신을 시작하시겠습니까?'"
                    :ok-text="$t('label.ok')"
                    :cancel-text="$t('label.cancel')"
                    @confirm="favoriteHandle(vm.id, !vm.favorite)"
                  >
                    <a-tooltip placement="bottom">
                      <template #title>{{  vm.state == 'Running' ? $t("가상머신 종료") : $t("가상머신 시작")}}</template>
                      <CaretRightOutlined
                        v-if="vm.state == 'Running' ? true : false"
                      />
                      <PoweroffOutlined v-else />
                    </a-tooltip>
                  </a-popconfirm>

                </template>
                <br /><br />

                <a-tooltip placement="bottom">
                  <template #title>{{ vm.handshake_status }}</template>
                  <a-badge
                    class="head-example"
                    :color="vm.handshake_status == 'Ready' ? 'green' : 'red'"
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
</template>

<script>
import { defineComponent, ref, reactive } from "vue";
import Actions from "@/components/Actions";

export default defineComponent({
  name: "UserDesktop",
  components: {
    Actions,
  },
  setup() {
    return {
      //desktopIconStyle: ref([ {"fontSize": '40px', "color": '#cc0000'} ]),
      favPopTitle: ref(""),
      actionFrom: "VirtualMachineList",
      loading: ref(true),
      dataList: [
        {
          workspace: "워크스페이스 1111111111111",
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
            {
              id: "11111111-111111-1111111-1-asdi",
              name: "vm11",
              state: "Running",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: true,
            },
            {
              id: "11111111-111111-1111111-1-asf3",
              name: "vm125",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: true,
            },
            {
              id: "11111111-111111-1111111-1-33g3",
              name: "vm513",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
            {
              id: "11111111-111111-1111111-1-r5h4",
              name: "vm55",
              state: "Running",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
            {
              id: "11111111-111111-1111111-1-hj5j",
              name: "vm52",
              state: "Stopped",
              handshake_status: "Ready",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
          ],
        },
        {
          workspace: "워크스페이스 22222222222222222",
          instanceList: [
            {
              id: "11111111-111111-1111111-1-123123",
              name: "vm10",
              state: "Stopped",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
            {
              id: "11111111-111111-1111111-1-asdfasdf",
              name: "vm11",
              state: "Running",
              ostype: "Windows 10 (64bit)",
              favorite: true,
            },
            {
              id: "11111111-111111-1111111-1-asdfasdf3",
              name: "vm125",
              state: "Stopped",
              ostype: "Windows 10 (64bit)",
              favorite: true,
            },
            {
              id: "11111111-111111-1111111-1-33g3g3",
              name: "vm513",
              state: "Stopped",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
            {
              id: "11111111-111111-1111111-1-sdhr5h4",
              name: "vm55",
              state: "Running",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
            {
              id: "11111111-111111-1111111-1-hj565hj5j",
              name: "vm52",
              state: "Stopped",
              ostype: "Windows 10 (64bit)",
              favorite: false,
            },
          ],
        },
      ],
    };
  },
  created() {
    setTimeout(() => {
      this.loading = false;
      this.fetchData();
    }, 500);
  },
  methods: {
    fetchData() {},
    favoriteHandle(uuid, bool) {
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
</style>
