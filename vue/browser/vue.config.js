/**
 * 配置参考:
 * https://cli.vuejs.org/zh/config/
 */
const path = require("path");

function resolve(dir) {
  return path.join(__dirname, dir);
}

let url = "http://localhost:8080";

module.exports = {
  publicPath: "./", // 打包基本路径
  // outputDir: "dist", // 编译后输出路径
  // assetsDir: "static", // 放置静态资源目录
  // indexPath: "index.html", // html输出路径 (相对于 outputDir)
  // filenameHashing: true, // 文件哈希名字（这个选项设为 false 来关闭文件名哈希）
  configureWebpack: {
    devtool: process.env.NODE_ENV === "production" ? "false" : "source-map",
  },
  // lintOnSave是否在保存的时候使用 `eslint-loader` 进行检查,有效的值：`ture` | `false` | `"error"`  当设置为 `"error"` 时，检查出的错误会触发编译失败
  lintOnSave: false,
  productionSourceMap: false, // 生产环境是否需要打包map
  chainWebpack: (config) => {
    // 配置路径别名
    config.resolve.alias
      .set("@", resolve("src"))
      .set("@components", resolve("src/components"))
      .set("@api", resolve("src/api"))
      .set("@assets", resolve("src/assets"))
      .set("@server", resolve("src/server"));
    const entry = config.entry("app");
    entry.add("babel-polyfill").end();
    entry.add("classlist-polyfill").end();
  },
  configureWebpack: (config) => {
    if (process.env.NODE_ENV === "production") {
      // 为生产环境修改配置...
      return {
        plugins: [
          // new CompressionWebpackPlugin({
          //   filename: "[path].gz[query]",
          //   algorithm: "gzip",
          //   test: new RegExp(
          //     "\\.(" + productionGzipExtensions.join("|") + ")$"
          //   ),
          //   threshold: 1024, // 只有大小大于该值的资源会被处理,当前配置为对于超过1k的数据进行处理，不足1k的可能会越压缩越大
          //   minRatio: 0.99, // 只有压缩率小于这个值的资源才会被处理
          //   deleteOriginalAssets: true // 删除原文件
          // })
        ],
      };
    } else {
      // 为开发环境修改配置...
    }
  },
  // 配置转发代理cli 3.5 以后不需要再配置
  devServer: {
    disableHostCheck: true,
    port: 8085,
    // open: true, // 配置浏览器自动打开
    host: 'localhost', // 默认是localhost
    proxy: {
      "^/api": {
        target: url,
        ws: false, // 需要websocket 开启
        pathRewrite: {
          "^/api": "/api",
        },
        changeOrigin: true,
      },
    },
  },
};
