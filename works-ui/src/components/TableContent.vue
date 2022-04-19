<template>
  <a-button
    v-if="state.addButtonBoolean"
    type="dashed"
    block
    style="padding-bottom: 10px"
    :disabled="state.statusReadyBool"
    @click="changeModal(state.callTapName, true)"
  >
    <PlusOutlined />
    {{ state.addButtontext }}
  </a-button>

  <a-alert
    v-if="state.statusReadyBool"
    :message="
      $t('message.workspace.test.status') + ' : ' + state.statusReadyMessage
    "
    :description="state.statusReadyDescription"
    type="info"
    show-icon
  />
  <div
    style="padding-left: 20px; padding-top: 1px; text-align: left; height: 33px"
  >
    <Actions
      v-if="vis"
      :action-from="actionFrom"
      :multi-select-list="state.selectedRows"
      :ws-name="workspaceInfo.name"
      @fetchData="parentRefresh"
    />
  </div>
  <a-table
    :loading="loading"
    size="small"
    class="ant-table-striped"
    style="margin-left: 5px"
    :columns="columns"
    :data-source="dataList"
    :pagination="pagination"
    :row-selection="{
      selectedRowKeys: state.selectedRowKeys,
      onChange: onSelectChange,
    }"
    >{{ actionFrom }}
    <template #bodyCell="{ column, text, record }">
      <template v-if="column.dataIndex === 'name'">
        <span
          v-if="actionFrom !== undefined && actionFrom === 'VirtualMachineList'"
        >
          <router-link :to="{ path: '/virtualMachineDetail/' + record.uuid }">
            {{ record.name }}
          </router-link>
        </span>
        <span
          v-else-if="
            actionFrom !== undefined && actionFrom === 'WorkspaceUserList'
          "
        >
          <router-link :to="{ path: '/accountDetail/' + record.name }">
            {{ record.name }}
          </router-link>
        </span>
        <span v-else>
          {{ record.name }}
        </span>
      </template>

      <template v-if="column.dataIndex === 'action'">
        <a-Popover v-if="actionRef === 'VirtualMachineList'" placement="bottom">
          <template #content>
            <ASpace direction="horizontal">
              <Actions
                :action-from="actionFrom"
                :vm-info="record"
                :ws-name="workspaceInfo.name"
                @fetchData="parentRefresh"
              />
            </ASpace>
          </template>
          <MoreOutlined />
        </a-Popover>

        <a-popconfirm
          v-if="actionFrom === 'WorkspaceUserList'"
          :title="$t('message.delete.confirm')"
          :ok-text="$t('label.ok')"
          :cancel-text="$t('label.cancel')"
          @confirm="delUserFromWorkspace(record.name)"
        >
          <a-tooltip placement="bottom">
            <template #title>{{ $t("tooltip.destroy") }}</template>
            <a-button type="primary" shape="circle" danger>
              <DeleteFilled />
            </a-button>
          </a-tooltip>
        </a-popconfirm>
      </template>

      <template v-if="column.dataIndex === 'owner_account_id'">
        <router-link
          :to="{ path: '/accountDetail/' + record.owner_account_id }"
        >
          {{ record.owner_account_id }}
        </router-link>
      </template>
      <template v-if="column.dataIndex === 'mold_status'">
        <a-badge
          class="head-example"
          :color="
            record.mold_status === 'Running'
              ? 'green'
              : record.mold_status === 'Stopping' ||
                record.mold_status === 'Starting'
              ? 'blue'
              : record.mold_status === 'Stopped'
              ? 'red'
              : ''
          "
          :text="
            record.mold_status === 'Running'
              ? $t('label.vm.status.running')
              : record.mold_status === 'Starting'
              ? $t('label.vm.status.starting')
              : record.mold_status === 'Stopping'
              ? $t('label.vm.status.stopping')
              : record.mold_status === 'Stopped'
              ? $t('label.vm.status.stopped')
              : ''
          "
        />
      </template>

      <template v-if="column.dataIndex === 'status'">
        <a-tooltip placement="bottom">
          <template #title>{{ record.handshake_status }}</template>
          <a-badge
            class="head-example"
            :color="
              record.handshake_status === 'Not Ready' ||
              record.handshake_status === 'Pending'
                ? 'red'
                : record.handshake_status === 'Joining' ||
                  record.handshake_status === 'Joined'
                ? 'yellow'
                : record.handshake_status === 'Ready'
                ? 'green'
                : 'red'
            "
            :text="
              record.handshake_status === 'Not Ready' ||
              record.handshake_status === 'Pending'
                ? $t('label.vm.status.initializing') +
                  '(' +
                  record.handshake_status +
                  ')'
                : record.handshake_status === 'Joining' ||
                  record.handshake_status === 'Joined'
                ? $t('label.vm.status.configuring') +
                  '(' +
                  record.handshake_status +
                  ')'
                : record.handshake_status === 'Ready'
                ? $t('label.vm.status.ready')
                : $t('label.vm.status.notready')
            "
          />
        </a-tooltip>
      </template>
      <!-- <template #userRender="{ record }">
        <router-link
          :to="{ path: '/accountDetail/' + record.owner_account_id }"
        >
          {{ record.owner_account_id }}
        </router-link>
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
          :color="
            record.mold_status === 'Running'
              ? 'green'
              : record.mold_status === 'Stopping' ||
                record.mold_status === 'Starting'
              ? 'blue'
              : record.mold_status === 'Stopped'
              ? 'red'
              : ''
          "
          :text="
            record.mold_status === 'Running'
              ? $t('label.vm.status.running')
              : record.mold_status === 'Starting'
              ? $t('label.vm.status.starting')
              : record.mold_status === 'Stopping'
              ? $t('label.vm.status.stopping')
              : record.mold_status === 'Stopped'
              ? $t('label.vm.status.stopped')
              : ''
          "
        />
      </template>
      <template #vmReadyStateRender="{ record }">
        <a-tooltip placement="bottom">
          <template #title>{{ record.handshake_status }}</template>
          <a-badge
            class="head-example"
            :color="
              record.handshake_status === 'Not Ready' ||
              record.handshake_status === 'Pending'
                ? 'red'
                : record.handshake_status === 'Joining' ||
                  record.handshake_status === 'Joined'
                ? 'yellow'
                : record.handshake_status === 'Ready'
                ? 'green'
                : 'red'
            "
            :text="
              record.handshake_status === 'Not Ready' ||
              record.handshake_status === 'Pending'
                ? $t('label.vm.status.initializing') +
                  '(' +
                  record.handshake_status +
                  ')'
                : record.handshake_status === 'Joining' ||
                  record.handshake_status === 'Joined'
                ? $t('label.vm.status.configuring') +
                  '(' +
                  record.handshake_status +
                  ')'
                : record.handshake_status === 'Ready'
                ? $t('label.vm.status.ready')
                : $t('label.vm.status.notready')
            "
          />
        </a-tooltip>
      </template>
      
      <template #accountNameRender="{ record }">
        <router-link :to="{ path: '/accountDetail/' + record.name }">{{
          record.name
        }}</router-link>
      </template>
      <template #deleteRender="{ record }">
        <a-popconfirm
          :title="$t('message.delete.confirm')"
          :ok-text="$t('label.ok')"
          :cancel-text="$t('label.cancel')"
          @confirm="delUserFromWorkspace(record.name)"
        >
          <a-tooltip placement="bottom">
            <template #title>{{ $t("tooltip.destroy") }}</template>
            <a-button type="primary" shape="circle" danger>
              <DeleteFilled />
            </a-button>
          </a-tooltip>
        </a-popconfirm>
      </template> -->
    </template>
  </a-table>
  <a-modal
    v-model:title="state.addButtontext"
    v-model:visible="state.addVmModalBoolean"
    width="500px"
    :ok-text="$t('label.ok')"
    :cancel-text="$t('label.cancel')"
    @ok="putVmToWorksapce"
  >
    <p>
      {{
        $t("modal.confirm.workspace.add.vm", {
          name: workspaceName,
        })
      }}
    </p>
    <a-input-number
      id="inputNumber"
      v-model:value="selectedVmQuantity"
      style="width: 100%; margin-top: 7px"
      :title="'Desktop Quantity'"
      :min="1"
    />
  </a-modal>

  <a-modal
    v-model:title="state.addButtontext"
    v-model:visible="state.addUserModalBoolean"
    width="500px"
    :ok-text="$t('label.ok')"
    :cancel-text="$t('label.cancel')"
    @ok="putUserToWorksapce"
  >
    <p>
      {{
        $t("modal.confirm.workspace.add.user", {
          name: workspaceName,
        })
      }}
    </p>
    <a-select
      v-model:value="selectedUser"
      :placeholder="$t('tooltip.user')"
      mode="multiple"
      style="width: 98%"
      option-filter-prop="label"
      class="addmodal-aform-item-div"
      :options="
        filteredOptions.map((item) => ({
          value: item.name,
          label:
            item.name === 'Guest'
              ? '👤&nbsp;&nbsp;&nbsp;&nbsp;' + item.name
              : '👤&nbsp;&nbsp;&nbsp;&nbsp;' + item.name,
        }))
      "
    >
      <!--👥 <a-select-option
        v-for="option in addAbleUserList"
        :key="option.name"
        :value="option.name"
        :label="option.name"
        >
        {{ option.name }}
      </a-select-option> -->
    </a-select>
  </a-modal>
