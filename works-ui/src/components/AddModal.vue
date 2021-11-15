<template>
  <a-modal
    v-model:visible="state.modalVisible"
    :footer="null"
    style="width: 100vw; top: 40px"
  >
    <template #title>{{ title }}</template>
    <a-form layout="vertical">
      <!--===============워크스페이스 추가 시작=======================-->
      <a-space
        direction="vertical"
        v-if="state.workspaceBoolean"
        class="addmodal-aform-item-div"
      >
        <!--워크스페이스 이름 start-->
        <a-form-item class="addmodal-aform-item">
          <slot name="label">
            {{ $t("label.name") }}
            <a-tooltip placement="top">
              <template #title>
                {{ $t("tooltip.workspace.name") }}
              </template>
              <InfoCircleOutlined class="addmodal-info-icon" />
            </a-tooltip>
          </slot>
          <a-input v-bind:placeholder="$t('tooltip.workspace.name')" />
        </a-form-item>
        <!--워크스페이스 이름 end-->
        <!--워크스페이스 설명 start-->
        <a-form-item class="addmodal-aform-item">
          <slot name="label">
            {{ $t("label.description") }}
            <a-tooltip placement="top">
              <template #title>
                {{ $t("tooltip.workspace.description") }}
              </template>
              <InfoCircleOutlined class="addmodal-info-icon" />
            </a-tooltip>
          </slot>
          <a-input
            v-bind:placeholder="$t('tooltip.workspace.description')"
            class="addmodal-aform-item-div"
          />
        </a-form-item>
        <!--워크스페이스 설명 end-->
        <!--워크스페이스 타입 start-->
        <a-form-item class="addmodal-aform-item">
          <slot name="label">
            <span>{{ $t("label.type") }}</span>
            <a-tooltip placement="top">
              <template #title>
                {{ $t("tooltip.workspace.type") }}
              </template>
              <InfoCircleOutlined class="addmodal-info-icon" />
            </a-tooltip>
          </slot>
          <a-select
            v-model:value="workspaceType"
            v-bind:placeholder="$t('tooltip.workspace.type')"
            show-search
            class="addmodal-aform-item-div"
            @change="workspaceTypeChange"
          >
            <a-select-option value="desktop">Desktop</a-select-option>
            <a-select-option value="application">Application</a-select-option>
          </a-select>
        </a-form-item>
        <!--워크스페이스 타입 end-->
        <!--워크스페이스 템플릿 start-->
        <a-form-item class="addmodal-aform-item">
          <slot name="label">
            {{ $t("label.template") }}
            <a-tooltip placement="top">
              <template #title>
                {{ $t("tooltip.workspace.template") }}
              </template>
              <InfoCircleOutlined class="addmodal-info-icon" />
            </a-tooltip>
          </slot>
          <a-select
            v-model:value="templateId"
            v-bind:placeholder="$t('tooltip.workspace.type')"
            show-search
            class="addmodal-aform-item-div"
          >
            <a-select-option value="템플릿1">템플릿1</a-select-option>
            <a-select-option value="템플릿2">템플릿2</a-select-option>
          </a-select>
        </a-form-item>
        <!--워크스페이스 템플릿 end-->
        <!--워크스페이스 전용 여부 start-->
        <a-form-item v-show="state.desktopBoolean" class="addmodal-aform-item">
          <slot name="label">
            {{ $t(state.switchLabel) }}
            <!--          <span v-show>{{ $t('label.shared') }}</span>-->
            <a-tooltip placement="top">
              <template #title>
                {{ $t("tooltip.workspace.type") }}
              </template>
              <InfoCircleOutlined class="addmodal-info-icon" />
            </a-tooltip>
          </slot>
          <br />
          <a-switch
            v-model:checked="workspacePublicBoolean"
            @change="dedicatedChange"
          />
        </a-form-item>
        <!--워크스페이스 전용 여부 end-->
        <!--워크스페이스 컴퓨트 오퍼링 start-->
        <a-form-item class="addmodal-aform-item">
          <slot name="label">
            <span>{{ $t("label.compute.offering") }}</span>
            <a-tooltip placement="top">
              <template #title>
                {{ $t("tooltip.workspace.compute.offering") }}
              </template>
              <InfoCircleOutlined class="addmodal-info-icon" />
            </a-tooltip>
          </slot>
          <a-select
            v-model:value="computeOfferingId"
            v-bind:placeholder="$t('tooltip.workspace.compute.offering')"
            show-search
            class="addmodal-aform-item-div"
          >
            <a-select-option value="4Core-8GB">4Core-8GB</a-select-option>
            <a-select-option value="8COre-16GB">8COre-16GB</a-select-option>
          </a-select>
        </a-form-item>
        <!--워크스페이스 컴퓨트 오퍼링 end-->
        <!--워크스페이스 디스크 오퍼링 start-->
        <a-form-item class="addmodal-aform-item">
          <slot name="label">
            <span>{{ $t("label.disk.offering") }}</span>
            <a-tooltip placement="top">
              <template #title>
                {{ $t("tooltip.workspace.disk.offering") }}
              </template>
              <InfoCircleOutlined class="addmodal-info-icon" />
            </a-tooltip>
          </slot>
          <a-select
            v-model:value="diskOfferingId"
            v-bind:placeholder="$t('tooltip.workspace.disk.offering')"
            show-search
            class="addmodal-aform-item-div"
          >
            <a-select-option value="RBD-100GB">RBD-100GB</a-select-option>
            <a-select-option value="RBD-200GB">RBD-200GB</a-select-option>
          </a-select>
        </a-form-item>
        <!--워크스페이스 디스크 오퍼링 end-->
        <!--데스크탑 생성 수량 start-->
        <a-form-item v-show="state.desktopBoolean" class="addmodal-aform-item">
          <slot name="label">
            <span>{{ $t("label.quantity") }}</span>
            <a-tooltip placement="top">
              <template #title>
                {{ $t("tooltip.workspace.quantity") }}
              </template>
              <InfoCircleOutlined class="addmodal-info-icon" />
            </a-tooltip>
          </slot>
          <a-input-number
            v-model:value="desktopQuantity"
            :min="1"
            :max="10"
            class="addmodal-aform-item-div"
          />
        </a-form-item>
        <!--데스크탑 생성 수량 end-->
        <a-form-item class="addmodal-aform-item-button">
          <a-button @click="changeAddModal(false)" style="margin-left: 7px">
            {{ $t("label.cancel") }}
          </a-button>
          <a-button
            type="primary"
            @click="putWorkspace()"
            style="margin-left: 7px"
            >{{ $t("label.ok") }}
            </a-button>
        </a-form-item>
      </a-space>
      <!--===============워크스페이스 추가  끝========================-->
      <!--===============유저 추가 시작=======================-->
      <a-space
        v-if="state.userBoolean"
        direction="vertical"
        class="addmodal-aform-item-div"
      >
        <!--워크스페이스 이름 start-->
        <a-form-item class="addmodal-aform-item">
          <slot name="label">
            {{ $t("label.name") }}
            <a-tooltip placement="top">
              <template #title>
                {{ $t("tooltip.workspace.name") }}
              </template>
              <InfoCircleOutlined class="addmodal-info-icon" />
            </a-tooltip>
          </slot>
          <a-input
            v-bind:placeholder="$t('tooltip.workspace.name')"
            class="addmodal-aform-item-div"
          />
        </a-form-item>
        <!--워크스페이스 이름 end-->
        <!--워크스페이스 설명 start-->
        <a-form-item class="addmodal-aform-item">
          <slot name="label">
            {{ $t("label.description") }}
            <a-tooltip placement="top">
              <template #title>
                {{ $t("tooltip.workspace.description") }}
              </template>
              <InfoCircleOutlined class="addmodal-info-icon" />
            </a-tooltip>
          </slot>
          <a-input v-bind:placeholder="$t('tooltip.workspace.description')" />
        </a-form-item>
        <!--워크스페이스 설명 end-->
        <!--워크스페이스 비밀번호 start-->
        <a-form-item class="addmodal-aform-item">
          <slot name="label">
            <span>{{ "비밀번호" }}</span>
            <a-tooltip placement="top">
              <template #title>
                {{ "비밀번호" }}
              </template>
              <InfoCircleOutlined class="addmodal-info-icon" />
            </a-tooltip>
          </slot>
          <a-input v-bind:placeholder="$t('tooltip.workspace.description')" />
        </a-form-item>
        <!--워크스페이스 비밀번호 end-->
        <!--워크스페이스 비밀번호 확인 start-->
        <a-form-item class="addmodal-aform-item">
          <slot name="label">
            {{ "비밀번호 확인" }}
            <a-tooltip placement="top">
              <template #title>
                {{ "비밀번호 확인" }}
              </template>
              <InfoCircleOutlined class="addmodal-info-icon" />
            </a-tooltip>
          </slot>
          <a-input
            v-bind:placeholder="$t('tooltip.workspace.description')"
            class="addmodal-aform-item-div"
          />
        </a-form-item>
        <!--워크스페이스 비밀번호 확인 end-->
        <a-form-item class="addmodal-aform-item-button">
          <a-button @click="changeAddModal(false)" style="margin-left: 7px">
            {{ $t("label.cancel") }}
          </a-button>
          <a-button
            type="primary"
            @click="putUser()"
            style="margin-left: 7px"
            >{{ $t("label.ok") }}
            </a-button>
        </a-form-item>
      </a-space>
      <!--===============유저 추가  끝========================-->

    </a-form>
  </a-modal>
