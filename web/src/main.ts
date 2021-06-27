import {createApp} from 'vue'
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
import App from './App.vue'

import router from './router'
import service from "./utils/request";
const app = createApp(App)

app.config.globalProperties.$https = service;
app.use(Antd)
app.use(router)
app.mount('#app')