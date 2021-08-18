<template>
  <a-row class="user-menu">
    <a-col :span="12">
<!--      <MenuFoldOutlined class="header-notice-icon" />-->
    </a-col>
    <a-col :span="12" style="float:right; text-align: right; padding-right: 14px;">
      <a-popover placement="bottom">
        <template #content>
          <a-button type="text" @click="$i18n.locale = 'ko'">
            한국어
          </a-button>
          <br />
          <a-button type="text" @click="$i18n.locale = 'en'">
            English
          </a-button>
        </template>
        <a-button type="text" shape="circle" class="header-notice-button">
          <template #icon>
            <font-awesome-icon
                :icon="['fas', 'language']"
                size="4x"
                style="color: #666; padding-top: 3px;"
                class="login-icon"
            />
          </template>
        </a-button>
      </a-popover>
      <a-button type="text" shape="circle" class="header-notice-button">
        <template #icon>
          <BellOutlined class="header-notice-icon" />
        </template>
      </a-button>
      <a-popover style="text-align: left;">
        <template #content>
          <a-card size="small" style="width: 200px" :bordered="false" :bodyStyle="{padding: '0', margin: '0'}">
            <a-button type="text" class="header-popover-button">
              <UserOutlined />
              {{$t("label.profile")}}
            </a-button>
            <a-button type="text" class="header-popover-button logout-button" @click="logoutSubmit">
              <LogoutOutlined />
              {{$t("label.logout")}}
            </a-button>
          </a-card>
        </template>
        <a-button type="text" shape="circle" class="header-notice-button">
          <template #icon>
            <UserOutlined class="header-notice-icon" />
          </template>
          {{ state.userID }}
        </a-button>
      </a-popover>
    </a-col>
  </a-row>
</template>

<script>
import { reactive } from "vue";
import { axiosUserDetail, axiosLogout } from "../../../api";
import router from "../../../router";
import {message} from "ant-design-vue";

export default {
  name: "AdminHeader",
  setup(){
    let res
    const state = reactive({
      userID: "",
    })
    return{
      state,
    }
  },
  methods: {
    async logoutSubmit() {
      const res = await axiosLogout()
      if(res.status === 200){
        await router.push("/login")
        message.success("로그아웃되었습니다.", 2);
      }
    }
  },
  async mounted() {
    const res = await axiosUserDetail()
    if(res.status === 200){
      this.$store.dispatch("loginCommit",res.data)
      this.state.userID = res.data.result.name
    }
  },
};
</script>

<style scoped>
.header-notice-icon {
  font-size: 18px;
}
.header-notice-button {
  margin-right: 10px;
}

.user-menu {
  background: white;
  height: 100%;
}
.login-icon {
  font-size: 27px;
}
.header-popover-button{
  width: 100%;
  text-align: left;
  padding: 0;
}
.logout-button{
  margin-top: 14px;
  border-top: 1px solid #e8e8e8;
}
</style>
