---
title: "Css Variables"
date: 2021-08-26T14:56:22+08:00
draft: false
---

## CSS 变量

### 基本用法

声明一个自定义属性，属性名需要以两个减号（--）开始，属性值则可以是任何有效的CSS值。和其他属性一样，自定义属性也是写在规则集之内的，如下：

```css
element {
  --main-bg-color: brown;
}
```

注意，规则集所指定的选择器定义了自定义属性的可见作用域。通常的最佳实践是定义在根伪类 :root 下，这样就可以在HTML文档的任何地方访问到它了：

```css
:root {
  --main-bg-color: brown;
}
```

然而这条规则不是绝对的，如果有理由去限制你的自定义属性，那么就应该限制。

> 注意：自定义属性名是大小写敏感的，--my-color 和 --My-color 会被认为是两个不同的自定义属性。

如前所述，使用一个局部变量时用 var() 函数包裹以表示一个合法的属性值：

```css
element {
  background-color: var(--main-bg-color);
}
```

### 使用自定义属性的第一步

我们从这个简单的CSS代码开始，它将相同的颜色应用在了不同class的元素上：
```css
.one {
  color: white;
  background-color: brown;
  margin: 10px;
  width: 50px;
  height: 50px;
  display: inline-block;
}

.two {
  color: white;
  background-color: black;
  margin: 10px;
  width: 150px;
  height: 70px;
  display: inline-block;
}
.three {
  color: white;
  background-color: brown;
  margin: 10px;
  width: 75px;
}
.four {
  color: white;
  background-color: brown;
  margin: 10px;
  width: 100px;
}

.five {
  background-color: brown;
}
```

应用在如下HTML上：
```html
<div>
  <div class="one">1:</div>
  <div class="two">2: Text <span class="five">5 - more text</span></div>
  <input class="three">
  <textarea class="four">4: Lorem Ipsum</textarea>
</div>
```

注意到在CSS代码中的重复：背景色 brown 被多处设置。对于一些CSS声明，是可以在级联关系更高的位置设置，通过CSS继承自然地解决这个重复的问题。但在一般项目中，是不可能通过这样的方式去解决。通过在 :root 伪类上设置自定义属性，然后在整个文档需要的地方使用，可以减少这样的重复性：

```css
:root {
  --main-bg-color: brown;
}

.one {
  color: white;
  background-color: var(--main-bg-color);
  margin: 10px;
  width: 50px;
  height: 50px;
  display: inline-block;
}

.two {
  color: white;
  background-color: black;
  margin: 10px;
  width: 150px;
  height: 70px;
  display: inline-block;
}
.three {
  color: white;
  background-color: var(--main-bg-color);
  margin: 10px;
  width: 75px;
}
.four {
  color: white;
  background-color: var(--main-bg-color);
  margin: 10px;
  width: 100px;
}

.five {
  background-color: var(--main-bg-color);
}
```

这里呈现的结果和前面的例子是一致的，但允许对所需属性值进行一个规范的声明。

### 自定义属性的继承性

自定义属性会继承。这意味着如果在一个给定的元素上，没有为这个自定义属性设置值，在其父元素上的值会被使用。看这一段HTML：

```html
<div class="one">
  <div class="two">
    <div class="three"></div>
    <div class="four"></div>
  </div>
</div>
```

配套的CSS：

```css
.two {
  --test: 10px;
}

.three {
  --test: 2em;
}
```

在这个情况下， var(--test) 的结果分别是：

- 对于元素 class="two" ：10px
- 对于元素 class="three" ：2em
- 对于元素 class="four" ：10px （继承自父属性）
- 对于元素 class="one" ：非法值，会变成自定义属性的默认值

注意，这些是自定义属性，并不是你在其他编程语言中遇到的实际的变量。这些值仅当需要的时候才会计算，而并不会按其他规则进行保存。比如，你不能为元素设置一个属性，然后让它从兄弟或旁支子孙规则上获取值。属性仅用于匹配当前选择器及其子孙，这和通常的CSS是一样的。

### 自定义属性备用值

用 var() 函数可以定义多个备用值(fallback value)，当给定值未定义时将会用备用值替换。这对于 Custom Elements 和 Shadow DOM 都很有用。

> 备用值并不是用于实现浏览器兼容性的。如果浏览器不支持CSS自定义属性，备用值也没什么用。它仅对支持CSS自定义属性的浏览器提供了一个备份机制，该机制仅当给定值未定义或是无效值的时候生效。

函数的第一个参数是自定义属性的名称。如果提供了第二个参数，则表示备用值，当自定义属性值无效时生效。第二个参数可以嵌套，但是不能继续平铺展开下去了，例如：

```css
.two {
  color: var(--my-var, red); /* Red if --my-var is not defined */
}

.three {
  background-color: var(--my-var, var(--my-background, pink)); /* pink if --my-var and --my-background are not defined */
}

.three {
  background-color: var(--my-var, --my-background, pink); /* Invalid: "--my-background, pink" */
}
```


第二个例子展示了如何处理一个以上的 fallback。该技术可能会导致性能问题，因为它花了更多的时间在处理这些变量上。

> 注意：像自定义属性这些 fallback 语法允许使用逗号。比如 var(--foo, red, blue) 定义了一个 red, blue 的备用值——从第一个逗号到最后的全部内容，都会被作为备用值的一部分。

### 有效性和值

传统的CSS概念里，有效性和属性是绑定的，这对自定义属性来说并不适用。当自定义属性值被解析，浏览器不知道它们什么时候会被使用，所以必须认为这些值都是有效的。

不幸的是，即便这些值是有效的，但当通过 var() 函数调用时，它在特定上下文环境下也可能不会奏效。属性和自定义变量会导致无效的CSS语句，这引入了一个新的概念：计算时有效性。

### 无效变量会导致什么？

当浏览器遇到无效的 var() 时，会使用继承值或初始值代替。

考虑如下代码：

HTML

```html
<p>This paragraph is initial black.</p> 
```

CSS

```css
:root { --text-color: 16px; }
p { color: blue; }
p { color: var(--text-color); }
```


毫不意外，浏览器将 --text-color 的值替换给了 var(--text-color)，但是 16px 并不是 color 的合法属性值。代换之后，该属性不会产生任何作用。浏览器会执行如下两个步骤：

1. 检查属性 color 是否为继承属性。是，但是 <p> 没有任何父元素定义了 color 属性。转到下一步。
2. 将该值设置为它的默认初始值，比如 black。


段落颜色并不是蓝色，因为无效代换导致了它被替换成了默认初始值的黑色。如果你直接写n color: 16px 的话，则会导致语法错误，而前面的定义则会生效（段落显示为蓝色）。

> 注意：当CSS属性-值对中存在语法错误，该行则会被忽略。然而如果自定义属性的值无效，它并不会被忽略，从而会导致该值被覆盖为默认值。

### JavaScript 中的值

在 JavaScript 中获取或者修改 CSS  变量和操作普通 CSS 属性是一样的：

```js
// 获取一个 Dom 节点上的 CSS 变量
element.style.getPropertyValue("--my-var");

// 获取任意 Dom 节点上的 CSS 变量
getComputedStyle(element).getPropertyValue("--my-var");

// 修改一个 Dom 节点上的 CSS 变量
element.style.setProperty("--my-var", jsVar + 4);
```