<template>
  <div id="content-layout">
    <a-layout>
      <a-layout-header id="content-header">
        <div id="content-header-body">
          <a-row id="content-header-row">
            <!-- 오른쪽 경로 -->
            <a-col id="content-path" :span="12">
              <Apath :paths="[$t('label.group.policy')]" />
              <a-button
                shape="round"
                style="margin-left: 20px"
                size="small"
                @click="refresh()"
              >
                <template #icon>
                  <ReloadOutlined /> {{ $t("label.refresh") }}
                </template>
              </a-button>
            </a-col>

            <!-- 왼쪽 액션 -->
            <!-- <a-col id="content-action" :span="12">
              <div>
                <Actions :actionFrom="actionFrom" v-if="actionFrom === 'GroupPolicyList'"/>
                <a-button type="primary" shape="round" style="margin-left: 10px;" @click="showModal(true)">{{ addModalTitle }}
                  <template #icon>
                      <PlusOutlined />
                    </template>
                </a-button>
              </div>
            </a-col> -->
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
          <GroupPolicyList
            ref="listRefreshCall"
            @actionFromChange="actionFromChange"
          />
        </div>
      </a-layout-content>
    </a-layout>
    <!-- ADD WORKSPACE MODAL START  -->
    <!-- <a-modal v-model:visible="visible" :title="addModalTitle" >
      <template #footer>
        <a-button key="close" @click="showModal(false)">{{$t("label.cancel")}}</a-button>
        <a-button key="submit" type="primary" :loading="loading" @click="putUser">{{$t("label.ok")}}</a-button>
      </template>
      <a-form
        ref="formRef"
        :model="formState"
        :rules="rules"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
        layout="vertical"
      >
        <a-form-item has-feedback name="account" :label="$t('label.account')">
          <a-input
            v-model:value="formState.account"
            :placeholder="$t('tooltip.user.account')"/>
        </a-form-item>
        <a-form-item has-feedback name="firstName" :label="$t('label.firstname')">
          <a-input
            v-model:value="formState.firstName"
            :placeholder="$t('tooltip.user.firstname')"/>
        </a-form-item>
        <a-form-item has-feedback name="lastName" :label="$t('label.lastname')">
          <a-input
            v-model:value="formState.lastName"
            :placeholder="$t('tooltip.user.lastname')"
            class="addmodal-aform-item-div"
          />
        </a-form-item>
        <a-form-item has-feedback name="password" :label="$t('label.password')" autocomplete="off">
          <a-input
            type="password"
            v-model:value="formState.password"
            :placeholder="$t('tooltip.user.password')"/>
        </a-form-item>
        <a-form-item has-feedback name="passwordCheck" :label="$t('label.passwordCheck')" autocomplete="off">
          <a-input
            type="password"
            v-model:value="formState.passwordCheck"
            :placeholder="$t('tooltip.user.passwordCheck')"/>
        </a-form-item>
        <a-form-item has-feedback name="email" :label="$t('label.email')">
          <a-input 
            v-model:value="formState.email"
            :placeholder="$t('tooltip.user.email')"
            class="addmodal-aform-item-div"
          />
        </a-form-item> -->
    <!-- <a-form-item has-feedback name="userGroup" :label="$t('label.userGroup')">
          <a-input
            v-model:value="formState.userGroup"
            :placeholder="$t('tooltip.user.userGroup')"
            class="addmodal-aform-item-div"
          />
        </a-form-item>
        <a-form-item name="department" :label="$t('label.department')">
          <a-input
            v-model:value="formState.department"
            :placeholder="$t('tooltip.user.department')"
            class="addmodal-aform-item-div"
          />
        </a-form-item>
        <a-form-item name="phone" :label="$t('label.phone')" >
          <a-input
            v-model:value="formState.phone"
            :placeholder="$t('tooltip.user.phone')"
            class="addmodal-aform-item-div"
          />
        </a-form-item>-->

    <!--워크스페이스 비밀번호 확인 end-->
    <!-- </a-form>
    </a-modal> -->
    <!-- ADD WORKSPACE MODAL END  -->
  </div>
</template>

<script>
//import Actions from "@/components/Actions";
import Apath from "@/components/Apath";
import GroupPolicyList from "./GroupPolicyList";
import { defineComponent, ref, reactive } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";

