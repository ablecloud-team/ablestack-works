import { createStore } from "vuex";

export default createStore({
  state: {
    user: {
      isLogin: false,
      isAdmin: false,
      userID: null,
      token: null,
    },
    workspace: {
      status: "",
      name: "",
    },
    dashboard: {
      works: false,
      mold: false,
      dc: false,
      samba: false,
    },
  },
  getters: {},
  mutations: {
    loginSuccess(state, payload) {
      state.user.isLogin = true;
      state.user.isAdmin = payload.result.isAdmin;
      state.user.userID = payload.result.username;
      state.user.token = payload.result.token;
    },
    logoutSuccess(state) {
      state.user.isLogin = false;
      state.user.isAdmin = false;
    },
    serverCheckSuccess(state, payload) {
      state.dashboard.works = payload.status;
      state.dashboard.mold = payload.data.result["Mold"];
      state.dashboard.dc = payload.data.result["Works-DC"];
      state.dashboard.samba = payload.data.result["Works-Samba"];
    }, Î
  },
  actions: {
    loginCommit({ commit }, payload) {
      commit("loginSuccess", payload);
    },
    logoutCommit({ commit }) {
      commit("logoutSuccess");
      sessionStorage.clear();
    },
    serverCheckCommit({ commit }, payload) {
      commit("serverCheckSuccess", payload);
    },
  },
  modules: {},
});
