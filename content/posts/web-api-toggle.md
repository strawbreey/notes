---
title: "Web Api Toggle"
date: 2020-11-16T19:20:23+08:00
draft: false
tags: ['webApi', 'toggle']
---


[DOMTokenList](/posts/web-api-domtokenlist) 接口的 toggle() 方法从列表中删除一个给定的标记 并返回 false 。 如果标记 不存在，则添加并且函数返回 true。

```js
tokenList.toggle(token, force);

```
参数列表

- token 标记列表中你想探查并切换的 DOMString .
- force (可选) 一个Boolean 值, 设置后会将方法变成单向操作. 如设置为false, 则会删除标记列表中匹配的给定标记，且不会再度添加. 如设置为 true, 则将在标记列表中添加给定标记，且不会再度删除。


example： 

```js
// <span class="a b">classList is 'a b'</span>

var span = document.querySelector("span");
var classes = span.classList;
span.onclick = function() {
  // 如果有就添加, 没有删除
  var result = classes.toggle("c");
  if(result) {
    span.textContent = "'c' added; classList is now '" + classes + "'.";
  } else {
    span.textContent = "'c' removed; classList is now '" + classes + "'.";
  }
}
```