import Vue from "vue"
import VueRouter from "vue-router"

Vue.use(VueRouter)

const routes = [
  {
    path: "/",
    name: "home",
    component: () => import(/* webpackChunkName: "about" */ "../views/HomeView.vue")
  },
  {
    path: "/about",
    name: "about",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ "../views/AboutView.vue")
  },
  {
    path: "/detail",
    name: "detail",
    component: () => import("../views/DetailView.vue")
  },
  {
    path: "/login",
    name: "login",
    component: () => import("../views/LogInView.vue")
  },
  {
    path: "/map",
    name: "map",
    component: () => import("../views/MapView.vue")
  },
  {
    path: "/changeuserinfo",
    name: "changeuserinfo",
    component: () => import("../views/ChangeUserInfoView.vue")
  },
]

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
})

export default router;
