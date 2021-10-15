---
title: "Switchmap"
date: 2021-10-08T11:13:36+08:00
draft: false
---

### switchMap

switchMap 和其他打平操作符的主要区别是它具有取消效果。在每次发出时，会取消前一个内部 observable (你所提供函数的结果) 的订阅，然后订阅一个新的 observable 。你可以通过短语切换成一个新的 observable来记忆它。


1. 示例 1: 每5秒重新启动 interval
```js
import { timer, interval } from 'rxjs';
import { switchMap } from 'rxjs/operators';

// 立即发出值， 然后每5秒发出值
const source = timer(0, 5000);
// 当 source 发出值时切换到新的内部 observable，发出新的内部 observable 所发出的值
const example = source.pipe(switchMap(() => interval(500)));
// 输出: 0,1,2,3,4,5,6,7,8,9...0,1,2,3,4,5,6,7,8
const subscribe = example.subscribe(val => console.log(val));
```

2. 示例 2: 每次点击时重置

[StackBlitz](https://stackblitz.com/edit/typescript-s4pvix?file=index.ts&devtoolsheight=100)

```js
// RxJS v6+
import { interval, fromEvent } from 'rxjs';
import { switchMap, mapTo } from 'rxjs/operators';

// 发出每次点击
const source = fromEvent(document, 'click');
// 如果3秒内发生了另一次点击，则消息不会被发出
const example = source.pipe(
  switchMap(val => interval(3000).pipe(mapTo('Hello, I made it!')))
);
// (点击)...3s...'Hello I made it!'...(点击)...2s(点击)...
const subscribe = example.subscribe(val => console.log(val));
```