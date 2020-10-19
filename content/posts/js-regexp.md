---
title: "Js Regexp"
date: 2020-10-13T17:41:49+08:00
draft: true
---

## 正则表达式

在 JavaScript 中，正则表达式通过内置的“RegExp”类的对象来实现，并与字符串集成。

创建正则表达式有两个方式

```js
// new RegExp
regexp = new RegExp("pattern", "flags");

// 表达式
regexp = /pattern/; // 没有修饰符
regexp = /pattern/gmi; // 伴随修饰符 g、m 和 i（后面会讲到）
```

用法

字符串搜索

```js
let str = "I love Javascript";
let substr = 'love';
alert(str.search(substr));
```

所以搜索 /love/ 与搜索 "love" 是等价的。

通常我们使用的都是简短语法 /.../。但是它不接受任何变量插入，所以我们必须在写代码的时候就知道确切的 regexp。

另一方面，new RegExp 允许从字符串中动态地构造模式。

所以我们可以找出需要搜索的字段，然后根据搜索字段创建 new RegExp：

```js

// 弹窗提示框, 默认输入值love
let search = prompt("What you want to search?", "love"); // love

// 
let regexp = new RegExp(search);

// 找到用户想要的任何东西
alert( "I love JavaScript".search(regexp));


```

#### 修饰符

正则表达式的修饰符可能会影响搜索结果。

在 JavaScript 中，有 5 个修饰符：

`i` 使用此修饰符后，搜索时不区分大小写: A 和 a 没有区别（具体看下面的例子）。

```js
let str = "I love JavaScript!";

alert( str.search(/LOVE/) ); // -1（没找到）
alert( str.search(/LOVE/i) ); // 2

// 第一个搜索返回的是 -1（也就是没找到），因为搜索默认是区分大小写的。
// 使用修饰符 /LOVE/i，在字符串的第 2 个位置上搜索到了 love。
```

`g` 使用此修饰符后，搜索时会查找所有的匹配项，而不只是第一个（在下一章会讲到）。

`m` 多行模式（详见章节 文章 "regexp-multiline" 未找到）。

`u` 开启完整的 unicode 支持。该修饰符能够修正对于代理对的处理。更详细的内容见章节 
Unicode：修饰符 “u” 和 class \p{...}。

`y` 粘滞模式（详见 下一章节）


### 字符类

1. 数字类 /d

```js
let str = "+7(903)-123-45-67";

let regexp = /\d/;

alert( str.match(regexp) ); // 7
```

如果没有标志 g，则正则表达式仅查找第一个匹配项，即第一个数字 \d。

让我们添加 g标志来查找所有数字

```js
let str = "+7(903)-123-45-67";

let regexp = /\d/g;

alert( str.match(regexp) ); // array of matches: 7,9,0,3,1,2,3,4,5,6,7

// let's make the digits-only phone number of them:
alert( str.match(regexp).join('') ); // 79035419441
```


\d（“d” 来自 “digit”） 数字：从 0 到 9 的字符。

\s（“s” 来自 “space”） 空格符号：包括空格，制表符 \t，换行符 \n 和其他少数稀有字符，例如 \v，\f 和 \r。

\w（“w” 来自 “word”） “单字”字符：拉丁字母或数字或下划线 _。非拉丁字母（如西里尔字母或印地文）不属于 \w。

```js
let str = "Is there CSS4?";
let regexp = /CSS\d/

alert( str.match(regexp) ); // CSS4

// 
alert( "I love HTML5!".match(/\s\w\w\w\w\d/) ); // ' HTML5'
```

2. 反向类

对于每个字符类，都有一个“反向类”，用相同的字母表示，但要以大写书写形式。


\D 非数字：除 \d 以外的任何字符，例如字母。

\S 非空格符号：除 \s 以外的任何字符，例如字母。

\W 非单字字符：除 \w 以外的任何字符，例如非拉丁字母或空格。

```js

let str = "+7(903)-123-45-67";

alert( str.match(/\d/g).join('') ); // 79031234567
```