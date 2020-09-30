---
title: "Js Array1"
date: 2020-09-28T16:15:53+08:00
draft: false
---

### indexOf 

法返回在数组中可以找到一个给定元素的第一个索引，如果不存在，则返回-1。

indexOf 使用strict equality (无论是 ===, 还是 triple-equals操作符都基于同样的方法)进行判断 searchElement与数组中包含的元素之间的关系。

```js
const beasts = ['ant', 'bison', 'camel', 'duck', 'bison'].indexOf('bison')
// expected output: 1

// start from index 2
['ant', 'bison', 'camel', 'duck', 'bison'].indexOf('bison', 2)
// expected output: 4

['ant', 'bison', 'camel', 'duck', 'bison'].indexOf('giraffe')
// expected output: -1

var array = [2, 5, 9];
array.indexOf(2);     // 0
array.indexOf(7);     // -1
array.indexOf(9, 2);  // 2
array.indexOf(2, -1); // -1
array.indexOf(2, -3); // 0
```

### join

 方法将一个数组（或一个类数组对象）的所有元素连接成一个字符串并返回这个字符串。如果数组只有一个项目，那么将返回该项目而不使用分隔符。

 ```js
const elements = ['Fire', 'Air', 'Water'];

elements.join() // expected output: "Fire,Air,Water"

elements.join('') // expected output: "FireAirWater"

elements.join('-') // expected output: "Fire-Air-Water"
 ```

### keys

keys() 方法返回一个包含数组中每个索引键的Array Iterator对象。

```js
const array1 = ['a', 'b', 'c'];
const iterator = array1.keys();

for (const key of iterator) {
  console.log(key);
}

// expected output: 0
// expected output: 1
// expected output: 2
```