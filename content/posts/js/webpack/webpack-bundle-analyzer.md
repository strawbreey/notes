---
title: "Webpack Bundle Analyzer"
date: 2021-08-24T17:44:01+08:00
draft: false
---

Webpack Bundle Analyzer (分析代码)

对于 Web app 来说，高性能总是最高优先级，对于 Angular 也不例外。但是随着应用复杂度的不断增长，我们如何才能知道哪些内容打包到了应用中呢？如何跟踪包的尺寸？我们不希望一次发送太多的 JavaScript ，以至于拖慢应用的速度。

过大尺寸的 JavaScript 包是丧失用户欢心的良药。不仅是拖慢了下载效率，而且要花费更多的时间在浏览器中分析然后执行。为了保持应用的速度，我们需要确保包尺寸足够小（250k 或更小），并在适当的时间加载。

在这篇文章中，我们将使用 Angular CLI 和一些简单的命令来获得我们发布产品的 Angular 应用包的详细报告。


1. 安装 `webpack-bundle-analyzer`

```bash
# NPM
npm install --save-dev webpack-bundle-analyzer
# Yarn
yarn add -D webpack-bundle-analyzer
```


2. 生成stats.json 文件

```bash
ng build --prod --stats-json

# 查看 stats.json 的位置 /dist/stats.json
```

3. 配置npm 脚本


为方便使用，安装之后，在 package.json 中的 scripts 部分，添加下面的行来创建自定义的 npm 命令 
```json
"bundle-report": "webpack-bundle-analyzer dist/stats.json"
```

4. 执行脚本
```bash
npm run bundle-report
```
isset()