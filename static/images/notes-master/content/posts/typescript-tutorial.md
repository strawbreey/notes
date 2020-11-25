---
title: "Typescript Tutorial"
date: 2020-09-27T10:24:02+08:00
draft: false
---

TypeScript 是 JavaScript 的类型的超集，它可以编译成纯 JavaScript。编译出来的 JavaScript 可以运行在任何浏览器上。TypeScript 编译工具可以运行在任何服务器和任何系统上。TypeScript 是开源的。

### Install Typescript

```ts
function sayHello(person: string) {
    return 'Hello, ' + person;
}

let user = 'Tom';
console.log(sayHello(user));
```

```shell
npm install -g typescript

tsc hello.ts
```

### data type

JavaScript 的数据类型分为两种: 基本数据类型和复杂数据类型

基本数据类型:

- Undefined
- Null
- Boolean
- Number
- String
- Symbol (ES6) symbol 是ES6中新增的一种特殊的、不可变的基本数据类型（primitive data type），可以作为对象属性的标识符使用

- Tuple (TS) 元组 (元组类型允许表示一个已知元素数量和类型的数组，各元素的类型不必相同)
- Enum (TS) 枚举
- Any (TS) 我们会想要为那些在编程阶段还不清楚类型的变量指定一个类型
- Void (TS) 与 Any 相反 (当一个函数没有返回值时，你通常会见到其返回值类型是 void)
- Never (TS) never类型表示的是那些永不存在的值的类型


复杂数据类型
- Object
