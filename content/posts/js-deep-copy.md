---
title: "Js Deep Copy"
date: 2020-09-27T19:38:49+08:00
draft: false
---

### JSON.parse/stringify

if you do not use Date, function, undefined, Infinity, RegFexps, Maps, Set Blobs, FileList，ImageData， sparseArray,Typed Array or other complex type within you object。

如果不使用复制的数据类型, JSON序列化是最快的复制方法
```js
JSON.parse(JSON.stringify(object))

const a = {
  string: 'string',
  number: 123,
  bool: false,
  nul: null,
  date: new Date(),  // stringified
  undef: undefined,  // lost
  inf: Infinity,  // forced to 'null'
  re: /.*/,  // lost
}
console.log(a);
console.log(typeof a.date);  // Date object
const clone = JSON.parse(JSON.stringify(a));
console.log(clone);
console.log(typeof clone.date);  // result of .toISOString()
```

### 使用常见库中clone函数

Since cloning objects is not trivial (complex types, circular references, function etc.), most major libraries provide function to clone objects. Don't reinvent the wheel - if you're already using a library, check if it has an object cloning function. For example,

- [lodash deepclone](https://github.com/lodash/lodash/blob/4.17.15/lodash.js#L11087) - cloneDeep; can be imported separately via the lodash.clonedeep module and is probably your best choice if you're not already using a library that provides a deep cloning function

- AngularJS - angular.copy

- vuejs [deepClone](https://github.com/vuejs/vue/blob/52719ccab8fccffbdf497b96d3731dc86f04c1ce/src/server/bundle-renderer/create-bundle-runner.js)

- Angular - 

- jQuery - jQuery.extend(true, { }, oldObject); .clone() only clones DOM elements


### ES6

For completeness, note that ES6 offers two shallow copy mechanisms: Object.assign() and the spread syntax. which copies values of all enumerable own properties from one object to another. For example:
```js
// 浅复制
var A1 = {a: "2"};
var A2 = Object.assign({}, A1);
var A3 = {...A1};
```