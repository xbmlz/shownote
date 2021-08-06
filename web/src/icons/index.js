import { createApp } from 'vue'

import SvgIcon from '@/components/SvgIcon'  // 引入自定义组件

createApp().component('svg-icon', SvgIcon)  // 注册到全局

const requireAll = requireContext => requireContext.keys().map(requireContext)
const req = require.context('./svgs', false, /\.svg$/)
requireAll(req)
