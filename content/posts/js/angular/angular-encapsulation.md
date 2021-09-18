---
title: "Angular Encapsulation"
date: 2021-09-13T20:42:08+08:00
draft: false
---

### 视图封装模式

在 Angular 中，组件的 CSS 样式被封装进了自己的视图中，而不会影响到应用程序的其它部分。

通过在组件的元数据上设置`视图封装模式`，你可以分别控制每个组件的封装模式。 可选的封装模式一共有如下几种：

- ShadowDom 模式使用浏览器原生的 Shadow DOM 实现（参阅 MDN 上的 Shadow DOM）来为组件的宿主元素附加一个 Shadow DOM。组件的视图被附加到这个 Shadow DOM 中，组件的样式也被包含在这个 Shadow DOM 中。(译注：不进不出，没有样式能进来，组件样式出不去。)

- Emulated 模式（默认值）通过预处理（并改名）CSS 代码来模拟 Shadow DOM 的行为，以达到把 CSS 样式局限在组件视图中的目的。 更多信息，见附录 1。(译注：只进不出，全局样式能进来，组件样式出不去)

- None 意味着 Angular 不使用视图封装。 Angular 会把 CSS 添加到全局样式中。而不会应用上前面讨论过的那些作用域规则、隔离和保护等。 从本质上来说，这跟把组件的样式直接放进 HTML 是一样的。(译注：能进能出。)

通过组件元数据中的 encapsulation 属性来设置组件封装模式：

```js
src/app/quest-summary.component.ts
content_copy
// warning: few browsers support shadow DOM encapsulation at this time
encapsulation: ViewEncapsulation.ShadowDom
```
ShadowDom 模式只适用于提供了原生 Shadow DOM 支持的浏览器（参阅 Can I use 上的 Shadow DOM v1 部分）。 它仍然受到很多限制，这就是为什么仿真 (Emulated) 模式是默认选项，并建议将其用于大多数情况。