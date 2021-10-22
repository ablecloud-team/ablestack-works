<template>
  <!-- <a-list item-layout="horizontal">
      <a-list-item v-for="(value, key) in dataInfo">
      <strong v-if="item === obj">{{ $t('label.' + String(key).toLowerCase()) }}</strong>
      <br>
      {{value}}
    </a-list-item> 
  </a-list>-->
  <a-list v-if="actionFrom === 'VirtualMachineDetail'" item-layout="horizontal">
    <a-list-item>
      <strong>{{ $t("label.name") }}</strong>
      <br />
      {{ dataInfo.displayname }}
    </a-list-item>
    <a-list-item>
      <strong>{{ $t("label.workspace") }}</strong>
      <br />
      {{ vmDbDataInfo.workspace_name }}
    </a-list-item>
    <a-list-item>
      <strong>{{ $t("label.vm.state") }}</strong>
      <br />
      {{
        dataInfo.state == "Running"
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
      {{ dataInfo.domain }}
    </a-list-item>
    <a-list-item>
      <strong>{{ $t("label.hypervisor") }}</strong>
      <br />
      {{ dataInfo.hypervisor }}
    </a-list-item>
    <a-list-item>
      <strong>{{ $t("label.pooltype") }}</strong>
      <br />
      {{ dataInfo.pooltype }}
    </a-list-item>
  </a-list>
  <a-list v-if="actionFrom === 'UserDetail'" item-layout="horizontal">
    <a-list-item>
      <strong>{{ $t("label.account") }}</strong
      ><br />
      {{ dataInfo.username }}
    </a-list-item>
    <a-list-item>
      <strong>{{ $t("label.country") }}</strong
      ><br />
      {{ dataInfo.co }}
    </a-list-item>
    <a-list-item>
      <strong>{{ $t("label.countryCode") }}</strong
      ><br />
      {{ dataInfo.countryCode }}
    </a-list-item>
    <a-list-item>
      <strong>{{ $t("label.title") }}</strong
      ><br />
      {{ dataInfo.title }}
    </a-list-item>
    <a-list-item>
      <strong>{{ $t("label.email") }}</strong
      ><br />
      {{ dataInfo.mail }}
    </a-list-item>
    <a-list-item>
      <strong>{{ $t("label.isAdmin") }}</strong
      ><br />
      {{ dataInfo.isAdmin }}
    </a-list-item>
    <a-list-item>
      <strong>{{ $t("label.telephoneNumber") }}</strong
      ><br />
      {{ dataInfo.telephoneNumber }}
    </a-list-item>
    <a-list-item>
      <strong>{{ $t("label.userPrincipalName") }}</strong
      ><br />
      {{ dataInfo.userPrincipalName }}
    </a-list-item>
    <a-list-item>
      <strong>{{ $t("label.distinguishedName") }}</strong
      ><br />
      {{ dataInfo.distinguishedName }}
    </a-list-item>
  </a-list>
  <a-list v-if="actionFrom === 'GroupPolicyDetail'" item-layout="horizontal">
    <a-list-item>
      <strong>{{ $t("label.account") }}</strong
      ><br />
      {{ dataInfo.username }}
    </a-list-item>
  </a-list>
</template>

<script>
export default {
  name: "DetailContent",
  props: {
    dataInfo: {
      type: Object,
      required: true,
      default: null,
    },
    vmDbDataInfo: {
      type: Object,
      required: true,
      default: null,
    },
    actionFrom: {
      type: String,
      required: true,
      default: null,
    },
  },
};
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
