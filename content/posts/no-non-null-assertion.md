---
title: "No Non Null Assertion"
date: 2021-04-01T11:10:57+08:00
draft: false
---

Disallows non-null assertions using the ! postfix operator.

Rationale
Using non-null assertion cancels the benefits of the strict null checking mode.

Instead of assuming objects exist:
```js
function foo(instance: MyClass | undefined) {
    instance!.doWork();
}
```
Either inform the strict type system that the object must exist:
```js
function foo(instance: MyClass) {
    instance.doWork();
}
```
Or verify that the instance exists, which will inform the type checker:
```js
function foo(instance: MyClass | undefined) {
    if (instance !== undefined) {
        instance.doWork();
    }
}
```
