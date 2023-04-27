---
title: "Angular Input Get With Set"
date: 2021-11-22T10:12:17+08:00
draft: false
---

如果您主要对仅对 setter实现逻辑感兴趣：

```js
import { Component, Input, OnChanges, SimpleChanges } from '@angular/core';

// [...]

export class MyClass implements OnChanges {
  @Input() allowDay: boolean;

  ngOnChanges(changes: SimpleChanges): void {
    if(changes['allowDay']) {
      this.updatePeriodTypes();
    }
  }
}
```
SimpleChanges如果更改哪个输入属性无关紧要，或者如果您只有一个输入属性，则不需要导入。

Angular 文档：OnChanges

除此以外：

```js
private _allowDay: boolean;

@Input() set allowDay(value: boolean) {
  this._allowDay = value;
  this.updatePeriodTypes();
}
get allowDay(): boolean {
  // other logic
  return this._allowDay;
}
```