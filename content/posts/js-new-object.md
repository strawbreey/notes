---
title: "Js New Object"
date: 2020-11-02T16:44:19+08:00
draft: true
---

常规的 {} 语法允许创建一个对象。但是我们经常需要创建许多类似的对象，例如多个用户或菜单项等。

这可以使用构造函数和 "new" 操作符来实现。

### 构造函数

构造函数在技术上是常规函数。不过有两个约定：

- 它们的命名以大写字母开头。
- 它们只能由 "new" 操作符来执行。

```js
function User(name) {
  this.name = name;
  this.isAdmin = false;
}

let user = new User("Jack");

alert(user.name); // Jack
alert(user.isAdmin); // false
```

当一个函数被使用 new 操作符执行时，它按照以下步骤：

- 一个新的空对象被创建并分配给 this。
- 函数体执行。通常它会修改 this，为其添加新的属性。
- 返回 this 的值。

```js
function User(name) {
  // this = {};（隐式创建）

  // 添加属性到 this
  this.name = name;
  this.isAdmin = false;

  // return this;（隐式返回）
}
```

所以 new User("Jack") 的结果是相同的对象：

```js
let user = {
  name: "Jack",
  isAdmin: false
};
```

现在，如果我们想创建其他用户，我们可以调用 new User("Ann")，new User("Alice") 等。比每次都使用字面量创建要短得多，而且更易于阅读。

*这是构造器的主要目的 —— 实现可重用的对象创建代码。*

 从技术上讲，任何函数都可以用作构造器。即：任何函数都可以通过 new 来运行，它会执行上面的算法。“首字母大写”是一个共同的约定，以明确表示一个函数将被使用 new 来运行。

### 构造器的 return

通常，构造器没有 return 语句。它们的任务是将所有必要的东西写入 this，并自动转换为结果。

但是，如果这有一个 return 语句，那么规则就简单了：

- 如果 return 返回的是一个对象，则返回这个对象，而不是 this。
- 如果 return 返回的是一个原始类型，则忽略。

换句话说，带有对象的 return 返回该对象，在所有其他情况下返回 this。

例如，这里 return 通过返回一个对象覆盖 this：

```js
function BigUser() {

  this.name = "John";

  return { name: "Godzilla" };  // <-- 返回这个对象
}

console.log( new BigUser().name );  // Godzilla，得到了那个对象

function SmallUser() {

  this.name = "John";

  return; // <-- 返回 this
}

console.log( new SmallUser().name );  // John
```

### 总结
- 构造函数，或简称构造器，就是常规函数，但大家对于构造器有个共同的约定，就是其命名首字母要大写。
- 构造函数只能使用 new 来调用。这样的调用意味着在开始时创建了空的 this，并在最后返回填充了值的 this。