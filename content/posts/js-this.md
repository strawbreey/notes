---
title: "Js This"
date: 2020-10-29T11:15:16+08:00
draft: false
---

在 JavaScript 中，this 关键字与其他大多数编程语言中的不同。JavaScript 中的 this 可以用于任何函数。

this 的值是在代码运行时计算出来的，它取决于代码上下文。

```js
let user = { name: "John" };
let admin = { name: "Admin" };

function sayHi() {
  alert( this.name );
}

// 在两个对象中使用相同的函数
user.f = sayHi;
admin.f = sayHi;

// 这两个调用有不同的 this 值
// 函数内部的 "this" 是“点符号前面”的那个对象
user.f(); // John（this == user）
admin.f(); // Admin（this == admin）

admin['f'](); // Admin（使用点符号或方括号语法来访问这个方法，都没有关系。）
```

这个规则很简单：如果 obj.f() 被调用了，则 this 在 f 函数调用期间是 obj。所以在上面的例子中 this 先是 user，之后是 admin。


在没有对象的情况下调用：this == undefined

```js
function sayHi() {
  alert(this);
}

sayHi(); // undefined
```

在这种情况下，严格模式下的 this 值为 undefined。如果我们尝试访问 this.name，将会报错。

在非严格模式的情况下，this 将会是 全局对象（浏览器中的 window，我们稍后会在 全局对象 一章中学习它）。这是一个历史行为，"use strict" 已经将其修复了。

如果你经常使用其他的编程语言，那么你可能已经习惯了“绑定 this”的概念，即在对象中定义的方法总是有指向该对象的 this。

在 JavaScript 中，this 是“自由”的，它的值是在调用时计算出来的，它的值并不取决于方法声明的位置，而是取决于在“点符号前”的是什么对象

在运行时对 this 求值的这个概念既有优点也有缺点。一方面，函数可以被重用于不同的对象。另一方面，更大的灵活性造成了更大的出错的可能。

#### 箭头函数没有自己的 “this”

箭头函数有些特别：它们没有自己的 this。如果我们在这样的函数中引用 this，this 值取决于外部“正常的”函数。

```js
let user = {
  firstName: "Ilya",
  sayHi () {
    console.log(this.firstName);  //Ilya
    let arrow = () => console.log(this.firstName); // Ilya
    arrow();
  }
};

let user = {
  firstName: "Ilya",
  sayHi: () =>{
    console.log(this.firstName);  // undefined 严格模式下, this 指向为undefined, 否则为上下文
    let arrow = () => console.log(this.firstName); // undefined
    arrow();
  }
};

```

### 总结

- 存储在对象属性中的函数被称为“方法”。
- 方法允许对象进行像 object.doSomething() 这样的“操作”。
- 方法可以将对象引用为 this。

this 的值是在程序运行时得到的。

- 一个函数在声明时，可能就使用了 this，但是这个 this 只有在函数被调用时才会有值。
- 可以在对象之间复制函数。
- 以“方法”的语法调用函数时：object.method()，调用过程中的 this 值是 object。

请注意箭头函数有些特别：它们没有 this。在箭头函数内部访问到的 this 都是从外部获取的。

```js
function identify() {
	return this.name.toUpperCase();
}

function speak() {
	var greeting = "Hello, I'm " + identify.call( this );
	console.log( greeting );
}

var me = {
	name: "Kyle"
};

var you = {
	name: "Reader"
};

identify.call( me ); // KYLE
identify.call( you ); // READER

speak.call( me ); // Hello, I'm KYLE
speak.call( you ); // Hello, I'm READER
```


与使用 this 相反地，你可以明确地将环境对象传递给 identify() 和 speak()。

```js
function identify(context) {
	return context.name.toUpperCase();
}

function speak(context) {
	var greeting = "Hello, I'm " + identify( context );
	console.log( greeting );
}

identify( you ); // READER
speak( me ); // Hello, I'm KYLE
```

然而，this 机制提供了更优雅的方式来隐含地“传递”一个对象引用，导致更加干净的API设计和更容易的复用。

你的使用模式越复杂，你就会越清晰地看到：将执行环境作为一个明确参数传递，通常比传递 this 执行环境要乱。当我们探索对象和原型时，你将会看到一组可以自动引用恰当执行环境对象的函数是多么有用。


### this 指向

- 第一种常见的倾向是认为 this 指向函数自己

```js
function foo(num) {
	console.log( "foo: " + num );
	// 追踪 `foo` 被调用了多少次
	this.count++;
}

foo.count = 0;

var i;

for (i=0; i<10; i++) {
	if (i > 5) {
		foo( i );
	}
}
// foo: 6
// foo: 7
// foo: 8
// foo: 9

// `foo` 被调用了多少次？
console.log( foo.count ); // 0 -- 这是怎么回事……？
```

当代码执行 foo.count = 0 时，它确实向函数对象 foo 添加了一个 count 属性。但是对于函数内部的 this.count 引用，this 其实 根本就不 指向那个函数对象，即便属性名称一样，但根对象也不同，因而产生了混淆。

