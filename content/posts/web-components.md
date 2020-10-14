---
title: "Web Components"
date: 2020-10-14T21:45:06+08:00
draft: true
---

众所周知，开发复杂软件的原则是：不要让软件复杂

如果某个部分变得复杂了 —— 将其拆分成更简单的部分，再以最简明的方式组合起来。

只有让复杂的事情简单化的架构才是好架构。

一个组件有：

自己的 JavaScript 类。
DOM 结构，并且只由自己的类管理，无法被外部代码操作。（「封装」原则）。
CSS 样式，作用在这个组件上。
API：事件，类方法等等，让组件可以与其他组件交互。

Custom elements —— 用于自定义 HTML 元素.
Shadow DOM —— 为组件创建内部 DOM，它对外部是不可见的。
CSS Scoping —— 申明仅应用于组件的 Shadow DOM 内的样式。
Event retargeting 以及更多的小东西，让自定义组件更适用于开发工作。


### Custom elements (自定义标签)

```js
class MyElement extends HTMLElement {
  constructor() {
    super();
    // 元素在这里创建
  }

  connectedCallback() {
    // 在元素被添加到文档之后，浏览器会调用这个方法
    //（如果一个元素被反复添加到文档／移除文档，那么这个方法会被多次调用）
  }

  disconnectedCallback() {
    // 在元素从文档移除到时候，浏览器会调用这个方法
    // （如果一个元素被反复添加到文档／移除文档，那么这个方法会被多次调用）
  }

  static get observedAttributes() {
    return [/* 属性数组，这些属性的变化会被被监视 */];
  }

  attributeChangedCallback(name, oldValue, newValue) {
    // 当上面数组里面的属性变化的时候，这个方法会被调用
  }

  adoptedCallback() {
    // 在元素被移动到新的文档的时候，这个方法会被调用
    // （document.adoptNode 会用到, 非常少见）
  }

  // 还可以添加更多的元素方法和属性
}
```

```js
// 让浏览器知道我们新定义的类是为 <my-element> 服务的
customElements.define("my-element", MyElement);
```

    Custom element 名称必须包括一个短横线 -
    Custom element 名称必须包括一个短横线 -, 比如 my-element 和 super-button 都是有效的元素名，但 myelement 并不是。

    这是为了确保 custom element 和内置 HTML 元素之间不会发生命名冲突。

创建一个自定义时间标签

```html
<script>
class TimeFormatted extends HTMLElement { 
  // 在元素被添加到文档之后，浏览器会调用这个方法 => 相当于生命周期的mount

  connectedCallback() {
    let date = new Date(this.getAttribute('datetime') || Date.now());

    this.innerHTML = new Intl.DateTimeFormat("default", {
      year: this.getAttribute('year') || undefined,
      month: this.getAttribute('month') || undefined,
      day: this.getAttribute('day') || undefined,
      hour: this.getAttribute('hour') || undefined,
      minute: this.getAttribute('minute') || undefined,
      second: this.getAttribute('second') || undefined,
      timeZoneName: this.getAttribute('time-zone-name') || undefined,
    }).format(date);
  }

}

customElements.define("time-formatted", TimeFormatted);
</script>

<!-- (3) -->
<time-formatted 
    datetime="2019-12-01"
    year="numeric" 
    month="long" 
    day="numeric"
    hour="numeric" 
    minute="numeric" 
    second="numeric"
    time-zone-name="short"
></time-formatted>

```


在 connectedCallback 中渲染，而不是 constructor 中
在上面的例子中，元素里面的内容是在 connectedCallback 中渲染（创建）的。

为什么不在 constructor 中渲染？

原因很简单：在 constructor 被调用的时候，还为时过早。虽然这个元素实例已经被创建了，但还没有插入页面。在这个阶段，浏览器还没有处理／创建元素属性：调用 getAttribute 将会得到 null。所以我们并不能在那里渲染元素。

而且，如果你仔细考虑，这样作对于性能更好 —— 推迟渲染直到真正需要的时候。

在元素被添加到文档的时候，它的 connectedCallback 方法会被调用。这个元素不仅仅是被添加为了另一个元素的子元素，同样也成为了页面的一部分。因此我们可以构建分离的 DOM，创建元素并且让它们为之后的使用准备好。它们只有在插入页面的时候才会真的被渲染。


### 监视属性

我们目前的 <time-formatted> 实现中，在元素渲染以后，后续的属性变化并不会带来任何影响。这对于 HTML 元素来说有点奇怪。通常当我们改变一个属性的时候，比如 a.href，我们会预期立即看到变化。我们将会在下面修正这一点。

