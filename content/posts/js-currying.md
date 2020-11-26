---
title: "Js Currying"
date: 2020-10-26T19:21:37+08:00
draft: false
---

## 函数柯里化

函数柯里化概念： 柯里化（Currying）是把接受多个参数的函数转变为接受一个单一参数的函数，并且返回接受余下的参数且返回结果的新函数的技术

柯里化（Currying）是一种关于函数的高阶技术。它不仅被用于 JavaScript，还被用于其他编程语言。

柯里化是一种函数的转换，它是指将一个函数从可调用的 f(a, b, c) 转换为可调用的 f(a)(b)(c)。

柯里化不会调用函数。它只是对函数进行转换。
```js
function add (a) {
	return function (b) {
		return function (c) {
		    return a + b + c;
		}
	}
}

// 输入
add(1)(2)(3); // 6

// 柯里化解决方案
const curry = (fn) =>
(judge = (...args) => args.length === fn.length ? fn(...args) : (...arg) => judge(...args, ...arg));

const add = (a, b, c) => a + b + c;

const curryAdd = curry(add);

console.log(curryAdd(1)(2)(3)); // 6
console.log(curryAdd(1, 2)(3)); // 6
console.log(curryAdd(1)(2, 3)); // 6

// 参数长度不固定
function add (...args) {
    //求和
    return args.reduce((a, b) => a + b)
}

function currying (fn) {
    let args = []
    return function temp (...newArgs) {
        if (newArgs.length) {
            args = [...args, ...newArgs]
            return temp
        } else {
            let val = fn.apply(this, args)
            args = [] //保证再次调用时清空
            return val
        }
    }
}

let addCurry = currying(add)
console.log(addCurry(1)(2)(3)(4, 5)())  //15
console.log(addCurry(1)(2)(3, 4, 5)())  //15
console.log(addCurry(1)(2, 3, 4, 5)())  //15
```

### Lodash 中的 curry

柯里化更高级的实现，例如 lodash 库的 _.curry，会返回一个包装器，该包装器允许函数被正常调用或者以偏函数（partial）的方式调用：

```js
function sum(a, b) {
  return a + b;
}

let curriedSum = _.curry(sum); // 使用来自 lodash 库的 _.curry

curriedSum(1, 2); // 3，仍可正常调用
curriedSum(1)(2); // 3，以偏函数的方式调用
```

### 柯里化？目的是什么？

例如，我们有一个用于格式化和输出信息的日志（logging）函数 log(date, importance, message)。在实际项目中，此类函数具有很多有用的功能，例如通过网络发送日志（log），在这儿我们仅使用 alert

```js

function log(date, importance, message) {
  alert(`[${date.getHours()}:${date.getMinutes()}] [${importance}] ${message}`);
}

// 让函数柯里话
log = _.curry(log);

// 柯里化之后，log 仍正常运行：
log(new Date(), "DEBUG", "some debug"); // log(a, b, c)

// 但是也可以以柯里化形式运行 => 疑惑, 会不会给使用者造成负担
log(new Date())("DEBUG")("some debug"); // log(a)(b)(c)

// 我们可以轻松地为当前日志创建便捷函数：

let logNow = log(new Date());

// 使用它
logNow("INFO", "message"); // [HH:mm] INFO message
```

优点:

- 柯里化之后，我们没有丢失任何东西：log 依然可以被正常调用。
- 我们可以轻松地生成偏函数，例如用于生成今天的日志的偏函数。


### 高级柯里化实现

```js
function curry(func) {
  return function curried(...args) {
    if (args.length >= func.length) {
      return func.apply(this, args);
    } else {
      return function(...args2) {
        return curried.apply(this, args.concat(args2));
      }
    }
  };
}

// 使用
function sum(a, b, c) {
  return a + b + c;
}

let curriedSum = curry(sum);

curriedSum(1, 2, 3); // 6，仍然可以被正常调用
curriedSum(1)(2,3); // 6，对第一个参数的柯里化
curriedSum(1)(2)(3); // 6，全柯里化
```

说明: 当我们运行它时，这里有两个 if 执行分支：

- 现在调用：如果传入的 args 长度与原始函数所定义的（func.length）相同或者更长，那么只需要将调用传递给它即可。
- 获取一个偏函数：否则，func 还没有被调用。取而代之的是，返回另一个包装器 pass，它将重新应用 curried，将之前传入的参数与新的参数一起传入。然后，在一个新的调用中，再次，我们将获得一个新的偏函数（如果参数不足的话），或者最终的结果。


### 总结

柯里化 是一种转换，将 f(a,b,c) 转换为可以被以 f(a)(b)(c) 的形式进行调用。JavaScript 实现通常都保持该函数可以被正常调用，并且如果参数数量不足，则返回偏函数。

柯里化让我们能够更容易地获取偏函数。就像我们在日志记录示例中看到的那样，普通函数 log(date, importance, message) 在被柯里化之后，当我们调用它的时候传入一个参数（如 log(date)）或两个参数（log(date, importance)）时，它会返回偏函数。

### 参考链接 
- [柯里化（Currying）](https://zh.javascript.info/currying-partials)
- [实现 add(1)(2)(3)](https://github.com/lgwebdream/FE-Interview/issues/21)
