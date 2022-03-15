import { createApp } from "vue";
import VueCryptojs from "vue-cryptojs";
import Antd from "ant-design-vue";
import "ant-design-vue/dist/antd.css";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import i18n from "./locales";
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { fas } from "@fortawesome/free-solid-svg-icons";
import { far } from "@fortawesome/free-regular-svg-icons";
import { fab } from "@fortawesome/free-brands-svg-icons";
library.add(far);
library.add(fas);
library.add(fab);
import {
  BackwardFilled,
  HomeFilled,
  DeleteFilled,
  UserOutlined,
  PoweroffOutlined,
  UserAddOutlined,
  UserDeleteOutlined,
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
  SearchOutlined,
  InfoCircleOutlined,
  PauseOutlined,
  ClusterOutlined,
  DashboardOutlined,
  CloudOutlined,
  DesktopOutlined,
  TeamOutlined,
  ReconciliationOutlined,
  BarChartOutlined,
  CoffeeOutlined,
  MenuUnfoldOutlined,
  DownOutlined,
  StarOutlined,
  AppstoreAddOutlined,
  DownloadOutlined,
  Html5Outlined,
  UsergroupAddOutlined,
  ArrowUpOutlined,
  ArrowDownOutlined,
  LoginOutlined,
  WindowsFilled,
  StarFilled,
  StarTwoTone,
  CloudDownloadOutlined,
  CodeFilled,
  SettingOutlined,
  CheckCircleOutlined,
  CloseCircleOutlined,
  SmileOutlined,
  FrownOutlined,
  UploadOutlined,
  CaretDownFilled,
  MinusOutlined,
  PlusCircleOutlined,
} from "@ant-design/icons-vue";

const app = createApp(App);
app.component("font-awesome-icon", FontAwesomeIcon);
app.component("BackwardFilled", BackwardFilled);
app.component("HomeFilled", HomeFilled);
app.component("DeleteFilled", DeleteFilled);
app.component("UserOutlined", UserOutlined);
app.component("PoweroffOutlined", PoweroffOutlined);
app.component("UserAddOutlined", UserAddOutlined);
app.component("UserDeleteOutlined", UserDeleteOutlined);
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
app.component("SearchOutlined", SearchOutlined);
app.component("InfoCircleOutlined", InfoCircleOutlined);
app.component("PauseOutlined", PauseOutlined);
app.component("ClusterOutlined", ClusterOutlined);
app.component("DashboardOutlined", DashboardOutlined);
app.component("CloudOutlined", CloudOutlined);
app.component("DesktopOutlined", DesktopOutlined);
app.component("TeamOutlined", TeamOutlined);
app.component("ReconciliationOutlined", ReconciliationOutlined);
app.component("BarChartOutlined", BarChartOutlined);
app.component("CoffeeOutlined", CoffeeOutlined);
app.component("MenuUnfoldOutlined", MenuUnfoldOutlined);
app.component("DownOutlined", DownOutlined);
app.component("StarOutlined", StarOutlined);
app.component("AppstoreAddOutlined", AppstoreAddOutlined);
app.component("DownloadOutlined", DownloadOutlined);
app.component("Html5Outlined", Html5Outlined);
app.component("UsergroupAddOutlined", UsergroupAddOutlined);
app.component("ArrowUpOutlined", ArrowUpOutlined);
app.component("ArrowDownOutlined", ArrowDownOutlined);
app.component("LoginOutlined", LoginOutlined);
app.component("WindowsFilled", WindowsFilled);
app.component("StarFilled", StarFilled);
app.component("StarTwoTone", StarTwoTone);
app.component("CloudDownloadOutlined", CloudDownloadOutlined);
app.component("CodeFilled", CodeFilled);
app.component("SettingOutlined", SettingOutlined);
app.component("CheckCircleOutlined", CheckCircleOutlined);
app.component("CloseCircleOutlined", CloseCircleOutlined);
app.component("SmileOutlined", SmileOutlined);
app.component("FrownOutlined", FrownOutlined);
app.component("UploadOutlined", UploadOutlined);
app.component("CaretDownFilled", CaretDownFilled);
app.component("MinusOutlined", MinusOutlined);
app.component("PlusCircleOutlined", PlusCircleOutlined);

app.use(VueCryptojs).use(i18n).use(store).use(router).use(Antd).mount("#app");
