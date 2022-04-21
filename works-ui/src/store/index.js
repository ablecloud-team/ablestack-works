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
      works: 0,
      mold: 0,
      dc: 0,
      samba: 0,
      guac: 0,
    },
    menukey: 0,
    client: {
      scale: 1,
      minScale: 0,
      maxScale: 3,
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
      if (payload.status === 200) {
        state.dashboard.mold = payload.data.result["Mold"];
        state.dashboard.dc = payload.data.result["Works-DC"];
        state.dashboard.samba = payload.data.result["Works-Samba"];
        state.dashboard.guac = payload.data.result["Guacamole"];
      } else {
        state.dashboard.mold = 0;
        state.dashboard.dc = 0;
        state.dashboard.samba = 0;
        state.dashboard.guac = 0;
      }
    },
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