</template>

<script>
import { defineComponent, onMounted, reactive, ref } from "vue";
export default defineComponent({
  name: "AddModal",
  components: {},
  props: {
    actionFrom: {
      type: String,
      requires: true,
      default: "",
    },
    title: {
      type: String,
      requires: true,
      default: "",
    },
  },
  emits: ["changeAddModal"],
  setup(props) {
    // AddModal의 기본값 설정
    //console.log("AddModal.vue setup");
    //console.log(props.add);
    const state = reactive({
      switchLabel: ref("label.dedicated"),
      desktopBoolean: ref(false),
      modalVisible: ref(false),
      workspaceBoolean: ref(false),
      userBoolean: ref(false),
    });
    onMounted(() => {
      if (props.actionFrom === "Workspace") {
        state.workspaceBoolean = true;
      } else if (props.actionFrom === "User") {
        state.userBoolean = true;
      }
    });
    // Type select box 의 변경에 따른 action start
    const workspaceTypeChange = (value) => {
      if (`${value}` === "desktop") {
        state.desktopBoolean = ref(true);
      } else {
        state.desktopBoolean = ref(false);
      }
    }; // Type select box 의 변경에 따른 action start
    // switch 변경에 따른 라벨 변경 start
    const dedicatedChange = (value) => {
      if (`${value}` === "true") {
        state.switchLabel = ref("label.shared");
        console.log("true");
      } else {
        state.switchLabel = ref("label.dedicated");
        console.log("false");
      }
    }; // switch 변경에 따른 라벨 변경 ens
    return {
      state,
      workspaceTypeChange,
      dedicatedChange,
      desktopQuantity: ref(1),
      workspaceType: ref(undefined),
      templateId: ref(undefined),
      computeOfferingId: ref(undefined),
      diskOfferingId: ref(undefined),
      workspacePublicBoolean: ref(false),
    };
  },
  methods: {
    changeAddModal (value) {
      this.$emit("changeAddModal", value);
    },
    putWorkspace() {
      alert("워크스페이스 생성!!!!!!!!!!!!!!!!!!!");
      // worksApi
      //   .get("/api/v1/workspace")
      //   .then((response) => {
      //     if (response.status == 200) {
      //       this.dataList = response.data.result.list;
      //     } else {
      //       console.log(response.message);
      //     }
      //   })
      //   .catch((error) => {
      //     console.log(error);
      //   });
    },
    putUser() {
      alert("유저 생성!!!!!!!!!!!!!!!!!!!");
      // worksApi
      //   .get("/api/v1/workspace")
      //   .then((response) => {
      //     if (response.status == 200) {
      //       this.dataList = response.data.result.list;
      //     } else {
      //       console.log(response.message);
      //     }
      //   })
      //   .catch((error) => {
      //     console.log(error);
      //   });
    },
  },
});
</script>

<style scoped>
.addmodal-aform-item {
  margin-bottom: 14px;
  width: 100%;
}
.addmodal-aform-item-button {
  /*margin-bottom: 14px;*/
  text-align: right;
}
.addmodal-info-icon {
  margin-left: 5px;
}
.addmodal-aform-item-div {
  width: 100%;
}
</style>
