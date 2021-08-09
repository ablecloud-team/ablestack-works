import { createStore } from "vuex";

export default createStore({
  state: {
    user: [{ isLogin: false, isAdmin: false, userID: "" }],
  },
  mutations: {
    loginSuccess(state, params) {
      state.user.isLogin = true;
      state.user.isAdmin = params.user.isAdmin;
      state.user.userID = params.user.userID;
    },
  },
  actions: {},
  modules: {},
});
