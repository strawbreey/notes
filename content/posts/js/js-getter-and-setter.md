---
title: "Js Getter and Setter"
date: 2020-10-23T11:58:24+08:00
draft: false
---

有两种类型的对象属性。

第一种是 数据属性。我们已经知道如何使用它们了。到目前为止，我们使用过的所有属性都是数据属性。

第二种类型的属性是新东西。它是 访问器属性（accessor properties）。它们本质上是用于获取和设置值的函数，但从外部代码来看就像常规属性。


### Getter 和 setter

访问器属性由 “getter” 和 “setter” 方法表示。在对象字面量中，它们用 get 和 set 表示：

```js
let obj = {
  get propName() {
    // 当读取 obj.propName 时，getter 起作用
  },

  set propName(value) {
    // 当执行 obj.propName = value 操作时，setter 起作用
  }
};
```

从外表看，访问器属性看起来就像一个普通属性。这就是访问器属性的设计思想。我们不以函数的方式 调用 user.fullName，我们正常 读取 它：getter 在幕后运行。

### 访问器描述符

访问器属性的描述符与数据属性的不同。

对于访问器属性，没有 value 和 writable，但是有 get 和 set 函数。

所以访问器描述符可能有：

- get —— 一个没有参数的函数，在读取属性时工作，
- set —— 带有一个参数的函数，当属性被设置时调用，
- enumerable —— 与数据属性的相同，
- configurable —— 与数据属性的相同。

### 更聪明的 getter/setter

Getter/setter 可以用作“真实”属性值的包装器，以便对它们进行更多的控制。

例如，如果我们想禁止太短的 user 的 name，我们可以创建一个 setter name，并将值存储在一个单独的属性 _name 中：

```js

let user = {
  get name() {
    return this._name;
  },

  set name(value) {
    if (value.length < 4) {
      alert("Name is too short, need at least 4 characters");
      return;
    }
    this._name = value;
  }
};

user.name = "Pete";
alert(user.name); // Pete

user.name = ""; // Name 太短了……
```

所以，name 被存储在 _name 属性中，并通过 getter 和 setter 进行访问。

从技术上讲，外部代码可以使用 user._name 直接访问 name。但是，这儿有一个众所周知的约定，即以下划线 "_" 开头的属性是内部属性，不应该从对象外部进行访问。

### 兼容性

### 参考链接 

- [属性的 getter 和 setter](https://zh.javascript.info/property-accessors)