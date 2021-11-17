---
title: "Js Modules"
date: 2021-11-17T17:57:03+08:00
draft: false
---

## javascript 模块化

Javascript模块化编程 CommonJS、AMD、CMD、ES6

## CommonJS

CommonJS定义的模块分为: 模块引用(require) 模块输出(exports) 模块标识(module)

CommonJS Modules有1.0、1.1、1.1.1三个版本:

Node.js、SproutCore实现了 Modules 1.0
SeaJS、AvocadoDB、CouchDB等实现了Modules 1.1.1
SeaJS、FlyScript实现了Modules/Wrappings

这里的CommonJS规范指的是CommonJS Modules/1.0规范。

CommonJS是一个更偏向于服务器端的规范。NodeJS采用了这个规范。CommonJS的一个模块就是一个脚本文件。require命令第一次加载该脚本时就会执行整个脚本，然后在内存中生成一个对象。


{
  id: '...',
  exports: { ... },
  loaded: true,
  ...
}

id是模块名，exports是该模块导出的接口，loaded表示模块是否加载完毕。此外还有很多属性，这里省略了。

以后需要用到这个模块时，就会到exports属性上取值。即使再次执行require命令，也不会再次执行该模块，而是到缓存中取值。

```js
// math.js
exports.add = function(a, b) {
  return a + b;
}
var math = require('math');
math.add(2, 3); // 512

```

> 由于CommonJS是同步加载模块，这对于服务器端不是一个问题，因为所有的模块都放在本地硬盘。等待模块时间就是硬盘读取文件时间，很小。但是，对于浏览器而言，它需要从服务器加载模块，涉及到网速，代理等原因，一旦等待时间过长，浏览器处于”假死”状态。

所以在浏览器端，不适合于CommonJS规范。所以在浏览器端又出现了一个规范—AMD(AMD是RequireJs在推广过程中对模块定义的规范化产出)。

## AMD
CommonJS解决了模块化的问题，但这种同步加载方式并不适合于浏览器端。

AMD是”Asynchronous Module Definition”的缩写，即”异步模块定义”。它采用异步方式加载模块，模块的加载不影响它后面语句的运行。
这里异步指的是不堵塞浏览器其他任务（dom构建，css渲染等），而加载内部是同步的（加载完模块后立即执行回调）。

AMD也采用require命令加载模块，但是不同于CommonJS，它要求两个参数：

require([module], callback);1

第一个参数[module]，是一个数组，里面的成员是要加载的模块，callback是加载完成后的回调函数。如果将上述的代码改成AMD方式：

require(['math'], function(math) {
  math.add(2, 3);
})

其中，回调函数中参数对应数组中的成员（模块）。

requireJS加载模块，采用的是AMD规范。也就是说，模块必须按照AMD规定的方式来写。

具体来说，就是模块书写必须使用特定的define()函数来定义。如果一个模块不依赖其他模块，那么可以直接写在define()函数之中。

define(id?, dependencies?, factory);12
id：模块的名字，如果没有提供该参数，模块的名字应该默认为模块加载器请求的指定脚本的名字；
dependencies：模块的依赖，已被模块定义的模块标识的数组字面量。依赖参数是可选的，如果忽略此参数，它应该默认为 ["require", "exports", "module"]。然而，如果工厂方法的长度属性小于3，加载器会选择以函数的长度属性指定的参数个数调用工厂方法。
factory：模块的工厂函数，模块初始化要执行的函数或对象。如果为函数，它应该只被执行一次。如果是对象，此对象应该为模块的输出值。
假定现在有一个math.js文件，定义了一个math模块。那么，math.js书写方式如下：

// math.js
define(function() {
  var add = function(x, y) {
    return x + y;
  }

  return  {
    add: add
  }
})
加载方法如下：

// main.js
require(['math'], function(math) {
  alert(math.add(1, 1));
})
如果math模块还依赖其他模块，写法如下：

