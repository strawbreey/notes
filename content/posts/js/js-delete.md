---
title: "Js Delete"
date: 2020-11-16T19:42:00+08:00
draft: false
---

`delete` 操作符用于删除对象的某个属性；如果没有指向这个属性的引用，那它最终会被释放。

语法

```js
delete expression

delete object.property 
delete object['property']
```
参数

- object 对象的名称，或计算结果为对象的表达式。
- property 要删除的属性。

返回值

对于所有情况都是true，除非属性是一个自身的 不可配置的属性，在这种情况下，非严格模式返回 false。


```js
const Employee = {
  firstname: 'John',
  lastname: 'Doe'
};

console.log(Employee.firstname);
// expected output: "John"

delete Employee.firstname;

console.log(Employee.firstname);
// expected output: undefined
```


但是，以下情况需要重点考虑：

- 如果你试图删除的属性不存在，那么delete将不会起任何作用，但仍会返回true

- 如果对象的原型链上有一个与待删除属性同名的属性，那么删除属性之后，对象会使用原型链上的那个属性（也就是说，delete操作只会在自身的属性上起作用）

- 任何使用 var 声明的属性不能从全局作用域或函数的作用域中删除。
    - 这样的话，delete操作不能删除任何在全局作用域中的函数（无论这个函数是来自于函数声明或函数表达式）
    - 除了在全局作用域中的函数不能被删除，在对象(object)中的函数是能够用delete操作删除的。

- 任何用let或const声明的属性不能够从它被声明的作用域中删除。

- 不可设置的(Non-configurable)属性不能被移除。这意味着像Math, Array, Object内置对象的属性以及使用Object.defineProperty()方法设置为不可设置的属性不能被删除。