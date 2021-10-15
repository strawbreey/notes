---
title: "User Select"
date: 2021-10-14T14:19:26+08:00
draft: false
---

### CSS user-select 属性


防止选取 <div> 元素的文本：

```css
div {
  -webkit-user-select: none; /* Safari */
  -ms-user-select: none; /* IE 10+ and Edge */
  user-select: none; /* Standard syntax */
}
```


### 定义和用法
user-select 属性规定是否能选取元素的文本。

在 web 浏览器中，如果您在文本上双击，文本会被选取或高亮显示。此属性用于阻止这种行为。


### JavaScript 

语法：object.style.userSelect="none"


### CSS 语法

user-select: auto|none|text|all;

### 属性值

- auto	默认。如果浏览器允许，则可以选择文本。
- none	防止文本选取。
- text	文本可被用户选取。
- all	单击选取文本，而不是双击。