// math.js
define(['dependenceModule'], function(dependenceModule) {
    // ...
})
当require()函数加载math模块的时候，就会先加载dependenceModule模块。当有多个依赖时，就将所有的依赖都写在define()函数第一个参数数组中，所以说AMD是依赖前置的。这不同于CMD规范，它是依赖就近的。

## CMD

CMD推崇依赖就近，延迟执行。可以把你的依赖写进代码的任意一行，如下：

define(factory)
factory为函数时，表示是模块的构造方法。执行该构造方法，可以得到模块向外提供的接口。factory 方法在执行时，默认会传入三个参数：require、exports 和 module.

```js
// CMD
define(function(require, exports, module) {
  var a = require('./a');
  a.doSomething();
  var b = require('./b');
  b.doSomething();
})
```
如果使用AMD写法，如下：

```js
// AMD
define(['a', 'b'], function(a, b) {
  a.doSomething();
  b.doSomething();
})
```
这个规范实际上是为了Seajs的推广然后搞出来的。那么看看SeaJS是怎么回事儿吧，基本就是知道这个规范了。

同样Seajs也是预加载依赖js跟AMD的规范在预加载这一点上是相同的，明显不同的地方是调用，和声明依赖的地方。AMD和CMD都是用difine和require，但是CMD标准倾向于在使用过程中提出依赖，就是不管代码写到哪突然发现需要依赖另一个模块，那就在当前代码用require引入就可以了，规范会帮你搞定预加载，你随便写就可以了。但是AMD标准让你必须提前在头部依赖参数部分写好（没有写好？ 倒回去写好咯）。这就是最明显的区别。

sea.js通过sea.use()来加载模块。

seajs.use(id, callback?)

## ES6

es6模块特性，推荐参看阮一峰老师的：ECMAScript 6 入门 - Module 的语法

说起 ES6 模块特性，那么就先说说 ES6 模块跟 CommonJS 模块的不同之处。

- ES6 模块输出的是值的引用，输出接口动态绑定，而 CommonJS 输出的是值的拷贝
- ES6 模块编译时执行，而 CommonJS 模块总是在运行时加载

## ES6 模块和 CommonJS 模块的区别

### CommonJS 输出值的拷贝

CommonJS 模块输出的是值的拷贝(原始值的拷贝)，也就是说，一旦输出一个值，模块内部的变化就影响不到这个值。

```js
// a.js
var b = require('./b');
console.log(b.foo);
setTimeout(() => {
  console.log(b.foo);
  console.log(require('./b').foo);
}, 1000);

// b.js
let foo = 1;
setTimeout(() => {
  foo = 2;
}, 500);
module.exports = {
  foo: foo,
};
// 执行：node a.js
// 执行结果：
// 1
// 1
// 1
```
上面代码说明，b 模块加载以后，它的内部 foo 变化就影响不到输出的 exports.foo 了。这是因为 foo 是一个原始类型的值，会被缓存。所以如果你想要在 CommonJS 中动态获取模块中的值，那么就需要借助于函数延时执行的特性。

```js
// a.js
var b = require('./b');
console.log(b.foo());
setTimeout(() => {
  console.log(b.foo());
  console.log(require('./b').foo());
}, 1000);

// b.js
let foo = 1;
setTimeout(() => {
  foo = 2;
}, 500);
module.exports = {
  foo: () => {
    return foo;
  },
};
// 执行：node a.js
// 执行结果：
// 1
// 2
// 2
```

所以我们可以总结一下：

- CommonJS 模块重复引入的模块并不会重复执行，再次获取模块直接获得暴露的 module.exports 对象
- 如果你要处处获取到模块内的最新值的话，也可以你每次更新数据的时候每次都要去更新 module.exports 上的值
- 如果你暴露的 module.exports 的属性是个对象，那就不存在这个问题了
  
所以如果你要处处获取到模块内的最新值的话，也可以你每次更新数据的时候每次都要去更新 module.exports 上的值，比如：

