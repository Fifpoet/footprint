import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";
import "./assets/css/iconfonts/iconfont.css";

Vue.config.productionTip = false;

Vue.use(ElementUI, { size: "big" });

new Vue({
	router,
	render: (h) => h(App),
}).$mount("#app");

import BaiduMap from 'vue-baidu-map'
Vue.use(BaiduMap, { ak: 'c4Iwa3iL5rzVF9Wn21B3ckGPXPQoyNWB' })


