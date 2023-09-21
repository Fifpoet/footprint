import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";
import "./assets/css/iconfonts/iconfont.css";
import axios from "axios";


// 创建一个 Axios 实例
const instance = axios.create({
	baseURL: "localhost:8080", // 设置全局请求路径
  });
  
  Vue.prototype.$http = instance; // 将实例挂载到 Vue 的原型上
Vue.config.productionTip = false;

Vue.use(ElementUI, { size: "big" });

new Vue({
	router,
	render: (h) => h(App),
}).$mount("#app");
