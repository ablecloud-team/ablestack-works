<template>
  <a-row :gutter="16" type="flex" justify="center" align="top">
    <a-col flex="215px">
      <h3>1. 입력 방법</h3>
      <a-radio-group
        v-model:value="setText"
        button-style="solid"
        @change="
          (selected) => {
            $emit('inputModeChange', selected.target.value);
          }
        "
      >
        <a-tooltip
          placement="bottom"
          title="사용중인 입력 방법이 없습니다. 키보드 입력은 열결된 물리적 키보드에서 받아들여집니다."
        >
          <a-radio-button value="none">없음</a-radio-button>
        </a-tooltip>
        <a-tooltip
          placement="bottom"
          title="텍스트 입력을 허용하고, 입력된 텍스트를 바탕으로 키보드가 없는 휴대폰과 같은 장치에 필요합니다."
        >
          <a-radio-button value="text">텍스트 입력</a-radio-button>
        </a-tooltip>
        <!-- <a-tooltip
          placement="bottom"
          title="내장된 데스크톱 화상 키보드의 입력을 표시하고 허용합니다. 화상 키보드는 다른 방법으로는 불가능할 수 있는 키 조합을 입력할 수 있습니다. (Ctrl-Alt-Del 등)"
        >
          <a-radio-button value="osk">화상 키보드</a-radio-button>
        </a-tooltip> -->
      </a-radio-group>
    </a-col>
    <a-col flex="215px">
      <h3>2. 마우스 에뮬레이션 모드</h3>
      <a-radio-group
        v-model:value="setMouse"
        button-style="solid"
        @change="
          (selected) => {
            $emit('mouseModeChange', selected.target.value);
          }
        "
      >
        <a-tooltip
          placement="bottom"
          title="탭하여 클릭합니다. 클릭은 터치 위치에서 발생합니다."
        >
          <a-radio-button :value="true">터치 스크린</a-radio-button>
        </a-tooltip>
        <a-tooltip
          placement="bottom"
          title="드래그하여 마우스 포인터를 움직이고 탭하여 클릭합니다. 클릭은 마우스 포인터 위치에서 발생합니다."
        >
          <a-radio-button :value="false">터치 패드</a-radio-button>
        </a-tooltip>
      </a-radio-group>
    </a-col>
    <a-col flex="215px">
      <h3>3. 디스플레이</h3>
      <a-button @click="zoomHandle(-10)" shape="circle">
        <template #icon><MinusOutlined /></template>
      </a-button>
      &nbsp;
      <a-input
        v-model:value="scale"
        style="width: 80px"
        :max-length="3"
        @blur="scaleHandle"
      />%&nbsp;
      <a-button @click="zoomHandle(10)" shape="circle">
        <template #icon><PlusOutlined /></template>
      </a-button>
    </a-col>
  </a-row>
</template>
<script>
import { ref } from "vue";
import store from "@/store/index";
export default {
  props: {
    client: {
      type: Object,
      requires: true,
      default: null,
    },
    display: {
      type: Object,
      requires: true,
      default: null,
    },
    mouse: {
      type: Object,
      requires: true,
      default: null,
    },
    keyboard: {
      type: Object,
      requires: true,
      default: null,
    },
  },
  setup() {
    const value = ref("");
    return {
      value,
    };
  },
  data(props) {
    return {
      client: ref(props.client),
      display: ref(props.display),
      mouse: ref(props.mouse),
      keyboard: ref(props.keyboard),
      setText: ref("none"),
      setMouse: ref(true),
      scale: ref(store.state.client.scale * 100),
    };
  },
  computed: {},
  created() {},
  mounted() {},
  watch: {},
  methods: {
    zoomHandle(inout) {
      this.scale = parseInt(this.scale) + inout;
      this.scaleHandle();
    },
    scaleHandle() {
      if (store.state.client.minScale * 100 > this.scale)
        this.scale = store.state.client.minScale * 100;
      else if (store.state.client.maxScale * 100 < this.scale)
        this.scale = store.state.client.maxScale * 100;

      this.display.scale(parseInt(this.scale) / 100);
      if (this.scale === store.state.client.minScale)
        document.getElementById("app").style.overflow = "hidden";
      else document.getElementById("app").style.overflow = "auto";
    },
  },
};
</script>
<style scoped>
.modal {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 5px;
  padding: 1rem;
}
.rct {
  text-decoration: underline;
  cursor: pointer;
}
</style>
