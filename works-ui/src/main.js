import { createApp } from 'vue';
import Antd from 'ant-design-vue';
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
    BellOutlined, FontSizeOutlined,
} from '@ant-design/icons-vue';
import 'ant-design-vue/dist/antd.css';
import App from './App';
import router from './router'
import store from './store'
import i18n from './locales'

const app = createApp(App);//.use(loadLocaleMessages);
app.addComponents = function () {
    for(let i in arguments){
        const component = arguments[i];
        this.component(component.name, component)
    }
}
app.config.productionTip = true;
app.addComponents(
    BackwardFilled,
    HomeFilled,
    DeleteFilled,
    UserOutlined,
    PoweroffOutlined,
    ReloadOutlined, SyncOutlined,
    CameraOutlined,
    PaperClipOutlined,
    LockOutlined,
    BellOutlined,
    FontSizeOutlined,
)
// app.ad
app.use(store);
app.use(router);
app.use(i18n);
app.use(Antd);
app.mount('#app');

