---
title: "Webpack Remove Console"
date: 2020-11-30T17:36:09+08:00
draft: false
---

去除 console 是交由 UglifyJS 处理的.webpack 内置了 webpack.optimize.UglifyJsPlugin, 也提供了 UglifyjsWebpackPlugin 插件形式调用, 二者的区别是: webpack.optimize.UglifyJSPlugin 基于的是 UglifyjsWebpackPlugin v0.4.6 版本, 而单独使用插件的版本可以获得更高的版本.  个人比较追新, 所以说一下 UglifyjsWebpackPlugin 的使用方式, 配置项比内置的丰富了些:在你的 webpack.config.js 中 (只列出关键点):

```js
const UglifyJsPlugin = require('uglifyjs-webpack-plugin');
module.exports = {
  plugins: [
    new UglifyJsPlugin({
      uglifyOptions: {
        compress: {
          drop_console: true
        }
      }
    })
  ]
}

```

### 安装webpack 插件
babel-plugin-transform-remove-console


### 去掉console 

  window.console.log = function () { };
  window.console.info = function () { };
  window.console.warn = function () { };
  window.console.error = function () { };

### 参考资料

- [uglifyjs-webpack-plugin](https://github.com/webpack-contrib/uglifyjs-webpack-plugin#uglifyoptions)
- [babel-plugin-transform-remove-console](https://github.com/Riokai/babel-plugin-transform-remove-console)