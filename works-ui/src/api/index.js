import axios from "axios";
// import store from "../store/index.js";

const instance = axios.create({
  baseURL: process.env.VUE_APP_API_URL,
  headers: {
    Authorization: "XMLHttpRequest",
  },
});

function axiosLogin(userData) {
  return instance.post("/api/login", userData, {
    withCredentials: true,
    origin: true,
  });
}

function axiosUserDetail() {
  return instance.get("/api/v1/user", {
    withCredentials: true,
  });
  // return instance.get("/api/test/user/", { withCredentials: true });
}

function axiosLogout() {
  return instance.get("/api/v1/logout", {
    withCredentials: true,
    origin: true,
  });
  // return instance.get("/api/test/user/", { withCredentials: true });
}

export { axiosLogin, axiosUserDetail, axiosLogout };
