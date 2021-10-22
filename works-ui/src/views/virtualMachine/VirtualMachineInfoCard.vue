<template>
  <div class="resource-details">
    <div class="resource-details__name">
      <a-avatar shape="square" :size="60">
        <template #icon>
          <DesktopOutlined />
        </template>
      </a-avatar>
      <h4 style="margin-left: 20px">
        {{ vmDbDataInfo.name }}
      </h4>
    </div>
  </div>
  <div id="Status" class="CardItem">
    <div class="ItemName">{{ $t("label.vm.ready.state") }}</div>
    <div class="Item">
      <a-badge
        class="head-example"
        :color="vmDbDataInfo.checked === true ? 'green' : 'red'"
        :text="
          vmDbDataInfo.checked === true
            ? $t('label.vm.status.ready')
            : $t('label.vm.status.notready')
        "
      />
    </div>
  </div>
  <div id="ID" class="CardItem">
    <div class="ItemName">{{ $t("label.uuid") }}</div>
    <div class="Item">
      <a-button shape="circle" type="dashed">
        <BarcodeOutlined />
      </a-button>
      {{ vmDbDataInfo.uuid }}
    </div>
  </div>
  <div class="CardItem">
    <div class="ItemName">{{ $t("label.createdate") }}</div>
    <div class="Item">{{ vmDbDataInfo.create_date }}</div>
  </div>
  <div class="CardItem">
    <div class="ItemName">{{ $t("label.vm.ostype") }}</div>
    <div class="Item">{{ vmMoldDataInfo.osdisplayname }}</div>
  </div>
  <div class="CardItem">
    <div class="ItemName">{{ $t("label.vm.cpu.size") }}</div>
    <div class="Item">{{ vmMoldDataInfo.cputotal }}</div>
    <a-progress :percent="cpuused" size="small" status="active" />
  </div>
  <div class="CardItem">
    <div class="ItemName">{{ $t("label.vm.memory.size") }}</div>
    <div class="Item">{{ vmMoldDataInfo.memory }} MB</div>
  </div>
  <div class="CardItem">
    <div class="ItemName">{{ $t("label.vm.disk.size") }}</div>
    <div class="Item">
      {{ vmDiskInfo.sizegb }}<br />
      <a-tag>
        {{
          $t("label.read") +
          " " +
          (vmMoldDataInfo.diskkbsread / 1048576).toFixed(2)
        }}
        GB</a-tag
      >
      <a-tag>
        {{
          $t("label.write") +
          " " +
          (vmMoldDataInfo.diskkbswrite / 1048576).toFixed(2)
        }}
        GB</a-tag
      ><br />
      <a-tag>
        {{ $t("label.read.io") + " " + vmMoldDataInfo.diskioread }}</a-tag
      >
      <a-tag>
        {{ $t("label.write.io") + " " + vmMoldDataInfo.diskiowrite }}</a-tag
      >
    </div>
  </div>
  <div class="CardItem">
    <div class="ItemName">{{ $t("label.network") }}</div>
    <div class="Item">
      <a-tag>
        <ArrowUpOutlined /> RX {{ vmMoldDataInfo.networkkbsread }} KB</a-tag
      >
      <a-tag>
        <ArrowDownOutlined /> TX {{ vmMoldDataInfo.networkkbswrite }} KB</a-tag
      ><br />
      {{ vmNetworkInfo.networkname }}
    </div>
  </div>
  <div class="CardItem">
    <div class="ItemName">{{ $t("label.vm.network.ip") }}</div>
    <div class="Item">{{ vmNetworkInfo.ipaddress }}</div>
  </div>
  <div class="CardItem">
    <div class="ItemName">{{ $t("label.template") }}</div>
    <div class="Item">{{ vmMoldDataInfo.templatename }}</div>
  </div>
  <div class="CardItem">
    <div class="ItemName">{{ $t("label.vm.serviceOffering") }}</div>
    <div class="Item">{{ vmMoldDataInfo.serviceofferingname }}</div>
  </div>
</template>

<script>
import { defineComponent, reactive, ref } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";

export default defineComponent({
  components: {},
  props: {
    // vmDbDataInfo: {
    //   type: Object,
    //   required: true,
    //   default: null,
    // },
    // vmMoldDataInfo: {
    //   type: Object,
    //   required: true,
    //   default: null,
    // },
  },
  setup() {
    return {};
  },
  data() {
    return {
      vmDbDataInfo: ref([]),
      vmMoldDataInfo: ref([]),
      vmNetworkInfo: ref([]),
      vmDiskInfo: ref([]),
      cpuused: ref(0),
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      worksApi
        .get("/api/v1/instance/detail/" + this.$route.params.vmUuid)
        .then((response) => {
          if (response.status === 200) {
            //console.log(response.data.result.instanceDBInfo);
            this.vmDbDataInfo = response.data.result.instanceDBInfo;
            this.vmMoldDataInfo =
              response.data.result.instanceMoldInfo.virtualmachine[0];
            this.vmNetworkInfo = this.vmMoldDataInfo.nic[0];
            this.vmDiskInfo =
              response.data.result.instanceInstanceVolumeInfo.volume[0];
            this.cpuused = this.vmMoldDataInfo.cpuused.split("%")[0];
          } else {
            message.error(this.$t("message.response.data.fail"));
            //console.log("데이터를 정상적으로 가져오지 못했습니다.");
          }
        })
        .catch(function (error) {
          console.log(error);
        });
    },
  },
});
</script>

<style lang="scss" scoped>
.CardItem {
  margin-bottom: 20px;
}
::v-deep(.ant-badge-status-dot) {
  width: 12px;
  height: 12px;
}
.ItemName {
  font-size: 14px;
  font-weight: bold;
}
.resource-details {
  text-align: center;
  margin-bottom: 20px;

  &__name {
    display: flex;
    align-items: center;
    align-content: center;
    text-align: center;

    .ant-avata {
      margin-right: 20px;
      overflow: hidden;
      min-width: 50px;
      cursor: pointer;
      img {
        height: 100%;
        width: 100%;
      }
    }
    h4 {
      font-size: 18px;
    }
  }
}
</style>
