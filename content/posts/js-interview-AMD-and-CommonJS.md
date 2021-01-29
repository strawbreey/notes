---
title: "Js Interview AMD and CommonJS"
date: 2021-01-27T10:33:36+08:00
draft: false
---

## 说说你对 AMD 和 CommonJS 的了解。

它们都是实现模块体系的方式，直到 ES2015 出现之前，JavaScript 一直没有模块体系。CommonJS 是同步的，而 AMD（Asynchronous Module Definition）从全称中可以明显看出是异步的。CommonJS 的设计是为服务器端开发考虑的，而 AMD 支持异步加载模块，更适合浏览器。

我发现 AMD 的语法非常冗长，CommonJS 更接近其他语言 import 声明语句的用法习惯。大多数情况下，我认为 AMD 没有使用的必要，因为如果把所有 JavaScript 都捆绑进一个文件中，将无法得到异步加载的好处。此外，CommonJS 语法上更接近 Node 编写模块的风格，在前后端都使用 JavaScript 开发之间进行切换时，语境的切换开销较小。

我很高兴看到 ES2015 的模块加载方案同时支持同步和异步，我们终于可以只使用一种方案了。虽然它尚未在浏览器和 Node 中完全推出，但是我们可以使用代码转换工具进行转换。

### 参考资料

- [relation-between-commonjs-amd-and-requirejs](https://stackoverflow.com/questions/16521471/relation-between-commonjs-amd-and-requirejs)