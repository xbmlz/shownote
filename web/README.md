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

template

```html
<div id="vditor"></div>
```

ts

```javascript
import { ref, defineComponent, onMounted } from "vue";
import Vditor from "vditor";

export default defineComponent({
  name: "App",
  setup() {
    let vditor = ref(null);
    onMounted(() => {
      vditor = new Vditor("vditor", {
        height: 360,
        toolbarConfig: {
          pin: true,
        },
        cache: {
          enable: false,
        },
        after: () => {
          if (vditor) vditor.setValue("hello, Vditor + Vue!");
        },
      });
    });
    const getValue = () => {
      alert(vditor.getValue())
    };
    return {
      getValue,
    };
  },
});
```