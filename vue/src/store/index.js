// store/index.js
import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

const store = new Vuex.Store({
	state: {
		// 定义你的状态
		counter: 0,
		username: "",
		isLogIn: false,
	},
	mutations: {
		// 定义修改状态的方法
		increment(state) {
			state.counter++;
		},
		isLogInTrue(state) {
			state.isLogIn = true;
		},
		isLogInFalse(state) {
			state.isLogIn = false;
		},
		changeUserName(state, name) {
			state.username = name;
		},
	},
	actions: {
		// 定义异步操作和触发 mutations 的方法
		incrementAsync(context) {
			setTimeout(() => {
				context.commit("increment");
			}, 1000);
		},
	},
	getters: {
		// 定义计算属性
		doubledCounter(state) {
			return state.counter * 2;
		},
	},
});

export default store;
