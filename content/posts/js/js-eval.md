---
title: "Js Eval"
date: 2021-10-21T10:03:14+08:00
draft: false
---

```js
  execFunction = (name: string) => (new Function( 'return ' + name))();

  execEval = (code: string) => eval('(' + code + ')')
```


### eval()

以 str 的方式运行 JavaScript 代码，比如：

```js
var a = 12;
eval( 'a+5' ); // 17
```

注意语句上下文 eval()的解析：

```js
eval('{ foo: 123 }');  // 123
eval('({ foo: 123 })'); // { foo: 123 }
```

1. 严格模式下的 eval()

对于 eval()，理应当在严格模式下使用。在松散模式下运行代码会在当前的作用域中创建局部变量：
```js
function f(){ 
    eval( 'var foo = 1' );
    console.log( foo ); // 1
} 
```
严格模式下就不会出现该情况。但是，运行代码仍然具有读写当前作用域中变量的权限。你需要通过间接调用 eval() 来阻止这种权限。


2. 全局作用域下间接执行 eval() 有两种调用 eval() 的方式：

- 直接方式：通过直接调用名为 "eval" 的函数。
- 间接方式：使用其他的一些方式。（通过 call 调用，将其以其他名字作为 window 下的一个方法存储，在那里进行调用） 之前已经看过，在当前作用域直接使用 eval 执行代码

```js
   var x = 'global'; 
   function directEval(){
       'use strict';
        var x = 'local';
        console.log(eval('x')); // local
   } 
```

反之，在全局作用域中间接调用 eval。 

```js
    var x = 'global'; 
    function indirectEval(){
        'use strict';
        var x = 'local';
        // 不同方式调用 call 
        console.log( eval.call(null, 'x') ); // global
        console.log( window.eval('x') ); // global
        console.log( (1,eval)('x') ); // global (1)
        var xeval = eval;
        console.log( xeval('x') ); // global
        var obj = { eval: eval }
        console.log( obj.eval('x') ); // global
    } 
```

说明：当你通过一个名称来引用一个变量的时候，其初始值为一个所谓的引用，数据结构分为两部分： 
1. 基础是指向存储变量的值的数据结构。 
2. 引用名是变量的名称 
   
在一个函数调用 eval 的时候，该函数的调用操作符（括号）遇到一个对 eval 的引用可以确定被调用函数的名称。所以此时函数会触发一个直接的 eval 调用。当然你可以不给调用操作符引用来强制间接调用 eval。这是由于在操作符运行之前获取引用的值来实现的。在 （1）标注的那一行，逗号操作符为我们实现的这点。这个运算符运行了第一个运算元并返回了第二个运算元的结果。运算结果总是返回 值 的，意味着引用已经被解析。 间接的运行代码总是松散的。这是由于代码被独立的在当前环境中运行的结果。 

```js
function strictFunc(){ 
    'use strict';
    var code = '(function(){ return this; }())';
    var result = eval.call( null, code );
    console.log( result !== undefined ); // true sloppy mode
}
```

### new Function()

Function 构造函数的函数签名。

`new Function( param1, param2, ..., paramN,funcBody );`

它创建一个包含0个或者过个参数名为 param1 等的函数，函数体为 funcBody。相当于如下方式创建函数：
```js

    function( (param1), (param2), ..., (paramN) ){ 
        (funcBody)
    } 

    // > var f = new Function('x', 'y', 'return x+y'); 
    // > f( 3, 4 ) 
```
例如： 

与间接 eval 调用类似，new Function() 创建的函数作用域也是全局的。 

```js
var x = 'global'; 
function strictFunc(){
   'use strict';
   var x = 'local';
   var f = new Function('return x');
   console.log( f() ); //global
} 
```
以下的函数也是默认松散模式 

```js
function strictFunc(){ 
    'use strict';
    var sl = new Function( 'return this' );
    console.log( sl() !== undefined ); // true, sloppy mode
    var st = new Function( '"use strict"; return this;' );
    console.log( st() !== undefined ); // true, strict mode
} 
```

### eval() 对比 new Function()

一般来说，使用 new Function() 运行代码比 eval() 更为好一些：函数的参数提供了清晰的接口来运行代码，而没有必要使用较为笨拙的语法来间接的调用 eval() 确保代码只能访问自己的和全局的变量。

### 最佳实践

通常：避免使用 eval() 和 new Function() 。动态运行代码不但速度较慢，还有潜在的安全风险。一般都可以找到更好地替代方案。
如，Brendan Eich 最近在 tweete发了一个程序猿需要访问某一个属性，其属性名被存储在名为propName的变量中的反模式：
```js
var value = eval( 'obj.'+propName );

// => 以下方式会更OK： 
var value = obj[ propName ]; 
```

你也不应该使用 eval() 或者 new Function() 来解析 JSON格式的数据。那也是不安全的。要么使用 ECMAScript 5 内置的对JSON的支持方法，要么使用一个类库。 
    
合理使用实例。依旧有一些较为合理，对 eval() 和 new Function() 使用较为高级的：配置函数数据（JSON 不允许），模板库，解析，命令行和模块系统。 

### 总结 
 
这是对于动态运行代码较高层次的概述，如果想更深的了解可以看一下 kanhax的文章 “[全局eval，何去何从？](http://perfectionkills.com/global-.-what-are-the-options/)” 
 

### 参考文献 
【1】 [JavaScript 的 表达式和语句](http://www.2ality.com/2012/09/.s-vs-statements.html)  
【2】 [JavaScript](http://www.2ality.com/2011/01/.s-strict-mode-summary.html)[ 严格模式：综述](http://www.2ality.com/2011/01/.s-strict-mode-summary.html)   
【3】 [JavaScript的 JSON API](http://www.2ality.com/2011/08/json-api.html)   

