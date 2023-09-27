import Vue from "vue";
import './plugins/axios'
import App from "./App.vue";
import router from "./router";
import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";
import "./assets/css/iconfonts/iconfont.css";
// import axios from "axios";
import store from "./store/index";
import BaiduMap from 'vue-baidu-map';

Vue.use(ElementUI, { size: "big" });

Vue.use(BaiduMap, { ak: 'c4Iwa3iL5rzVF9Wn21B3ckGPXPQoyNWB' });
axios.defaults.baseURL = "http://localhost:8889/";

Vue.prototype.$http = axios;
Vue.config.productionTip = false;

new Vue({
	router,
	store, // 将 Vuex Store 实例挂载到 Vue 实例中
	render: (h) => h(App),
}).$mount("#app");


