<template>
  <div>
    <el-container>
      <!-- 侧边栏 -->
      <el-aside :width="asideWidth" style="width:'200px'; height: 100vh; background-color: #eee; border: 1px solid gray;">
        <!-- logo区域 -->
        <div
          style="height: 200px; width: 200px; display: flex; align-items: center; justify-content: center; background-color: lightpink"
          class="logo">
          <img src="../assets/孙尚香时之恋人.jpg" alt="logo" style="height: 120px; width: 120px;">
        </div>

        <!-- 菜单区域 -->
        <el-menu :collapse="isCollapse" router background-color="#eee" text-color="rgb(0,0,0,0.75)"
          :default-active="$route.path"><!-- 动态绑定路由 -->
          <el-menu-item index="/">
            <template slot="title">
              <i class="el-icon-house"></i>
              <span>系统首页</span>
            </template>
          </el-menu-item>
          <el-menu-item index="/map">
            <template slot="title">
              <i class="el-icon-map-location"></i>
              <span>地图</span>
            </template>
          </el-menu-item>
          <el-menu-item index="/about">
            <template slot="title">
              <i class="el-icon-warning-outline"></i>
              <span>关于</span>
            </template>
          </el-menu-item>
          <!-- 
          <el-submenu>
            <template slot="title">
              <i class="el-icon-menu"></i>
              <span>信息管理</span>
            </template>
            <el-menu-item index="/about">
              <template slot="title">
                <i class="el-icon-user"></i>
                <span>用户信息</span>
              </template>
            </el-menu-item>
          </el-submenu> -->

        </el-menu>

      </el-aside>

      <!-- 主体 -->
      <el-container>
        <!-- 头部区域 -->
        <el-header
          style="height: 60px; box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2); text-align: right; background-color: #fefefe;">
          <div style="height: 60px;">
            <el-button type="text" @click="toLogInView" v-if="!this.$store.state.isLogIn">立即登录</el-button>

            <el-dropdown @command="handleDropdownCommand" v-if="this.$store.state.isLogIn">
              <div style="display: flex; align-items: center;">
                <img src="../assets/孙尚香时之恋人.jpg" alt="logo" style="height: 40px; width: 40px; margin: 0 20px;">
                <span>{{ this.user.username }}</span>
              </div>
                
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item command="changeUserInfo">修改个人信息</el-dropdown-item>
                <el-dropdown-item command="exitLogIn">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </div>
        </el-header>

        <!-- 主体区域 -->
        <el-main>
          主体区域
        </el-main>

      </el-container>
    </el-container>



  </div>
</template>

<script>
export default {
  name: "HomeView",
  data() {
    return {
      isCollapse: false,  //默认不收缩
      asideWidth: "200px",
      user: {
        username: "时之恋人",
      },
    };
  },
  methods: {
    toLogInView() {
      this.$router.push("/login");
    },
    handleDropdownCommand(command) {
      if (command === "changeUserInfo") {   //修改个人信息
        this.$router.push("/changeuserinfo");
      } else {
        if (command === "exitLogIn") {   //退出登录      

          this.$store.commit("isLogInFalse");
          this.$message.success("退出登录成功！");
        }
      }
    }
  },
}
</script>

<style scoped>
.el-menu-item:hover,
.el-menu-item:hover i,
.el-submenu__title:hover,
.el-submenu__title:hover i {
  color: #000 !important;
  background-color: rgb(173, 216, 230, 0.5) !important;
  border-radius: 8px;
  margin: 0 1px;
}

.el-menu-item.is-active,
.el-submenu__title.is-active {
  color: #000 !important;
  background-color: lightblue !important;
  border-radius: 8px;
  margin: 0 1px;
}

.el-page-header {
  margin: 20px 0;
}
</style>