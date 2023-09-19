<template>
  <div id="/">
    <!-- 页头 -->
    <el-page-header @back="goback" :content="pageHeader" style="margin-top: 20px; margin-bottom: 30px;"></el-page-header>
    <router-view />
  </div>
</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>

<script>
export default {
  name: "App",
  data() {
    return {
      pageHeader: "主界面",
    }
  },
  mounted() {
    this.pageHeader = this.getPageHeader(this.$route.name);
  },
  watch: {
    "$route"(to) {
      this.pageHeader = this.getPageHeader(to.name);
    }
  },
  methods: {
    goback() {
      history.back();
    },
    getPageHeader(routeName) {
      console.log(routeName);
      if (routeName === "home" || routeName === "login") {
        return "主界面";
      } else if (routeName === "element") {
        return "element界面";
      } else if (routeName === "about") {
        return "关于页面";
      } else if (routeName === "map") {
        return "地图界面";
      }
    },
  },
}
</script>