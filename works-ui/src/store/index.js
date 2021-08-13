import { createStore } from "vuex";

export default createStore({
  state: {
    user: [{ isLogin: false, isAdmin: false, userID: null, accessToken: null }],
  },
  getters: {},
  mutations: {
    loginSuccess(state, payload) {
      state.user.isLogin = true;
      state.user.isAdmin = payload.result.isAdmin;
      state.user.userID = payload.result.username;
      state.user.accessToken = payload.accessToken;
      // console.log("store index loginSuccess");
      // console.log(state);
    },
    logoutSuccess(state) {
      state.user.isLogin = false;
      state.user.isAdmin = false;
    },
  },
  actions: {
    loginCommit({ commit }, payload) {
      // console.log("store index loginCommit");
      // console.log(payload);
      commit("loginSuccess", payload);
      // localStorage.setItem("token", payload.accessToken);
    },
    logoutCommit({ commit }) {
      commit("logoutSuccess");
    },
  },
  modules: {},
});
