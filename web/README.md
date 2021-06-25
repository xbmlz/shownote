# Vue 3 + Typescript + Vite

## Ant Design Vue

```bash
npm i --save ant-design-vue@next --registry https://registry.npm.taobao.org
```

main.ts

```bash
import {createApp} from 'vue'
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
import App from './App.vue'

createApp(App)
    .use(Antd) 
    .mount('#app')
```

## vditor

```bash
npm install vditor --save --registry https://registry.npm.taobao.org
```