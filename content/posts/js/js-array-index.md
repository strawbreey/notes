---
title: "Js Array Index"
date: 2020-09-11T19:54:40+08:00
draft: false
---

## JS 中数组字符串索引和数值索引研究

Javascript的数组其实不像PHP或者其他一些语言一样拥有真正的字符串下标，当我们试图为一个js数组添加字符串下标的时候，其实就相当于为该数组对象添加了一个属性，属性名称就是我们所谓的“字符串下标”。由于为数组对象添加属性不会影响到同为该对象属性的length的值，因此该值将始终为零。同样地，.pop()和.shift()等作用于数组元素的方法也不能够作用于这些对象属性。因此，如果要使用的是一个完全由“字符串下标”组成的数组，那还是将其声明为一个Object类型的对象要更好一些。


 总结：

1、push()：在数组末端添加项；

2、shift()：把数组中的第一个元素删除；

3、pop()：把数组中的最后一个元素删除；

4、unshift()：在数组前端添加项；

5、push(),unshift()在推入多个项时，各个项之间的顺序不变

6、push(),unshift()将数组的长度+1并返回的是数组的长度，pop(),shift()将数组length-1并返回的是移除的项