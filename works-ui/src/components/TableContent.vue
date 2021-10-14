<template>
  <a-button
    v-if="state.addButtonBoolean"
    type="dashed"
    block
    style="margin-bottom: 14px"
    @click="changeModal(state.callTapName, true)"
    :disabled="buttonDisable"
  >
    <PlusOutlined />
    {{ state.addButtontext }}
  </a-button>

  <a-table
    :loading="loading"
    size="small"
    class="ant-table-striped"
    style="overflow: auto; margin-left: 5px"
    :columns="columns"
    :data-source="dataList"
    :pagination="pagination"
  >
    <template #nameRender="{ record }">
      <span v-if="actionFrom !== undefined && actionFrom === 'VirtualMachineList'">
        <router-link :to="{ path: '/virtualMachineDetail/' + record.uuid }">{{
          record.name
        }}</router-link>
      </span>
      <span v-else>
        {{ record.name }}
      </span>
    </template>

    <template #actionRender="{ record }">
      <a-Popover placement="bottom">
        <template #content>
          <ASpace direction="horizontal">
            <Actions
              v-if="actionFrom !== undefined && actionFrom === 'VirtualMachineList'"
              :action-from="actionFrom"
              :vm-uuid="record.uuid"
              :workspace="record.workspace_name"
              :allocate-status="record.owner_account_id"
              :status="record.status"
              @fetchData="fetchData"
            />
            <Actions v-else :action-from="actionFrom" />
          </ASpace>
        </template>
        <MoreOutlined />
      </a-Popover>
    </template>
    <template #userRender="{ record }">
      <!-- {{
        record.owner_account_id == ""
          ? $t("label.owner.account.false")
          : record.owner_account_id
      }} -->
      {{record.owner_account_id}}
    </template>
    <template #connRender="{ record }">
      {{
        record.connected == false
          ? $t("label.connect.false")
          : $t("label.connect.true")
      }}
    </template>
    <template #vmStateRender="{ record }">
      <a-badge
        class="head-example"
        :color="record.mold_status == 'Running' ? 'green' : 'red'"
        :text="record.mold_status == 'Running' ? $t('label.vm.status.running') : $t('label.vm.status.stopped')"
      />
    </template>
    <template #vmReadyStateRender="{ record }">
      <a-badge
        class="head-example"
        :color="record.checked === true ? 'green' : 'red'"
        :text="record.checked ===true ? $t('label.vm.status.ready') : $t('label.vm.status.notready')"
      />
    </template>
    <template #stateRender="{ record }">
      {{ record.state }}
    </template>
    <template #deleteRender="{ record }">
      <a-popconfirm
        :title="$t('message.delete.confirm')"
        :ok-text="$t('label.ok')"
        :cancel-text="$t('label.cancel')"
        @confirm="onUserDelete(record.name)"
      >
        <a-tooltip placement="bottom">
          <template #title>{{ $t("tooltip.destroy") }}</template>
          <a-button type="primary" shape="circle" danger>
            <DeleteFilled />
          </a-button>
        </a-tooltip>
      </a-popconfirm>
    </template>
  </a-table>

  <a-modal
    v-model:title="state.addButtontext"
    v-model:visible="state.addVmModalBoolean"
    width="400px"
    :ok-text="$t('label.ok')"
    :cancel-text="$t('label.cancel')"
    @ok="putVmToWorksapce"
  >
    <a-input-number
      id="inputNumber"
      v-model:value="selectedVmQuantity"
      style="width: 100%; margin-top: 7px"
      :title="'Desktop Quantity'"
      :min="1"
      :max="10"
    />
  </a-modal>

  <a-modal
    v-model:title="state.addButtontext"
    v-model:visible="state.addUserModalBoolean"
    width="400px"
    :ok-text="$t('label.ok')"
    :cancel-text="$t('label.cancel')"
    @ok="putUserToWorksapce"
  >
    <a-select
      v-model:value="selectedUser"
      :placeholder="$t('tooltip.user')"
      show-search
      style="width: 100%; margin-top: 7px"
      option-filter-prop="label"
      class="addmodal-aform-item-div"
    >
      <a-select-option
        v-for="option in userDataList"
        :key="option.name"
        :value="option.name"
        :label="option.name"
      >
        {{ option.name }}
      </a-select-option>
    </a-select>
  </a-modal>
</template>

<script>
import { defineComponent, reactive, ref } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";
import Actions from "../components/Actions";

