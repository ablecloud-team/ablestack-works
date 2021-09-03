import { createStore } from "vuex";

export default createStore({
  state: {
    user: [{ isLogin: false, isAdmin: false, userID: null, accessToken: null }],
  },
  getters: {},
  mutations: {
    loginSuccess(state, payload) {
      state.user.isLogin = true;
      //state.user.isAdmin = payload.result.isAdmin;
      state.user.userID = payload.result.username;
      state.user.token = payload.result.token;
    },
    logoutSuccess(state) {
      state.user.isLogin = false;
      state.user.isAdmin = false;
    },
  },
  actions: {
    loginCommit({ commit }, payload) {
      // console.log("store index loginCommit");
      //console.log(payload);
      commit("loginSuccess", payload);
    },
    logoutCommit({ commit }) {
      commit("logoutSuccess");
      localStorage.setItem("token", "");
    },
  },
  modules: {},
});
