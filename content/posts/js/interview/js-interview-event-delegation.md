---
title: "Js Interview Event Delegation"
date: 2021-01-25T16:33:23+08:00
draft: false
tags: ['interview']
---

请解释事件委托（event delegation）

事件委托是将事件监听器添加到父元素，而不是每个子元素单独设置事件监听器。当触发子元素时，事件会冒泡到父元素，监听器就会触发。这种技术的好处是：

- 内存占用减少，因为只需要一个父元素的事件处理程序，而不必为每个后代都添加事件处理程序。
- 无需从已删除的元素中解绑处理程序，也无需将处理程序绑定到新元素上。


###  参考资料

- [What is DOM Event delegation?](https://stackoverflow.com/questions/1687296/what-is-dom-event-delegation)
