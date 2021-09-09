import axios from "axios";
// import store from "../store/index.js";

const worksApi = axios.create({
  baseURL: process.env.VUE_APP_API_URL,
  headers: {
    Authorization: "works" + localStorage.getItem("token"),
  },
});

function axiosLogin(userData) {
  return worksApi.post("/api/login", userData, { withCredentials: true });
}

function axiosTokenAuth() {
  return worksApi.get("/api/v1/token", { withCredentials: true });
}

function axiosLogout() {
  return worksApi.get("/api/v1/logout", { withCredentials: true });
}

export { worksApi, axiosLogin, axiosLogout, axiosTokenAuth };
