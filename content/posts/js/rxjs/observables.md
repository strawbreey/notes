---
title: "Observables"
date: 2021-10-15T10:57:08+08:00
draft: false
---

### 创建 Observables

Rx.Observable.create 是 Observable 构造函数的别名，它接收一个参数：subscribe 函数。

下面的示例创建了一个 Observable，它每隔一秒会向观察者发送字符串 'hi' 。

```js
var observable = Rx.Observable.create(function subscribe(observer) {
  var id = setInterval(() => {
    observer.next('hi')
  }, 1000);
});
```

Observables 可以使用 create 来创建, 但通常我们使用所谓的创建操作符, 像 of、from、interval、等等。

在上面的示例中，subscribe 函数是用来描述 Observable 最重要的一块。我们来看下订阅是什么意思。

### 订阅 Observables
示例中的 Observable 对象 observable 可以订阅，像这样：

```js
observable.subscribe(x => console.log(x));
```
`observable.subscribe` 和 `Observable.create(function subscribe(observer) {...})` 中的 subscribe 有着同样的名字，这并
不是一个巧合。在库中，它们是不同的，但从实际出发，你可以认为在概念上它们是等同的。

这表明 subscribe 调用在同一 Observable 的多个观察者之间是不共享的。当使用一个观察者调用 observable.subscribe 时，Observable.create(function subscribe(observer) {...}) 中的 subscribe 函数只服务于给定的观察者。对 observable.subscribe 的每次调用都会触发针对给定观察者的独立设置。

订阅 Observable 像是调用函数, 并提供接收数据的回调函数。

这与像 addEventListener / removeEventListener 这样的事件处理方法 API 是完全不同的。使用 observable.subscribe，在 Observable 中不会将给定的观察者注册为监听器。Observable 甚至不会去维护一个附加的观察者列表。

subscribe 调用是启动 “Observable 执行”的一种简单方式， 并将值或事件传递给本次执行的观察者。

### 执行 Observables
Observable.create(function subscribe(observer) {...}) 中...的代码表示 “Observable 执行”，它是惰性运算，只有在每个观察者订阅后才会执行。随着时间的推移，执行会以同步或异步的方式产生多个值。

Observable 执行可以传递三种类型的值：

"Next" 通知： 发送一个值，比如数字、字符串、对象，等等。
"Error" 通知： 发送一个 JavaScript 错误 或 异常。
"Complete" 通知： 不再发送任何值。
"Next" 通知是最重要，也是最常见的类型：它们表示传递给观察者的实际数据。"Error" 和 "Complete" 通知可能只会在 Observable 执行期间发生一次，并且只会执行其中的一个。

这些约束用所谓的 Observable 语法或合约表达最好，写为正则表达式是这样的：

next*(error|complete)?
在 Observable 执行中, 可能会发送零个到无穷多个 "Next" 通知。如果发送的是 "Error" 或 "Complete" 通知的话，那么之后不会再发送任何通知了。

下面是 Observable 执行的示例，它发送了三个 "Next" 通知，然后是 "Complete" 通知：

var observable = Rx.Observable.create(function subscribe(observer) {
  observer.next(1);
  observer.next(2);
  observer.next(3);
  observer.complete();
});
Observable 严格遵守自身的规约，所以下面的代码不会发送 "Next" 通知 4：

var observable = Rx.Observable.create(function subscribe(observer) {
  observer.next(1);
  observer.next(2);
  observer.next(3);
  observer.complete();
  observer.next(4); // 因为违反规约，所以不会发送
});
在 subscribe 中用 try/catch 代码块来包裹任意代码是个不错的主意，如果捕获到异常的话，会发送 "Error" 通知：

var observable = Rx.Observable.create(function subscribe(observer) {
  try {
    observer.next(1);
    observer.next(2);
    observer.next(3);
    observer.complete();
  } catch (err) {
    observer.error(err); // 如果捕获到异常会发送一个错误
  }
});
### 清理 Observable 执行

因为 Observable 执行可能会是无限的，并且观察者通常希望能在有限的时间内中止执行，所以我们需要一个 API 来取消执行。因为每个执行都是其对应观察者专属的，一旦观察者完成接收值，它必须要一种方法来停止执行，以避免浪费计算能力或内存资源。

当调用了 observable.subscribe ，观察者会被附加到新创建的 Observable 执行中。这个调用还返回一个对象，即 Subscription (订阅)：
```js
var subscription = observable.subscribe(x => console.log(x));
```
Subscription 表示进行中的执行，它有最小化的 API 以允许你取消执行。想了解更多订阅相关的内容，请参见 Subscription 类型。使用 subscription.unsubscribe() 你可以取消进行中的执行：

```js
var observable = Rx.Observable.from([10, 20, 30]);
var subscription = observable.subscribe(x => console.log(x));
// 稍后：
subscription.unsubscribe();
```

当你订阅了 Observable，你会得到一个 Subscription ，它表示进行中的执行。只要调用 unsubscribe() 方法就可以取消执行。

当我们使用 create() 方法创建 Observable 时，Observable 必须定义如何清理执行的资源。你可以通过在 function subscribe() 中返回一个自定义的 unsubscribe 函数。

举例来说，这是我们如何清理使用了 setInterval 的 interval 执行集合：

```js
var observable = Rx.Observable.create(function subscribe(observer) {
  // 追踪 interval 资源
  var intervalID = setInterval(() => {
    observer.next('hi');
  }, 1000);

  // 提供取消和清理 interval 资源的方法
  return function unsubscribe() {
    clearInterval(intervalID);
  };
});
```
正如 observable.subscribe 类似于 Observable.create(function subscribe() {...})，从 subscribe 返回的 unsubscribe 在概念上也等同于 subscription.unsubscribe。事实上，如果我们抛开围绕这些概念的 ReactiveX 类型，保留下来的只是相当简单的 JavaScript 。

```js
function subscribe(observer) {
  var intervalID = setInterval(() => {
    observer.next('hi');
  }, 1000);

  return function unsubscribe() {
    clearInterval(intervalID);
  };
}

var unsubscribe = subscribe({next: (x) => console.log(x)});
```
// 稍后：
unsubscribe(); // 清理资源
为什么我们要使用像 Observable、Observer 和 Subscription 这样的 Rx 类型？原因是保证代码的安全性(比如 Observable 规约)和操作符的可组合性。