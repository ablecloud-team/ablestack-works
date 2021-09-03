<template>
    
    <a-row class="user-menu">
    <a-col :span="12">
      <!--<MenuFoldOutlined class="header-notice-icon" />-->
    </a-col>
    <a-col :span="12" style="float:right; text-align: right; padding-right: 14px">
      <a-popover placement="bottom">
        <template #content>
          <a-button type="text" @click="$i18n.locale = 'ko'">한국어</a-button>
          <br />
          <a-button type="text" @click="$i18n.locale = 'en'">English</a-button>
        </template>
        <a-button type="text" shape="circle" class="header-notice-button">
          <template #icon>
            <font-awesome-icon
              :icon="['fas', 'language']"
              size="4x"
              style="color: #666; padding-top: 3px"
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
      <!-- <a-popover style="text-align: left">
        <template #content>
          <a-card size="small" style="width: 150px" :bordered="false" >
            <a-button type="text" class="">
              <UserOutlined />{{ $t("label.profile") }}
            </a-button>
            <a-divider />
            <a-button type="text" class="" @click="logoutSubmit">
              <LogoutOutlined />{{ $t("label.logout") }}
            </a-button>
          </a-card>
        </template>
        <a-button type="text" shape="circle" class="header-notice-button">
          <template #icon>
            <UserOutlined class="header-notice-icon" />
          </template>
          {{ state.userID }}
        </a-button>
      </a-popover> -->

      <a-dropdown :trigger="['click']">
        <a class="ant-dropdown-link" @click.prevent>
          <UserOutlined class="header-notice-icon" />
        </a>
        <template #overlay>
          <a-menu style="width: 200px; padding: 0;">
            <a-menu-item key="0" >
              <a-button type="text" class="">
                <UserOutlined />{{ $t("label.profile") }}
              </a-button>
            </a-menu-item>
            <a-menu-divider />
            <a-menu-item key="1">
              <a-button type="text" class="" @click="logoutSubmit">
                <LogoutOutlined />{{ $t("label.logout") }}
              </a-button>
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
    </a-col>
  </a-row>
</template>

<script>
import {
  MenuUnfoldOutlined,
  MenuFoldOutlined,
  DownOutlined,
} from '@ant-design/icons-vue';
import { defineComponent, reactive, ref } from "vue";
import { message } from "ant-design-vue";
import { axiosUserDetail, axiosLogout } from "../../../api";
import store from "../../../store/index";
import router from "../../../router";

export default defineComponent({
  name: "AdminHeader",
  components: {
    DownOutlined,
    MenuUnfoldOutlined,
    MenuFoldOutlined,
  },
  props: {
    collapsed: Boolean,
  },
  setup(props) {
    let res;
    const state = reactive({
      userID: "",
      collapsed: ref(props.collapsed),
    });
    return {
      state,
    };
  },
  async mounted() {
    // const res = await axiosUserDetail()
    // if(res.status === 200){
    //   this.$store.dispatch("loginCommit",res.data);
    //   this.state.userID = res.data.result.name;
    // }
  },
  methods: {
    async logoutSubmit() {
      const res = await axiosLogout();
      if (res.data.result.login === false || res.data.result.status > 200) {
        await store.dispatch("logoutCommit");
        await router.push({ name: "Login" });
        message.success("로그아웃되었습니다.", 2);
      }
    },
  },
});
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
.header-popover-button {
  width: 100%;
  text-align: left;
  padding: 0;
}
.logout-button {
  margin-top: 14px;
  border-top: 1px solid #e8e8e8;
}

</style>
