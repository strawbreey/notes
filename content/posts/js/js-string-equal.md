---
title: "Js String Equal"
date: 2020-11-16T19:36:59+08:00
draft: false
---

判断是不是字符串：

1 基本方法：

```js
typeof(str)=='string'
```

2 利用js原生函数：

```js
Object.prototype.toString  // ƒ toString() { [native code] }

Object.prototype.toString.call(str)=="[object String]"

```

（1）在Object.prototype这个this（上下文环境）中执行toString原生函数，会把里边的环境变量类型打印出来。

```js
Object.prototype.toString() -->执行结果-->"[object Object]"
```

（2）如果我们改变this（上下文环境），就能打印出当前环境变量类型，根据这个类型来判断。

```js
Object.prototype.toString.call(str) -->执行结果-->"[object String]"
```