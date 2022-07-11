<template>
  <div id="content-layout">
    <a-layout>
      <a-layout-header id="content-header">
        <div id="content-header-body">
          <a-row id="content-header-row">
            <!-- 오른쪽 경로 -->
            <a-col id="content-path" :span="12">
              <Apath :paths="[$t('label.users')]" />
              <a-button
                shape="round"
                style="margin-left: 20px"
                size="small"
                @click="refresh()"
              >
                <template #icon><ReloadOutlined /></template>
                {{ $t("label.refresh") }}
              </a-button>
            </a-col>
            <!-- 왼쪽 액션 -->
            <a-col id="content-action" :span="12">
              <div>
                <Actions
                  v-if="actionFrom === 'ACList'"
                  :action-from="actionFrom"
                  :multi-select-list="multiSelectList"
                  @fetchData="refresh"
                />
                <!-- <a-tooltip placement="bottom">
                  <template #title>{{
                    $t("tooltip.account.download.csv")
                  }}</template>
                  <a-button
                    v-show="dataLoadSucc"
                    shape="circle"
                    style="margin-left: 10px"
                    @click="downloadCSV"
                  >
                    <DownloadOutlined />
                  </a-button>
                </a-tooltip>
                <a-tooltip placement="bottom">
                  <template #title>{{
                    $t("tooltip.account.upload.csv")
                  }}</template>
                  <a-button
                    shape="circle"
                    style="margin-left: 10px"
                    @click="showUploadModal(true)"
                  >
                    <UploadOutlined />
                  </a-button>
                </a-tooltip> -->
                <a-button
                  type="primary"
                  shape="round"
                  style="margin-left: 10px"
                  @click="showAddModal(true)"
                  >{{ $t("label.user.add") }}
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
          <AccountList
            ref="listRefreshCall"
            @actionFromChange="actionFromChange"
            @downloadListSetting="downloadListSetting"
          />
        </div>
      </a-layout-content>
    </a-layout>
    <!-- ADD MODAL START  -->
    <a-modal v-model:visible="addVisible" :title="$t('label.user.add')">
      <template #footer>
        <a-button key="close" @click="showAddModal(false)">{{
          $t("label.cancel")
        }}</a-button>
        <a-button
          key="submit"
          type="primary"
          :loading="addLoading"
          @click="putUser"
          >{{ $t("label.ok") }}</a-button
        >
      </template>
      <a-form
        ref="formRef"
        :model="formState"
        :rules="rules"
        layout="vertical"
        autocomplete="off"
      >
        <a-form-item has-feedback name="account" :label="$t('label.account')">
          <a-input
            v-model:value="formState.account"
            :placeholder="$t('tooltip.user.account')"
          />
        </a-form-item>
        <a-row :gutter="12">
          <a-col :md="24" :lg="12">
            <a-form-item
              has-feedback
              name="lastName"
              :label="$t('label.lastname')"
            >
              <a-input
                v-model:value="formState.lastName"
                :placeholder="$t('tooltip.user.lastname')"
                class="addmodal-aform-item-div"
              />
            </a-form-item>
          </a-col>
          <a-col :md="24" :lg="12">
            <a-form-item
              has-feedback
              name="firstName"
              :label="$t('label.firstname')"
            >
              <a-input
                v-model:value="formState.firstName"
                :placeholder="$t('tooltip.user.firstname')"
              />
            </a-form-item>
          </a-col>
        </a-row>
        <a-row :gutter="12">
          <a-col :md="24" :lg="12">
            <a-form-item
              has-feedback
              name="password"
              :label="$t('label.password')"
              :label-col="{ span: 12 }"
            >
              <a-input-password
                v-model:value="formState.password"
                :placeholder="$t('tooltip.user.password')"
              />
            </a-form-item>
          </a-col>
          <a-col :md="24" :lg="12">
            <a-form-item
              has-feedback
              name="passwordCheck"
              :label="$t('label.passwordCheck')"
              :label-col="{ span: 12 }"
            >
              <a-input-password
                v-model:value="formState.passwordCheck"
                :placeholder="$t('tooltip.user.passwordCheck')"
              />
            </a-form-item>
          </a-col>
        </a-row>
        <a-form-item name="email" :label="$t('label.email')">
          <a-input
            v-model:value="formState.email"
            :placeholder="$t('tooltip.user.email')"
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
        <a-form-item name="department" :label="$t('label.department')">
          <a-input
            v-model:value="formState.department"
            :placeholder="$t('tooltip.user.department')"
            class="addmodal-aform-item-div"
          />
        </a-form-item>
        <a-form-item name="title" :label="$t('label.title')">
          <a-input
            v-model:value="formState.title"
            :placeholder="$t('tooltip.user.title')"
            class="addmodal-aform-item-div"
          />
        </a-form-item>
      </a-form>
    </a-modal>
    <!-- ADD MODAL END  -->
    <a-modal
      v-model:visible="uploadVisible"
      width="800px"
      :title="$t('label.user.csv.upload')"
      :ok-text="$t('label.ok')"
      :cancel-text="$t('label.cancel')"
      @cancel="showUploadModal(false)"
      @ok="handleSubmit()"
    >
      <a-alert
        :message="modalConfirm"
        :description="modalDescription"
        type="info"
        show-icon
      />
      <a-upload-dragger
        :file-list="fileList"
        @remove="handleRemove"
        :before-upload="beforeUpload"
      >
        <p class="ant-upload-drag-icon">
          <inbox-outlined />
        </p>
        <p class="ant-upload-text">
          {{ $t("input.file.upload.dragdrop") }}
        </p>
      </a-upload-dragger>
      <br />
      <a-table
        size="small"
        :loading="uploadCsvTableLoading"
        :columns="UserListColumns"
        :pagination="pagination"
        :data-source="dataList"
        :scroll="{ y: 780 }"
      >
      </a-table>
    </a-modal>
  </div>
