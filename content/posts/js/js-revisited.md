---
title: "Js Revisited"
date: 2020-12-14T11:15:19+08:00
draft: false
---

让我们深入研究一下箭头函数。

箭头函数不仅仅是编写简洁代码的“捷径”。它还具有非常特殊且有用的特性。

JavaScript 充满了我们需要编写在其他地方执行的小函数的情况。

例如：

- arr.forEach(func) —— forEach 对每个数组元素都执行 func。
- setTimeout(func) —— func 由内建调度器执行。


### 箭头函数没有 “this”

### 箭头函数没有 “arguments”