```js
// a.js
var b = require('./b');
console.log(b.foo);
setTimeout(() => {
  console.log(b.foo);
  console.log(require('./b').foo);
}, 1000);

// b.js
module.exports.foo = 1;   // 同 exports.foo = 1 
setTimeout(() => {
  module.exports.foo = 2;
}, 500);

// 执行：node a.js
// 执行结果：
// 1
// 2
// 2
```

### ES6 输出值的引用

然而在 ES6 模块中就不再是生成输出对象的拷贝，而是动态关联模块中的值。

### ES6 静态编译，CommonJS 运行时加载

关于第二点，ES6 模块编译时执行会导致有以下两个特点：

import 命令会被 JavaScript 引擎静态分析，优先于模块内的其他内容执行。

export 命令会有变量声明提前的效果。

import 优先执行:

从第一条来看，在文件中的任何位置引入 import 模块都会被提前到文件顶部。

```js
// a.js
console.log('a.js')
import { foo } from './b';

// b.js
export let foo = 1;
console.log('b.js 先执行');

// 执行结果:
// b.js 先执行
// a.js
```
从执行结果我们可以很直观地看出，虽然 a 模块中 import 引入晚于 console.log('a')，但是它被 JS 引擎通过静态分析，提到模块执行的最前面，优于模块中的其他部分的执行。


由于 import 是静态执行，所以 import 具有提升效果即 import 命令在模块中的位置并不影响程序的输出。

```js
/ a.js
import { foo } from './b';
console.log('a.js');
export const bar = 1;
export const bar2 = () => {
  console.log('bar2');
}
export function bar3() {
  console.log('bar3');
}

// b.js
export let foo = 1;
import * as a from './a';
console.log(a);

// 执行结果:
// { bar: undefined, bar2: undefined, bar3: [Function: bar3] }
// a.js
```

从上面的例子可以很直观地看出，a 模块引用了 b 模块，b 模块也引用了 a 模块，export 声明的变量也是优于模块其它内容的执行的，但是具体对变量赋值需要等到执行到相应代码的时候。(当然函数声明和表达式声明不一样，这一点跟 JS 函数性质一样，这里就不过多解释)



好了，讲完了 ES6 模块和 CommonJS 模块的不同点之后，接下来就讲讲相同点：

## ES6 模块和 CommonJS 模块的相同点

### 模块不会重复执行

这个很好理解，无论是 ES6 模块还是 CommonJS 模块，当你重复引入某个相同的模块时，模块只会执行一次。

CommonJS 模块循环依赖

```js

// a.js
console.log('a starting');
exports.done = false;
const b = require('./b');
console.log('in a, b.done =', b.done);
exports.done = true;
console.log('a done');

// b.js
console.log('b starting');
exports.done = false;
const a = require('./a');
console.log('in b, a.done =', a.done);
exports.done = true;
console.log('b done');

// node a.js
// 执行结果：
// a starting
// b starting
// in b, a.done = false
// b done
// in a, b.done = true
// a done

```
结合之前讲的特性很好理解，当你从 b 中想引入 a 模块的时候，因为 node 之前已经加载过 a 模块了，所以它不会再去重复执行 a 模块，而是直接去生成当前 a 模块吐出的 module.exports 对象，因为 a 模块引入 b 模块先于给 done 重新赋值，所以当前 a 模块中输出的 module.exports 中 done 的值仍为 false。而当 a 模块中输出 b 模块的 done 值的时候 b 模块已经执行完毕，所以 b 模块中的 done 值为 true。



从上面的执行过程中，我们可以看到，在 CommonJS 规范中，当遇到 require() 语句时，会执行 require 模块中的代码，并缓存执行的结果，当下次再次加载时不会重复执行，而是直接取缓存的结果。正因为此，出现循环依赖时才不会出现无限循环调用的情况。虽然这种模块加载机制可以避免出现循环依赖时报错的情况，但稍不注意就很可能使得代码并不是像我们想象的那样去执行。因此在写代码时还是需要仔细的规划，以保证循环模块的依赖能正确工作。



