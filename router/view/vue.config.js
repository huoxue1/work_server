const AutoImport = require("unplugin-auto-import/webpack");
const Components = require("unplugin-vue-components/webpack");
const { ElementPlusResolver } = require("unplugin-vue-components/resolvers");


module.exports = {
    // 配置网页的路径
  publicPath: '/static/view/dist/',
    // 配置生成的html名字
  indexPath: 'home.html',
    configureWebpack: {
        plugins: [
            AutoImport({
                resolvers: [ElementPlusResolver()],
            }),
            Components({
                resolvers: [ElementPlusResolver()],
            }),
        ],
    },
  chainWebpack: config => {
    config
        .plugin('html')
        .tap(args => {
            // 配置前端title
          args[0].title= 'work'
          return args
        })
  },
}
