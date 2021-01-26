---
title: "Js Interview This"
date: 2021-01-25T19:55:22+08:00
draft: false
---

请简述JavaScript中的this。

JS 中的this是一个相对复杂的概念，不是简单几句能解释清楚的。粗略地讲，函数的调用方式决定了this的值。我阅读了网上很多关于this的文章，Arnav Aggrawal 写的比较清楚。this取值符合以下规则：

在调用函数时使用new关键字，函数内的this是一个全新的对象。
如果apply、call或bind方法用于调用、创建一个函数，函数内的 this 就是作为参数传入这些方法的对象。
当函数作为对象里的方法被调用时，函数内的this是调用该函数的对象。比如当obj.method()被调用时，函数内的 this 将绑定到obj对象。
如果调用函数不符合上述规则，那么this的值指向全局对象（global object）。浏览器环境下this的值指向window对象，但是在严格模式下('use strict')，this的值为undefined。
如果符合上述多个规则，则较高的规则（1 号最高，4 号最低）将决定this的值。
如果该函数是 ES2015 中的箭头函数，将忽略上面的所有规则，this被设置为它被创建时的上下文。