所以有什么办法可以出现循环依赖的时候避免自己出现混乱呢？一种解决方式便是将每个模块先写 exports 语法，再写 requre 语句，利用 CommonJS 的缓存机制，在 require() 其他模块之前先把自身要导出的内容导出，这样就能保证其他模块在使用时可以取到正确的值。比如：

```js

// a.js
exports.done = true;
let b = require('./b');
console.log(b.done)

// b.js
exports.done = true;
let a = require('./a');
console.log(a.done)

```

这种写法简单明了，缺点是要改变每个模块的写法，而且大部分同学都习惯了在文件开头先写 require 语句。

### ES6 模块循环依赖

跟 CommonJS 模块一样，ES6 不会再去执行重复加载的模块，又由于 ES6 动态输出绑定的特性，能保证 ES6 在任何时候都能获取其它模块当前的最新值。

```js

// a.js
console.log('a starting')
import {foo} from './b';
console.log('in b, foo:', foo);
export const bar = 2;
console.log('a done');

// b.js
console.log('b starting');
import {bar} from './a';
export const foo = 'foo';
console.log('in a, bar:', bar);
setTimeout(() => {
  console.log('in a, setTimeout bar:', bar);
})
console.log('b done');

// babel-node a.js
// 执行结果：
// b starting
// in a, bar: undefined
// b done
// a starting
// in b, foo: foo
// a done
// in a, setTimeout bar: 2

```


### 动态 import()

ES6 模块在编译时就会静态分析，优先于模块内的其他内容执行，所以导致了我们无法写出像下面这样的代码：

```js
if(some condition) {
  import a from './a';
}else {
  import b from './b';
}

// or 
import a from (str + 'b');

```
因为编译时静态分析，导致了我们无法在条件语句或者拼接字符串模块，因为这些都是需要在运行时才能确定的结果在 ES6 模块是不被允许的，所以 动态引入 import() 应运而生。



import() 允许你在运行时动态地引入 ES6 模块，想到这，你可能也想起了 require.ensure 这个语法，但是它们的用途却截然不同的。

require.ensure 的出现是 webpack 的产物，它是因为浏览器需要一种异步的机制可以用来异步加载模块，从而减少初始的加载文件的体积，所以如果在服务端的话 require.ensure 就无用武之地了，因为服务端不存在异步加载模块的情况，模块同步进行加载就可以满足使用场景了。 CommonJS 模块可以在运行时确认模块加载。
而 import() 则不同，它主要是为了解决 ES6 模块无法在运行时确定模块的引用关系，所以需要引入 import()
我们先来看下它的用法：

动态的 import() 提供一个基于 Promise 的 API
动态的import() 可以在脚本的任何地方使用
import() 接受字符串文字，你可以根据你的需要构造说明符

举个简单的使用例子：

```js

// a.js
const str = './b';
const flag = true;
if(flag) {
  import('./b').then(({foo}) => {
    console.log(foo);
  })
}
import(str).then(({foo}) => {
  console.log(foo);
})

// b.js
export const foo = 'foo';

// babel-node a.js
// 执行结果
// foo
// foo

```
当然，如果在浏览器端的 import() 的用途就会变得更广泛，比如 按需异步加载模块，那么就和 require.ensure 功能类似了。



因为是基于 Promise 的，所以如果你想要同时加载多个模块的话，可以是 Promise.all 进行并行异步加载。

```js
Promise.all([
  import('./a.js'),
  import('./b.js'),
  import('./c.js'),
]).then(([a, {default: b}, {c}]) => {
    console.log('a.js is loaded dynamically');
    console.log('b.js is loaded dynamically');
    console.log('c.js is loaded dynamically');
});

```
还有 Promise.race 方法，它检查哪个 Promise 被首先 resolved 或 reject。我们可以使用import()来检查哪个CDN速度更快：

