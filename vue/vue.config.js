const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
	transpileDependencies: true,
  devServer: {
    port: 8080,
  },
  chainWebpack: config =>{
    config.plugin('html')
    .tap(args=>{
      args[0].title = "时之恋人";
      return args;
    })
  },
});
