import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './router'

import BMap from 'baidu-map'

export default {
  mounted() {
    // 初始化地图
    const map = new BMap.Map('map')
    // 设置地图中心点
    const point = new BMap.Point(116.404, 39.915)
    map.centerAndZoom(point, 15)
  }
}


const app = createApp(App)

app.use(router)
app.use(ElementPlus)
app.mount('#app')


