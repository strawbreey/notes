---
title: "Js Microtask Queue"
date: 2020-10-18T23:38:54+08:00
draft: false
---

浏览器中 JavaScript 的执行流程和 Node.js 中的流程都是基于 事件循环 的。

理解事件循环的工作方式对于代码优化很重要，有时对于正确的架构也很重要。


## 事件循环

事件循环 的概念非常简单。它是一个在 JavaScript 引擎等待任务，执行任务和进入休眠状态等待更多任务这几个状态之间转换的无限循环。

引擎的一般算法：

1. 当有任务时：
   - 从最先进入的任务开始执行。
2. 休眠直到出现任务，然后转到第 1 步。

两个细节：

1. 引擎执行任务时永远不会进行渲染（render）。如果任务执行需要很长一段时间也没关系。仅在任务完成后才会绘制对 DOM 的更改。

2. 如果一项任务执行花费的时间过长，浏览器将无法执行其他任务，无法处理用户事件，因此，在一定时间后浏览器会在整个页面抛出一个如“页面未响应”之类的警报，建议你终止这个任务。这种情况常发生在有大量复杂的计算或导致死循环的程序错误时。


### 拆分CPU过载任务

例如，语法高亮（用来给本页面中的示例代码着色）是相当耗费 CPU 资源的任务。为了高亮显示代码，它执行分析，创建很多着了色的元素，然后将它们添加到文档中 —— 对于文本量大的文档来说，需要耗费很长时间。

当引擎忙于语法高亮时，它就无法处理其他 DOM 相关的工作，例如处理用户事件等。它甚至可能会导致浏览器“中断（hiccup）”甚至“挂起（hang）”一段时间，这是不可接受的。

我们可以通过将大任务拆分成多个小任务来避免这个问题。高亮显示前 100 行，然后使用 setTimeout（延时参数为 0）来安排（schedule）后 100 行的高亮显示，依此类推

### 进度指示


## 宏任务和微任务

除了本章中所讲的 宏任务（macrotask） 外，还有在 微任务（Microtask） 一章中提到的 微任务（microtask）。

微任务仅来自于我们的代码。它们通常是由 promise 创建的：对 .then/catch/finally 处理程序的执行会成为微任务。微任务也被用于 await 的“幕后”，因为它是 promise 处理的另一种形式。

每个宏任务之后，引擎会立即执行微任务队列中的所有任务，然后再执行其他的宏任务，或渲染，或进行其他任何操作。

```js
setTimeout(() => alert("timeout"));

// 微任务
Promise.resolve()
  .then(() => alert("promise"));

alert("code");
```
这里的执行顺序是怎样的？

- code 首先显示，因为它是常规的同步调用。
- promise 第二个出现，因为 then 会通过微任务队列，并在当前代码之后执行。
- timeout 最后显示，因为它是一个宏任务。

微任务会在执行任何其他事件处理，或渲染，或执行任何其他宏任务之前完成。

## 总结

事件循环的更详细的算法（尽管与 规范 相比仍然是简化过的）：

1. 从 宏任务 队列（例如 “script”）中出队（dequeue）并执行最早的任务。
2. 执行所有 微任务：
    - 当微任务队列非空时：
    - 出队（dequeue）并执行最早的微任务。
3. 执行渲染，如果有。
4. 如果宏任务队列为空，则休眠直到出现宏任务。
5. 转到步骤 1。