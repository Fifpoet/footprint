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
