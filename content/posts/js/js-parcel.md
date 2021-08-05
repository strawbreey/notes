---
title: "Js Parcel"
date: 2021-08-05T10:44:47+08:00
draft: false
---

Parcel 是 Web 应用打包工具，适用于经验不同的开发者。 不同于 `webpack` 它利用多核处理提供了极快的速度，并且不需要任何配置。 

首先通过 Yarn 或者 npm 安装 Parcel ：

Yarn:

`yarn global add parcel-bundler`

npm:

`npm install -g parcel-bundler`

在你正在使用的项目目录下创建一个 package.json 文件：
```
yarn init -y
or

npm init -y
```

Parcel 可以使用任何类型的文件作为入口，但是最好还是使用 HTML 或 JavaScript 文件。如果在 HTML 中使用相对路径引入主要的 JavaScript 文件，Parcel 也将会对它进行处理将其替换为相对于输出文件的 URL 地址。

接下来，创建一个 index.html 和 index.js 文件。

<html>
  <body>
    <script src="./index.js"></script>
  </body>
</html>

console.log('hello world')
Parcel 内置了一个当你改变文件时能够自动重新构建应用的开发服务器，而且为了实现快速开发，该开发服务器支持热模块替换。只需要在入口文件指出：

parcel index.html
现在在浏览器中打开 http://localhost:1234/。如果模块热重载没有生效，你可能需要配置你的编辑器。你也可以使用 -p <port number> 选项覆盖默认的端口。 如果没有自己的服务器可使用开发服务器，或者你的应用程序完全由客户端呈现。如果有自己的服务器，你可以在watch 模式下运行 Parcel 。当文件改变它仍然会自动重新构建并支持热替换，但是不会启动 web 服务。

parcel watch index.html
你也能使用createapp.dev在浏览器中创建一个 Parcel 项目。选择你需要的特性列如 React， Vue，Typescript 和 CSS，然后你将会看到项目实时生成。你能通过这个工具去学习如何怎么建立一个新的项目并且你也能下载这个项目作为一个 zip 文件然后立即开始写代码。

### 参考资料

- https://zh.parceljs.org/how_it_works.html