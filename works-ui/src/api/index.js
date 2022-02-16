import store from "@/store/index.js";
import axios from "axios";

function setInterceptors(instance) {
  instance.interceptors.request.use(
    function (config) {
      // Do something before request is sent
      config.headers = {
        "Authorization": "works" + (sessionStorage.getItem("token") == "" ? store.state.user.token : sessionStorage.getItem("token")),
      };
      return config;
    },
    function (error) {
      // Do something with request error
      return Promise.reject(error);
    }
  );
  // Add a response interceptor
  instance.interceptors.response.use(
    function (response) {
      // Any status code that lie within the range of 2xx cause this function to trigger
      // Do something with response data
      return response;
    },
    function (error) {
      // Any status codes that falls outside the range of 2xx cause this function to trigger
      // Do something with response error
      return Promise.reject(error);
    }
  );
  return instance;
}

function worksApiAuth() {
  const instance = axios.create({
    //baseURL: process.env.VUE_APP_API_URL,
    baseURL: process.env.VUE_APP_API_URL == "" ? "http://" + window.location.hostname + ":8082" : process.env.VUE_APP_API_URL,
    withCredentials: true,
    timeout: 5000,
  });
  return setInterceptors(instance);
}
export const worksApi = worksApiAuth();

// const worksApi = axios.create({
//   baseURL: process.env.VUE_APP_API_URL,
//   withCredentials: true,
//   headers: {
//     Authorization: "works" + store.state.user.token == null ? localStorage.getItem("token") : store.state.user.token,
//   },
// });

// function axiosLogin(userData) {
//   return worksApi.post("/api/login", userData);
// }

// function axiosTokenAuth() {
//   return worksApi.get("/api/v1/token");
// }

function axiosLogout() {
  return worksApi.get("/api/v1/logout");
}

export {
  axiosLogout,
  // axiosLogin,
  // axiosTokenAuth,
};
