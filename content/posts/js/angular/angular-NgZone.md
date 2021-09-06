---
title: "Angular NgZone"
date: 2021-08-31T11:05:06+08:00
draft: false
---

## NgZone

一种用于在 Angular Zone 内部或外部执行任务的可注入服务。

此服务最常见的用途是在启动包含一个或多个不需要 Angular 处理的 UI 更新或错误处理的异步任务的工作时优化性能。可以通过 runOutsideAngular 启动此类任务，如果需要，这些任务可以通过 run 重新进入 Angular zone。

```ts

class NgZone {
  static isInAngularZone(): boolean
  static assertInAngularZone(): void
  static assertNotInAngularZone(): void
  constructor(__0)
  hasPendingMacrotasks: boolean
  hasPendingMicrotasks: boolean
  isStable: boolean
  onUnstable: EventEmitter<any>
  onMicrotaskEmpty: EventEmitter<any>
  onStable: EventEmitter<any>
  onError: EventEmitter<any>
  run<T>(fn: (...args: any[]) => T, applyThis?: any, applyArgs?: any[]): T
  runTask<T>(fn: (...args: any[]) => T, applyThis?: any, applyArgs?: any[], name?: string): T
  runGuarded<T>(fn: (...args: any[]) => T, applyThis?: any, applyArgs?: any[]): T
  runOutsideAngular<T>(fn: (...args: any[]) => T): T
}

```