</template>

<script>
import Actions from "@/components/Actions";
import Apath from "@/components/Apath";
import AccountList from "@/views/account/AccountList";
import { defineComponent, ref, reactive } from "vue";
export default defineComponent({
  components: {
    AccountList,
    Apath,
    Actions,
  },
  props: {},
  setup() {
    const addVisible = ref(false);
    const uploadVisible = ref(false);
    const checkDupl = ref(false);
    const showAddModal = (state) => {
      addVisible.value = state;
    };
    const showUploadModal = (state) => {
      uploadVisible.value = state;
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
      phone: "",
      title: "",
      department: "",
    });
    let validateAccount = async (rule, value) => {
      let lengthCheck = value.length > rule.max ? true : false; //길이체크
      // let containsEng = /[a-zA-Z]/.test(value); // 대소문자
      // let containsEngUpper = /[A-Z]/.test(value); 대문자
      // let containsNumber = /[0-9]/.test(value);
      let containsSpecial = /[~!@#$%^&*()\-_+|<>?:{}]/.test(value);
      let containsHangle = /[ㄱ-ㅎ|ㅏ-ㅣ|가-힣]/.test(value);
      let containsEngNum = /^[a-zA-Z0-9]+$/.test(value);
      if (
        value === "" ||
        containsHangle ||
        lengthCheck ||
        containsSpecial ||
        !containsEngNum
      ) {
        return Promise.reject();
      } else {
        return Promise.resolve();
      }
    };
    let validateName = async (rule, value) => {
      let lengthCheck = value.length > rule.max ? true : false; //길이체크
      // let containsEng = /[a-zA-Z]/.test(value); // 대소문자
      // let containsEngUpper = /[A-Z]/.test(value); 대문자
      // let containsNumber = /[0-9]/.test(value);
      // let containsSpecial = /[~!@#$%^&*()\-_+|<>?:{}]/.test(value);
      // let containsHangle = /[ㄱ-ㅎ|ㅏ-ㅣ|가-힣]/.test(value);
      let containsHanEngNum = /^[ㄱ-ㅎㅏ-ㅣ가-힣a-zA-Z0-9]+$/.test(value);
      if (value === "" || lengthCheck || !containsHanEngNum) {
        return Promise.reject();
      } else {
        return Promise.resolve();
      }
    };
    let validatePass = async (rule, value) => {
      let lengthCheck = value.length >= rule.min ? true : false; //길이체크
      let containsEng = /[a-zA-Z]/.test(value); // 대소문자
      //let containsEngUpper = /[A-Z]/.test(value); 대문자
      let containsNumber = /[0-9]/.test(value);
      let containsSpecial = /[~!@#$%^&*()\-_+|<>?:{}]/.test(value);
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
        return Promise.reject();
      }
    };
    const rules = {
      account: [
        { required: true, validator: validateAccount, trigger: "change" },
        { max: 32 },
      ],
      firstName: [
        { required: true, validator: validateName, trigger: "change" },
        { max: 32 },
      ],
      lastName: [
        { required: true, validator: validateName, trigger: "change" },
        { max: 32 },
      ],
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
        required: false,
        type: "email",
      },
      phone: {
        required: false,
        pattern: /^\d{2,3}-\d{3,4}-\d{4}$/,
      },
      title: [{ required: false }, { max: 32 }],
      department: [{ required: false }, { max: 32 }],
    };
    return {
      formRef,
      formState,
      rules,
      addVisible,
      uploadVisible,
      showAddModal,
      showUploadModal,
      checkDupl,
    };
  },
  data() {
    return {
      addLoading: ref(false),
      actionFrom: ref("AC"),
      multiSelectList: ref([]),
      fileList: ref([]),
      dataList: ref([]),
      downloadList: ref([]),
      dataLoadSucc: ref(false),
      uploadCsvTableLoading: ref(false),
      modalConfirm: this.$t("message.account.file.upload.confirm"),
      modalDescription: this.$t("message.account.file.upload.description"),
      pagination: {
        // pageSize: 10,
        showSizeChanger: true, // display can change the number of pages per page
        pageSizeOptions: ["10", "20", "50", "100", "200"], // number of pages per option
        showTotal: (total) =>
          this.$t("label.total") + ` ${total}` + this.$t("label.items"), // show total
        // showSizeChange: (current, pageSize) => (this.pageSize = pageSize), // update display when changing the number of pages per page
      },
      UserListColumns: [
        {
          title: this.$t("label.account"),
          dataIndex: "계정",
          key: "계정",
          width: "15%",
        },
        {
          title: this.$t("label.lastname"),
          dataIndex: "이름(성)",
          key: "이름(성)",
          width: "10%",
        },
        {
          title: this.$t("label.firstname"),
          dataIndex: "이름(명)",
          key: "이름(명)",
          width: "10%",
        },
        {
          title: this.$t("label.title"),
          dataIndex: "직급",
          key: "직급",
          width: "10%",
        },
        {
          title: this.$t("label.department"),
          dataIndex: "부서",
          key: "부서",
          width: "15%",
        },
        {
          title: this.$t("label.phone"),
          dataIndex: "전화번호",
          key: "전화번호",
          width: "20%",
        },
        {
          title: this.$t("label.email"),
          dataIndex: "이메일",
          key: "이메일",
          width: "20%",
        },
      ],
    };
  },
  created() {
    this.rules.account[0].message = this.$t("input.user.account");
    this.rules.firstName[0].message = this.$t("input.user.firstname");
    this.rules.lastName[0].message = this.$t("input.user.lastname");
    this.rules.password.message = this.$t("input.user.password");
    this.rules.passwordCheck.message = this.$t("input.user.passwordCheck");
    this.rules.email.message = this.$t("input.user.email");
    this.rules.phone.message = this.$t("input.user.phone");
    this.rules.title[0].message = this.$t("input.user.title");
    this.rules.department[0].message = this.$t("input.user.department");

    this.rules.account[1].message = this.$t("input.max.32");
    this.rules.title[1].message = this.$t("input.max.32");
    this.rules.firstName[1].message = this.$t("input.max.32");
    this.rules.lastName[1].message = this.$t("input.max.32");
    this.rules.department[1].message = this.$t("input.max.32");
  },
  methods: {
    handleRemove(file) {
      const index = this.fileList.indexOf(file);
      const newFileList = this.fileList.slice();
      newFileList.splice(index, 1);
      this.fileList = newFileList;
      this.dataList = null;
    },
    beforeUpload(file) {
      this.dataList = null;
      this.uploadCsvTableLoading = true;

      this.fileList = [file];
      let thisJson = "";
      const reader = new FileReader();
      reader.readAsText(file);
      reader.onload = () => {
        let lines = reader.result.split("\r");
        for (let i = 0; i < lines.length; i++) {
          lines[i] = lines[i].replace(/\s/, ""); //빈 공간 삭제하기
        }
        let result = [];
        let headers = lines[0].split(",");

        for (let i = 1; i < lines.length; i++) {
          let obj = {};
          let currentline = lines[i].split(",");

          for (let j = 0; j < headers.length; j++) {
            obj[headers[j]] = currentline[j];
          }
          result.push(obj);
        }
        thisJson = result;
      };

      setTimeout(() => {
        this.dataList = thisJson;
        this.uploadCsvTableLoading = false;
      }, 2000);

      return false;
    },
    downloadListSetting(obj) {
      this.downloadList = obj;
      this.dataLoadSucc = true;
    },
    downloadCSV() {
      const date = new Date();

      // "\ufeff" => 한글 깨짐 방지
      let csv = "\ufeff" + "계정,이름(성),이름(명),직급,부서,전화번호,이메일\n";
      this.downloadList.forEach((el) => {
        var line =
          el["name"] +
          "," +
          el["givenName"] +
          "," +
          el["sn"] +
          "," +
          el["title"] +
          "," +
          el["department"] +
          "," +
          el["telephoneNumber"] +
          "," +
          el["mail"] +
          "\n";
        csv += line;
      });
      var blob = new Blob([csv], { type: "text/csv;charset=utf-8;" });
      let link = document.createElement("a");
      link.href = window.URL.createObjectURL(blob);
      link.download =
        "사용자 목록(" +
        new Date(+new Date() + 3240 * 10000)
          .toISOString()
          .replace("T", " ")
          .replace(/\..*/, "") +
        ").csv";
      link.click();
    },
    refresh() {
      this.$refs.listRefreshCall.fetchRefresh();
    },
    actionFromChange(val, obj) {
      this.actionFrom = "AC";
      setTimeout(() => {
        this.actionFrom = val;
        this.multiSelectList = obj;
      }, 100);
    },
    putUser() {
      let params = new URLSearchParams();
      params.append("username", this.formState.account);
      params.append("firstName", this.formState.firstName);
      params.append("lastName", this.formState.lastName);
      params.append("password", this.formState.password);
      params.append("email", this.formState.email);
      params.append("phone", this.formState.phone);
      params.append("title", this.formState.title);
      params.append("department", this.formState.department);

      //console.log(params);
      this.formRef
        .validate()
        .then(() => {
          this.addLoading = true;
          this.$worksApi //이름 중복 확인 체크
            .get("/api/v1/user/" + this.formState.account)
            .then((response) => {
              if (response.status === 200) {
                //중복일 때
                this.$message.error(this.$t("message.name.dupl"));
              }
            })
            .catch((error) => {
              //중복 이름 없을 때
              this.$message.loading(this.$t("message.user.createing"), 1);
              this.$worksApi
                .put("/api/v1/user", params)
                .then((response) => {
                  //console.log(response.status);
                  if (response.status === 200) {
                    this.$message.loading(
                      this.$t("message.user.create.success"),
                      1
                    );
                  } else {
                    this.$message.error(this.$t("message.user.create.fail"));
                  }
                })
                .catch((error) => {
                  this.$message.error(this.$t("message.user.create.fail"));
                  console.log("error", error);
                })
                .finally(() => {
                  this.fetch();
                });
            });
        })
        .catch((error) => {
          console.log("error", error);
          //message.error(error);
        });
    },
    fetch() {
      this.addLoading = false;
      this.showAddModal(false);
      this.showUploadModal(false);
      this.$refs.listRefreshCall.fetchRefresh();
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
