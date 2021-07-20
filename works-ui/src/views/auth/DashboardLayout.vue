<template>
  <a-image
      class="user-layout-logo"
      src="@/assets/ablestack-logo.png"
  />
  <div class="user-layout desktop">
      <a-form
          layout="horizontal"
          :model="formState"
          @finish="handleFinish"
          @finishFailed="handleFinishFailed"
          class="user-layout-login"
      >
        <a-form-item>
          <a-input v-model:value="formState.user" placeholder="Username" size="large">
            <template #prefix><UserOutlined style="color: rgba(0, 0, 0, 0.25)" /></template>
          </a-input>
        </a-form-item>
        <a-form-item>
          <a-input-password v-model:value="formState.password" type="password" placeholder="Password" size="large">
            <template #prefix><LockOutlined style="color: rgba(0, 0, 0, 0.25)" /></template>
          </a-input-password>
        </a-form-item>
        <a-form-item>
          <a-button
              type="primary"
              @click="$router.push({name: 'Home'})"
              block
          >
            Log in
          </a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<script>
import { defineComponent, reactive } from 'vue';

export default defineComponent({
  name: "DashboardLayout",
  setup(){
    const formState = reactive({
      user: '',
      password: '',
    });

    const handleFinish = values => {
      console.log(values, formState);
    };

    const handleFinishFailed = errors => {
      console.log(errors);
    };
    return {
      formState,
      handleFinish,
      handleFinishFailed,
    };
  },
});
</script>

<style lang="less" scoped>
.user-layout{
  background: #42b983;
  height: 100%;

  button.login-button {
    margin-top: 8px;
    padding: 0 15px;
    font-size: 16px;
    height: 40px;
    width: 100%;
  }

  .user-login-other {
    text-align: left;
    margin-top: 24px;
    line-height: 22px;

    .item-icon {
      font-size: 24px;
      color: rgba(0, 0, 0, 0.2);
      margin-left: 16px;
      vertical-align: middle;
      cursor: pointer;
      transition: color 0.3s;

      &:hover {
        color: #1890ff;
      }
    }

    .register {
      float: right;
    }
  }
}
.user-layout-container{
  padding: 3rem 0;
  width: 100%;
}

.user-layout-login{
  min-width: 260px;
  width: 368px;
  margin: 0 auto;
}

.user-layout-logo{
  width: 600px;
  height: 80px;
  margin: 0 auto 2rem;
  border-style: none;
  display: block;
}
</style>
