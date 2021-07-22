import {createApp} from 'vue'
import {ElButton, ElInput, ElSelect, ElAvatar, ElEmpty, ElTree, ElDialog, ElMessage , ElLoading} from 'element-plus';
import 'element-plus/packages/theme-chalk/src/base.scss';
import lang from 'element-plus/lib/locale/lang/zh-cn'
import 'dayjs/locale/zh-cn'
import locale from 'element-plus/lib/locale'
import '@/styles/index.scss'

import App from './App.vue'

import router from './router'
import service from "./utils/request";

// 设置语言
locale.use(lang)

const app = createApp(App)

app.config.globalProperties.$https = service;
app.use(ElButton)
app.use(ElInput)
app.use(ElSelect)
app.use(ElAvatar)
app.use(ElEmpty)
app.use(ElTree)
app.use(ElDialog)
app.use(ElMessage)
app.use(ElLoading)
app.use(router)
app.mount('#app')
