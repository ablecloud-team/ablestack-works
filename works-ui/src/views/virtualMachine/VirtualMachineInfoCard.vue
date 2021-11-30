<template>
  <a-spin :spinning="spinning">
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
      <div class="ItemName">{{ $t("label.vm.state") }}</div>
      <div class="Item">
        <a-tooltip placement="bottom">
          <template #title>{{ vmDbDataInfo.mold_status }}</template>
          <a-badge
            class="head-example"
            :color="
              vmDbDataInfo.mold_status === 'Running'
                ? 'green'
                : vmDbDataInfo.mold_status === 'Stopping' ||
                  vmDbDataInfo.mold_status === 'Starting'
                ? 'blue'
                : vmDbDataInfo.mold_status === 'Stopped'
                ? 'red'
                : ''
            "
            :text="
              vmDbDataInfo.mold_status === 'Running'
                ? $t('label.vm.status.running')
                : vmDbDataInfo.mold_status === 'Starting'
                ? $t('label.vm.status.starting')
                : vmDbDataInfo.mold_status === 'Stopping'
                ? $t('label.vm.status.stopping')
                : vmDbDataInfo.mold_status === 'Stopped'
                ? $t('label.vm.status.stopped')
                : ''
            "
          />
        </a-tooltip>
      </div>
    </div>

    <div id="Status" class="CardItem">
      <div class="ItemName">{{ $t("label.vm.ready.state") }}</div>
      <div class="Item">
        <a-tooltip placement="bottom">
          <template #title>{{ vmDbDataInfo.handshake_status }}</template>
          <a-badge
            class="head-example"
            :color="vmDbDataInfo.handshake_status === 'Ready' ? 'green' : 'red'"
            :text="
              vmDbDataInfo.handshake_status === 'Not Ready' ||
              vmDbDataInfo.handshake_status === 'Pending'
                ? $t('label.vm.status.initializing')
                : vmDbDataInfo.handshake_status === 'Joining' ||
                  vmDbDataInfo.handshake_status === 'Joined'
                ? $t('label.vm.status.configuring')
                : vmDbDataInfo.handshake_status === 'Ready'
                ? $t('label.vm.status.ready')
                : $t('label.vm.status.notready')
            "
          />
        </a-tooltip>
      </div>
    </div>
    <div id="ID" class="CardItem">
      <div class="ItemName">{{ $t("label.uuid") }}</div>
      <div class="Item">
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
          <ArrowDownOutlined /> TX
          {{ vmMoldDataInfo.networkkbswrite }} KB</a-tag
        ><br />
        {{ vmNetworkInfo.networkname }} ({{ vmNetworkInfo.type }})
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
  </a-spin>
</template>

<script>
import { defineComponent, ref } from "vue";
import { worksApi } from "@/api/index";
import { message } from "ant-design-vue";

export default defineComponent({
  components: {},
  props: {
    vmDbDataInfo: {
      type: Object,
      required: false,
      default: null,
    },
    vmMoldDataInfo: {
      type: Object,
      required: false,
      default: null,
    },
    vmNetworkInfo: {
      type: Object,
      required: false,
      default: null,
    },
    vmDiskInfo: {
      type: Object,
      required: false,
      default: null,
    },
    cpuused: {
      type: Number,
      required: false,
      default: 0,
    },
  },
  setup() {
    return {};
  },
  data(props) {
    return {
      spinning: ref(true),
    };
  },
  created() {
    this.fetchRefresh();
  },
  methods: {
    fetchRefresh() {
      this.spinning = true;
      setTimeout(() => {
        this.spinning = false;
      }, 500);
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
