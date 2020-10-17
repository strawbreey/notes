---
title: "Js Optional Chaining"
date: 2020-10-17T17:52:12+08:00
draft: true
---

## 可选链 "?."


可选链 ?. 是一种访问嵌套对象属性的防错误方法。即使中间的属性不存在，也不会出现错误。

```js

let user = {}; // 这个 user 恰巧没有 address

alert(user?.address); // undefined!

alert(user?.address.street); // undefined!

```

- 不要过度使用可选链
我们应该只将 ?. 使用在一些东西可以不存在的地方。

例如，如果根据我们的代码逻辑，user 对象必须存在，但 address 是可选的，那么 user.address?.street 会更好。

所以，如果 user 恰巧因为失误变为 undefined，我们会知道并修复这个失误。否则，代码中的 error 在不恰当的地方被消除了，这会导致调试更加困难。

?. 前的变量必须已声明
如果未声明变量 user，那么 user?.anything 会触发一个错误：

// ReferenceError: user is not defined
user?.address;

### 短路效应

正如前面所说的，如果 ?. 左边部分不存在，就会立即停止运算（“短路效应”）。

所以，如果后面有任何函数调用或者副作用，它们均不会执行：

### 其它情况：?.()，?.[]

可选链 ?. 不是一个运算符，而是一个特殊的语法结构。它还可以与函数和方括号一起使用。

例如，将 ?.() 用于调用一个可能不存在的函数。

在下面这段代码中，有些用户具有 admin 方法，而有些没有：

```js
let user1 = {
  admin() {
    alert("I am admin");
  }
}

let user2 = {};

user1.admin?.(); // I am admin
user2.admin?.();
```

### 总结
可选链 ?. 语法有三种形式：

- obj?.prop —— 如果 obj 存在则返回 obj.prop，否则返回 undefined。
- obj?.[prop] —— 如果 obj 存在则返回 obj[prop]，否则返回 undefined。
- obj?.method() —— 如果 obj 存在则调用 obj.method()，否则返回 undefined。

正如我们所看到的，这些语法形式用起来都很简单直接。?. 检查左边部分是否为 null/undefined，如果不是则继续运算。

?. 链使我们能够安全地访问嵌套属性。

但是，我们应该谨慎地使用 ?.，仅在当左边部分不存在也没问题的情况下使用为宜。

以保证在代码中有编程上的 error 出现时，也不会对我们隐藏。
