<template>
  <div style="width: 256px; height: 100%;">
    <Logo @click="toggleCollapsed"/>
    <a-menu
        mode="inline"
        theme="light"
        :inline-collapsed="state.collapsed"
        v-model:selectedKeys="state.selectedKeys"
        style="padding-top: 14px"
    >
      <a-menu-item
          key="1"
          @click="$router.push({name: 'Dashboard'})"
      >
        <template #icon>
          <DashboardOutlined />
        </template>
        <span>{{ $t('label.dashboard') }}</span>
      </a-menu-item>
      <a-menu-item key="2"
                   @click="$router.push({name: 'Workspaces'})">
        <template #icon>
          <CloudOutlined />
        </template>
        <span>{{ $t('label.workspace') }}</span>
      </a-menu-item>
      <a-menu-item key="3"
                   @click="$router.push({name: 'Virtualmachine'})">
        <template #icon>
          <DesktopOutlined />
        </template>
        <span>{{ $t('label.vm') }}</span>
      </a-menu-item>
      <a-menu-item key="4"
                   @click="$router.push({name: 'Users'})">
        <template #icon>
          <TeamOutlined />
        </template>
        <span>{{ $t('label.users') }}</span>
      </a-menu-item>
      <a-menu-item key="5"
                   @click="$router.push({name: 'GroupPolicy'})">
        <template #icon>
          <ReconciliationOutlined />
        </template>
        <span>{{ $t('label.group.policy') }}</span>
      </a-menu-item>
      <a-menu-item key="6">
        <template #icon>
          <BarChartOutlined />
        </template>
        <span>{{ $t('label.audit') }}</span>
      </a-menu-item>
      <a-menu-item key="7">
        <template #icon>
          <CoffeeOutlined />
        </template>
        <span>{{ $t('label.community') }}</span>
      </a-menu-item>
    </a-menu>
  </div>
</template>
<script>
import {defineComponent, reactive, ref, watch} from 'vue';
import {
  DashboardOutlined,
  CloudOutlined,
  DesktopOutlined,
  TeamOutlined,
  ReconciliationOutlined,
  BarChartOutlined,
  CoffeeOutlined,
} from '@ant-design/icons-vue';
import Logo from "./Logo";
export default defineComponent({
  components: {
    Logo,
    DashboardOutlined,
    CloudOutlined,
    DesktopOutlined,
    TeamOutlined,
    ReconciliationOutlined,
    BarChartOutlined,
    CoffeeOutlined,
  },
  props:{
    collapsed: Boolean,
  },
  emits: [
    'changeToggleCollapsed'
  ],
  setup(props) {
    const state = reactive({
      collapsed: ref(props.collapsed),
      selectedKeys: ['1'],
      openKeys: ['sub1'],
      preOpenKeys: ['sub1'],
    });
    watch(
        () => state.openKeys,
        (val, oldVal) => {
          state.preOpenKeys = oldVal;
        },
    );
    return {
      state,
    };
  },
  methods:{
    toggleCollapsed: function (){
      this.$emit('changeToggleCollapsed');
    }
  }
});
</script>

<style>
#components-layout-demo-custom-trigger .trigger {
  font-size: 18px;
  line-height: 64px;
  padding: 0 24px;
  cursor: pointer;
  transition: color 0.3s;
}

#components-layout-demo-custom-trigger .trigger:hover {
  color: #1890ff;
}

#components-layout-demo-custom-trigger .logo {
  height: 32px;
  background: rgba(255, 255, 255, 0.3);
  margin: 16px;
}

.site-layout .site-layout-background {
  background: #fff;
}
</style>