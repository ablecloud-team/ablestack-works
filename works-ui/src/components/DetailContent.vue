<template>
  <a-spin :spinning="spinning">
    <a-list
      v-if="actionFrom === 'VirtualMachineDetail'"
      item-layout="horizontal"
    >
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
        <br />{{ vmDbDataInfo.state }}

        {{
          vmDbDataInfo.mold_status === "Running"
            ? $t("label.vm.status.running")
            : vmDbDataInfo.mold_status === "Starting"
            ? $t("label.vm.status.starting")
            : vmDbDataInfo.mold_status === "Stopping"
            ? $t("label.vm.status.stopping")
            : vmDbDataInfo.mold_status === "Stopped"
            ? $t("label.vm.status.stopped")
            : ""
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
    <a-list
      v-if="actionFrom === 'UserDetail' || actionFrom === 'AccountDetail'"
      item-layout="horizontal"
    >
      <a-list-item>
        <strong>{{ $t("label.account") }}</strong>
        <br />
        {{ accountInfo.username }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.country") }}</strong>
        <br />
        {{ accountInfo.co }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.countryCode") }}</strong>
        <br />
        {{ accountInfo.countryCode }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.title") }}</strong>
        <br />
        {{ accountInfo.title }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.department") }}</strong>
        <br />
        {{ accountInfo.department }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.email") }}</strong>
        <br />
        {{ accountInfo.mail }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.isAdmin") }}</strong
        ><br />
        {{ accountInfo.isAdmin }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.telephoneNumber") }}</strong>
        <br />
        {{ accountInfo.telephoneNumber }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.userPrincipalName") }}</strong>
        <br />
        {{ accountInfo.userPrincipalName }}
      </a-list-item>
      <a-list-item>
        <strong>{{ $t("label.distinguishedName") }}</strong>
        <br />
        {{ accountInfo.distinguishedName }}
      </a-list-item>
    </a-list>
    <a-list v-if="actionFrom === 'GroupPolicyDetail'" item-layout="horizontal">
      <a-list-item>
        <strong>{{ $t("label.account") }}</strong
        ><br />
        {{ accountInfo.username }}
      </a-list-item>
    </a-list>
  </a-spin>
</template>

<script>
import { defineComponent, ref } from "vue";

export default defineComponent({
  name: "DetailContent",
  props: {
    actionFrom: {
      type: String,
      required: true,
      default: null,
    },
    vmDbDataInfo: {
      type: Object,
      required: false,
      default: null,
    },
    vmMoldDataInfo: {
      type: Object,
      required: false,
      default: null,
    },
    accountInfo: {
      type: Object,
      required: false,
      default: null,
    },
  },
  setup(props) {
    return {
      actionFrom: ref(props.actionFrom),
    };
  },
  data() {
    return {
      spinning: ref(true),
      userDataInfo: ref([]),
    };
  },
  created() {
    this.fetchRefresh();
  },
  methods: {
    fetchRefresh() {
      this.fetchData();
      this.spinning = true;
    },
    fetchData() {
      setTimeout(() => {
        this.spinning = false;
      }, 500);
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
