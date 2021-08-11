import {createApp} from 'vue'
import {
    ElButton,
    ElInput,
    ElSelect,
    ElAvatar,
    ElEmpty,
    ElTree,
    ElDialog,
    ElMessage,
    ElLoading,
    ElMessageBox,
    ElDropdown,
    ElDropdownItem,
    ElDropdownMenu,
} from 'element-plus';
import 'element-plus/packages/theme-chalk/src/base.scss';
import lang from 'element-plus/lib/locale/lang/zh-cn'
import 'dayjs/locale/zh-cn'
import locale from 'element-plus/lib/locale'
import '@/styles/index.scss'

import App from './App.vue'

import router from './router'
import service from "./utils/request";

// 设置语言
// locale.use(lang)

import 'vite-plugin-svg-icons/register';
// 需要全局引入再添加
import svgIcon from './components/SvgIcon/index.vue' // 全局svg图标组件


const app = createApp(App)

app.component('svg-icon', svgIcon)

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
app.use(ElMessageBox)
app.use(ElDropdown)
app.use(ElDropdownMenu)
app.use(ElDropdownItem)
app.use(router)
app.mount('#app')
