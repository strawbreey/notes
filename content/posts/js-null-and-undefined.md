---
title: "Js Null and Undefined"
date: 2021-01-27T14:19:43+08:00
draft: false
---

## null、undefined和未声明变量之间有什么区别？如何检查判断这些状态值？

当你没有提前使用var、let或const声明变量，就为一个变量赋值时，该变量是未声明变量（undeclared variables）。未声明变量会脱离当前作用域，成为全局作用域下定义的变量。在严格模式下，给未声明的变量赋值，会抛出ReferenceError错误。和使用全局变量一样，使用未声明变量也是非常不好的做法，应当尽可能避免。要检查判断它们，需要将用到它们的代码放在try/catch语句中。

```js
function foo() {
  x = 1; // 在严格模式下，抛出 ReferenceError 错误
}

foo();
console.log(x); // 1
```

当一个变量已经声明，但没有赋值时，该变量的值是undefined。如果一个函数的执行结果被赋值给一个变量，但是这个函数却没有返回任何值，那么该变量的值是undefined。要检查它，需要使用严格相等（===）；或者使用typeof，它会返回'undefined'字符串。请注意，不能使用非严格相等（==）来检查，因为如果变量值为null，使用非严格相等也会返回true。
```js
var foo;
console.log(foo); // undefined
console.log(foo === undefined); // true
console.log(typeof foo === 'undefined'); // true

console.log(foo == null); // true. 错误，不要使用非严格相等！

function bar() {}
var baz = bar();
console.log(baz); // undefined
```

null只能被显式赋值给变量。它表示空值，与被显式赋值 undefined 的意义不同。要检查判断null值，需要使用严格相等运算符。请注意，和前面一样，不能使用非严格相等（==）来检查，因为如果变量值为undefined，使用非严格相等也会返回true。

```js
var foo = null;
console.log(foo === null); // true
console.log(foo == undefined); // true. 错误，不要使用非严格相等！
```

作为一种个人习惯，我从不使用未声明变量。如果定义了暂时没有用到的变量，我会在声明后明确地给它们赋值为null

### 参考资料

- https://stackoverflow.com/questions/15985875/effect-of-declared-and-undeclared-variables