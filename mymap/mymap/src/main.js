import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(router)

app.mount('#app')

//import BMap from 'baidu-map'

export default {
  mounted() {
    // 初始化地图
    const map = new BMap.Map('map')
    // 设置地图中心点
    const point = new BMap.Point(116.404, 39.915)
    map.centerAndZoom(point, 15)
  }
}

