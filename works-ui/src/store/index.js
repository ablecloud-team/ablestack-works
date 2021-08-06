import { createStore } from "vuex";

export default createStore({
  state: {
    user: [{ isLogin: false, isAdmin: false, userName: "" }],
  },
  mutations: {
    loginSuccess(state, params) {
      state.user.isLogin = true;
      state.user.isAdmin = params.user.isAdmin;
    },
  },
  actions: {},
  modules: {},
});
