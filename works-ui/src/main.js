import { createApp } from "vue";
import Antd from "ant-design-vue";
import "ant-design-vue/dist/antd.css";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import i18n from "./locales";

import {
  faCameraRetro,
  faUserSecret,
  faLanguage,
} from "@fortawesome/free-solid-svg-icons";
import { faUbuntu, faCentos } from "@fortawesome/free-brands-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import {
  BackwardFilled,
  HomeFilled,
  DeleteFilled,
  UserOutlined,
  PoweroffOutlined,
  ReloadOutlined,
  SyncOutlined,
  CameraOutlined,
  PaperClipOutlined,
  LockOutlined,
  BellOutlined,
  FontSizeOutlined,
  EditOutlined,
  HomeOutlined,
  BarcodeOutlined,
  PlusOutlined,
  MoreOutlined,
  CaretRightOutlined,
  GlobalOutlined,
  MenuFoldOutlined,
  LogoutOutlined,
} from "@ant-design/icons-vue";

library.add(faCameraRetro);
library.add(faUbuntu);
library.add(faCentos);
library.add(faUserSecret);
library.add(faLanguage);

const app = createApp(App); //.use(loadLocaleMessages);

app.component("font-awesome-icon", FontAwesomeIcon);
app.component("BackwardFilled", BackwardFilled);
app.component("HomeFilled", HomeFilled);
app.component("DeleteFilled", DeleteFilled);
app.component("UserOutlined", UserOutlined);
app.component("PoweroffOutlined", PoweroffOutlined);
app.component("ReloadOutlined", ReloadOutlined);
app.component("SyncOutlined", SyncOutlined);
app.component("CameraOutlined", CameraOutlined);
app.component("PaperClipOutlined", PaperClipOutlined);
app.component("LockOutlined", LockOutlined);
app.component("BellOutlined", BellOutlined);
app.component("FontSizeOutlined", FontSizeOutlined);
app.component("EditOutlined", EditOutlined);
app.component("HomeOutlined", HomeOutlined);
app.component("BarcodeOutlined", BarcodeOutlined);
app.component("PlusOutlined", PlusOutlined);
app.component("MoreOutlined", MoreOutlined);
app.component("CaretRightOutlined", CaretRightOutlined);
app.component("GlobalOutlined", GlobalOutlined);
app.component("MenuFoldOutlined", MenuFoldOutlined);
app.component("LogoutOutlined", LogoutOutlined);

app.use(i18n).use(store).use(router).use(Antd).mount("#app");