为了监视这些属性，我们可以在 observedAttributes() static getter 中提供属性列表。当这些属性发生变化的时候，attributeChangedCallback 会被调用。出于性能优化的考虑，其他属性变化的时候并不会触发这个回调方法。

以下是 <time-formatted> 的新版本，它会在属性变化的时候自动更新：

```html

<script>
class TimeFormatted extends HTMLElement {

  render() { // (1)
    let date = new Date(this.getAttribute('datetime') || Date.now());

    this.innerHTML = new Intl.DateTimeFormat("default", {
      year: this.getAttribute('year') || undefined,
      month: this.getAttribute('month') || undefined,
      day: this.getAttribute('day') || undefined,
      hour: this.getAttribute('hour') || undefined,
      minute: this.getAttribute('minute') || undefined,
      second: this.getAttribute('second') || undefined,
      timeZoneName: this.getAttribute('time-zone-name') || undefined,
    }).format(date);
  }

  connectedCallback() { // (2)
    if (!this.rendered) {
      this.render();
      this.rendered = true;
    }
  }

    // 设置监听属性
  static get observedAttributes() { // (3)
    return ['datetime', 'year', 'month', 'day', 'hour', 'minute', 'second', 'time-zone-name'];
  }

  attributeChangedCallback(name, oldValue, newValue) { // (4)
    this.render();
  }

}

customElements.define("time-formatted", TimeFormatted);
</script>

<time-formatted id="elem" hour="numeric" minute="numeric" second="numeric"></time-formatted>

<script>
setInterval(() => elem.setAttribute('datetime', new Date()), 1000); // (5)
</script>
```

渲染逻辑被移动到了 render() 这个辅助方法里面。
这个方法在元素被插入到页面的时候调用。
attributeChangedCallback 在 observedAttributes() 里的属性改变的时候被调用。
…… 然后重渲染元素。
最终，一个计时器就这样被我们轻松地实现了。


### 渲染顺序

在 HTML 解析器构建 DOM 的时候，会按照先后顺序处理元素，先处理父级元素再处理子元素。例如，如果我们有 <outer><inner></inner></outer>，那么 <outer> 元素会首先被创建并接入到 DOM，然后才是 <inner>。

这对 custom elements 产生了重要影响。

比如，如果一个 custom element 想要在 connectedCallback 内访问 innerHTML，它什么也拿不到:

```html
<script>
customElements.define('user-info', class extends HTMLElement {
  connectedCallback() {
    alert(this.innerHTML); // empty (*)
  }
});
</script>

<user-info>John</user-info>
```

如果我们要给 custom element 传入信息，我们可以使用元素属性。它们是即时生效的。

或者，如果我们需要子元素，我们可以使用延迟时间为零的 setTimeout 来推迟访问子元素。

这样是可行的：


```html
<script>
customElements.define('user-info', class extends HTMLElement {

  connectedCallback() {
    setTimeout(() => alert(this.innerHTML)); // John (*)
  }

});
</script>

<user-info>John</user-info>
```

让我们用一个例子来说明：

```html
<script>
customElements.define('user-info', class extends HTMLElement {
  connectedCallback() {
    alert(`${this.id} 已连接。`);
    setTimeout(() => alert(`${this.id} 初始化完成。`));
  }
});
</script>

<user-info id="outer">
  <user-info id="inner"></user-info>
</user-info>
```

输出顺序：

outer 已连接。
inner 已连接。
outer 初始化完成。
inner 初始化完成。

### SEO

我们创建的 <time-formatted> 这些新元素，并没有任何相关的语义。搜索引擎并不知晓它们的存在，同时无障碍设备也无法处理它们。

```html
<script>
// 这个按钮在被点击的时候说 "hello"
class HelloButton extends HTMLButtonElement {
  constructor() {
    super();
    this.addEventListener('click', () => alert("Hello!"));
  }
}

customElements.define('hello-button', HelloButton, {extends: 'button'});
</script>

<button is="hello-button">Click me</button>

<button is="hello-button" disabled>Disabled</button>
```

### 总结

有两种 custom element

“Autonomous” —— 全新的标签，继承 HTMLElement。

定义方式：

```js
class MyElement extends HTMLElement {
  constructor() { super(); /* ... */ }
  connectedCallback() { /* ... */ }
  disconnectedCallback() { /* ... */  }
  static get observedAttributes() { return [/* ... */]; }
  attributeChangedCallback(name, oldValue, newValue) { /* ... */ }
  adoptedCallback() { /* ... */ }
 }
customElements.define('my-element', MyElement);
/* <my-element> */

```

