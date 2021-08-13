<template>
  <div class="user-layout desktop">
    <div class="user-layout-container">
      <div class="user-layout-container">
        <img
            src="../../assets/ablestack-logo.png"
            alt="logo"
            class="user-layout-logo"
        />
      </div>
      <a-form
          ref="formRef"
          layout="horizontal"
          :model="formState"
          :rules="rules"
          class="user-layout-login"
      >
        <a-form-item name="id">
          <a-input
              v-model:value="formState.id"
              :placeholder="$t('label.user.id')"
              size="large"
          >
            <template #prefix>
              <UserOutlined style="color: rgba(0, 0, 0, 0.25)" />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item name="password">
          <a-input-password
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
          <a-button type="primary" @click="onSubmit" block class="login-button">
            {{ $t("label.login") }}
          </a-button>
        </a-form-item>
        <!--   언어변환 버튼 start     -->
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
          <a-button type="text">
            <template #icon>
              <font-awesome-icon
                  :icon="['fas', 'language']"
                  size="4x"
                  style="color: #666"
                  class="login-ico"
              />
            </template>
          </a-button>
        </a-popover>
        <!--   언어변환 버튼 끝     -->
      </a-form>
    </div>
  </div>
</template>

<script>
import { defineComponent, reactive, ref } from "vue";
import { message } from "ant-design-vue";
import { axiosLogin } from "../../api/index";
import store from "../../store/index"
import router from "../../router";

export default defineComponent({
  name: "Login",
  setup() {
    const formRef = ref();
    const formState = reactive({
      id: "administrator",
      password: "Ablecloud1!",
      // id: "",
      // password: "",
    });
    const rulesIdMeassage = ref();
    const rulesPasswordMeassage = ref();
    const rules = {
      id: {
        required: true,
        message: "message.login.failure",
      },
      password: {
        required: true,
        message: "Please input name",
      },
    };

    return {
      formRef,
      formState,
      rules,
      rulesIdMeassage,
      rulesPasswordMeassage
    };
  },
  data() {
    let rulesIdMeassage = this.$t(`message.please.enter.your.id`)
    let rulesPasswordMeassage = this.$t("message.please.enter.your.password")
    return {
      rulesIdMeassage,
      rulesPasswordMeassage,
    };
  },
  computed: {

  },
  methods: {
    onSubmit() {
      this.rules.id.message = this.rulesIdMeassage;
      this.rules.password.message = this.rulesPasswordMeassage;
      let params = new URLSearchParams();
      let res
      this.formRef
          .validate()
          .then(async () => {
            message.loading(this.$t("message.logging"), 0);
            params.append("id", this.formState.id);
            params.append("password", this.formState.password);
            try {
              res = await axiosLogin(params)
              // console.log(res);
              if (res.status === 200) {
                message.destroy();
                message.success(this.$t("message.login.completed"), 2);
                await store.dispatch("loginCommit", res.data)
                await router.push({name: "home"})
                // console.log("res.data.result");
                // console.log(res.data);
                // console.log(store.state.user.accessToken);
              }
            }catch (error){
              message.destroy();
              //TODO i18n 적용
              console.log(error)
              message.error("로그인에 실패했습니다. 관리자에게 문의해주세요")
            }
          })
          .catch((error) => {
            console.log("error", error);
          });
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
