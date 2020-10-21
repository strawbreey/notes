---
title: "Js Promise"
date: 2020-10-21T23:01:58+08:00
draft: true
---

Promise 对象的构造器（constructor）语法如下：

```js
let promise = new Promise(function(resolve, reject) {
  // executor（生产者代码，“歌手”）
});
```

当 executor 获得了结果，无论是早还是晚都没关系，它应该调用以下回调之一：

resolve(value) — 如果任务成功完成并带有结果 value。
reject(error) — 如果出现了 error，error 即为 error 对象。

Promise 是一个对象，它代表了一个异步操作的最终完成或者失败。

```js
// 成功的回调函数
function successCallback(result) {
  console.log("音频文件创建成功: " + result);
}

// 失败的回调函数
function failureCallback(error) {
  console.log("音频文件创建失败: " + error);
}

createAudioFileAsync(audioSettings, successCallback, failureCallback)
```

```js
const promise = createAudioFileAsync(audioSettings); 
promise.then(successCallback, failureCallback);

// 简写
createAudioFileAsync(audioSettings).then(successCallback, failureCallback);
```

### 约定

不同于“老式”的传入回调，在使用 Promise 时，会有以下约定：

- 在本轮 事件循环 运行完成之前，回调函数是不会被调用的。
- 即使异步操作已经完成（成功或失败），在这之后通过 then() 添加的回调函数也会被调用。
- 通过多次调用 then() 可以添加多个回调函数，它们会按照插入顺序进行执行。

Promise 很棒的一点就是链式调用（chaining）。

### 链式调用

连续执行两个或者多个异步操作是一个常见的需求，在上一个操作执行成功之后，开始下一个的操作，并带着上一步操作所返回的结果。

```js

const promise = doSomething();
const promise2 = promise.then(successCallback, failureCallback);
// 或者
const promise2 = doSomething().then(successCallback, failureCallback);
```

避免地狱回调

```js
doSomething().then(function(result) {
  return doSomethingElse(result);
})
.then(function(newResult) {
  return doThirdThing(newResult);
})
.then(function(finalResult) {
  console.log('Got the final result: ' + finalResult);
})
.catch(failureCallback);

// 箭头函数简化
doSomething()
.then(result => doSomethingElse(result))
.then(newResult => doThirdThing(newResult))
.then(finalResult => {
  console.log(`Got the final result: ${finalResult}`);
})
.catch(failureCallback);
```

    注意: 一定要有返回值，否则，callback 将无法获取上一个 Promise 的结果。(如果使用箭头函数，() => x 比 () => { return x; } 更简洁一些，但后一种保留 return 的写法才支持使用多个语句。）。

### Catch 的后续链式操作

有可能会在一个回调失败之后继续使用链式操作，即，使用一个 catch，这对于在链式操作中抛出一个失败之后，再次进行新的操作会很有用。请阅读下面的例子：

```js
new Promise((resolve, reject) => {
    console.log('初始化');
    resolve();
})
.then(() => {
    throw new Error('有哪里不对了');
    console.log('执行「这个」”');
})
.catch(() => {
    console.log('执行「那个」');
})
.then(() => {
    console.log('执行「这个」，无论前面发生了什么');
});
```
输出结果如下：

    初始化
    执行“那个”
    执行“这个”，无论前面发生了什么


### 错误传递
在之前的回调地狱示例中，你可能记得有 3 次 failureCallback 的调用，而在 Promise 链中只有尾部的一次调用。

```js

doSomething()
.then(result => doSomethingElse(result))
.then(newResult => doThirdThing(newResult))
.then(finalResult => console.log(`Got the final result: ${finalResult}`))
.catch(failureCallback);

```

通常，一遇到异常抛出，浏览器就会顺着 Promise 链寻找下一个 onRejected 失败回调函数或者由 .catch() 指定的回调函数。这和以下同步代码的工作原理（执行过程）非常相似。

```js
try {
  let result = syncDoSomething();
  let newResult = syncDoSomethingElse(result);
  let finalResult = syncDoThirdThing(newResult);
  console.log(`Got the final result: ${finalResult}`);
} catch(error) {
  failureCallback(error);
}
```

### Promise 拒绝事件

### 参考资料

- [使用 Promise](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Guide/Using_promises)
- [Promise then](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/then)
- [async function](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/async_function)
- [Promise](https://zh.javascript.info/promise-basics)
