---
title: "Js CodePointAt"
date: 2020-11-22T23:29:51+08:00
draft: false
---

codePointAt() 方法返回 一个 Unicode 编码点值的非负整数。

如果在指定的位置没有元素则返回 undefined 。如果在索引处开始没有UTF-16 代理对，将直接返回在那个索引处的编码单元。

Surrogate Pair是UTF-16中用于扩展字符而使用的编码方式，是一种采用四个字节(两个UTF-16编码)来表示一个字符，称作代理对。

```js
str.codePointAt(pos)
```

```js
'ABC'.codePointAt(1);          // 66
'\uD800\uDC00'.codePointAt(0); // 65536

'XYZ'.codePointAt(42); // undefined
```