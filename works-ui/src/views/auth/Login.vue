<template>
  <div class="user-layout desktop">
    <span>
      <a-alert
        v-if="serverStatus"
        :message="statusMessage"
        :description="statusDescription"
        :type="statusType"
        banner
        show-icon
      >
        <!-- <template #icon>
          <smile-outlined />
          <smile-outlined v-if="statusType == 'success'" />
          <frown-outlined v-else /> -->
        <!-- </template> -->
      </a-alert>
    </span>
    <div class="user-layout-container">
      <div class="user-layout-container">
        <img
          src="@/assets/ablestack-logo.png"
          alt="logo"
          class="user-layout-logo"
        />
        <div v-if="disabled" style="text-align: center">
          <a-spin />
        </div>
      </div>

      <a-form
        ref="formRef"
        layout="horizontal"
        :model="formState"
        :rules="rules"
        @finish="onSubmit"
        class="user-layout-login"
      >
        <a-form-item has-feedback name="id">
          <a-input
            :disabled="disabled"
            v-model:value="formState.id"
            :placeholder="$t('label.user.id')"
            size="large"
          >
            <template #prefix>
              <UserOutlined style="color: rgba(0, 0, 0, 0.25)" />
            </template>
          </a-input>
        </a-form-item>

        <a-form-item has-feedback name="password">
          <a-input-password
            :disabled="disabled"
            v-model:value="formState.password"
            type="password"
            :placeholder="$t('label.password')"
            size="large"
          >
            <template #prefix>
              <LockOutlined style="color: rgba(0, 0, 0, 0.25)" />
            </template>
          </a-input-password>
        </a-form-item>
        <a-form-item style="margin-bottom: 0">
          <a-button
            :disabled="disabled"
            type="primary"
            block
            class="login-button"
            html-type="submit"
          >
            {{ $t("label.login") }}
          </a-button>
        </a-form-item>
        <!--   언어변환 버튼 start     -->
        <a-dropdown>
          <a-button type="text" shape="circle" class="header-notice-button">
            <a class="ant-dropdown-link" @click.prevent>
              <font-awesome-icon
                :icon="['fas', 'language']"
                size="2x"
                style="color: #666"
                class="login-icon"
              />
              <!-- <GlobalOutlined /> -->
            </a>
          </a-button>
          <template #overlay>
            <a-menu :selected-keys="[language]" @click="setLocaleClick">
              <a-menu-item key="ko" value="koKR"> 한국어 </a-menu-item>
              <a-menu-item key="en" value="enUS"> English </a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
        <!--   언어변환 버튼 끝     -->
      </a-form>
    </div>
  </div>
  <!-- </a-spin> -->
</template>

<script>
import { SmileOutlined, FrownOutlined } from "@ant-design/icons-vue";
import { defineComponent, reactive, ref } from "vue";
import { message } from "ant-design-vue";
import { worksApi } from "@/api/index";
import store from "@/store/index";
import router from "@/router";

