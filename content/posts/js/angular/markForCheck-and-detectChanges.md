---
title: "MarkForCheck and DetectChanges"
date: 2021-12-09T19:48:58+08:00
draft: false
---

### markForCheck()、OnPush、zone.js 与 detectChanges()

Angular 更新 HTML 的原理

首先，总结一下 Angular 的 change detection 的原理。一般来说，Angular 会在以下情况下更新 HTML：

- 初始化 Component 。在最开始 bootstrapping Angular Application 的时候，Angular 会载入 bootstrap component 并调用 ApplicationRef.tick() 方法来触发 change detection 以及 View Rendering。

- 监听事件。如果 DOM 的 event listener 更新了 Angular component 里面的 data，那么 change detection 就会被触发。

- HTTP 数据请求。如果在 Angular component 里更新了请求的数据，那么 change detection 就会被触发。
MacroTasks。如果在 MacroTasks 的 callback 方法里更新了 Angular component 的数据，那么 change detection 就会被触发。比如 setTimeout(), setInterval()。

- MicroTasks。例如 Promise.then() 里更新了 Angular component 的数据，那么 change detection 就会被触发。
其他的 async operations。例如：

```
WebSocket.onmessage();
```

### 关于 zone.js

简单来说，在 Zone 里面的 asynchronous operations 的 execution context 会维持不变。Angular 自带了 ngZone 的 service，这个 service 会创建一个叫 ‘angular’ 的 Zone。在这个 Angular Zone 里，当没有 schedule MicroTasks 或是当执行 sync/async function 的时候会自动触发 change detection。如果 Angular 的检测策略是的 default 的，那么所有的 asynchronous operations 都是在 Angular Zone 里，并且都会按以上条件自动触发 change detection。



当不想触发 change detections 的时候，ngZone 也提供了一种方法：runOutsideAngular()。例如：

```js
this.ngZone.runOutsideAngular(() => {
      setTimeout(() => {
        // do something
        // and codes in here won't trigger change detections
      });
```
如上面代码所示，在 runOutsideAngular() 里面的异步事件并不会触发 change detections。

当然，虽然 zone.js 帮助 Angular 管理了什么时候触发 change detections 的问题，但是它也可能导致 change detections 运行过多，或是意外地跳过一部分 change detections。如果程序员对于 change detections 由很深入的了解，可以选择 disable 掉 zone.js 并自己管理何时触发 change detections，或是选择 extends ReactiveComponent （ ReactiveComponent 里包含用 markDirty 触发 change detections ）。想要 disable 掉 zone.js，首先需要 comment 掉 src/polyfills.ts 文件里面的 zone.js import ：

/***************************************************************************************************
 * Zone JS is required by default for Angular itself.
 */
// import 'zone.js/dist/zone';  // Included with Angular CLI.
接着在 src/main.ts 文件里，将 bootstrap Angular 的 ngZone 设置为 noop：

```js
platformBrowserDynamic().bootstrapModule(AppModule, { ngZone: 'noop' })
  .catch(err => console.error(err));
```

### 关于 detectChanges()

虽然 Zone 可以处理大多数的 asynchronous APIs，但是还是有一些 Zone 无法处理的第三方APIs。当这个第三方 APIs 的方法 update data 的时候，这个 update 是在 Angular Zone 之外的，所以 Angular 并不会知道这个 update 的发生，因此也不会触发 change detection。在这个情况下，就可以调用 ChangeDetectorRef.detectChanges() 来手动触发 change detection。除了这个方法外，也可以把调用第三方 APIs 的 codes 放入 NgZone 的 zone.run() 方法或是放入 MacroTasks 的 setTimeOut() 里也都可以手动触发 change detection。除此之外，还有一种适用 detectChanges() 的情况是当触发了 change detection 之后又发生了 update 。总而言之，detectChanges() 只是用于手动触发 change detection 的方法之一。

### 关于 OnPush 和 markForCheck()

如果 Angular 的检测策略被设置为 OnPush，那么只有以下条件下 change detection 才会被触发：

- 当带有 @Input property 的引用完全变化或是完全被新的值替换的时候。
- 当当前 component 或是 child component 触发 event 的时候，例如 Click，KeyUp，Subscription 等等。
- 当有了以上条件限制，很多时候 update 就不会触发 change detection。这时就必须通过调用 markForCheck() 

来手动告诉 Angular 去检查并更新 view。所以总体来说，markForCheck() 基本上只在 Angular 的检测策略被设置为 OnPush 的时候会被需要使用到。

### markForCheck() 和 detectChanges() 的核心区别

最后，总结一下 markForCheck() 和 detectChanges() 的核心区别。detectChanges() 会真正触发 Angular 的 change detection，而 markForCheck() 则不会。


### 参考文章

- https://stackoverflow.com/questions/41364386/whats-the-difference-between-markforcheck-and-detectchanges