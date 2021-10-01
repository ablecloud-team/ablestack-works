<template>
  <div id="ContentTab">
    <a-tabs v-model:activeKey="activeKey" :tab-position="tabPosition">
      <a-tab-pane key="1" :tab="$t('label.users')">
        <a-button
          type="dashed"
          block
          style="margin-bottom: 14px"
          @click="changeModal(state.callTapName, true)"
          v-if="state.addButtonBoolean"
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
          :data-source="userDataList"
          :pagination="pagination"
        >
          <template #nameRender="{ record }">
            <span v-if="comp !== undefined && comp === 'VirtualMachineDetail'">
              <router-link
                :to="{ path: '/virtualMachineDetail/' + record.uuid }"
                >{{ record.name }}</router-link
              >
            </span>
            <span v-else>
              {{ record.name }}
            </span>
          </template>

          <template #actionRender>
            <a-Popover placement="top">
              <template #content>
                <ASpace direction="horizontal">
                  <Actions :actionFrom="actionFrom" />
                </ASpace>
              </template>
              <MoreOutlined />
            </a-Popover>
          </template>
          <template #userRender="{ record }">
            {{
              record.owner_account_id.String == ""
                ? "No"
                : record.owner_account_id.String
            }}
          </template>
          <template #connRender="{ record }">
            {{ record.connected == "true" ? "Connected" : "Disconnect" }}
          </template>
        </a-table>

        <a-modal
          v-model:title="state.addButtontext"
          v-model:visible="state.addVmModalBoolean"
          @ok="putVmToWorksapce"
          width="400px"
        >
          <a-input-number
            v-model:value="addDesktopQuantity"
            id="inputNumber"
            style="width: 100%; margin-top: 7px"
            :title="'Desktop Quantity'"
            :min="1"
            :max="10"
          />
        </a-modal>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script>
import { defineComponent, reactive, ref } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";
import Actions from "@/components/Actions";

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
  props: {
    tapName: String,
    //data: Object,
    //columns: Object,
    bordered: Boolean,
    comp: String,
    actionFrom: String,
  },
  setup(props) {
    //console.log("TableContent.vue props.tapName");
    //console.log(props.tapName);
    const state = reactive({
      addVmModalBoolean: ref(false),
      addUserModalBoolean: ref(false),
      callTapName: ref(props.tapName),
      addButtonBoolean: ref(false),
      addButtontext: ref(""),
      callModal: ref("desktop"),
    });
    const addDesktopQuantity = ref(1);
    return {
      rowSelection,
      state,
      addDesktopQuantity,
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
  created() {
    this.fetchData();
  },
  data() {
    return {
      dataList: [],
      loading: ref(true),
      workspaceUserList: [],
      userDataList: [],
      selectedUser: ref(""),
      selectedVmQuantity: ref(1),
      columns: [
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
          slots: { customRender: "actionRender" },
        },
        // {
        //   title: this.$t("label.state"),
        //   dataIndex: "state",
        //   key: "state",
        //   sorter: (a, b) => (a.state < b.state ? -1 : a.state > b.state ? 1 : 0),
        //   sortDirections: ["descend", "ascend"],
        // },
        {
          title: this.$t("label.allocateddesktop"),
          dataIndex: "desktop",
          key: "desktop",
          sorter: (a, b) =>
            a.desktop < b.desktop ? -1 : a.desktop > b.desktop ? 1 : 0,
          sortDirections: ["descend", "ascend"],
        },
      ],
    };
  },
  components: {
    Actions,
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
      this.loading = ref(true);

      //해당 워크스페이스에 추가 된 사용자 목록 조회
      worksApi
        .get("/api/v1/group/" + this.$route.params.groupName)
        .then((response) => {
          if (response.status == 200) {
            //console.log(response.data.result.member);
            const temp =
              response.data.result.member == undefined
                ? ""
                : response.data.result.member;
            for (let str of temp) {
              this.userDataList.push({
                name: str.split(",")[0].split("CN=")[1],
              });
            }
          } else {
            message.error(this.t("message.response.data.fail"));
            //console.log(response.message);
          }
        })
        .catch(function (error) {
          message.error(error);
          //console.log(error);
        });

      // worksApi
      //   .get("/api/v1/group/"+this.$route.params.groupName)
      //   .then((response) => {
      //     if (response.status == 200) {
      //       this.userDataList = response.data.result.member;
      //       //console.log(response.data.result.member);
      //     } else {
      //       message.error(this.$t('message.response.data.fail'));
      //       //console.log("데이터를 정상적으로 가져오지 못했습니다.");
      //     }
      //   })
      //   .catch(function (error) {
      //     console.log(error);
      //   });
      setTimeout(() => {
        this.loading = ref(false);
      }, 1000);

      if (this.state.callTapName === "user") {
        this.state.addButtonBoolean = ref(true);
        this.state.addButtontext = this.$t("label.desktop.user.add");
      }
    },

    putUserToWorksapce() {
      let params = new URLSearchParams();
      params.append("username", this.selectedUser);
      //  worksApi
      //   .put("/api/v1/workspace/user", params)
      //   .then((response) => {
      //     if (response.status === 200) {
      //       message.loading(this.$t("message.workspace.user.add"), 1);

      //     } else {
      //       message.error(response.data.result.createuserresponse.errortext);
      //     }
      //     changeModal('user', false)
      //     //this.$refs.listRefleshCall.fetchData();
      //   })
      //   .catch(function (error) {
      //     message.error(error.message);
      //   //console.log(error);
      //   });
      this.changeModal("user", false);
    },
  },
});
</script>

<style scoped>
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
