import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";
import "./assets/css/iconfonts/iconfont.css";
import axios from "axios";
import store from "./store/index";
import { mapState } from "vuex";
import { createStore } from "vuex";

Vue.use(ElementUI, { size: "big" });

// 创建一个 Axios 实例
const instance = axios.create({
	baseURL: "localhost:8080", // 设置全局请求路径
});

Vue.prototype.$https = instance; // 将实例挂载到 Vue 的原型上
Vue.config.productionTip = false;

new Vue({
	router,
	store, // 将 Vuex Store 实例挂载到 Vue 实例中
	render: (h) => h(App),
}).$mount("#app");
