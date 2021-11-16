---
title: "Js Uitls"
date: 2021-11-04T10:53:03+08:00
draft: false
---

### 获取文件名和后缀名称

1. 使用subtring() 截取字符串，对于文件名中会出现多个点的很有用，从最后一个点的地方截取

```js
// 获取文件名
getFileName (name) {
  return name.substring(0, name.lastIndexOf("."))
},
// 获取 .后缀名
getExtension (name) {
  return name.substring(name.lastIndexOf("."))
}
// 只获取后缀名
getExtension (name) {
  return name.substring(name.lastIndexOf(".")+1)
}
```　　

1. 使用正则，对只会出现一个点的适用

```js
  const a="c:\\windows\\abc.txt";
  const reg = /([^\\/]+)\.([^\\/]+)/i;
  reg.test(a);
```