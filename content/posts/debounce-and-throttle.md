---
title: "Debounce and Throttle"
date: 2020-09-07T14:51:32+08:00
draft: false
---

### 防抖

> 根据一个选择器函数，舍弃掉在两次输出之间小于指定时间的发出值。

- 实现方式：每次触发事件时设置一个延迟调用方法，并且取消之前的延时调用方法
- 缺点：如果事件在规定的时间间隔内被不断的触发，则调用方法会被不断的延迟

```js
// RxJS v6+
import { of, timer } from 'rxjs';
import { debounce } from 'rxjs/operators';

// 发出四个字符串
const example = of('WAIT', 'ONE', 'SECOND', 'Last will display');
/*
  只有在最后一次发送后再经过一秒钟，才会发出值，并抛弃在此之前的所有其他值
*/

const debouncedExample = example.pipe(debounce(() => timer(1000)));
/*
    在这个示例中，所有的值都将被忽略，除了最后一个
    输出: 'Last will display'
*/
const subscribe = debouncedExample.subscribe(val => console.log(val));



// lodash 4.17
_.debounce(func, [wait=0], [options={}])


```

#### 实现方式:

```js

```


#### 参考资料

- [rxjs debounce](https://rxjs-cn.github.io/learn-rxjs-operators/operators/filtering/debounce.html)
- [debounce](https://cn.rx.js.org/class/es6/Observable.js~Observable.html#instance-method-debounce)
- [lodash debounce source code](https://github.com/lodash/lodash/blob/4.17.15/lodash.js#L10304)





### 截流 

以某个时间间隔为阈值，在 durationSelector 完成前将抑制新值的发出

- 实现方式：每次触发事件时，如果当前周期内有等待执行的函数，则取消当前事件

```js
import { interval } from 'rxjs';
import { throttle } from 'rxjs/operators';

// 每1秒发出值
const source = interval(1000);

// 节流2秒后才发出最新值
const example = source.pipe(throttle(val => interval(2000)));
// 输出: 0...3...6...9
const subscribe = example.subscribe(val => console.log(val));
```

#### 实现方式:

```js

```

#### 参考资料

- [Official documents throttle](https://cn.rx.js.org/class/es6/Observable.js~Observable.html#instance-method-throttle)
- [rxjs throttle operators](https://rxjs-cn.github.io/learn-rxjs-operators/operators/filtering/throttle.html)
- [throttle source code](https://github.com/ReactiveX/rxjs/blob/master/src/internal/operators/throttle.ts)