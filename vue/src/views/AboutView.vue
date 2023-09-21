<template>
  <div class="about">
    <!-- 登录窗口 -->
    <el-dialog title="关于" :before-close="closeAboutView" :visible.sync="loginDialogVisible" class="login-container" :close-on-click-modal="false" center="true" style="width: 800px; text-align: center; margin: auto;">
      <div style="text-indent: 2em;">
        该网站由空信田展源队开发，成员是赵策、郝鹏洲、孙凯、崔意臻、韦凯五个人。该网站是用于记录您旅行的足迹，方便您打卡，让您能随时回顾自己曾经的岁月。您喜欢该网站就是对我们最大的认可。
      </div>

      <div style="margin-top: 40px; text-align: center;">
        <el-button type="primary" @click="closeAboutView">确定</el-button>
      </div>

    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "AboutView",
  data() {
    return {
      loginDialogVisible: "true",
    };
  },
  methods: {
    closeAboutView() {
      this.loginDialogVisible = "false";
      this.$router.push("/");
    },
  },
}
</script>