---
title: "Angular RunOutSideAngular"
date: 2021-08-31T11:00:49+08:00
draft: false
---

### 利用ngzone-runOutsideAngular优化

Angular依赖`NgZone`来监听异步操作，并从根部执行变化检测。换句话说，我们代码中的每一个 addEventListener都会触发脏检查。但是如果我们非常明确，有些addEventListener要执行的东西，不会（或者说可以忽略）影响数据结果，不想然他触发脏检查。比如监测scroll，监测鼠标事件等。

针对这种情况， 我们可以使用zone提供的runOutsideAngular，让这些事件不触发脏检查。

```ts
this.zone.runOutsideAngular(() => {
    window.document.addEventListener('mousemove', this.bindMouse);
});
```