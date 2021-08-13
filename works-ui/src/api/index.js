import axios from "axios";
// import store from "../store/index.js";

const instance = axios.create({
  baseURL: process.env.VUE_APP_API_URL,
  headers: {
    // Authorization: store.state.token, // store에서 불러오기
  },
});

function axiosLogin(userData) {
  return instance.post("/api/login", userData, { withCredentials: true });
}

function axiosUserDetail() {
  return instance.get("/api/v1/user", { withCredentials: true });
  // return instance.get("/api/test/user/", { withCredentials: true });
}

function axiosLogout() {
  return instance.get("/api/v1/logout", { withCredentials: true });
  // return instance.get("/api/test/user/", { withCredentials: true });
}

export { axiosLogin, axiosUserDetail, axiosLogout };