const rowSelection = {
  onChange: (selectedRowKeys, selectedRows) => {
    // console.log(
    //   `selectedRowKeys: ${selectedRowKeys}`,
    //   "selectedRows: ",
    //   selectedRows
    // );
  },
  onSelect: (record, selected, selectedRows) => {
    //console.log(record, selected, selectedRows);
  },
  onSelectAll: (selected, selectedRows, changeRows) => {
    //console.log(selected, selectedRows, changeRows);
  },
};
export default defineComponent({
  components: {
    Actions,
  },
  props: {
    tapName: {
      type: String,
      required: true,
      default: "",
    },
    bordered: {
      type: Boolean,
      required: false,
      default: false,
    },
    actionFrom: {
      type: String,
      required: false,
      default: "",
    },
  },
  setup(props) {
    const state = reactive({
      addVmModalBoolean: ref(false),
      addUserModalBoolean: ref(false),
      callTapName: ref(props.tapName),
      addButtonBoolean: ref(false),
      addButtontext: ref(""),
      modalConfirm: ref(""),
      modalTitle: ref(""),
      confirmModalView: ref(false),
      userAllocateVmModalBoolean: ref(false),
    });
    return {
      rowSelection,
      state,
      pagination: {
        pageSize: 10,
        showSizeChanger: true, // display can change the number of pages per page
        pageSizeOptions: ["10", "20", "30", "50"], // number of pages per option
        showTotal: (total) => `Total ${total} items`, // show total
        showSizeChange: (current, pageSize) => (this.pageSize = pageSize), // update display when changing the number of pages per page
      },
      //actionFrom: ref(props.actionFrom),
    };
  },
  data() {
    return {
      dataList: ref([]),
      loading: ref(false),
      workspaceUserList: ref([]),
      userDataList: ref([]),
      selectedUser: ref(""),
      selectedVmQuantity: ref(1),
      columns: ref([]),
      confirmModalView: ref(false),
      deleteUser: ref(""),
      timer: ref(null),
      buttonDisable: ref(false),
      workspaceName: ref(""),
    };
  },
  created() {
    this.fetchData();
    this.timer = setInterval(() => { //30초 자동 갱신
      this.fetchData();
    }, 30000);
  },
  unmounted() {
    clearInterval(this.timer);
  },
  methods: {
    changeModal(target, value) {
      if (target == "desktop") {
        this.state.addVmModalBoolean = ref(value);
      } else if (target == "user") {
        this.state.addUserModalBoolean = ref(value);
      }
    },
    fetchData() {
      //console.log(this.state.callTapName);
      this.dataList = []; //초기화
      this.loading = ref(true);
      if (this.state.callTapName === "desktop" || this.state.callTapName === "user") {
        worksApi
        .get("/api/v1/workspace/" + this.$route.params.workspaceUuid)
        .then((response) => {
          if (response.status == 200) {
            this.workspaceName = response.data.result.workspaceInfo.name;
            if(response.data.result.workspaceInfo.template_ok_check !== "AgentOK") {
              this.buttonDisable = true;//워크스페이스 Agent상태가 OK일때 데스크톱가상머신추가 버튼 활성화
            }
            if(this.state.callTapName === "desktop"){
              this.fetchDesktop();
            }else if(this.state.callTapName === "user"){
              this.fetchUser();
            }
          }
        })
        .catch(function (error) {
          console.log(error);
        });
      } else if (this.state.callTapName === "policy") {
        this.fetchPolicy();
      } else if (this.state.callTapName === "network") {
        this.fetchNetwork();
      } else if (this.state.callTapName === "datadisk") {
        this.fetchDatadisk();
      }
      setTimeout(() => {
        this.loading = ref(false);
      }, 200);
    },

    fetchDesktop() {
      this.state.addButtonBoolean = ref(true);
      this.state.addButtontext = this.$t("label.desktop.vm.add");
      this.columns = [
        {
          dataIndex: "name",
          key: "name",
          slots: { customRender: "nameRender" },
          title: this.$t("label.name"),
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: "",
          key: "action",
          dataIndex: "action",
          align: "right",
          width: "5%",
          slots: { customRender: "actionRender" },
        },
        {
          title: this.$t("label.users"),
          dataIndex: "owner_account_id",
          key: "owner_account_id",
          sorter: (a, b) =>
            a.owner_account_id < b.owner_account_id
              ? -1
              : a.owner_account_id > b.owner_account_id
              ? 1
              : 0,
          sortDirections: ["descend", "ascend"],
          slots: { customRender: "userRender" },
        },
        {
          title: this.$t("label.vm.state"),
          dataIndex: "mold_status",
          key: "mold_status",
          sorter: (a, b) =>
            a.mold_status < b.mold_status ? -1 : a.mold_status > b.mold_status ? 1 : 0,
          sortDirections: ["descend", "ascend"],
          slots: { customRender: "vmStateRender" },
        },
        {
          title: this.$t("label.vm.ready.state"),
          dataIndex: "status",
          key: "status",
          sorter: (a, b) =>
            a.status < b.status ? -1 : a.status > b.status ? 1 : 0,
          sortDirections: ["descend", "ascend"],
          slots: { customRender: "vmReadyStateRender" },
        },

        // {
        //   title: this.$t("label.desktop.connect.boolean"),
        //   dataIndex: "connected",
        //   key: "connected",
        //   sorter: (a, b) =>
        //     a.connected < b.connected
        //       ? -1
        //       : a.connected > b.connected
        //       ? 1
        //       : 0,
        //   sortDirections: ["descend", "ascend"],
        //   slots: { customRender: "connRender" },
        // },
      ];

      worksApi
        .get("/api/v1/instance/" + this.$route.params.workspaceUuid)
        .then((response) => {
          if (response.status == 200) {
            this.dataList = response.data.result.instanceInfo;
          } else {
            //message.error(this.$t("message.response.data.fail"));
            //console.log("데이터를 정상적으로 가져오지 못했습니다.");
          }
        })
        .catch(function (error) {
          console.log(error);
        });
    },
    fetchUser() {
      this.state.addButtonBoolean = ref(true);
      this.state.addButtontext = this.$t("label.desktop.user.add");

      this.columns = [
        {
          dataIndex: "name",
          key: "name",
          slots: { customRender: "nameRender" },
          title: this.$t("label.name"),
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: "",
          key: "action",
          dataIndex: "action",
          align: "right",
          width: "5%",
          slots: { customRender: "deleteRender" },
        },
      ];

      //해당 워크스페이스에 추가 된 사용자 목록 조회
      worksApi
        .get("/api/v1/group/" + this.workspaceName)
        .then((response) => {
          if (response.status == 200) {
            //console.log(response.data.result.member);
            const temp =
              response.data.result.member == undefined
                ? ""
                : response.data.result.member;
            for (let str of temp) {
              this.dataList.push({ name: str.split(",")[0].split("CN=")[1] });
            }
          } else {
            //message.error(this.t("message.response.data.fail"));
            //console.log(response.message);
          }
        })
        .catch(function (error) {
          message.error(error);
          //console.log(error);
        });

      // 워크스페이스에 추가 할 사용자 목록 조회
      worksApi
        .get("/api/v1/user")
        .then((response) => {
          if (response.data.result.status == 200) {
            this.userDataList = response.data.result.result;
          } else {
            //message.error(this.t("message.response.data.fail"));
            //console.log(response.message);
          }
        })
        .catch(function (error) {
          message.error(error);
          //console.log(error);
        });
    },
    fetchPolicy() {
      this.columns = [
        {
          dataIndex: "name",
          key: "name",
          slots: { customRender: "nameRender" },
          title: this.$t("label.name"),
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        // {
        //   title: "",
        //   key: "action",
        //   dataIndex: "action",
        //   align: "right",
        //   width: "5px",
        //   slots: { customRender: "actionRender" },
        // },
        {
          title: this.$t("label.state"),
          dataIndex: "state",
          key: "state",
          sorter: (a, b) =>
            a.state < b.state ? -1 : a.state > b.state ? 1 : 0,
          sortDirections: ["descend", "ascend"],
          slots: { customRender: "stateRender" },
        },
      ];
      // 워크스페이스 정책 리스트 조회
      worksApi
        .get("/api/v1/workspace/" + this.$route.params.workspaceUuid)
        .then((response) => {
          if (response.status == 200) {
            //console.log(response.data.result.networkInfo.listnetworksresponse.network[0]);
            this.dataList = null;
          } else {
            //message.error(this.$t("message.response.data.fail"));
            //console.log("데이터를 정상적으로 가져오지 못했습니다.");
          }
        })
        .catch(function (error) {
          console.log(error);
        });
    },
    fetchNetwork() {
      this.columns = [
        {
          dataIndex: "name",
          key: "name",
          slots: { customRender: "nameRender" },
          title: this.$t("label.name"),
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        // {
        //   title: "",
        //   key: "action",
        //   dataIndex: "action",
        //   align: "right",
        //   width: "5px",
        //   slots: { customRender: "actionRender" },
        // },
        {
          title: this.$t("label.state"),
          dataIndex: "state",
          key: "state",
          sorter: (a, b) =>
            a.state < b.state ? -1 : a.state > b.state ? 1 : 0,
          sortDirections: ["descend", "ascend"],
          slots: { customRender: "stateRender" },
        },
      ];
      // 워크스페이스 네트워크 조회
      worksApi
        .get("/api/v1/workspace/" + this.$route.params.workspaceUuid)
        .then((response) => {
          if (response.status == 200) {
            this.dataList =
              response.data.result.networkInfo.network;
          } else {
            //message.error(this.$t("message.response.data.fail"));
            //console.log("데이터를 정상적으로 가져오지 못했습니다.");
          }
        })
        .catch(function (error) {
          console.log(error);
        });
    },
    fetchDatadisk() {
      this.columns = [
        {
          dataIndex: "name",
          key: "name",
          slots: { customRender: "nameRender" },
          title: this.$t("label.name"),
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        // {
        //   title: "",
        //   key: "action",
        //   dataIndex: "action",
        //   align: "right",
        //   width: "5px",
        //   slots: { customRender: "actionRender" },
        // },
        {
          title: this.$t("label.type"),
          dataIndex: "type",
          key: "type",
          sorter: (a, b) =>
            a.type < b.type ? -1 : a.type > b.type ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.size"),
          dataIndex: "sizegb",
          key: "sizegb",
          sorter: (a, b) =>
            a.sizegb < b.sizegb ? -1 : a.sizegb > b.sizegb ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
      ];
      // 가상머신 데이터디스크 조회
      worksApi
      .get("/api/v1/instance/detail/"+this.$route.params.vmUuid)
      .then((response) => {
        if (response.status === 200) {
          this.dataList = response.data.result.instanceInstanceVolumeInfo.volume;
        } else {
          //message.error(this.$t('message.response.data.fail'));
        }
      })
      .catch(function (error) {
        console.log(error);
      });
    },

    putVmToWorksapce() { //데스크톱 가상머신 개수선택해 추가
      message.loading(this.$t("message.workspace.vm.adding"), 20);
      let params = new URLSearchParams();
      params.append("uuid", this.$route.params.workspaceUuid);
      params.append("quantity", this.selectedVmQuantity);
      worksApi
        .put("/api/v1/instance", params)
        .then((response) => {
          if (response.status === 200) {
            this.loading = ref(true);
            setTimeout(() => {
              this.fetchData();
              message.destroy();
              message.success(this.$t("message.workspace.vm.add"), 1);
            }, 20000);
          } else {
            message.error(this.$t("message.workspace.vm.add.fail"));
          }
          this.changeModal("desktop", false);
        })
        .catch(function (error) {
          message.error(error);
        });
    },
    putUserToWorksapce() { //워크스페이스에 사용자 추가
      //console.log(this.selectedUser);
      if (!this.selectedUser) return false;
      worksApi
        .put("/api/v1/group/" + this.workspaceName + "/" + this.selectedUser)
        .then((response) => {
          if (response.status === 200) {
            message.success(this.$t("message.workspace.user.add"), 5);
            setTimeout(() => {
              this.fetchData();
            }, 1000);
          } else {
            message.error(this.$t("message.workspace.user.add.fail"));
          }
          this.changeModal("user", false);
        })
        .catch((error) => {
          message.error(this.$t("message.workspace.user.add.dupl"));
          //console.log(error);
        });
    },
    onUserDelete(val) {
      worksApi
        .delete("/api/v1/group/" + this.workspaceName + "/" + val)
        .then((response) => {
          if (response.status === 200) {
            message.success(this.$t("message.workspace.user.delete"), 1);
            setTimeout(() => {
              this.fetchData('desktop');
            }, 1000);
          } else {
            message.error(this.$t("message.workspace.user.delete.fail"));
          }
        })
        .catch(function (error) {
          message.error(error);
        });
    },
  },
});
</script>

<style scoped>
::v-deep(.ant-badge-status-dot) {
  width: 12px;
  height: 12px;
}

::v-deep(.ant-table-thead) {
  background-color: #f9f9f9;
}

::v-deep(.ant-table-small) > .ant-table-content > .ant-table-body {
  margin: 0;
}

::v-deep(.light-row) {
  background-color: #fff;
}

::v-deep(.dark-row) {
  background-color: #f9f9f9;
}
</style>

<style scoped lang="scss">
.shift-btns {
  display: flex;
}
.shift-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 20px;
  height: 20px;
  font-size: 12px;

  &:not(:last-child) {
    margin-right: 5px;
  }

  &--rotated {
    font-size: 10px;
    transform: rotate(90deg);
  }
}

.alert-notification-threshold {
  background-color: rgba(255, 231, 175, 0.75);
  color: #e87900;
  padding: 10%;
}

.alert-disable-threshold {
  background-color: rgba(255, 190, 190, 0.75);
  color: #f50000;
  padding: 10%;
}
</style>
