---
title: "Js Slice"
date: 2020-12-09T20:14:00+08:00
draft: false
---


### Array

1. slice() 方法可从已有的数组中返回选定的元素。
```js
arrayObject.slice(start,end)
```

- start	必需。规定从何处开始选取。如果是负数，那么它规定从数组尾部开始算起的位置。也就是说，-1 指最后一个元素，-2 指倒数第二个元素，以此类推。
- end	可选。规定从何处结束选取。该参数是数组片断结束处的数组下标。如果没有指定该参数，那么切分的数组包含从 start 到数组结束的所有元素。如果这个参数是负数，那么它规定的是从数组尾部开始算起的元素。

返回一个新的数组，包含从 start 到 end （不包括该元素）的 arrayObject 中的元素。

请注意，该方法并不会修改数组，而是返回一个子数组。如果想删除数组中的一段元素，应该使用方法 Array.splice()。



2. splice() 方法向/从数组中添加/删除项目，然后返回被删除的项目。

注释：该方法会改变原始数组。

```js
arrayObject.splice(index,howmany,item1,.....,itemX)
```

- index	必需。整数，规定添加/删除项目的位置，使用负数可从数组结尾处规定位置。
- howmany	必需。要删除的项目数量。如果设置为 0，则不会删除项目。
- item1, ..., itemX	可选。向数组添加的新项目。

Array	包含被删除项目的新数组，如果有的话。

splice() 方法可删除从 index 处开始的零个或多个元素，并且用参数列表中声明的一个或多个值来替换那些被删除的元素。

如果从 arrayObject 中删除了元素，则返回的是含有被删除的元素的数组。


## String

1. slice() 方法可提取字符串的某个部分，并以新的字符串返回被提取的部分。

```js
stringObject.slice(start,end)
```

- start	要抽取的片断的起始下标。如果是负数，则该参数规定的是从字符串的尾部开始算起的位置。也就是说，-1 指字符串的最后一个字符，-2 指倒数第二个字符，以此类推。
- end	紧接着要抽取的片段的结尾的下标。若未指定此参数，则要提取的子串包括 start 到原字符串结尾的字符串。如果该参数是负数，那么它规定的是从字符串的尾部开始算起的位置。

返回值

一个新的字符串。包括字符串 stringObject 从 start 开始（包括 start）到 end 结束（不包括 end）为止的所有字符。

说明

String 对象的方法 slice()、substring() 和 substr() （不建议使用）都可返回字符串的指定部分。slice() 比 substring() 要灵活一些，因为它允许使用负数作为参数。slice() 与 substr() 有所不同，因为它用两个字符的位置来指定子串，而 substr() 则用字符位置和长度来指定子串。

还要注意的是，String.slice() 与 Array.slice() 相似。


2. substring() 方法用于提取字符串中介于两个指定下标之间的字符。

```js
stringObject.substring(start,stop)
```

- start	必需。一个非负的整数，规定要提取的子串的第一个字符在 stringObject 中的位置。
- stop	可选。一个非负的整数，比要提取的子串的最后一个字符在 stringObject 中的位置多 1。

一个新的字符串，该字符串值包含 stringObject 的一个子字符串，其内容是从 start 处到 stop-1 处的所有字符，其长度为 stop 减 start。

substring() 方法返回的子串包括 start 处的字符，但不包括 stop 处的字符。

如果参数 start 与 stop 相等，那么该方法返回的就是一个空串（即长度为 0 的字符串）。如果 start 比 stop 大，那么该方法在提取子串之前会先交换这两个参数。
