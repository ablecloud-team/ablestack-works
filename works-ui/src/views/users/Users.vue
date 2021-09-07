<template>
  <div id="content-layout">
    <a-layout>
      <a-layout-header id="content-header">
        <div id="content-header-body">
          <a-row id="content-header-row">
            <!-- 오른쪽 경로 -->
            <a-col id="content-path" :span="12">
              <Apath v-bind:paths="[$t('label.users')]" />
            </a-col>

            <!-- 왼쪽 액션 -->
            <a-col id="content-action" :span="12">
              <div>
                <a-button type="primary" shape="round" @click="showModal(true)">{{ addModalTitle }}
                  <template #icon>
                      <PlusOutlined />
                    </template>
                </a-button>
              </div>
            </a-col>
          </a-row>
        </div>
      </a-layout-header>
      <a-layout-content>
        <div id="content-body">
          <UsersList
            :data="userDataList" :bordered="false" />
        </div>
      </a-layout-content>
    </a-layout>
    <!-- ADD WORKSPACE MODAL START  -->
    <a-modal v-model:visible="visible" :title="addModalTitle" >
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
        <a-form-item name="name" :label="$t('label.name')">
          <a-input
            v-model:value="formState.name"
            :placeholder="$t('tooltip.user.name')"/>
        </a-form-item>
        <!--워크스페이스 이름 end-->
        <!--워크스페이스 설명 start-->
        <a-form-item name="description" :label="$t('label.description')">
          <a-input
            v-model:value="formState.description"
            :placeholder="$t('tooltip.user.description')"
            class="addmodal-aform-item-div"
          />
        </a-form-item>
        <a-form-item has-feedback name="password" :label="$t('label.password')" autocomplete="off">
          <a-input
            type="password"
            v-model:value="formState.password"
            :placeholder="$t('tooltip.user.password')"/>
        </a-form-item>
        <a-form-item has-feedback name="passwordCheck" :label="$t('label.passwordCheck')" autocomplete="off" >
          <a-input
            type="password"
            v-model:value="formState.passwordCheck"
            :placeholder="$t('tooltip.user.passwordCheck')"/>
        </a-form-item>
        <a-form-item name="email" :label="$t('label.email')">
          <a-input
            v-model:value="formState.email"
            :placeholder="$t('tooltip.user.email')"
            class="addmodal-aform-item-div"
          />
        </a-form-item>
        <a-form-item name="userGroup" :label="$t('label.userGroup')">
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
        <a-form-item name="phone" :label="$t('label.phone')">
          <a-input
            v-model:value="formState.phone"
            :placeholder="$t('tooltip.user.phone')"
            class="addmodal-aform-item-div"
          />
        </a-form-item>


        
        <!--워크스페이스 비밀번호 확인 end-->
      </a-form>
    </a-modal>
    <!-- ADD WORKSPACE MODAL END  -->
  </div>
</template>

<script>
import Apath from "@/components/Apath";
import UsersList from "@/views/users/UsersList";
import { defineComponent, ref, reactive } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";

export default defineComponent({
  props: {
  },
  components: { 
    UsersList,
    Apath
  },
  data(){
    return {
      addModalTitle: this.$t("label.user.add"),
      userDataList : [
          {"name":"user01", "uuid":"123123123123123123123123", "state":"Allocated", "desktop":"Desktop1"}
        ],
      
    }
  },
  setup() {
    const visible = ref(false);
    const showModal = (state) => {
      visible.value = state;
    };
    const formRef = ref();
    const formState = reactive({
      name: '',
      description: '',
      password: '',
      passwordCheck: '',
      email: '',
      userGroup: '',
      department: '',
      phone: '',
    });
    let validatePass = async (rule, value) => {
      if (value === '') {
        return Promise.reject();
      } else {
        if (formState.passwordCheck !== '') {
          formRef.value.validateFields('passwordCheck');
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
      name: { required: true, },
      description: { required: true, },
      password: { 
        required: true,
        validator: validatePass,
        trigger: 'change',
      },
      passwordCheck: {
        required: true,
        validator: validatePass2,
        trigger: 'change',
      },
      email: { required: true, },
      userGroup: { required: true, },
      department: { required: true, },
      phone: { required: true, },
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
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      this.rules.name.message = this.$t('input.user.name');
      this.rules.description.message = this.$t('input.user.description');
      this.rules.password.message = this.$t('input.user.password');
      this.rules.passwordCheck.message = this.$t('input.user.passwordCheck');
      this.rules.email.message = this.$t('input.user.email');
      this.rules.userGroup.message = this.$t('input.user.userGroup');
      this.rules.department.message = this.$t('input.user.department');
      this.rules.phone.message = this.$t('input.user.phone');

      // worksApi
      //   .get("/api/v1/users", { withCredentials: true })
      //   .then((response) => {
      //     if (response.data.result.status == 200) {
      //       this.userDataList = response.data.result.list;
      //     } else {
      //       message.error(this.t('message.response.data.fail'));
      //       //console.log(response.message);
      //     }
      //   })
      //   .catch(function (error) {
      //     message.error(error.message);
      //     //console.log(error);
      //   });
    },

    putUser() {


      let params = new URLSearchParams();
      params.append("name", this.formState.name);
      params.append("description", this.formState.description);
      params.append("password", this.formState.password);
      //console.log(params);
      this.formRef
          .validate()
          .then( () => {
          //console.log(toRaw(formState));
          message.loading(this.$t("message.workspace.createing"), 1);
          // try {
          //   worksApi
          //     .put("/api/v1/user", params, { withCredentials: true })
          //     .then((response) => {
          //       if (response.data.result.status === 200) {
          //         message.loading(this.$t("message.workspace.create.success"), 1);
          //         setTimeout(() => {
          //           location.reload();
          //         }, 1500);
          //       } else {
          //         message.error(response.data.result.deployvirtualmachineresponse.errortext);
          //       }
          //       this.showModal(false);
          //     })
          //     .catch(function (error) {
          //       message.error(error.message);
          //     //console.log(error);
          //     });
          // }catch (error){
          //   console.log(error)
          //   message.error(this.$t("message.workspace.create.fail"))
          // }
        })
        .catch(error => {
          console.log('error', error);
          //message.error(error);
          
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
  padding: 24px;
  height: auto;
}

#content-path {
  text-align: left;
  align-items: center;
  display: flex;
}

#content-action {
  text-align: right;
}
</style>