强制this指向

```js
function foo(num) {
	console.log( "foo: " + num );

	// 追踪 `foo` 被调用了多少次
	// 注意：由于 `foo` 的被调用方式（见下方），`this` 现在确实是 `foo`
	this.count++;
}

foo.count = 0;

var i;

for (i=0; i<10; i++) {
	if (i > 5) {
		// 使用 `call(..)`，我们可以保证 `this` 指向函数对象(`foo`)
		foo.call( foo, i );
	}
}
// foo: 6
// foo: 7
// foo: 8
// foo: 9

// `foo` 被调用了多少次？
console.log( foo.count ); // 4
```

this 作用域

明确地说，this 不会以任何方式指向函数的 词法作用域。作用域好像是一个将所有可用标识符作为属性的对象，这从内部来说是对的。但是 JavasScript 代码不能访问作用域“对象”。它是 引擎 的内部实现。

```
function foo() {
	var a = 2;
	this.bar();
}

function bar() {
	console.log( this.a );
}

foo(); //undefined
```

### 调用点和调用栈

为了理解 this 绑定，我们不得不理解调用点(一个函数是在哪里被调用的)：函数在代码中被调用的位置（不是被声明的位置）。我们必须考察调用点来回答这个问题：这个 this 指向什么？


考虑 调用栈（call-stack） （使我们到达当前执行位置而被调用的所有方法的堆栈）是十分重要的。我们关心的调用点就位于当前执行中的函数 之前 的调用。

我们来展示一下调用栈和调用点：

```js
function baz() {
    // 调用栈是: `baz`
    // 我们的调用点是 global scope（全局作用域）

    console.log( "baz" );
    bar(); // <-- `bar` 的调用点
}

function bar() {
    // 调用栈是: `baz` -> `bar`
    // 我们的调用点位于 `baz`

    console.log( "bar" );
    foo(); // <-- `foo` 的 call-site
}

function foo() {
    // 调用栈是: `baz` -> `bar` -> `foo`
    // 我们的调用点位于 `bar`

    console.log( "foo" );
}

baz(); // <-- `baz` 的调用点
```

你可以通过按顺序观察函数的调用链在你的大脑中建立调用栈的视图，就像我们在上面代码段中的注释那样。但是这很痛苦而且易错。另一种观察调用栈的方式是使用你的浏览器的调试工具。大多数现代的桌面浏览器都内建开发者工具，其中就包含 JS 调试器。在上面的代码段中，你可以在调试工具中为 foo() 函数的第一行设置一个断点，或者简单的在这第一行上插入一个 debugger 语句。当你运行这个网页时，调试工具将会停止在这个位置，并且向你展示一个到达这一行之前所有被调用过的函数的列表，这就是你的调用栈。所以，如果你想调查this 绑定，可以使用开发者工具取得调用栈，之后从上向下找到第二个记录，那就是你真正的调用点。

### this 绑定方式

#### 默认绑定（Default Binding）
```js
var a = 2;

function foo() {
	console.log( this.a );
}

foo(); // 2
// 非严格模式下, this 指向全局
```

```js
"use strict";
var a = 2;

function foo() {
	console.log( this.a );
}

foo(); // TypeError: `this` is `undefined`
```


#### 隐含绑定（Implicit Binding）

```js
function foo() {
	console.log( this.a );
}

var obj = {
	a: 2,
	foo: foo
};

obj.foo(); // 2
```

#### 隐含丢失（Implicitly Lost）
this 绑定最常让人沮丧的事情之一，就是当一个 隐含绑定 丢失了它的绑定，这通常意味着它会退回到 默认绑定， 根据 strict mode 的状态，其结果不是全局对象就是 undefined。

```js
function foo() {
	console.log( this.a );
}

var obj = {
	a: 2,
	foo: foo
};

var bar = obj.foo; // 函数引用！

var a = "oops, global"; // `a` 也是一个全局对象的属性

bar(); // "oops, global"
```

#### 明确绑定（Explicit Binding）

函数拥有 call(..) 和 apply(..) 方法

```js
function foo() {
	console.log( this.a );
}

var obj = {
	a: 2
};

foo.call( obj ); // 2
```

#### 硬绑定（Hard Binding)

```js
function foo() {
	console.log( this.a );
}

var obj = {
	a: 2
};

var bar = function() {
	foo.call( obj );
};

bar(); // 2
setTimeout( bar, 100 ); // 2

// `bar` 将 `foo` 的 `this` 硬绑定到 `obj`
// 所以它不可以被覆盖
bar.call( window ); // 2
```

创建了一个函数 bar()，在它的内部手动调用 foo.call(obj)，由此强制 this 绑定到 obj 并调用 foo。无论你过后怎样调用函数 bar，它总是手动使用 obj 调用 foo。这种绑定即明确又坚定，所以我们称之为 硬绑定（hard binding）


### 参考资料

- [你不懂JS: this 与对象原型] (https://github.com/getify/You-Dont-Know-JS/blob/1ed-zh-CN/this%20%26%20object%20prototypes/ch2.md)