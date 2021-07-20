import {createApp} from 'vue'
import {ElButton, ElSelect, ElAvatar, ElEmpty, ElInput} from 'element-plus';
import App from './App.vue'

import router from './router'
import service from "./utils/request";

const app = createApp(App)

app.config.globalProperties.$https = service;
app.use(ElButton)
app.use(ElSelect)
app.use(ElAvatar)
app.use(ElEmpty)
app.use(ElInput)
app.use(router)
app.mount('#app')
