import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import styleImport from 'vite-plugin-style-import'
const path = require('path')

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        styleImport({
            libs: [{
                libraryName: 'element-plus',
                esModule: true,
                ensureStyleFile: true,
                resolveStyle: (name) => {
                    name = name.slice(3)
                    return `element-plus/packages/theme-chalk/src/${name}.scss`;
                },
                resolveComponent: (name) => {
                    return `element-plus/lib/${name}`;
                },
            }]
        })
    ],
    server: {
        host: '0.0.0.0'
    },
    resolve: {
        // 导入文件夹别名
        alias: {
            "@": path.resolve(__dirname, "src"),
        }
    }
})
