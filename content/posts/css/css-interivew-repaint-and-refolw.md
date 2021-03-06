---
title: "Css Interivew Repaint and Refolw"
date: 2020-09-03T10:48:27+08:00
draft: false
---

## 重绘与回流(repaint-and-refolw)

![avatar](/images/1.png)

从上面这个图上，我们可以看到，浏览器渲染过程如下：

- 解析HTML，生成DOM树

- 解析CSS，生成CSSOM树

- 将DOM树和CSSOM树结合，生成渲染树(Render Tree)

- Layout(回流): 根据生成的渲染树，进行回流(Layout)，得到节点的几何信息（位置，大小） 

- Painting(重绘): 根据渲染树以及回流得到的几何信息，得到节点的绝对像素

- Display:将像素发送给GPU，展示在页面上。


> 注意：回流一定会触发重绘，而重绘不一定会回流, 前端性能优化，一般是尽量避免回流


产生回流的场景如下: 

- 添加或删除可见的DOM元素
- 元素的位置发生变化
- 元素的尺寸发生变化（包括外边距、内边框、边框大小、高度和宽度等）
- 内容发生变化，比如文本变化或图片被另一个不同尺寸的图片所替代。
- 页面一开始渲染的时候（这肯定避免不了）
- 浏览器的窗口尺寸变化（因为回流是根据视口的大小来计算元素的位置和大小的）


针对回流和重绘的优化方案:

1. 合并多次对DOM和样式的修改
2. 批量修改DOM
3. 使元素脱离文档流,对其进行多次修改,将元素带回到文档中。
4. 对于复杂动画效果,使用绝对定位让其脱离文档流
5. 比起考虑如何减少回流重绘，我们更期望的是，根本不要回流重绘。

有三种方式可以让DOM脱离文档流：

- 隐藏元素，应用修改，重新显示
- 使用文档片段(document fragment)在当前DOM之外构建一个子树，再把它拷贝回文档。
- 将原始元素拷贝到一个脱离文档的节点中，修改节点后，再替换原始的元素。

### 回流 (重新计算DOM节点)

前面我们通过构造渲染树，我们将可见DOM节点以及它对应的样式结合起来，可是我们还需要计算它们在设备视口(viewport)内的确切位置和大小，这个计算的阶段就是回流。


### 什么节点是不可见的。不可见的节点包括：

一些不会渲染输出的节点，比如script、meta、link等。
一些通过css进行隐藏的节点。比如display:none。注意，利用visibility和opacity隐藏的节点，还是会显示在渲染树上的。只有display:none的节点才不会显示在渲染树上。

> 注意：渲染树只包含可见的节点


## 重绘 (重新构造渲染树)

我们通过构造渲染树和回流阶段，我们知道了哪些节点是可见的，以及可见节点的样式和具体的几何信息(位置、大小)，那么我们就可以将渲染树的每个节点都转换为屏幕上的实际像素，这个阶段就叫做重绘节点。


## 浏览器的优化机制

现代的浏览器都是很聪明的，由于每次重排都会造成额外的计算消耗，因此大多数浏览器都会通过队列化修改并批量执行来优化重排过程。浏览器会将修改操作放入到队列里，直到过了一段时间或者操作达到了一个阈值，才清空队列。

- 当你获取布局信息的操作的时候，会强制队列刷新，比如当你访问以下属性或者使用以下方法：
offsetTop、offsetLeft、offsetWidth、offsetHeight
scrollTop、scrollLeft、scrollWidth、scrollHeight
clientTop、clientLeft、clientWidth、clientHeight
getComputedStyle()
getBoundingClientRect


## 参考链接 

- [渲染树构建、布局及绘制](https://developers.google.com/web/fundamentals/performance/critical-rendering-path/render-tree-construction?hl=zh-cn)