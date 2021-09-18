---
title: "Js ConcatAll"
date: 2021-08-24T14:26:42+08:00
draft: false
---

示例 1: 使用 observable 来进行 concatAll 

```js
// RxJS v6+
import { map, concatAll } from 'rxjs/operators';
import { of, interval } from 'rxjs';

// 每2秒发出值
const source = interval(2000);
const example = source.pipe(
  // 为了演示，增加10并作为 observable 返回
  map(val => of(val + 10)),
  // 合并内部 observables 的值
  concatAll()
);
// 输出: 'Example with Basic Observable 10', 'Example with Basic Observable 11'...
const subscribe = example.subscribe(val =>
  console.log('Example with Basic Observable:', val)
);
```