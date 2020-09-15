---
title: "Css Interview BFC"
date: 2020-09-03T10:02:39+08:00
draft: false
---

> Describe BFC (Block Formatting Context) and how it works.

## BFC(Block formatting context) 块级格式化上下文

A block formatting context is a part of a visual CSS rendering of a web page. It's the region in which the layout of block boxes occurs and in which floats interact with other elements.

BFC是Web页面的CSS渲染的一部分，是块盒子的布局过程发生的区域，也是浮动元素与其他元素交互的区域

## what
A block formatting context is created by at least one of the following:

- The root element of the document (<html>). 
- Floats (elements where float isn't none).
- Absolutely positioned elements (elements where position is absolute or fixed).
- Inline-blocks (elements with display: inline-block).
- Table cells (elements with display: table-cell, which is the default for HTML table cells).
- Table captions (elements with display: table-caption, which is the default for HTML table captions).
- Anonymous table cells implicitly created by the elements with display: table, table-row, table-row-group, table-header-group, table-footer-group (which is the default for HTML tables, table rows, table bodies, table  headers, and table footers, respectively), or inline-table.
- Block elements where overflow has a value other than visible.
- display: flow-root.
- Elements with contain: layout, content, or paint.
- Flex items (direct children of the element with display: flex or inline-flex) if they are neither flex nor grid nor table containers themselves.
- Grid items (direct children of the element with display: grid or inline-grid) if they are neither flex nor grid nor table containers themselves.
- Multicol containers (elements where column-count or column-width isn't auto, including elements with column-count: 1).
- column-span: all should always create a new formatting context, even when the column-span: all element isn't contained by a multicol container (Spec change, Chrome bug).

Formatting contexts affect layout, but typically, we create a new block formatting context for the positioning and clearing floats rather than changing the layout, because an element that establishes a new block formatting context will:


## Margin collapsing
Creating a new BFC to avoid the margin collapsing between two neighbor div:

## quote
[Block formatting context](https://developer.mozilla.org/en-US/docs/Web/Guide/CSS/Block_formatting_context)


## meaning

```
region [ˈriːdʒən]
n: (通常界限不明的)地区，区域，地方;行政区;(一国除首都以外的)全部地区，所有区域

occurs [əˈkɜːrz]
v. 发生;出现;存在于;出现在

interact [ˌɪntərˈækt]
v. 交流;沟通;合作;相互影响;相互作用

at least one of
至少有一个

Anonymous [əˈnɑːnɪməs]
adj. 不知姓名的; 名字不公开的; 匿名的; 不具名的; 没有特色的
adj. (of a person) not identified by name; of unknown name.

collapsing [kəˈlæpsɪŋ]
v: (突然)倒塌，坍塌;(尤指因病重而)倒下，昏倒，晕倒;(尤指工作劳累后)坐下，躺下放松

implicitly [ɪmˈplɪsətli]
adv. 暗中地；不明显地；含蓄地；蕴涵地;无疑问地；无保留地；绝对地
adv. in a way that is not directly expressed; tacitly.

typically  [ˈtɪpɪkli]
adv. 通常;一般;典型地;具有代表性地;不出所料;果然
adv. in most cases; usually.
```