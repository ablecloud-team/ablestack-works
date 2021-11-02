<template>
  <div id="content-layout">
    <a-layout>
      <a-layout-header id="content-header">
        <div id="content-header-body">
          <a-row id="content-header-row">
            <!-- 왼쪽 경로 -->
            <a-col id="content-path" :span="12">
              <Apath :paths="[$t('label.workspace')]" />
              <a-button
                shape="round"
                style="margin-left: 20px; height: 30px"
                @click="reflesh()"
              >
                <template #icon>
                  <ReloadOutlined /> {{ $t("label.reflesh") }}
                </template>
              </a-button>
            </a-col>
            <!-- 우측 액션 -->
            <a-col id="content-action" :span="12">
              <div>
                <Actions
                  v-if="actionFrom === 'WorkspaceList'"
                  :action-from="actionFrom"
                />
                <a-button
                  type="primary"
                  shape="round"
                  style="margin-left: 10px"
                  @click="showModal(true)"
                  >{{ addModalTitle }}
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
          <WorkSpaceList
            ref="listRefleshCall"
            @actionFromChange="actionFromChange"
          />
        </div>
      </a-layout-content>
    </a-layout>
    <!-- ADD WORKSPACE MODAL START  -->
    <a-modal v-model:visible="visible" :title="addModalTitle">
      <template #title> </template>
      <template #footer>
        <a-button key="close" @click="showModal(false)">{{
          $t("label.cancel")
        }}</a-button>
        <a-button
          key="submit"
          type="primary"
          :loading="loading"
          @click="putWorkspace"
          >{{ $t("label.ok") }}</a-button
        >
      </template>
      <a-form
        ref="formRef"
        :model="formState"
        :rules="rules"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
        layout="vertical"
      >
        <!--워크스페이스 이름 start-->
        <a-form-item has-feedback name="name" :label="$t('label.name')">
          <a-input
            v-model:value="formState.name"
            :placeholder="$t('tooltip.workspace.name')"
          />
        </a-form-item>
        <!--워크스페이스 이름 end-->
        <!--워크스페이스 설명 start-->
        <a-form-item
          has-feedback
          name="description"
          :label="$t('label.description')"
        >
          <a-input
            v-model:value="formState.description"
            :placeholder="$t('tooltip.workspace.description')"
            class="addmodal-aform-item-div"
          />
        </a-form-item>
        <!--워크스페이스 설명 end-->
        <!--워크스페이스 타입 start-->
        <a-form-item
          name="workspaceType"
          :label="$t('label.type')"
        >
          <!-- <a-select
            v-model:value="formState.workspaceType"
            :placeholder="$t('tooltip.workspace.type')"
            @change="workspaceTypeChange"
          >
            <a-select-option value="desktop">Desktop</a-select-option>
            <a-select-option value="application">Application</a-select-option>
          </a-select> -->

          <a-radio-group
            v-model:value="formState.workspaceType"
            button-style="solid"
            @change="selected => { workspaceTypeChange(selected) }"
          >
            <a-radio-button value="desktop">{{
              $t("label.desktop")
            }}</a-radio-button>
            <a-radio-button value="application">{{
              $t("label.application")
            }}</a-radio-button>
          </a-radio-group>
        </a-form-item>
        <!--워크스페이스 타입 end-->
        <!--워크스페이스 전용 여부 start-->
        <a-form-item
          v-show="formState.desktopBoolean"
          name="dedicatedOrSharedBoolean"
          :label="$t('label.dedicated.shared')"
        >
          <!-- {{ $t(switchLabel) }}
            <a-switch
              v-model:checked="formState.dedicatedOrSharedBoolean"
              @change="dedicatedChange"
            /> -->

          <a-radio-group
            v-model:value="formState.dedicatedOrSharedBoolean"
            button-style="solid"
          >
            <a-radio-button :value="false">{{
              $t("label.dedicated")
            }}</a-radio-button>
            <a-radio-button :value="true">{{
              $t("label.shared")
            }}</a-radio-button>
          </a-radio-group>
        </a-form-item>
        <!--워크스페이스 전용 여부 end-->
        <!--워크스페이스 템플릿 start-->
        <a-form-item
          has-feedback
          name="selectedMasterTemplateId"
          :label="$t('label.mastertemplate')"
        >
          <a-select
            v-model:value="formState.selectedMasterTemplateId"
            :placeholder="$t('tooltip.workspace.template')"
            show-search
            option-filter-prop="label"
            class="addmodal-aform-item-div"
            
          >
            <a-select-option
              v-for="option in templates"
              :key="option.name"
              :value="option.id"
              :label="option.name"
              
            >
              {{ option.name + '(' + option.version +')' }}
            </a-select-option>
          </a-select>
        </a-form-item>
        <!--워크스페이스 템플릿 end-->
        <!--워크스페이스 컴퓨트 오퍼링 start-->
        <a-form-item
          has-feedback
          name="selectedOfferingId"
          :label="$t('label.compute.offering')"
        >
          <a-select
            v-model:value="formState.selectedOfferingId"
            :placeholder="$t('tooltip.workspace.compute.offering')"
            show-search
            option-filter-prop="label"
            class="addmodal-aform-item-div"
          >
            <a-select-option
              v-for="option in offerings"
              :key="option.name"
              :value="option.id"
              :label="option.name"
            >
              {{ option.name }}
            </a-select-option>
          </a-select>
        </a-form-item>
        <!--워크스페이스 컴퓨트 오퍼링 end-->
      </a-form>
    </a-modal>
    <!-- ADD WORKSPACE MODAL END  -->
  </div>