“Customized built-in elements” —— 已有元素的扩展。

需要多一个 .define 参数，同时 is="..." 在 HTML 中：

```js
class MyButton extends HTMLButtonElement { /*...*/ }
customElements.define('my-button', MyElement, {extends: 'button'});
/* <button is="my-button"> */
```

请注意：

在元素被从文档移除的时候，我们会清除 setInterval 的 timer。这非常重要，否则即使我们不再需要它了，它仍然会继续计时。这样浏览器就不能清除这个元素占用和被这个元素引用的内存了。
我们可以通过 elem.date 属性得到当前时间。类所有的方法和属性天生就是元素的方法和属性。


## Shadow DOM

Shadow DOM 为封装而生。它可以让一个组件拥有自己的「影子」DOM 树，这个 DOM 树不能在主文档中被任意访问，可能拥有局部样式规则，还有其他特性。

内建 shadow DOM

浏览器在内部使用 DOM/CSS 来绘制它们。这个 DOM 结构一般来说对我们是隐藏的，但我们可以在开发者工具里面看见它。比如，在 Chrome 里，我们需要打开「Show user agent shadow DOM」选项。

然后 <input type="range"> 看起来会像这样：

你在 #shadow-root 下看到的就是被称为「shadow DOM」的东西。

我们不能使用一般的 JavaScript 调用或者选择器来获取内建 shadow DOM 元素。它们不是常规的子元素，而是一个强大的封装手段。

在上面的例子中，我们可以看到一个有用的属性 pseudo。这是一个因为历史原因而存在的属性，并不在标准中。我们可以使用它来给子元素加上 CSS 样式，像这样：

```html
<style>
/* 让滑块轨道变红 */
input::-webkit-slider-runnable-track {
  background: red;
}
</style>

<input type="range">
```

### Shadow tree

一个 DOM 元素可以有以下两类 DOM 子树：

Light tree（光明树） —— 一个常规 DOM 子树，由 HTML 子元素组成。我们在之前章节看到的所有子树都是「光明的」。
Shadow tree（影子树） —— 一个隐藏的 DOM 子树，不在 HTML 中反映，无法被察觉。


### 封装
Shadow DOM 被非常明显地和主文档分开：

Shadow DOM 元素对于 light DOM 中的 querySelector 不可见。实际上，Shadow DOM 中的元素可能与 light DOM 中某些元素的 id 冲突。这些元素必须在 shadow tree 中独一无二。
Shadow DOM 有自己的样式。外部样式规则在 shadow DOM 中不产生作用。


### 总结
Shadow DOM 是创建组件级别 DOM 的一种方法。

shadowRoot = elem.attachShadow({mode: open|closed}) —— 为 elem 创建 shadow DOM。如果 mode="open"，那么它通过 elem.shadowRoot 属性被访问。
我们可以使用 innerHTML 或者其他 DOM 方法来扩展 shadowRoot。
Shadow DOM 元素：

有自己的 id 空间。
对主文档的 JavaScript 选择器隐身，比如 querySelector。
只使用 shadow tree 内部的样式，不使用主文档的样式。
Shadow DOM，如果存在的话，会被浏览器渲染而不是所谓的 「light DOM」（普通子元素）。在 Shadow DOM 插槽，组成 章节中我们将会看到如何组织它们。

## 模板元素

内建的 <template> 元素用来存储 HTML 模板。浏览器将忽略它的内容，仅检查语法的有效性，但是我们可以在 JavaScript 中访问和使用它来创建其他元素。

从理论上讲，我们可以在 HTML 中的任何位置创建不可见元素来储存 HTML 模板。那 <template> 元素有什么优势？

首先，其内容可以是任何有效的HTML，即使它通常需要特定的封闭标签。

```html
<template>
  <tr>
    <td>Contents</td>
  </tr>
</template>

```

### 总结

<template> 的内容可以是任何语法正确的 HTML。
<template> 内容被视为“超出文档范围”，因此它不会产生任何影响。
我们可以在JavaScript 中访问 template.content ，将其克隆以在新组件中重复使用。
<template> 标签非常独特，因为：
浏览器将检查其中的HTML语法（与在脚本中使用模板字符串不同）。
但允许使用任何顶级 HTML 标签，即使没有适当包装元素的无意义的元素（例如 <tr> ）。
其内容是交互式的：插入其文档后，脚本会运行， <video autoplay> 会自动播放。
<template> 元素不具有任何迭代机制，数据绑定或变量替换的功能，但我们可以在其基础上实现这些功能。
