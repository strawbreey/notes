---
title: "Js Module"
date: 2020-10-16T00:13:48+08:00
draft: true
---

，JavaScript 很长一段时间内都没有语言级（language-level）的模块语法。这不是一个问题，因为最初的脚本又小又简单，所以没必要将其模块化。

但是最终脚本变得越来越复杂，因此社区发明了许多种方法来将代码组织到模块中，使用特殊的库按需加载模块。

[AMD](https://en.wikipedia.org/wiki/Asynchronous_module_definition) —— 最古老的模块系统之一，最初由 require.js 库实现。
[CommonJS](https://en.wikipedia.org/wiki/Asynchronous_module_definition) —— 为 Node.js 服务器创建的模块系统。
[UMD](https://en.wikipedia.org/wiki/Asynchronous_module_definition) —— 另外一个模块系统，建议作为通用的模块系统，它与 AMD 和 CommonJS 都兼容。

语言级的模块系统在 2015 年的时候出现在了标准（ES6）中，此后逐渐发展，现在已经得到了所有主流浏览器和 Node.js 的支持。因此，我们将从现在开始学习现代 JavaScript 模块（module）。

### 什么是模块？

- export 关键字标记了可以从当前模块外部访问的变量和函数。
- import 关键字允许从其他模块导入功能。


```js
//  sayHi.js
export function sayHi(user) {
  alert(`Hello, ${user}!`);
}

// 📁 main.js
import {sayHi} from './sayHi.js';

alert(sayHi); // function...
sayHi('John'); // Hello, John!
```

由于模块支持特殊的关键字和功能，因此我们必须通过使用 `<script type="module">` 特性（attribute）来告诉浏览器，此脚本应该被当作模块（module）来对待。

```html
<!doctype html>
<script type="module">
  import {sayHi} from './say.js';

  document.body.innerHTML = sayHi('John');
</script>
```

浏览器会自动获取并解析（evaluate）导入的模块（如果需要，还可以分析该模块的导入），然后运行该脚本。


    模块只通过 HTTP(s) 工作，在本地文件则不行
    如果你尝试通过 file:// 协议在本地打开一个网页，你会发现 import/export 指令不起作用。你可以使用本地 Web 服务器，例如 static-server，或者使用编辑器的“实时服务器”功能，例如 VS Code 的 Live Server Extension 来测试模块。

模块核心功能

- 始终使用 “use strict”

    模块始终默认使用 use strict，例如，对一个未声明的变量赋值将产生错误

- 模块级作用域

    每个模块都有自己的顶级作用域（top-level scope）。换句话说，一个模块中的顶级作用域变量和函数在其他脚本中是不可见的。
