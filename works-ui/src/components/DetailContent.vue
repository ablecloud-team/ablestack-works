<template>
   <!-- <a-list item-layout="horizontal">
      <a-list-item v-for="(value, key) in dataInfo">
      <strong v-if="item === obj">{{ $t('label.' + String(key).toLowerCase()) }}</strong>
      <br>
      {{value}}
    </a-list-item> 
  </a-list>-->
  <a-list item-layout="horizontal" v-if="actionFrom === 'vmDetail'"> 
    <a-list-item >
      <strong>{{ $t('label.name') }}</strong>
      <br>
      {{dataInfo.name}}
    </a-list-item>
  </a-list>
  <a-list item-layout="horizontal" v-if="actionFrom === 'userDetail'">
    <a-list-item >
      <strong>{{ $t('label.account') }}</strong><br>
      {{dataInfo.username}}
    </a-list-item>
    <a-list-item >
      <strong>{{ $t('label.country') }}</strong><br>
      {{dataInfo.co}}
    </a-list-item>
    <a-list-item >
      <strong>{{ $t('label.countrycode') }}</strong><br>
      {{dataInfo.countryCode}}
    </a-list-item>
    <a-list-item >
      <strong>{{ $t('label.distinguishedName') }}</strong><br>
      {{dataInfo.distinguishedName}}
    </a-list-item>
    <a-list-item >
      <strong>{{ $t('label.email') }}</strong><br>
      {{dataInfo.mail}}
    </a-list-item>
        <a-list-item >
      <strong>{{ $t('label.isAdmin') }}</strong><br>
      {{dataInfo.isAdmin}}
    </a-list-item>
    <a-list-item >
      <strong>{{ $t('label.telephoneNumber') }}</strong><br>
      {{dataInfo.telephoneNumber}}
    </a-list-item>
    <a-list-item >
      <strong>{{ $t('label.userPrincipalName') }}</strong><br>
      {{dataInfo.userPrincipalName}}
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
    actionFrom: {
      type: String,
      required: true,
      default: null,
    }
  },
};
</script>

<style scoped>
.detail-content-card {
  margin-bottom: 10px;
  minHeight: '0px';
  padding: '7px 24px 0px 24px';
  fontWeight: '700';
  border: '0px';
}
::v-deep(.ant-card-head-title) {
  padding: 0px;
}
</style>