</template>

<script>
import Actions from "@/components/Actions";
import Apath from "@/components/Apath";
import WorkSpaceList from "./WorkSpaceList";
import { defineComponent, onMounted, reactive, ref, toRaw } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";

export default defineComponent({
  name: "WorkSpace",
  components: {
    WorkSpaceList,
    Apath,
    Actions,
  },
  props: {},
  setup() {
    const visible = ref(false);
    const showModal = (state) => {
      visible.value = state;
    };
    const formRef = ref();
    const formState = reactive({
      name: ref(""),
      description: ref(""),
      selectedMasterTemplateId: ref(""),
      selectedOfferingId: ref(""),
      workspaceType: ref("desktop"),
      dedicatedOrSharedBoolean: ref(false),
      desktopBoolean: ref(true),
    });
    let validateName = async (rule, value) => {
      let lengthCheck = value.length > rule.max ? true : false; //길이체크
      //let containsEng = /[a-zA-Z]/.test(value); // 대소문자
      //let containsEngUpper = /[A-Z]/.test(value); 대문자
      //let containsNumber = /[0-9]/.test(value);
      //let containsSpecial = /[~!@#$%^&*()_+|<>?:{}]/.test(value);
      let containsHangle = /[ㄱ-ㅎ|ㅏ-ㅣ|가-힣]/.test(value);
      if (value === "" || containsHangle || lengthCheck) {
        return Promise.reject();
      } else {
        return Promise.resolve();
      }
    };
    const rules = {
      name: {
        required: true,
        max: 7,
        validator: validateName,
        trigger: "change",
      },
      description: {
        required: true,
        max: 32,
      },
      workspaceType: { required: true },
      dedicatedOrSharedBoolean: { required: true },
      selectedMasterTemplateId: { required: true },
      selectedOfferingId: { required: true },
    };
    return {
      templates: ref([]),
      templateDisabled: ref(true),
      desktopTemplates: ref([]),
      applicationTemplates: ref([]),
      offerings: ref([]),
      labelCol: { span: 10 },
      wrapperCol: { span: 40 },
      formRef,
      formState,
      rules,
      switchLabel: ref("label.dedicated"),
      visible,
      showModal,
    };
  },
  data() {
    return {
      addModalTitle: this.$t("label.workspace.add"),
      actionFrom: ref("Workspace"),
    };
  },
  created() {
    this.fetchOfferingsAndTemplates();
    this.rules.name.message = this.$t("input.workspace.name");
    this.rules.description.message = this.$t("input.workspace.description");
    this.rules.workspaceType.message = this.$t("input.workspace.workspaceType");
    this.rules.selectedMasterTemplateId.message = this.$t("input.workspace.selectedTemplateId");
    this.rules.selectedOfferingId.message = this.$t("input.workspace.selectedOfferingId");
  },
  methods: {
    reflesh() {
      this.$refs.listRefleshCall.fetchData();
    },
    actionFromChange(val) {
      this.actionFrom = ref(val);
    },
    dedicatedChange(value) {
      this.dedicatedOrSharedBoolean = value;
      if (this.dedicatedOrSharedBoolean) {
        this.switchLabel = ref("label.shared");
        //console.log("true");
      } else {
        this.switchLabel = ref("label.dedicated");
        //console.log("false");
      }
    },
    workspaceTypeChange(value) {
      this.formState.selectedMasterTemplateId = ref("");
      if (value.target.value === "application") {
        this.formState.desktopBoolean = ref(false);
        this.templates = this.applicationTemplates;
      } else {
        this.formState.desktopBoolean = ref(true);
        this.templates = this.desktopTemplates;
      }
    },
    putWorkspace() {
      this.rules.name.message = this.$t("input.workspace.name");
      this.rules.description.message = this.$t("input.workspace.description");
      this.rules.workspaceType.message = this.$t("input.workspace.workspaceType");
      this.rules.selectedMasterTemplateId.message = this.$t("input.workspace.selectedTemplateId");
      this.rules.selectedOfferingId.message = this.$t("input.workspace.selectedOfferingId");

      // 실제 template uuid 를 넘겨주기 위함.
      var realTemplateId = "";
      var masterTemplateName = "";
      for (let str of this.templates) {
        if(str.id === this.formState.selectedMasterTemplateId){
          realTemplateId = str.templateId;
          masterTemplateName =str.name + "(" +str.version + ")";
          break;
        }
      }
      //console.log(this.formState.selectedMasterTemplateId + " :: " + realTemplateId + " :: " + masterTemplateName);
      let params = new URLSearchParams();
      params.append("name", this.formState.name);
      params.append("description", this.formState.description);
      params.append("type", this.formState.workspaceType);
      params.append("shared", this.formState.dedicatedOrSharedBoolean);
      params.append("masterTemplateName", masterTemplateName);
      params.append("templateUuid", realTemplateId);
      params.append("computeOfferingUuid", this.formState.selectedOfferingId);
      //console.log(params);
      this.formRef
        .validate()
        .then(() => {
          worksApi
            .get("/api/v1/group/" + this.formState.name) //이름 중복 확인
            .then((response) => {
              if (response.status === 200) {
                //이름 중복일때 메시지 확인
                message.error(this.$t("message.name.dupl"));
              }
            })
            .catch((error) => {
              //이름 중복이 아닐때(status code = 401)
              message.loading(this.$t("message.workspace.createing"), 1);
              worksApi
                .put("/api/v1/workspace", params)
                .then((response) => {
                  if (response.status === 200) {
                    message.loading(
                      this.$t("message.workspace.create.success"),1
                    );
                  } else {
                    message.error(this.$t("message.workspace.create.fail"));
                  }
                  this.showModal(false);
                  setTimeout(() => {
                    this.$refs.listRefleshCall.fetchData();
                  }, 1500);
                })
                .catch(function (error) {
                  message.error(error);
                  //console.log(error);
                });
            });
        })
        .catch((error) => {
          console.log("error", error);
          //message.error(error);
        });
    },
    fetchOfferingsAndTemplates() {
      worksApi
        .get("/api/v1/offering")
        .then((response) => {
          if (response.status == 200) {
            this.offerings = response.data.serviceOfferingList.listserviceofferingsresponse.serviceoffering;
            const masterTemplates = response.data.templateList.listdesktopmasterversionsresponse.desktopmasterversion;

            for (let str of masterTemplates) {
              if(str.mastertemplatetype === "DESKTOP"){
                this.desktopTemplates.push({ id: str.id, templateId: str.templateid, name: str.name, version: str.version });
                this.templates = this.desktopTemplates; //기본값 desktop용 template 목록 세팅
              } else if(str.mastertemplatetype === "APP"){
                this.applicationTemplates.push({ id: str.id, templateId: str.templateid, name: str.name, version: str.version });
              }
            }
          } else {
            //console.log(response.message);
          }
        })
        .catch(function (error) {
          message.error(error.message);
          //console.log(error);
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
