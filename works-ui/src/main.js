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
    BellOutlined, FontSizeOutlined, EditOutlined, HomeOutlined,BarcodeOutlined,PlusOutlined
} from '@ant-design/icons-vue';
import 'ant-design-vue/dist/antd.css';
import App from './App';
import router from './router'
import store from './store'
import i18n from './locales'


import {faCameraRetro, faUserSecret} from '@fortawesome/free-solid-svg-icons'
import {faUbuntu, faCentos} from "@fortawesome/free-brands-svg-icons";
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { library } from '@fortawesome/fontawesome-svg-core'

library.add(faCameraRetro)
library.add(faUbuntu)
library.add(faCentos)
library.add(faUserSecret)

const app = createApp(App);//.use(loadLocaleMessages);
app.addComponents = function () {
    for(let i in arguments){
        const component = arguments[i];
        this.component(component.name, component)
    }
}
app.config.productionTip = false;
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
    EditOutlined,
    HomeOutlined,
    BarcodeOutlined,
    PlusOutlined

)

app.component('font-awesome-icon', FontAwesomeIcon)
app.use(store);
app.use(router);
app.use(i18n);
app.use(Antd);
app.mount('#app');

