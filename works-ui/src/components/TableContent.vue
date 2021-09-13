<template>
  <a-button
    type="dashed"
    block
    style="margin-bottom: 14px"
    @click="changeModal(state.callTapName, true)"
    v-if="state.addButtonBoolean">
    <PlusOutlined />
    {{ state.addButtontext }}
  </a-button>

  <a-table
    :loading="loading"
    size="small"
    class="ant-table-striped"
    style="overflow: auto; margin-left:5px;"
    :columns="columns"
    :data-source="data"
    :pagination="pagination">
    <template #nameRender="{ record }">
      <span v-if="comp !== undefined && comp ==='VirtualMachineDetail'">
        <router-link :to="{ path: '/virtualMachineDetail/'+record.uuid}">{{ record.name }}</router-link>
      </span>
      <span v-else>
        {{ record.name }}
      </span>
    </template>

    <template #actionRender>
      <a-Popover placement="top">
        <template #content>
          <ASpace direction="horizontal">
            <actions :actionFrom="actionFrom" />
          </ASpace>
        </template>
        <MoreOutlined />
      </a-Popover>
    </template>

    <!-- <template #tags="{ text: tags }">
      <span>
        <a-tag
          v-for="tag in tags"
          :key="tag"
          :color="
            tag === 'loser' ? 'volcano' : tag.length > 5 ? 'geekblue' : 'green'
          "
        >
          {{ tag.toUpperCase() }}
        </a-tag>
      </span>
    </template> -->
  </a-table>

  <a-modal
    v-model:title="state.addButtontext"
    v-model:visible="state.addVmModalBoolean"
    @ok="putVmToWorksapce"
    width="400px">
    <a-input-number
      v-model:value="addDesktopQuantity"
      id="inputNumber"
      style="width: 100%; margin-top: 7px"
      :title="'Desktop Quantity'"
      :min="1"
      :max="10"/>
  </a-modal>

  <a-modal
    v-model:title="state.addButtontext"
    v-model:visible="state.addUserModalBoolean"
    @ok="putUserToWorksapce"
    width="400px">
    <a-select
      v-model:value="selectedUser"
      :placeholder="$t('tooltip.user')"
      show-search
      style="width: 100%; margin-top: 7px"
      option-filter-prop="label"
      class="addmodal-aform-item-div">
      <a-select-option v-for="option in userDataList" :key="option.name" :value="option.name" :label="option.name">
        {{ option.name }}
      </a-select-option>
    </a-select>
  </a-modal>
</template>

<script>
// import {SmileOutlined} from '@ant-design/icons-vue';
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
  props: {
    tapName: String,
    data: Object,
    columns: Object,
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
        pageSizeOptions: ["10", "20", "50", "100"], // number of pages per option
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
      loading: ref(false),
      workspaceUserList: [],
      userDataList: [],
      selectedUser: ref(""),
      selectedVmQuantity: ref(1),
    };
  },
  components: {
    Actions,
  },
  methods: {
    changeModal(target, value) {
      if(target == 'desktop'){
        this.state.addVmModalBoolean = ref(value);
      }else if(target == 'user'){
        this.state.addUserModalBoolean = ref(value);
      }
      
    },
    fetchData() {
      this.loading = ref(true);
      if (this.state.callTapName === "desktop") {
        this.state.addButtonBoolean = ref(true);
        this.state.addButtontext = this.$t("label.desktop.vm.add");
      } else if (this.state.callTapName === "user") {
        this.state.addButtonBoolean = ref(true);
        this.state.addButtontext = this.$t("label.desktop.user.add");

        worksApi
          .get("/api/v1/user", { withCredentials: true })
          .then((response) => {
            if (response.data.result.status == 200) {
              this.userDataList = response.data.result.result;
            } else {
              message.error(this.t('message.response.data.fail'));
              //console.log(response.message);
            }
          })
          .catch(function (error) {
            message.error(error.message);
            //console.log(error);
          });


      } else if (this.state.callTapName === "datadisk") {
        this.state.addButtonBoolean = ref(false);
      } else if (this.state.callTapName === "network") {
        this.state.addButtonBoolean = ref(false);
      }

      this.loading = ref(false);
    },
    putVmToWorksapce() {
      //alert(this.$route.params.uuid + " ::: "+ this.selectedVmQuantity);
      let params = new URLSearchParams();
      params.append("uuid", this.$route.params.uuid);
      params.append("quantity", this.selectedVmQuantity);
      worksApi
        .put("/api/v1/instance", params, { withCredentials: true })
        .then((response) => {
          if (response.status === 200) {
            message.loading(this.$t("message.workspace.user.add"), 1);
            // setTimeout(() => {
            //   location.reload();
            // }, 1500);
          } else {
            message.error(response.data.result.createuserresponse.errortext);
          }
          this.changeModal('user', false);
        })
        .catch(function (error) {
          message.error(error);
        //console.log(error);
        });

      this.changeModal('desktop', false);
      this.$emit('tabReflesh');

    },
    putUserToWorksapce() {
      //alert(this.selectedUser);
      let params = new URLSearchParams();
      params.append("username", this.selectedUser);
      //  worksApi
      //   .put("/api/v1/workspace/user", params, { withCredentials: true })
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
        this.changeModal('user', false)
        this.$emit('tabReflesh');

    }
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