export default defineComponent({
  components: {
    GroupPolicyList,
    Apath,
    //Actions,
  },
  props: {},
  setup() {
    const visible = ref(false);
    const showModal = (state) => {
      visible.value = state;
    };
    const formRef = ref();
    const formState = reactive({
      account: "",
      firstName: "",
      lastName: "",
      password: "",
      passwordCheck: "",
      email: "",
      userGroup: "",
      department: "",
      phone: "",
    });
    let validatePass = async (rule, value) => {
      let lengthCheck = value.length >= rule.min ? true : false; //길이체크
      let containsEng = /[a-zA-Z]/.test(value); // 대소문자
      //let containsEngUpper = /[A-Z]/.test(value); 대문자
      let containsNumber = /[0-9]/.test(value);
      let containsSpecial = /[~!@#$%^&*()_+|<>?:{}]/.test(value);
      if (
        value === "" ||
        !containsEng ||
        !containsNumber ||
        !containsSpecial ||
        !lengthCheck
      ) {
        return Promise.reject();
      } else {
        if (formState.passwordCheck !== "") {
          formRef.value.validateFields("passwordCheck");
        }
        return Promise.resolve();
      }
    };

    let validatePass2 = async (rule, value) => {
      if (value !== formState.password) {
        return Promise.reject("Two inputs don't match!");
      }
    };
    const rules = {
      account: { required: true },
      firstName: { required: true },
      lastName: { required: true },
      password: {
        min: 7,
        required: true,
        validator: validatePass,
        trigger: "change",
      },
      passwordCheck: {
        required: true,
        validator: validatePass2,
        trigger: "change",
      },
      email: {
        required: true,
        type: "email",
      },
      // userGroup: { required: true, },
      // department: { required: false, },
      // phone: { required:  false, },
    };
    return {
      labelCol: { span: 10 },
      wrapperCol: { span: 40 },
      formRef,
      formState,
      rules,
      visible,
      showModal,
    };
  },
  data() {
    return {
      addModalTitle: this.$t("label.user.add"),
      actionFrom: ref("GroupPolicy"),
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    refresh() {
      this.$refs.listRefreshCall.fetchData();
    },
    actionFromChange(val) {
      //console.log(val);
      this.actionFrom = ref(val);
    },
    fetchData() {
      this.rules.account.message = this.$t("input.user.account");
      this.rules.firstName.message = this.$t("input.user.firstname");
      this.rules.lastName.message = this.$t("input.user.lastname");
      this.rules.password.message = this.$t("input.user.password");
      this.rules.passwordCheck.message = this.$t("input.user.passwordCheck");
      this.rules.email.message = this.$t("input.user.email");
      // this.rules.userGroup.message = this.$t('input.user.userGroup');
      // this.rules.department.message = this.$t('input.user.department');
      // this.rules.phone.message = this.$t('input.user.phone');
    },
    putUser() {
      let params = new URLSearchParams();
      params.append("username", this.formState.account);
      params.append("firstName", this.formState.firstName);
      params.append("lastName", this.formState.lastName);
      params.append("password", this.formState.password);
      params.append("email", this.formState.email);
      //console.log(params);
      this.formRef
        .validate()
        .then(() => {
          message.loading(this.$t("message.user.createing"), 1);
          try {
            worksApi
              .put("/api/v1/user", params)
              .then((response) => {
                if (response.status === 200) {
                  message.loading(this.$t("message.user.create.success"), 1);
                  // setTimeout(() => {
                  //   location.reload();
                  // }, 1500);
                } else {
                  message.error(
                    response.data.result.createuserresponse.errortext
                  );
                }
                this.showModal(false);
                this.$refs.listRefreshCall.fetchData();
              })
              .catch((error) => {
                message.error(this.$t("message.user.create.fail"));
                console.log(error);
              });
          } catch (error) {
            //console.log(error);
            //message.error(this.$t("message.user.create.fail"));
          }
        })
        .catch((error) => {
          console.log("error", error);
        });
    },
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#content-layout .ant-layout-content {
  background: #f0f2f5;
  padding: 10px;
}
#content-layout .ant-layout-header {
  text-align: left;
  background: white;
  border: 1px solid #e8e8e8;
  /*color: #fff;*/
  font-size: 14px;
  line-height: 1.5;
  padding: 20px;
  height: auto;
}

#content-path {
  text-align: left;
  align-items: center;
  display: flex;
  height: 32px;
}

#content-action {
  text-align: right;
}
</style>
