---
title: "Type challenge"
date: 2020-11-18T17:29:24+08:00
draft: true
---

# 获取函数返回类型 

不使用 [`ReturnType`](https://www.typescriptlang.org/docs/handbook/utility-types.html#returntypetype) 实现 TypeScript 的 `ReturnType<T>` 泛型。

例如：

```ts
const fn = (v: boolean) => {
  if (v)
    return 1
  else
    return 2
}

type a = MyReturnType<typeof fn> // 应推导出 "1 | 2"
```

解决方案:

```ts
type MyReturnType<T extends Function> =
  T extends (...args: any) => infer R
    ? R
    : never
```

# 实现 Omit 

不使用 `Omit` 实现 TypeScript 的 `Omit<T, K>` 泛型。

`Omit` 会创建一个省略 `K` 中字段的 `T` 对象。

例如：

```ts
interface Todo {
  title: string
  description: string
  completed: boolean
}

type TodoPreview = MyOmit<Todo, 'description' | 'title'>

const todo: TodoPreview = {
  completed: false,
}
```

解决方案:

```ts
type MyOmit<T, K extends keyof T> = {[P in keyof T as P extends K ? never: P] :T[P]}
```