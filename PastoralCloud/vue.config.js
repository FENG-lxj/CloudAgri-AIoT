const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  configureWebpack: {
         resolve: {
           fallback: {
           "url": require.resolve("url/")
          }
         }
       },
  transpileDependencies: true,

  devServer: {
    port: 8090, // 此处修改你想要的端口号
  },
})