```js
const CDNs = [
  {
    name: 'jQuery.com',
    url: 'https://code.jquery.com/jquery-3.1.1.min.js'
  },
  {
    name: 'googleapis.com',
    url: 'https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js'
  }
];

console.log(`------`);
console.log(`jQuery is: ${window.jQuery}`);

Promise.race([
  import(CDNs[0].url).then(()=>console.log(CDNs[0].name, 'loaded')),
  import(CDNs[1].url).then(()=>console.log(CDNs[1].name, 'loaded'))
]).then(()=> {
  console.log(`jQuery version: ${window.jQuery.fn.jquery}`);
});

```
当然，如果你觉得这样写还不够优雅，也可以结合 async/await 语法糖来使用。

```js
async function main() {
  const myModule = await import('./myModule.js');
  const {export1, export2} = await import('./myModule.js');
  const [module1, module2, module3] =
    await Promise.all([
      import('./module1.js'),
      import('./module2.js'),
      import('./module3.js'),
    ]);
}
```

动态 import() 为我们提供了以异步方式使用 ES 模块的额外功能。 根据我们的需求动态或有条件地加载它们，这使我们能够更快，更好地创建更多优势应用程序。

### export

一个模块就是一个独立的文件。该文件内部的所有变量，外部无法获取。如果希望外部文件能够读取该模块的变量，就需要在这个模块内使用export关键字导出变量。如：

```js
// profile.jsexport var a = 1;export var b = 2;export var c = 3;1234
下面的写法是等价的，这种方式更加清晰（在底部一眼能看出导出了哪些变量）：

var a = 1;var b = 2;var c = 3;
export {a, b, c}1234
import
// import命令可以导入其他模块通过export导出的部分。

var a = 1;var b = 2;var c = 3;
export {a, b, c}
//main.js
import {a, b, c} from './abc';
console.log(a, b, c);

```
如果想为导入的变量重新取一个名字，使用as关键字（也可以在导出中使用）。

import {a as aa, b, c};
console.log(aa, b, c)12
如果想在一个模块中先输入后输出一个模块，import语句可以和export语句写在一起。

import {a, b, c} form './abc';export {a, b,  c}// 使用连写, 可读性不好，不建议export {a, b, c} from './abc';12345
模块的整体加载
使用*关键字。

```js
import * from as abc form './abc';
export default
```

在export输出内容时，如果同时输出多个变量，需要使用大括号{}，同时导入也需要大括号。使用export defalut输出时，不需要大括号，而输入（import）export default输出的变量时，不需要大括号。

```js
// abc.jsvar a = 1, b = 2, c = 3;export {a, b};export default c;1234
import {a, b} from './abc';
import c from './abc'; // 不需要大括号console.log(a, b, c) // 1 2 3123

```

本质上，export default输出的是一个叫做default的变量或方法，输入这个default变量时不需要大括号。

```js
// abc.js
export {a as default};

// main.js
import a from './abc'; // 这样也是可以的
import {default as aa} from './abc'; // 这样也是可以的
console.log(aa);123456789
```

就到这里了吧。关于循环加载（模块相互依赖）没写，CommonJS和ES6处理方式不一样。

### 参考文章：

- [javascript模块化之CommonJS、AMD、CMD、UMD、ES6](https://blog.csdn.net/Real_Bird/article/details/54869066)
- 深入理解 ES6 模块机制
- 该如何理解AMD ，CMD，CommonJS规范–javascript模块化加载学习总结
- AMD/CMD与前端规范
- 前端模块化之旅（二）：CommonJS、AMD和CMD
- 研究一下javascript的模块规范（CommonJs/AMD/CMD）
- Javascript模块化编程（一）：模块的写法
- Javascript模块化编程（二）：AMD规范
- Javascript模块化编程（三）：require.js的用法
- ECMAScript 6入门

