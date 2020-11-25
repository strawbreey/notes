---
title: "Js Closure"
date: 2020-09-07T17:02:22+08:00
draft: false
---

函数和对其周围状态（lexical environment，词法环境）的引用捆绑在一起构成闭包（closure）。也就是说，闭包可以让你从内部函数访问外部函数作用域。在 JavaScript 中，每当函数被创建，就会在函数生成时生成闭包。

#### 词法作用域

```js
function init() {
    var name = "Mozilla";       // name 是一个被 init 创建的局部变量
    function displayName() {    // displayName() 是内部函数，一个闭包
        alert(name);            // 使用了父函数中声明的变量
    }
    displayName();
}
init();
```

init() 创建了一个局部变量 name 和一个名为 displayName() 的函数。displayName() 是定义在 init() 里的内部函数，并且仅在 init() 函数体内可用。请注意，displayName() 没有自己的局部变量。然而，因为它可以访问到外部函数的变量，所以 displayName() 可以使用父函数 init() 中声明的变量 name 。

#### 闭包
```js
function makeFunc() {
    var name = "Mozilla";
    function displayName() {
        alert(name);
    }
    return displayName;
}

var myFunc = makeFunc();
myFunc();

//  => 

function makeFunc() {
    var name = "Mozilla";
    function displayName() {
        alert(name);
    }
    return displayName;
}

var myFunc = makeFunc();

var myFunc =  displayName() {
        alert(name);
    }

(displayName() {
  alert(name);
})()

```


```js
function makeAdder(x) {
  return function(y) {
    return x + y;
  };
}

var add5 = makeAdder(5);
var add10 = makeAdder(10);

console.log(add5(2));  // 7
console.log(add10(2)); // 12

```

特点: 

add5 和 add10 都是闭包。它们`共享`相同的函数定义，但是`保存`了不同的词法环境。在 add5 的环境中，x 为 5。而在 add10 中，x 则为 10。