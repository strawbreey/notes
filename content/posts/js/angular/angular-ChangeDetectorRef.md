---
title: "Angular ChangeDetectorRef"
date: 2021-08-31T11:02:20+08:00
draft: false
---

### 手动控制脏检查 ChangeDetectorRef

Angular的ChangeDetectorRef实例上提供了可以绑定或解绑某个组件脏检查的方法。

```ts
class ChangeDetectorRef {
  markForCheck() : void     // 通知框架进行变化检查/Change Detection
  detach() : void           // 禁止脏检查
  detectChanges() : void    // 手工触发脏检查， 从该组件到各个子组件执行一次变化检测
  checkNoChanges() : void
  reattach() : void         // detach逆操作，启用脏检查
}
```