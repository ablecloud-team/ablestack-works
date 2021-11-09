<template>
  <a-spin :spinning="spinning">
    <a-list v-if="actionFrom === 'VirtualMachineDetail'" item-layout="horizontal">
      <a-list-item>
        <strong>{{ $t("label.name") }}</strong>
        <br />
        {{ vmMoldDataInfo.displayname }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.workspace") }}</strong>
        <br />
        {{ vmDbDataInfo.workspace_name }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.vm.state") }}</strong>
        <br />{{vmDbDataInfo.state}}
        
        {{
          vmMoldDataInfo.state == "Running"
            ? $t("label.vm.status.running")
            : $t("label.vm.status.stopped")
        }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.users") }}</strong>
        <br />
        {{
          vmDbDataInfo.owner_account_id == ""
            ? $t("label.owner.account.false")
            : vmDbDataInfo.owner_account_id
        }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.vm.session.count") }}</strong>
        <br />
        {{ vmDbDataInfo.connected }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.domain") }}</strong>
        <br />
        {{ vmMoldDataInfo.domain }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.hypervisor") }}</strong>
        <br />
        {{ vmMoldDataInfo.hypervisor }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.pooltype") }}</strong>
        <br />
        {{ vmMoldDataInfo.pooltype }}
      </a-list-item>
    </a-list>
    <a-list v-if="actionFrom === 'UserDetail' || actionFrom === 'AccountDetail'" item-layout="horizontal">
      <a-list-item>
        <strong>{{ $t("label.account") }}</strong
        ><br />
        {{ userDataInfo.username }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.country") }}</strong
        ><br />
        {{ userDataInfo.co }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.countryCode") }}</strong
        ><br />
        {{ userDataInfo.countryCode }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.title") }}</strong
        ><br />
        {{ userDataInfo.title }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.email") }}</strong
        ><br />
        {{ userDataInfo.mail }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.isAdmin") }}</strong
        ><br />
        {{ userDataInfo.isAdmin }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.telephoneNumber") }}</strong
        ><br />
        {{ userDataInfo.telephoneNumber }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.userPrincipalName") }}</strong
        ><br />
        {{ userDataInfo.userPrincipalName }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.distinguishedName") }}</strong
        ><br />
        {{ userDataInfo.distinguishedName }}
      </a-list-item>
    </a-list>
    <a-list v-if="actionFrom === 'GroupPolicyDetail'" item-layout="horizontal">
      <a-list-item>
        <strong>{{ $t("label.account") }}</strong
        ><br />
        {{ userDataInfo.username }}
      </a-list-item>
    </a-list>
  </a-spin>
</template>

<script>
import { defineComponent, reactive, ref } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";

export default defineComponent({
  name: "DetailContent",
  props: {
    actionFrom: {
      type: String,
      required: true,
      default: null,
    },
  },
  setup(props) {
     const state = reactive({
      actionFrom: ref(props.actionFrom),
    });
    return {
      state,
    };
  },
  data() {
    return {
      spinning: ref(true),
      vmDbDataInfo: ref([]),
      vmMoldDataInfo: ref([]),
      userDataInfo: ref([]),
    };
  },
  created() {
    this.reflesh();
  },
  methods: {
    reflesh() {
      this.fetchData();
      this.spinning = true;
      setTimeout(() => {
        this.spinning = false;
      }, 500);
    },
    fetchData() {
      // 가상머신 상세조회
      if (this.state.actionFrom == "VirtualMachineDetail") {
        worksApi
        .get("/api/v1/instance/detail/" + this.$route.params.vmUuid)
        .then((response) => {
          if (response.status === 200) {
            this.vmDbDataInfo = response.data.result.instanceDBInfo;
            this.vmMoldDataInfo =
              response.data.result.instanceMoldInfo.virtualmachine[0];
          } else {
            message.error(this.$t("message.response.data.fail"));
            //console.log("데이터를 정상적으로 가져오지 못했습니다.");
          }
        })
        .catch(function (error) {
          console.log(error);
        });
      }else if (this.state.actionFrom === "UserDetail" || this.state.actionFrom === "AccountDetail") {

        let apiUrl = this.state.actionFrom === "AccountDetail" ? "/api/v1/user/" + this.$route.params.userName : 
                  this.state.actionFrom === "UserDetail" ? "/api/v1/user/" + sessionStorage.getItem("username") : "";

        worksApi
          .get(apiUrl)
          .then((response) => {
            if (response.status == 200) {
              this.userDataInfo = response.data.result;
            } else {
              message.error(this.$t("message.response.data.fail"));
              //console.log("데이터를 정상적으로 가져오지 못했습니다.");
            }
          })
          .catch(function (error) {
            console.log(error);
          });
      }
    },
  },
});
</script>

<style scoped>
.detail-content-card {
  margin-bottom: 10px;
  padding: "7px 24px 0px 24px";
  border: "0px";
}
::v-deep(.ant-card-head-title) {
  padding: 0px;
}
</style>
