<template>
    <div>
      <el-container style="height: 100vh;">
        <el-header>Header</el-header>
  
        <el-container>
          <el-aside width='200px'>Aside</el-aside>
          <el-main>
            <el-scrollbar wrap-class="scrollbar-wrap">
              <!-- 大块的内容 -->
              <div class="content">
                <!-- Your content goes here -->
              </div>
            </el-scrollbar>
          </el-main>
        </el-container>
  
        <el-footer>Footer</el-footer>
      </el-container>
    </div>
  </template>
  
  <style>
  .scrollbar-wrap {
    height: 100%;
  }
  
  .content {
    /* 设置内容的样式 */
  }
  </style>
  
  <script>
  export default {
    name: "MapView",
  };
  </script>
  