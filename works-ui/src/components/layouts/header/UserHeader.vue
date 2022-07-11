<template>
  <a-row class="user-menu">
    <a-col :flex="1">
      <menu-unfold-outlined
        v-if="state.collapsed"
        class="menu-collapsed"
        @click="setCollapsed()"
      />
      <menu-fold-outlined
        v-else
        class="menu-collapsed"
        @click="setCollapsed()"
      />
    </a-col>
    <a-col :flex="4" class="header-col-right">
      <span>
        【 {{ $t("label.cluster") }} : {{ clusterName }} ㅣ
        {{ $t("label.domain") }} : {{ domainName }} 】
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
      </span>
      <a-dropdown placement="bottomLeft">
        <span>
          <TranslationOutlined class="header-col" />
        </span>
        <template #overlay>
          <a-menu
            v-model:selected-keys="state.language"
            class="header-dropdown-menu"
            @click="setLocaleClick"
          >
            <a-menu-item key="ko" value="koKR"> 한국어 </a-menu-item>
            <a-menu-item key="en" value="enUS"> English </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
      &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
      <a-dropdown placement="bottomLeft">
        <span>
          <UserOutlined class="header-col" />
          {{ userName }}
        </span>
        <template #overlay>
          <a-menu>
            <a-menu-item key="1" @click="userinfo">
              <UserOutlined /> {{ $t("label.profile") }}
            </a-menu-item>
            <a-menu-divider />
            <a-menu-item key="2" @click="logoutSubmit">
              <LogoutOutlined /> {{ $t("label.logout") }}
            </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
    </a-col>
  </a-row>
</template>

<script>
import { defineComponent, reactive, ref } from "vue";
import { axiosLogout } from "@/api/index";
export default defineComponent({
  name: "UserHeader",
  components: {},
  props: {
    collapsed: Boolean,
  },
  emits: ["setCollapsed"],
  setup(props) {
    const state = reactive({
      collapsed: ref(props.collapsed),
      language: [
        sessionStorage.getItem("locale") == null
          ? "ko"
          : sessionStorage.getItem("locale"),
      ],
    });
    return {
      state,
    };
  },
  data() {
    return {
      loadedLanguage: ref[""],
      username: ref(""),
    };
  },
  created() {
    // const res = await axiosUserDetail()
    // if(res.status === 200){
    //   this.$store.dispatch("loginCommit",res.data);
    //   this.state.userID = res.data.result.name;
    // }
    this.userName = sessionStorage.getItem("userName");
    this.clusterName = sessionStorage.getItem("clusterName");
    this.domainName = sessionStorage.getItem("domainName");

    this.setLocale(this.state.language[0]);
  },
  methods: {
    setCollapsed() {
      this.state.collapsed = !this.state.collapsed;
      this.$emit("setCollapsed");
    },
    setLocaleClick(e) {
      let localeValue = e.key;
      if (!localeValue) {
        localeValue = "ko";
      }
      this.setLocale(localeValue);
      // setTimeout(() => {
      //   location.reload();
      // }, 0);
    },
    setLocale(localeValue) {
      this.$locale = localeValue;
      this.$i18n.locale = localeValue;
      this.state.language = [localeValue];
      sessionStorage.setItem("locale", localeValue);
    },
    async logoutSubmit() {
      const res = await axiosLogout();
      if (res.data.result.login === false || res.data.result.status > 200) {
        this.$message.success(this.$t("message.logout.success"), 2);
        await this.$store.dispatch("logoutCommit");
        await this.$router.push({ name: "Login" });
      }
    },
    userinfo() {
      this.$router.push({ path: "/userDetail" });
    },
    // setLanguage(lang, message) {
    //   if (i18n) {
    //     i18n.locale = lang;
    //     if (message && Object.keys(message).length > 0) {
    //       i18n.setLocaleMessage(lang, message);
    //     }
    //   }
    //   if (!this.loadedLanguage.includes(lang)) {
    //     this.loadedLanguage.push(lang);
    //   }
    //   if (message && Object.keys(message).length > 0) {
    //     messages[lang] = message;
    //   }
    // },
  },
});
</script>

<style scoped>
</style>
