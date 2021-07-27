<template>
  <a-modal
      v-model:visible="modalVisible"
      closable
      :footer="null"
      style="width: 100vw; top: 40px"
    >
    <template #title>{{ add }}</template>
    <a-form layout="vertical">
      <!--워크스페이스 이름 start-->
      <a-form-item class="addmodal-aform-item">
        <template #label>
          <span>{{ $t('label.name') }}</span>
          <a-tooltip placement="top">
            <template #title>
              {{ $t('tooltip.workspace.name') }}
            </template>
            <InfoCircleOutlined class="addmodal-info-icon"/>
          </a-tooltip>
        </template>
        <a-input v-bind:placeholder="$t('tooltip.workspace.name')"/>
      </a-form-item>
      <!--워크스페이스 이름 end-->
      <!--워크스페이스 설명 start-->
      <a-form-item class="addmodal-aform-item">
        <template #label>
          <span>{{ $t('label.description') }}</span>
          <a-tooltip placement="top">
            <template #title>
              {{ $t('tooltip.workspace.description') }}
            </template>
            <InfoCircleOutlined class="addmodal-info-icon"/>
          </a-tooltip>
        </template>
        <a-input v-bind:placeholder="$t('tooltip.workspace.description')"/>
      </a-form-item>
      <!--워크스페이스 설명 end-->
      <!--워크스페이스 타입 start-->
      <a-form-item class="addmodal-aform-item">
        <template #label>
          <span>{{ $t('label.type') }}</span>
          <a-tooltip placement="top">
            <template #title>
              {{ $t('tooltip.workspace.type') }}
            </template>
            <InfoCircleOutlined class="addmodal-info-icon"/>
          </a-tooltip>
        </template>
        <a-select
            v-model:value="workspaceType"
            v-bind:placeholder="$t('tooltip.workspace.type')"
            show-search
            class="addmodal-aform-item-div"
            @change="workstpaceTypeChange"
        >
          <a-select-option value="desktop">Desktop</a-select-option>
          <a-select-option value="application">Application</a-select-option>
        </a-select>
      </a-form-item>
      <!--워크스페이스 타입 end-->
      <!--워크스페이스 템플릿 start-->
      <a-form-item class="addmodal-aform-item">
        <template #label>
          <span>{{ $t('label.template') }}</span>
          <a-tooltip placement="top">
            <template #title>
              {{ $t('tooltip.workspace.template') }}
            </template>
            <InfoCircleOutlined class="addmodal-info-icon"/>
          </a-tooltip>
        </template>
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
      <!--데스크탑 퍼블릭 여부 start-->
      <a-form-item
          class="addmodal-aform-item"
          v-model:style="state.publicSwitch"
      >
        <template #label>
          <span>{{ $t('label.workspace.public.boolean') }}</span>
          <a-tooltip placement="top">
            <template #title>
              {{ $t('tooltip.workspace.type') }}
            </template>
            <InfoCircleOutlined class="addmodal-info-icon"/>
          </a-tooltip>
        </template>
        <a-switch v-model:checked="workspacePublicBoolean" />
      </a-form-item>
      <!--데스크탑 퍼블릭 여부 end-->
      <!--워크스페이스 컴퓨트 오퍼링 start-->
      <a-form-item class="addmodal-aform-item">
        <template #label>
          <span>{{ $t('label.compute.offering') }}</span>
          <a-tooltip placement="top">
            <template #title>
              {{ $t('tooltip.workspace.compute.offering') }}
            </template>
            <InfoCircleOutlined class="addmodal-info-icon"/>
          </a-tooltip>
        </template>
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
        <template #label>
          <span>{{ $t('label.disk.offering') }}</span>
          <a-tooltip placement="top">
            <template #title>
              {{ $t('tooltip.workspace.disk.offering') }}
            </template>
            <InfoCircleOutlined class="addmodal-info-icon"/>
          </a-tooltip>
        </template>
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
      <a-form-item
          v-model:style="state.publicSwitch"
          class="addmodal-aform-item"
      >
        <template #label>
          <span>{{ $t('label.quantity') }}</span>
          <a-tooltip placement="top">
            <template #title>
              {{ $t('tooltip.workspace.quantity') }}
            </template>
            <InfoCircleOutlined class="addmodal-info-icon"/>
          </a-tooltip>
        </template>
        <a-input-number
            v-model:value="desktopQuantity"
            :min="1"
            :max="10"
            class="addmodal-aform-item-div"
        />
      </a-form-item>
      <!--데스크탑 생성 수량 end-->
      <a-form-item class="addmodal-aform-item-button">
        <a-button type="button" style="margin-left: 7px">{{ $t('label.cancel') }}</a-button>
        <a-button type="primary" style="margin-left: 7px">{{ $t('label.ok') }}</a-button>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script>
import {defineComponent, reactive, ref} from 'vue';
import {InfoCircleOutlined} from "@ant-design/icons-vue";

export default defineComponent({
  name: "AddModal",
  components: {
    InfoCircleOutlined,
  },
  props:{
    actionFrom: {
      String,
      required: true,
    },
    add: {
      String,
    }
  },
  emits: [
      'changeAddModal'
  ],
  setup(){
    const modalVisible = ref(false);
    // let publicSwitch = ref('display: none');
    const state = reactive({
      publicSwitch : ref('display: none')
    });
    const workstpaceTypeChange = value => {
      console.log(`${value}`);
      if (`${value}` === 'desktop'){
        state.publicSwitch = ref('display: block');
        console.log(state.publicSwitch.value);
      }else{
        state.publicSwitch = ref('display: none');
      }
    };
    return{
      modalVisible,
      desktopQuantity: ref(1),
      workstpaceTypeChange,
      workspaceType: ref(undefined),
      templateId: ref(undefined),
      computeOfferingId: ref(undefined),
      diskOfferingId: ref(undefined),
      workspacePublicBoolean: ref(false),
      state
    }
  },
});
</script>

<style scoped>
.addmodal-aform-item{
  margin-bottom: 14px;
}
.addmodal-aform-item-button{
  /*margin-bottom: 14px;*/
  text-align: right;
}
.addmodal-info-icon{
  margin-left: 5px;
}
.addmodal-aform-item-div{
  width: 100%;
}
</style>
