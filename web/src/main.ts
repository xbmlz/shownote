import {createApp} from 'vue'
import {ElButton, ElInput, ElSelect, ElAvatar, ElEmpty, ElTree, ElDialog} from 'element-plus';
import 'element-plus/packages/theme-chalk/src/base.scss';
import App from './App.vue'

import router from './router'
import service from "./utils/request";

const app = createApp(App)

app.config.globalProperties.$https = service;
app.use(ElButton)
app.use(ElInput)
app.use(ElSelect)
app.use(ElAvatar)
app.use(ElEmpty)
app.use(ElTree)
app.use(ElDialog)
app.use(router)
app.mount('#app')
