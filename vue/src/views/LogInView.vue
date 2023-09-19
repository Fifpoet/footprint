<template>
  <div>

    <!-- 登录窗口 -->
    <el-dialog title="用户登录" :visible.sync="loginDialogVisible" class="login-container" :close-on-click-modal="false"
      :center="true">
      <el-form :model="loginForm" label-width="80px" ref="loginForm">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="loginForm.username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input show-password type="password" v-model="loginForm.password" placeholder="请输入密码"
            auto-complete="off"></el-input>
        </el-form-item>
      </el-form>
      <template v-slot:footer>
        <div>
          <el-row>
            <el-button type="primary" @click="submitLoginForm">登录</el-button>
            <el-button @click="cancelLoginDialog">取消</el-button>
          </el-row>
          <el-row>
            <el-button type="text" @click="showRegisterDialog">立即注册</el-button>
            <!-- <el-button type="text" @click="forgetPassword" style="margin-left: 150px;">忘记密码</el-button> -->
          </el-row>
        </div>
      </template>

    </el-dialog>

    <!-- 注册窗口 -->
    <el-dialog v-model="registerDialogVisible" title="用户注册" :visible.sync="registerDialogVisible"
      class="register-container" :center="true">
      <el-form :model="registerForm" label-width="80px" ref="registerForm" :rules="registerRules">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="registerForm.username" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input show-password type="password" v-model="registerForm.password" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input show-password type="password" v-model="registerForm.confirmPassword" auto-complete="off"></el-input>
        </el-form-item>
      </el-form>
      <template v-slot:footer>
        <div class="dialog-footer">
          <el-button type="primary" @click="registerUser" :loading="registerSubmitting">注册</el-button>
          <el-button @click="closeRegisterDialog">取消</el-button>
        </div>
      </template>
    </el-dialog>

  </div>
</template>


<script>
export default {
  name: "LogInView",
  data() {
    return {
      loginForm: { //提交的表单
        username: "",
        password: "",
      },
      registerForm: { //注册新用户的表单
        username: "",
        password: "",
        confirmPassword: "",
      },
      loginDialogVisible: true,
      registerDialogVisible: false,
      registerRules: {
        username: [
          { required: true, message: "请输入用户名", trigger: ["blur", "change"] }
        ],
        password: [
          { required: true, message: "请输入密码", trigger: ["blur", "change"] }
        ],
        confirmPassword: [
          { required: true, message: "请再次输入密码", trigger: ["blur", "change"] },
          { validator: this.validateConfirmPassword, trigger: ["blur", "change"] },
        ],
      },
      registerSubmitting: false,
    };
  },
  methods: {
    cancelLoginDialog() {
      this.$router.push("/");
    },
    showRegisterDialog() {
      this.registerDialogVisible = true;
    },
    closeRegisterDialog() {
      this.registerDialogVisible = false;
    },
    submitLoginForm() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          const loginData = {
            username: this.loginForm.username,
            password: this.loginForm.password
          };

          // 调用登录接口，发送登录请求
          api.login(loginData)
            .then(response => {
              if (response.data.success) {
                // 登录成功
                this.$message.success("登录成功！");
                this.closeLoginDialog();
              } else {
                // 登录失败
                this.$message.error("登录失败！");
              }
            })
            .catch(error => {
              // 处理登录请求错误 
              console.log(error);
              this.$message.error("登录失败！");
            });
        } else {
          this.$message.error("登录失败！");
          return false;
        }
      });
    },
    registerUser() {
      //注册新用户
      this.$refs.registerForm.validate((valid) => {
        if (valid) {
          // 表单验证通过，开始注册
          this.registerSubmitting = true;
          setTimeout(() => {
            // 模拟注册逻辑
            this.registerSubmitting = false;
            this.registerDialogVisible = false;
            this.$message.success("注册成功！");
          }, 1500);
        } else {
          // 表单验证失败，提示错误信息
          this.$message.error("请填写完整的注册信息！");
        }
      });
    },
    validateConfirmPassword(rule, value, callback) {
      if (value !== this.registerForm.password) {
        callback(new Error("两次输入的密码不一致"));
      } else {
        callback();
      }
    },
    /* 没有忘记密码可言 */
    // forgetPassword() {
    //   //忘记密码
    // },
  },
}

</script>


<style scoped>
.login-container {
  margin: auto;
  width: 1000px;
  text-align: center;
}

.register-container {
  margin: auto;
  width: 1000px;
  text-align: center;
}

.dialog-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>