---
title: "Rxjs Unsubscribe"
date: 2021-10-18T15:13:24+08:00
draft: false
---


Angular 中引入Observable处理各种常用异步操作，如 `EventEmitter`、`HttpClient`、`Router`以及`响应式表单Form`，实际开发中也会自定义Observable对象，通过订阅Observable对象可以实现数据通信，并在组件注销时取消订阅。

### 为什么要取消订阅？

和Javascript中的addEventLister一样，Angular中一旦订阅了Observable对象，在组件destory时就应该及时取消订阅，否则会造成内存泄露或其他不必要的操作。

下面这个例子通过PageService订阅refresh事件，返回一个Observable对象，当发布refresh时，请求用户信息。这里为了演示，延迟了10秒钟，若此时通过路由跳转到其他页面，IndexComponent将销毁，但是10秒后，请求依然会发出。也就是说组件销毁了，组件内的订阅并没有自动销毁，需要手动进行取消订阅。

```js
export class IndexComponent implements OnInit {
    constructor(private page: PageService, user:UserService) {}
    ngOnInit() {
        this.page.refresh.delay(10000).subscribe(() => {
            this.user.getUser();
        });
    }
}
```


### 如何取消订阅？

组件销毁前，需要通过OnDestroy生命钩子进行取消订阅可观察对象，有3种方法可以实现：

1. unsubscribe()

```js
private subscription: Subscription;

constructor(private page: PageService, user:UserService) {}
ngOnInit() {
    this.subscription = this.page.refresh.subscribe(() => {
        this.user.getUser();
    });
}
ngOnDestroy(){
    this.subscription.unsubscribe();
}
```
当组件中存在多个subscription时，要一个一个进行unsubscribe未免有些麻烦，有2种方式可以简化代码。


2. 通过subscription.add()添加子subscription

subscription.add()通过将多个subscription添加为子subscription，当取消父subscription时，子subscription也将自动取消，该方法要求必须先执行第一个subscription。

```js
ngOnInit() {
    // 订阅第一个可观察对象-parent
    this.subscription = this.page.refresh.subscribe(() => {
      this.user.getUser();
    });
    // 将第二个可观察对象添加为子对象-child
    this.subscription.add(this.page.load.subscribe(() => {
      this.user.getData();
    }));
}
ngOnDestroy(){
      // 同时取消2次订阅
      this.subscription.unsubscribe();
}
```

3. 定义subscription数组，组件销毁时依次取消订阅

```js
private subscriptions: Subscription[]=[];
ngOnInit() {
    this.subscriptions.push(this.page.refresh.subscribe(() => {
      this.user.getUser();
    }));
    this.subscriptions.push(this.page.load.subscribe(() => {
      this.user.getData();
    }));
}
ngOnDestroy() {
    this.subscriptions.forEach(sub => sub.unsubscribe());
}
```

4. takeWhile()
Observable对象的[takeWhile]方法将对数据流进行操作，表示：持续发出值，当传入表达式结果为false时，停止发出值。例子中定义了一个组件销毁标识destroy，当destroy==true时，可观察对象发出的值将不再流向该组件。该方法较简洁，推荐使用。

```js
takeWhile(predicate: boolean): Observable;
destroy = false;  

ngOnInit() {
    this.page.refresh.takeWhile(() => !this.destroy).subscribe(() => {
        this.user.getUser();
    });
}
ngOnDestroy() {
    this.destroy = true;
}
```

5. takeUntil()
Observable对象的`takeUntil`方法将对数据流进行操作，表示：持续发出值，当传入 observable 发出值时，停止发出值。该方法需要传入另一个Observable，在组件销毁时发出值，进而取消订阅。

```js
takeUntil(notifier: Observable): Observable
private unsubscribe = new Subject<void>();
ngOnInit() {
    this.page.refresh.takeUntil(this.unsubscribe).subscribe(num => {
        this.user.getUser();
    });
}
ngOnDestroy() {
    this.unsubscribe.next();  // 所有subscription结束
    this.unsubscribe.complete();  // 传入observable结束
}
```

### 不需要手动取消的场景

Angular中有些场景已进行unsubscribe或通过Observable.complete()的方式结束订阅数据流，在开发中并不需要手动再进行处理订阅。

1) Async pipe
AsyncPipe 会订阅一个可观察对象或承诺，并返回其发出的最后一个值。当发出新值时，该管道就会把这个组件标记为需要进行变更检查的。当组件销毁时，自动取消订阅。如果只是模板中需要的变量，采用Async是最佳实践。

```js
@Component({
    selector: 'index',
    template: `<div>refreshTime: {{ refreshTime | async }}</div>`
})
export class IndexComponent implements OnInit {
    refreshTime: Observable<string>;
    constructor(private page: PageService) {}
    ngOnInit() {
        this.refreshTime = this.page.refresh();
    }
}
```
1) HostListener
通过@HostListener进行订阅的事件，和直接在模板里订阅事件一样，也是自动取消的。

3) EventEmitter
EventEmitter 类派生自 Observable，Angular 提供了一个 EventEmitter 类，它用来从组件的 @Output() 属性中发布一些值。

4) 有限的Observable
有限的Observable指的是发出的值是有限的，如timer。

5) Http
http在load事件完成之后进行responseObserver.complete()，自动结束数据流，下面是Http源码中的部分代码。

```js
const onLoad = (event) => {
    ...
    responseObserver.next(new Response(responseOptions));
    responseObserver.complete();
};
script.addEventListener('load', onLoad);
```

### 需要手动取消的场景
1) Router
Angular在组件销毁时并没有取消router的所有订阅事件，同样是延迟10秒，可以看到请求依然是会发出的。
```js
ngOnInit() {
    this.subscription = this.router.queryParamMap.delay(10000).subscribe((param) => {
        this.user.getUser(+param.get(id));
    });
}
ngOnDestroy() {
    this.subscription.unsubscribe();
}
```
2) Forms
表单中的valueChanges和statusChanges等Observable都需要手动取消。
```js
ngOnInit() {
    this.form = new FormGroup({...});
    this.subscription  = this.form.valueChanges.subscribe(() => {});
}
ngOnDestroy() {
    this.subscription.unsubscribe();
}
```
3) Renderer Service

```js
constructor(private renderer: Renderer2, private element : ElementRef) { }
    ngOnInit() {
    this.subscription = this.renderer.listen(this.element.nativeElement, "click", () => {});
}
ngOnDestroy() {
    this.subscription.unsubscribe();
}
```

4) fromEvent()、interval()等输出值可能为无限个的可观察对象

5) 
```js
destroy: boolean = false;  // 组件销毁标识
ngOnInit() {
    Observable.interval(1000).takeWhile(() => !this.destroy).subscribe(() => {}));
    Observable.fromEvent(this.element.nativeElement, 'click').takeWhile(() => !this.destroy).subscribe(() => {}));
}
ngOnDestroy() {
       this.destroy = true;
}
```
5) 自定义Observable

所有自定义Observable必须在组件销毁前手动取消。

在实践中，要时刻记得取消可观察对象的订阅，如果只是模板中需要的变量，采用Async是最佳实践，其次是通过takeWhile()或takeUntil()手动取消订阅。

### 参考资料

https://angular.cn/guide/observables-in-angular
https://netbasal.com/when-to-unsubscribe-in-angular-d61c6b21bad3
https://blog.angularindepth.com/rxjs-composing-subscriptions-b53ab22f1fd5
http://brianflove.com/2016/12/11/anguar-2-unsubscribe-observables/