</template>

<script>
import { defineComponent, reactive, ref, computed } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";
import Actions from "../components/Actions";

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
    workspaceInfo: {
      type: Object,
      required: false,
      default: null,
    },
    networkList: {
      type: Object,
      required: false,
      default: null,
    },
    vmDiskInfo: {
      type: Object,
      required: false,
      default: null,
    },
    vmList: {
      type: Object,
      required: false,
      default: null,
    },
    groupInfo: {
      type: Object,
      required: false,
      default: null,
    },
    workspacePolicyList: {
      type: Object,
      required: false,
      default: null,
    },
  },
  emits: ["parentRefresh"],
  setup(props) {
    const state = reactive({
      addVmModalBoolean: ref(false),
      addUserModalBoolean: ref(false),
      callTapName: ref(props.tapName),
      addButtonBoolean: ref(false),
      addButtontext: ref(""),
      statusReadyBool: ref(false),
      statusReadyMessage: ref(""),
      statusReadyDescription: ref(""),
      modalConfirm: ref(""),
      modalTitle: ref(""),
      selectedRowKeys: [],
      selectedRows: [],
    });
    return {
      state,
    };
  },
  data(props) {
    return {
      actionRef: ref(""),
      workspaceInfo: ref(props.workspaceInfo),
      workspaceName: ref(""),
      networkList: ref(props.networkList),
      vmDiskInfo: ref(props.vmDiskInfo),
      vmList: ref(props.vmList),
      groupInfo: ref(props.groupInfo),
      dataList: ref([]),
      loading: ref(false),
      addAbleUserList: ref([]),
      selectedUser: ref([]),
      selectedVmQuantity: ref(1),
      columns: ref([]),
      deleteUser: ref(""),
      vis: ref(false),
      succCnt: ref(0),
      failCnt: ref(0),
      pagination: {
        // pageSize: 10,
        showSizeChanger: true, // display can change the number of pages per page
        pageSizeOptions: ["10", "20", "50", "100", "200"], // number of pages per option
        showTotal: (total) =>
          this.$t("label.total") + ` ${total}` + this.$t("label.items"), // show total
        // showSizeChange: (current, pageSize) => (this.pageSize = pageSize), // update display when changing the number of pages per page
      },
      rowSelection: ref(null),
      filteredOptions: computed(() =>
        this.addAbleUserList.filter((o) => !this.selectedUser.includes(o.name))
      ),
    };
  },
  created() {
    this.fetchRefresh();
    if (this.workspaceInfo == null) {
      this.workspaceName = "";
    } else {
      this.workspaceName = this.workspaceInfo.name;
    }
  },
  methods: {
    parentRefresh() {
      this.$emit("parentRefresh");
    },
    changeModal(target, value) {
      if (target == "desktop") {
        this.state.addVmModalBoolean = value;
      } else if (target == "user") {
        this.state.addUserModalBoolean = value;
      }
    },
    onSelectChange(selectedRowKeys, selectedRows) {
      console.log(selectedRowKeys, selectedRows);
      this.vis = false;
      this.state.selectedRowKeys = selectedRowKeys;

      if (selectedRows.length > 0) {
        setTimeout(() => {
          this.state.selectedRows = selectedRows;
          //console.log(this.state.selectedRows);
          this.vis = true;
        }, 100);
      } else {
        this.vis = false;
      }
    },
    fetchRefresh(refreshClick) {
      if (refreshClick) this.loading = true;
      else this.loading = false;
      this.actionRef = "";
      this.vis = false;
      this.state.selectedRowKeys = [];
      this.state.selectedRows = [];

      this.fetchData();
    },
    fetchData() {
      if (this.state.callTapName === "desktop") {
        this.state.addButtonBoolean = ref(true);
        this.fetchDesktop();
        let stat = this.workspaceInfo.template_ok_check;
        if (stat === "Ready") {
          this.state.statusReadyBool = false; //워크스페이스 Agent상태가 OK일때 데스크톱가상머신추가 활성화
        } else {
          this.state.statusReadyBool = true;
          if (stat === "Not Ready") {
            this.state.statusReadyMessage = this.$t(
              "message.workspace.test.vmcreate"
            );
          } else if (stat === "Pending") {
            this.state.statusReadyMessage = this.$t(
              "message.workspace.test.initializing"
            );
          } else if (stat === "Joining" || stat === "Joined") {
            this.state.statusReadyMessage = this.$t(
              "message.workspace.test.configuring"
            );
          }
          this.state.statusReadyDescription = this.$t("message.workspace.test");
        }
      } else if (this.state.callTapName === "user") {
        this.state.addButtonBoolean = ref(true);
        this.fetchUser();
      } else if (this.state.callTapName === "policy") {
        this.fetchPolicy();
        // this.rowSelection = null;
      } else if (this.state.callTapName === "network") {
        this.fetchNetwork();
        // this.rowSelection = null;
      } else if (this.state.callTapName === "datadisk") {
        this.fetchDatadisk();
        // this.rowSelection = null;
      }

      setTimeout(() => {
        this.loading = false;
      }, 500);
    },
    fetchDesktop() {
      this.state.addButtontext = this.$t("label.desktop.vm.add");
      this.columns = [
        {
          dataIndex: "name",
          key: "name",
          // slots: { customRender: "nameRender" },
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
          // slots: { customRender: "actionRender" },
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
          // slots: { customRender: "userRender" },
        },
        {
          title: this.$t("label.vm.state"),
          dataIndex: "mold_status",
          key: "mold_status",
          sorter: (a, b) =>
            a.mold_status < b.mold_status
              ? -1
              : a.mold_status > b.mold_status
              ? 1
              : 0,
          sortDirections: ["descend", "ascend"],
          // slots: { customRender: "vmStateRender" },
        },
        {
          title: this.$t("label.vm.ready.state"),
          dataIndex: "status",
          key: "status",
          sorter: (a, b) =>
            a.status < b.status ? -1 : a.status > b.status ? 1 : 0,
          sortDirections: ["descend", "ascend"],
          // slots: { customRender: "vmReadyStateRender" },
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

      if (this.vmList !== undefined && this.vmList !== null) {
        this.dataList = this.vmList;
        this.dataList.forEach((value, index, array) => {
          this.dataList[index].key = index;
        });
      }

      // tooltip action버튼 리로드 시 시간차가 필요함.
      setTimeout(() => {
        this.actionRef = "VirtualMachineList";
      }, 100);

      // worksApi
      //   .get("/api/v1/instance/" + this.$route.params.workspaceUuid)
      //   .then((response) => {
      //     if (response.status == 200) {
      //       this.dataList = response.data.result.instanceInfo;
      //       this.dataList.forEach((value, index, array) => {
      //         this.dataList[index].key = index;
      //       });
      //     } else {
      //       message.error(this.$t("message.response.data.fail"));
      //     }
      //   })
      //   .catch((error) => {
      //     message.destroy();
      //     message.error(this.$t("message.response.data.fail"));
      //     console.log(error);
      //   })
      //   .finally(() => {
      //     this.loading = false;
      //   });
    },
    async fetchUser() {
      this.state.addButtontext = this.$t("label.desktop.user.add");
      this.columns = [
        {
          dataIndex: "name",
          key: "name",
          title: this.$t("label.name"),
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
          // slots: { customRender: "accountNameRender" },
        },
        {
          title: this.$t("label.lastname"),
          dataIndex: "givenName",
          key: "givenName",
          width: "10%",
          sorter: (a, b) =>
            a.givenName < b.givenName ? -1 : a.givenName > b.givenName ? 1 : 0,
          // sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.firstname"),
          dataIndex: "sn",
          key: "sn",
          width: "10%",
          sorter: (a, b) => (a.sn < b.sn ? -1 : a.sn > b.sn ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.title"),
          dataIndex: "title",
          key: "title",
          width: "15%",
          sorter: (a, b) =>
            a.title < b.title ? -1 : a.title > b.title ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.department"),
          dataIndex: "department",
          key: "department",
          width: "10%",
          sorter: (a, b) =>
            a.department < b.department
              ? -1
              : a.department > b.department
              ? 1
              : 0,
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.phone"),
          dataIndex: "telephoneNumber",
          key: "telephoneNumber",
          width: "15%",
          sorter: (a, b) =>
            a.telephoneNumber < b.telephoneNumber
              ? -1
              : a.telephoneNumber > b.telephoneNumber
              ? 1
              : 0,
          sortDirections: ["descend", "ascend"],
        },
        {
          title: "",
          key: "action",
          dataIndex: "action",
          align: "right",
          width: "5%",
          // slots: { customRender: "deleteRender" },
        },
      ];
      var userInWorkspaceList = [];
      if (
        this.groupInfo.member !== undefined &&
        this.groupInfo.member !== null
      ) {
        for (let str of this.groupInfo.member) {
          userInWorkspaceList.push({ name: str.split(",")[0].split("CN=")[1] });
        }
      }
      // 워크스페이스에 추가 할 사용자 목록 조회
      await worksApi
        .get("/api/v1/user")
        .then((response) => {
          if (response.status == 200) {
            if (
              response.data.result !== null &&
              response.data.result.length !== 0 &&
              response.data.result !== undefined
            ) {
              this.addAbleUserList = response.data.result.filter(function (o1) {
                return !userInWorkspaceList.some(function (o2) {
                  return o1.name == o2.name;
                });
              });
              userInWorkspaceList = response.data.result.filter(function (o1) {
                return userInWorkspaceList.some(function (o2) {
                  return o1.name == o2.name;
                });
              });
            } else {
              this.addAbleUserList = [];
            }
          } else {
            message.error(this.$t("message.response.data.fail"));
          }
        })
        .catch((error) => {
          message.destroy();
          message.error(this.$t("message.response.data.fail"));
          console.log(error);
        });

      this.dataList = userInWorkspaceList;
      this.dataList.forEach((value, index, array) => {
        this.dataList[index].key = index;
      });
    },
    fetchPolicy() {
      this.columns = [
        {
          dataIndex: "name",
          key: "name",
          // slots: { customRender: "nameRender" },
          title: this.$t("label.name"),
          sorter: (a, b) => (a.name < b.name ? -1 : a.name > b.name ? 1 : 0),
          sortDirections: ["descend", "ascend"],
        },
        {
          title: this.$t("label.description"),
          dataIndex: "description",
          key: "description",
          width: "60%",
          sorter: (a, b) =>
            a.description < b.description
              ? -1
              : a.description > b.description
              ? 1
              : 0,
          sortDirections: ["descend", "ascend"],
        },
      ];
      if (
        this.workspacePolicyList !== undefined &&
        this.workspacePolicyList !== null
      ) {
        this.dataList = this.workspacePolicyList;
        this.dataList.forEach((value, index, array) => {
          this.dataList[index].key = index;
        });
      }
    },
    fetchNetwork() {
      this.columns = [
        {
          dataIndex: "name",
          key: "name",
          // slots: { customRender: "nameRender" },
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
          // slots: { customRender: "stateRender" },
        },
      ];
      // 워크스페이스 네트워크 조회
      if (this.networkList !== undefined && this.networkList !== null)
        this.dataList = this.networkList;
    },
    fetchDatadisk() {
      this.columns = [
        {
          dataIndex: "name",
          key: "name",
          // slots: { customRender: "nameRender" },
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
          sorter: (a, b) => (a.type < b.type ? -1 : a.type > b.type ? 1 : 0),
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
      if (this.vmDiskInfo !== undefined && this.vmDiskInfo !== null)
        this.dataList = [this.vmDiskInfo];

      // worksApi
      //   .get("/api/v1/instance/detail/" + this.$route.params.vmUuid)
      //   .then((response) => {
      //     if (response.status === 200) {
      //       this.dataList =
      //         response.data.result.instanceInstanceVolumeInfo.volume;
      //     } else {
      //       message.error(this.$t("message.response.data.fail"));
      //     }
      //   })
      //   .catch((error) => {
      //     message.destroy();
      //     message.error(this.$t("message.response.data.fail"));
      //     console.log(error);
      //   })
      //   .finally(() => {
      //     this.loading = false;
      //   });
    },
    async funcDelay(delay) {
      return new Promise(function (resolve) {
        setTimeout(function () {
          resolve("delay call!");
        }, delay);
      });
    },
    funcEndMessage() {
      console.log(
        this.sucMessage,
        this.failMessage,
        this.succCnt,
        this.failCnt
      );
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
    async putVmToWorksapce() {
      //데스크톱 가상머신 개수선택해 추가
      message.loading(this.$t("message.workspace.vm.adding"), 100);
      let params = new URLSearchParams();
      params.append("uuid", this.$route.params.workspaceUuid);
      params.append("quantity", this.selectedVmQuantity);

      try {
        const response = await worksApi.put("/api/v1/instance", params);
        if (response.status == 200) {
          this.loading = ref(true);

          //생성 요청 후 20초 정도 딜레이 후 리프레시
          await this.funcDelay(20000);
          this.parentRefresh();

          //2초 후 메시지 표시
          await this.funcDelay(2000);
          this.changeModal("desktop", false);
          message.destroy();
          message.success(
            this.$t("message.workspace.vm.add.success", {
              count: this.selectedVmQuantity,
            }),
            5
          );
        } else {
          message.error(this.$t("message.workspace.vm.add.fail"));
        }
      } catch (error) {
        message.destroy();
        message.error(this.$t("message.workspace.vm.add.fail"));
        console.log(error);
      }
    },
    async putUserToWorksapce() {
      //워크스페이스에 사용자 추가
      //console.log(this.selectedUser);
      if (this.selectedUser.length == 0) return false;

      this.sucMessage = "message.workspace.user.add.ok";
      this.failMessage = "message.workspace.user.add.dupl";
      message.loading(this.$t("message.workspace.vm.user.allocateing"), 100);

      for (let val of this.selectedUser) {
        // console.log(val);
        try {
          const response = await worksApi.put(
            "/api/v1/group/" + this.workspaceInfo.name + "/" + val
          );

          if (response.status == 200) {
            this.succCnt = this.succCnt + 1;
          }
        } catch (error) {
          this.failCnt = this.failCnt + 1;
          console.log(error);
        }
      }
      //console.log(this.succCnt, this.failCnt);
      this.parentRefresh();
      await this.funcDelay(2000);
      this.changeModal("user", false);
      this.funcEndMessage();
      this.selectedUser = [];
    },
    async delUserFromWorkspace(val) {
      message.loading(this.$t("message.workspace.vm.user.deleting"), 100);

      try {
        const response = await worksApi.delete(
          "/api/v1/group/" + this.workspaceInfo.name + "/" + val
        );
        if (response.status === 200) {
          this.parentRefresh();

          setTimeout(() => {
            message.destroy();
            message.success(
              this.$t("message.workspace.user.delete.ok", { count: 1 }),
              5
            );
          }, 2000);
        }
      } catch (error) {
        message.destroy();
        message.error(
          this.$t("message.workspace.user.delete.fail", { count: 1 }),
          5
        );
        console.log(error);
      }
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

/* For demo */
.ant-carousel :deep(.slick-slide) {
  text-align: center;
  height: 100px;
  line-height: 100px;
  background: #c5e7fa;
  overflow: hidden;
}

.ant-carousel :deep(.slick-slide h3) {
  color: #3d3d3d;
}
</style>
