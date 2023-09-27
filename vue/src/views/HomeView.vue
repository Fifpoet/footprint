<template>
  <div>
    <el-container>
      <!-- 侧边栏 -->
      <el-aside :width="asideWidth" style="width: 150px; height: 100vh; background-color: #eee; border: 1px solid gray;">
        <!-- logo区域 -->
        <div
          style="height: 150px; width: 150px; display: flex; align-items: center; justify-content: center; background-color: white"
          class="logo">
          <svg aria-hidden="true" viewBox="0 0 16 16" version="1.1" data-view-component="true"
            class="octicon octicon-mark-github v-align-middle color-fg-default">
            <path
              d="M8 0c4.42 0 8 3.58 8 8a8.013 8.013 0 0 1-5.45 7.59c-.4.08-.55-.17-.55-.38 0-.27.01-1.13.01-2.2 0-.75-.25-1.23-.54-1.48 1.78-.2 3.65-.88 3.65-3.95 0-.88-.31-1.59-.82-2.15.08-.2.36-1.02-.08-2.12 0 0-.67-.22-2.2.82-.64-.18-1.32-.27-2-.27-.68 0-1.36.09-2 .27-1.53-1.03-2.2-.82-2.2-.82-.44 1.1-.16 1.92-.08 2.12-.51.56-.82 1.28-.82 2.15 0 3.06 1.86 3.75 3.64 3.95-.23.2-.44.55-.51 1.07-.46.21-1.61.55-2.33-.66-.15-.24-.6-.83-1.23-.82-.67.01-.27.38.01.53.34.19.73.9.82 1.13.16.45.68 1.31 2.69.94 0 .67.01 1.3.01 1.49 0 .21-.15.45-.55.38A7.995 7.995 0 0 1 0 8c0-4.42 3.58-8 8-8Z">
            </path>
          </svg>
        </div>

        <!-- 菜单区域 -->
        <el-menu :collapse="isCollapse" router background-color="#eee" text-color="rgb(0,0,0,0.75)"
          :default-active="$route.path"><!-- 动态绑定路由 -->
          <el-menu-item index="/">
            <template slot="title">
              <i class="el-icon-house"></i>
              <span class="button-text">系统首页</span>
            </template>
          </el-menu-item>
          <el-menu-item @click="toMap">
            <template slot="title">
              <i class="el-icon-map-location"></i>
              <span class="button-text">足迹地图</span>
            </template>
          </el-menu-item>
          <el-menu-item index="/about">
            <template slot="title">
              <i class="el-icon-warning-outline"></i>
              <span class="button-text">关于我们</span>
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
          style="height: 55px; box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2); text-align: right; background-color: #fefefe;">
          <div style="height: 60px;">
            <el-button type="text" @click="toLogInView" v-if="!this.$store.state.isLogIn">立即登录</el-button>

            <el-dropdown @command="handleDropdownCommand" v-if="this.$store.state.isLogIn">
              <div style="display: flex; align-items: center;">
                <svg style="width: 50px; height: 50px;" aria-hidden="true" viewBox="0 0 16 16" version="1.1"
                  data-view-component="true" class="octicon octicon-mark-github v-align-middle color-fg-default">
                  <path
                    d="M8 0c4.42 0 8 3.58 8 8a8.013 8.013 0 0 1-5.45 7.59c-.4.08-.55-.17-.55-.38 0-.27.01-1.13.01-2.2 0-.75-.25-1.23-.54-1.48 1.78-.2 3.65-.88 3.65-3.95 0-.88-.31-1.59-.82-2.15.08-.2.36-1.02-.08-2.12 0 0-.67-.22-2.2.82-.64-.18-1.32-.27-2-.27-.68 0-1.36.09-2 .27-1.53-1.03-2.2-.82-2.2-.82-.44 1.1-.16 1.92-.08 2.12-.51.56-.82 1.28-.82 2.15 0 3.06 1.86 3.75 3.64 3.95-.23.2-.44.55-.51 1.07-.46.21-1.61.55-2.33-.66-.15-.24-.6-.83-1.23-.82-.67.01-.27.38.01.53.34.19.73.9.82 1.13.16.45.68 1.31 2.69.94 0 .67.01 1.3.01 1.49 0 .21-.15.45-.55.38A7.995 7.995 0 0 1 0 8c0-4.42 3.58-8 8-8Z">
                  </path>
                </svg>
                <span style="width: 65px;">{{ this.$store.state.username }}</span>
              </div>

              <el-dropdown-menu slot="dropdown">
                <!-- <el-dropdown-item command="changeUserInfo">修改个人信息</el-dropdown-item> -->
                <el-dropdown-item command="exitLogIn">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </div>
        </el-header>

        <!-- 主体区域 -->
        <el-main>
          <!-- 帖子列表 -->
          <div class="post-list">
            <!-- 使用v-for循环渲染帖子 -->
            <div v-for="(post, index) in posts" :key="index" class="post-item">
              <div class="post-image">
                <img :src="post.imageUrl" alt="Post Image" />
              </div>
              <div class="post-details">
                <h2 class="post-title">{{ post.title }}</h2>
                <p class="post-description">{{ post.description }}</p>
              </div>
            </div>
          </div>


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
      posts: [
        {
          imageUrl: require('@/assets/logo.png'),
          title: '帖子标题1',
          description: '帖子正文段落1。这里可以放帖子的详细内容。',
        },
        {
          imageUrl: require('@/assets/logo.png'),
          title: '帖子标题2',
          description: '帖子正文段落2。这里可以放帖子的详细内容。',
        },
        {
          imageUrl: require('@/assets/logo.png'),
          title: '帖子标题3',
          description: '帖子正文段落3。这里可以放帖子的详细内容。',
        },
      ],

    };
  },
  methods: {
    toLogInView() {
      this.$router.push("/login");
    },
    handleDropdownCommand(command) {
      // if (command === "changeUserInfo") {   //修改个人信息
      //   this.$router.push("/changeuserinfo");
      // } else {
      if (command === "exitLogIn") {   //退出登录      

        this.$store.commit("isLogInFalse");
        this.$message.success("退出登录成功！");
      }
      //}
    },
    toMap() {
      if (this.$store.state.isLogIn) {
        this.$router.push("/map");
      } else {

        this.$router.push("/login");
      }
    }
  },
}
</script>

<style scoped>
.el-main {
  background: linear-gradient(135deg, #ccfaf9, #fadefa);
}
.el-icon {
  /* 设置字体大小 */
  font-size: 24px;
  /* 设置字体颜色 */
  color: #ff0000;
}

.button-text {
  font-family: Arial, sans-serif; /* 设置字体族 */
  font-size: 16px; /* 设置字体大小 */
  font-weight: bold; /* 设置字体粗细 */
  color: #5548e3; /* 设置字体颜色 */
}

.el-menu-item:hover,
.el-menu-item:hover i,
.el-submenu__title:hover,
.el-submenu__title:hover i {
  color: #000 !important;
  background-color: #37d6fa !important;
  border-radius: 8px;
  margin: 0 1px;
}

.el-menu-item.is-active,
.el-submenu__title.is-active {
  color: #000 !important;
  background-color: #acfaf9!important;
  border-radius: 8px;
  margin: 0 1px;
}

.el-page-header {
  margin: 20px 0;
}
</style>