export default defineComponent({
  name: "Login",
  components: { SmileOutlined, FrownOutlined },
  setup() {
    const formRef = ref();
    const formState = reactive({
      id: ref(""),
      password: ref(""),
    });
    const rules = {
      id: {
        required: true,
        trigger: "change",
      },
      password: {
        required: true,
        trigger: "change",
      },
    };
    return {
      formRef,
      formState,
      rules,
    };
  },
  data() {
    return {
      timer: ref(null),
      spinning: ref(true),
      language: ref(""),
      loadedLanguage: ref[""],
      disabled: ref(true),
      serverStatus: ref(true),
      statusMessage: ref(this.$t("message.status.check.worksapi")),
      statusDescription: ref(this.$t("message.status.check.desc.wait")),
      statusType: ref("info"),
    };
  },
  created() {
    this.serverStatus = true;
    this.statusType = "info";
    this.disabled = true;
    this.serverCheck();
    this.timer = setInterval(() => {
      this.serverCheck();
    }, 10000);

    this.rules.id.message = this.$t("message.please.enter.your.id");
    this.rules.password.message = this.$t("message.please.enter.your.password");
    this.onClear();
  },
  beforeUnmount() {
    clearInterval(this.timer);
  },
  methods: {
    setLocaleClick(e) {
      let localeValue = e.key;
      if (!localeValue) {
        localeValue = "ko";
      }
      this.setLocale(localeValue);
    },
    setLocale(localeValue) {
      this.$locale = localeValue;
      this.$i18n.locale = localeValue;
      this.language = localeValue;
      sessionStorage.setItem("locale", localeValue);
      //this.loadLanguageAsync(localeValue);
    },
    onClear() {
      sessionStorage.clear();
    },
    onSubmit() {
      let params = new URLSearchParams();
      this.formRef
        .validate()
        .then(() => {
          message.destroy();
          message.loading(this.$t("message.logging"), 60);
          params.append("id", this.formState.id);
          params.append("password", this.formState.password);

          worksApi
            .post("/api/login", params)
            .then((response) => {
              //console.log(response);
              if (
                response.status === 200 &&
                response.data.result.login === true
              ) {
                store.dispatch("loginCommit", response.data);
                sessionStorage.setItem("token", response.data.result.token);
                sessionStorage.setItem(
                  "username",
                  response.data.result.username
                );
                sessionStorage.setItem("isAdmin", response.data.result.isAdmin);
                if (
                  response.data.result.username.toLowerCase() ===
                  "administrator"
                ) {
                  router.push({ name: "Dashboard" });
                } else {
                  router.push({ name: "Favorite" });
                }
                message.destroy();
                message.success(this.$t("message.login.completed"));
              } else {
                message.destroy();
                message.error(this.$t("message.login.wrong"));
              }
            })
            .catch((error) => {
              message.destroy();
              console.log(error);
              if (error.response.status === 400) {
                message.error(this.$t("message.login.wrong.400"));
              } else if (error.response.status === 401) {
                message.error(this.$t("message.login.wrong.401"));
              } else {
                message.error(this.$t("message.login.wrong"));
              }
            });
        })
        .catch((error) => {
          console.log("error", error);
        });
    },
    async serverCheck() {
      await worksApi
        .get("/api/serverCheck")
        .then((response) => {
          this.statusMessage = this.$t("message.status.check.worksapi");
          setTimeout(() => {
            if (response.status === 200) {
              this.statusDescription = this.$t(
                "message.status.check.desc.worksapi.ok"
              );

              setTimeout(() => {
                this.statusMessage = this.$t("message.status.check.moldapi");
                if (response.data.result["Mold"] === 200) {
                  this.statusDescription = this.$t(
                    "message.status.check.desc.moldapi.ok"
                  );

                  setTimeout(() => {
                    this.statusMessage = this.$t("message.status.check.dc");
                    if (response.data.result["Works-DC"] === 200) {
                      this.statusDescription = this.$t(
                        "message.status.check.desc.dc.ok"
                      );
                      setTimeout(() => {
                        this.statusMessage = this.$t("message.status.check.ad");
                        if (response.data.result["Works-Samba"] === 200) {
                          this.statusDescription = this.$t(
                            "message.status.check.desc.ad.ok"
                          );

                          setTimeout(() => {
                            //정상일 경우
                            this.statusMessage = this.$t(
                              "message.status.check.all.ok"
                            );
                            this.statusDescription = this.$t(
                              "message.status.check.desc.all.ok"
                            );
                            this.statusType = "success";
                            clearInterval(this.timer);
                            setTimeout(() => {
                              this.disabled = false;
                              this.serverStatus = false;
                            }, 1500);
                          }, 500);
                        } else {
                          this.statusDescription = this.$t(
                            "message.status.check.desc.ad.fail"
                          );
                          this.statusType = "error";
                          return false;
                        }
                      }, 200);
                    } else {
                      this.statusDescription = this.$t(
                        "message.status.check.desc.dc.fail"
                      );
                      this.statusType = "error";
                      return false;
                    }
                  }, 200);
                } else {
                  this.statusDescription = this.$t(
                    "message.status.check.desc.moldapi.fail"
                  );
                  this.statusType = "error";
                  return false;
                }
              }, 200);
            } else {
              this.statusDescription = this.$t(
                "message.status.check.desc.worksapi.fail"
              );
              this.statusType = "error";
              return false;
            }
          }, 200);
        })
        .catch((error) => {
          this.statusMessage = this.$t("message.status.check.worksapi");
          this.statusDescription = this.$t(
            "message.status.check.desc.worksapi.fail"
          );
          this.statusType = "error";
          console.log(error.message);
        })
        .finally(() => {});
    },
  },
});
</script>

<style lang="less" scoped>
.user-layout {
  height: 100%;

  button.login-button {
    margin-top: 8px;
    padding: 0 15px;
    font-size: 16px;
    height: 40px;
    width: 100%;
  }

  .user-login-other {
    text-align: left;
    margin-top: 24px;
    line-height: 22px;

    .item-icon {
      font-size: 24px;
      color: rgba(0, 0, 0, 0.2);
      margin-left: 16px;
      vertical-align: middle;
      cursor: pointer;
      transition: color 0.3s;

      &:hover {
        color: #1890ff;
      }
    }

    .register {
      float: right;
    }
  }
}
.user-layout-container {
  padding: 3rem 0;
  width: 100%;
}

.user-layout-login {
  min-width: 260px;
  width: 368px;
  margin: 0 auto;
}

.user-layout-logo {
  width: 600px;
  height: 80px;
  margin: 0 auto 2rem;
  border-style: none;
  display: block;
}
// .user-layout-logo {
//   margin: 0 auto 2rem;
//   border-style: none;
//   display: block;
// }
.login-button {
  margin-top: 8px;
  padding: 0 15px;
  font-size: 16px;
  height: 40px;
  width: 100%;
}

.login-ico {
  font-size: 30px;
}
</style>
