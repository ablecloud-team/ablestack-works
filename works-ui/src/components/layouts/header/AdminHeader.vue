<template>
    <a-row class="user-menu">
    <a-col :span="12">
        <menu-unfold-outlined
          v-if="state.collapsed"
          class="trigger3"
          @click="setCollapsed()"
        />
        <menu-fold-outlined v-else class="trigger3" @click="setCollapsed()" />
    </a-col>
    <a-col :span="12" style="float:right; text-align: right; padding-right: 40px">
      <a-dropdown>
        <a-button type="text" shape="circle" class="header-notice-button">
          <a class="ant-dropdown-link" @click.prevent>
            <font-awesome-icon 
              :icon="['fas', 'language']"
              style="color: #666; margin-bottom: -2px;"
              class="login-icon"/>
            <!-- <GlobalOutlined /> -->
          </a>
        </a-button>
        <template #overlay>
          <a-menu
            :selectedKeys="[language]"
            @click="setLocaleClick">
            <a-menu-item key="ko" value="koKR"> 한국어 </a-menu-item>
            <a-menu-item key="en" value="enUS"> English </a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
      <a-button type="text" shape="circle" class="header-notice-button">
        <a class="ant-dropdown-link" @click.prevent>
          <BellOutlined class="header-notice-icon" />
        </a>
      </a-button>
      <a-dropdown>
        <a-button type="text" shape="circle" class="header-notice-button">
          <a class="ant-dropdown-link" @click.prevent>
            <UserOutlined class="header-notice-icon" />
          </a>
        </a-button>
        <template #overlay>
          <a-menu >
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
      <!-- <a-popover placement="bottom">
        <template #content>
          <a-button type="text" @click="setLocaleClick() $i18n.locale = 'ko'">한국어</a-button>
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
      </a-popover> -->

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


    </a-col>
  </a-row>
</template>

<script>
import {
  MenuUnfoldOutlined,
  MenuFoldOutlined,
  DownOutlined,
  StarOutlined
} from '@ant-design/icons-vue';
import { defineComponent, reactive, ref } from "vue";
import { message } from "ant-design-vue";
import { axiosUserDetail, axiosLogout } from "@/api/index";
import store from "@/store/index";
import router from "@/router";

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
  data() {
    return {
      language: '',
      loadedLanguage: [],
    }
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
  created() {
    // const res = await axiosUserDetail()
    // if(res.status === 200){
    //   this.$store.dispatch("loginCommit",res.data);
    //   this.state.userID = res.data.result.name;
    // }
    this.language = localStorage.getItem('locale');
    this.setLocale(this.language);
  },
  methods: {
    setCollapsed() {
      this.state.collapsed = !this.state.collapsed;
      this.$emit("setCollapsed");
    },
    setLocaleClick (e) {
      let localeValue = e.key;
      if (!localeValue) {
        localeValue = 'ko';
      }
      this.setLocale(localeValue);
    },
    setLocale (localeValue) {
      this.$locale = localeValue;
      this.$i18n.locale = localeValue;
      this.language = localeValue;
      localStorage.setItem('locale', localeValue);
      //this.loadLanguageAsync(localeValue);
    },
    async logoutSubmit() {
      const res = await axiosLogout();
      if (res.data.result.login === false || res.data.result.status > 200) {
        message.success(this.$t('message.logout.success'), 2);
        await store.dispatch("logoutCommit");
        await router.push({ name: "Login" });
      }
    },
    loadLanguageAsync(lang) {
      if (!lang) {
        lang = Vue.ls ? Vue.ls.get('LOCALE') || 'en' : 'en';
      }
      if (this.loadedLanguage.includes(lang)) {
        return Promise.resolve(this.setLanguage(lang));
      }

      return fetch(`locales/${lang}.json`)
        .then(response => response.json())
        .then(json => Promise.resolve(this.setLanguage(lang, json)))
    },
    setLanguage(lang, message) {
      if (i18n) {
        i18n.locale = lang;

        if (message && Object.keys(message).length > 0) {
          i18n.setLocaleMessage(lang, message);
        }
      }

      if (!this.loadedLanguage.includes(lang)) {
        this.loadedLanguage.push(lang);
      }

      if (message && Object.keys(message).length > 0) {
        messages[lang] = message;
      }
    }
  },
});
</script>

<style scoped>
.header-notice-icon {
  font-size: 18px;
}
.header-notice-button {
  margin-left: 10px;
  margin-right: 10px;
}

.user-menu {
  background: white;
  height: 100%;
}
.login-icon {
  font-size: 24px;
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

.trigger3 {
  font-size: 18px;
  line-height: 64px;
  padding: 0 24px;
  cursor: pointer;
  transition: color 0.3s;
}

.trigger3:hover {
  color: #1890ff;
}
</style>
