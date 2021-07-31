<template>
  <a-row class="user-menu">
    <a-col :span="12">
      <MenuFoldOutlined class="header-notice-icon" />
    </a-col>
    <a-col :span="12" style="text-align: right;">
      <a-space :align="end">
        <a-popover placement="bottom">
          <template #content>
            <a-button type="text" @click="setI18nLanguage('ko')">한국어</a-button>
            <br>
            <a-button type="text" @click="setI18nLanguage('en')">English</a-button>
          </template>
          <a-button type="text" shape="circle" class="header-notice-button">
            <template #icon>
              <GlobalOutlined class="header-notice-icon" />
            </template>
          </a-button>
        </a-popover>
        <a-button type="text" shape="circle" class="header-notice-button">
          <template #icon>
            <BellOutlined class="header-notice-icon" />
          </template>
        </a-button>
        <a-button type="text" shape="circle" class="header-notice-button">
          <template #icon>
            <UserOutlined class="header-notice-icon" />
          </template>
        </a-button>
      </a-space>
    </a-col>
  </a-row>
</template>

<script>
export default {
  name: "AdminHeader",
  methods: {
    setI18nLanguage(locale) {
      this.$i18n.locale = locale
      this.$router.push({
        params: { lang: locale }
      })
    }
  }
}
</script>

<style scoped>
.header-notice-icon{
  font-size: 22px;
  padding: 6px;
  margin: 5px;
}
.header-notice-button{
  font-size: 21px;
  padding: 6px;
  margin-right: 21px;
}

.user-menu{
  background: white;
  height: 100%;
  padding: 0 2px;
}
</style>
