---
title: "Js MouseEvent"
date: 2020-11-10T15:13:34+08:00
draft: false
---


```js
MouseEvent = {
  isTrusted: true
  screenX: 567
  screenY: 442
  clientX: 264, //  鼠标指针在点击元素（DOM）中的X坐标。
  clientY: 272, // 鼠标指针在点击元素（DOM）中的Y坐标。
  ctrlKey: false, //当鼠标事件触发时，如果 control 键被按下，则返回 true；
  shiftKey: false
  altKey: false
  metaKey: false, //当鼠标事件触发时，如果 meta键被按下，则返回 true；
  button: 0
  buttons: 0
  relatedTarget: null
  pageX: 264
  pageY: 297
  x: 264
  y: 272
  offsetX: 248, //鼠标指针相对于目标节点内边位置的X坐标
  offsetY: 69, //鼠标指针相对于目标节点内边位置的Y坐标
  movementX: 0, //鼠标指针相对于最后mousemove事件位置的X坐标。
  movementY: 0, //鼠标指针相对于最后mousemove事件位置的Y坐标。
  fromElement: null
  toElement: p
  layerX: 248, //鼠标指针相对于整个文档的X坐标；
  layerY: 125, //鼠标指针相对于整个文档的Y坐标；
  view: Window {parent: Window, opener: null, top: Window, length: 0, frames: Window, …}
  detail: 1
  sourceCapabilities: InputDeviceCapabilities {firesTouchEvents: true}
  which: 1
  type: "click"
  target: p
  currentTarget: null
  eventPhase: 0
  bubbles: true
  cancelable: true
  defaultPrevented: false
  composed: true
  timeStamp: 68556.15000007674
  srcElement: p
  returnValue: true
  cancelBubble: false
  path: (9) [p, div.post-content, div.post, div.content, div.container, body, html, document, Window]
  __proto__: MouseEvent
}
```


### 参考资料

- [MouseEvent](https://developer.mozilla.org/zh-CN/docs/Web/API/MouseEvent)

user agent stylesheet
a:-webkit-any-link {