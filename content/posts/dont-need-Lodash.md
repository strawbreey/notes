---
title: "Dont Need Lodash"
date: 2020-09-09T10:26:37+08:00
draft: false
---

Lodash and Underscore are great modern JavaScript utility libraries, and they are widely used by Front-end developers. However, when you are targeting modern browsers, you may find out that there are many methods which are already supported natively thanks to ECMAScript5 [ES5] and ECMAScript2015 [ES6]. If you want your project to require fewer dependencies, and you know your target browser clearly, then you may not need Lodash/Underscore.


## Quick Links

**[Array](#array)**

1. [chunk](#chunk)
1. [compact](#compact)
1. [concat](#concat)
1. [difference](#difference)
1. [drop](#drop)
1. [dropRight](#dropRight)
1. [fill](#fill)
1. [find](#find)
1. [findIndex](#findindex)
1. [first](#first)
1. [flatten](#flatten)
1. [flattenDeep](#flattendeep)
1. [fromPairs](#frompairs)
1. [head and _.tail](#head-and-_tail)
1. [indexOf](#indexof)
1. [intersection](#intersection)
1. [isArray](#isarray)
1. [isArrayBuffer](#isarraybuffer)
1. [join](#join)
1. [last](#last)
1. [lastIndexOf](#lastindexof)
1. [reverse](#reverse)
1. [slice](#slice)
1. [without](#without)
1. [initial](#initial)
1. [pull](#pull)

### chunk

Creates an array of elements split into groups the length of size.

```js
// Lodash
_.chunk(['a', 'b', 'c', 'd'], 2); // => [['a', 'b'], ['c', 'd']]
_.chunk(['a', 'b', 'c', 'd'], 3); // => [['a', 'b', 'c'], ['d']]

// Native
const chunk = (input, size) => {
  return input.reduce((arr, item, idx) => {
    return idx % size === 0
      ? [...arr, [item]]
      : [...arr.slice(0, -1), [...arr.slice(-1)[0], item]];
  }, []);
};

chunk(['a', 'b', 'c', 'd'], 2); // => [['a', 'b'], ['c', 'd']]

chunk(['a', 'b', 'c', 'd'], 3); // => [['a', 'b', 'c'], ['d']]
```

### compact

Creates an array with all falsy values removed.

创建一个新数组，包含原数组中所有的非假值元素。例如false, null, 0, "", undefined, 和 NaN 都是被认为是“假值”。

```js
// Lodash
_.compact([0, 1, false, 2, '', 3]);

// Native
[0, 1, false, 2, '', 3].filter(Boolean)
```

### concat

Creates a new array concatenating array with any additional arrays and/or values.

创建一个新数组，将array与任何数组 或 值连接在一起。

```js
// Lodash
var array = [1]
var other = _.concat(array, 2, [3], [[4]]) // output: [1, 2, 3, [4]]

// Native
var array = [1]
var other = array.concat(2, [3], [[4]]) // output: [1, 2, 3, [4]]
```

### difference

Similar to without, but returns the values from array that are not present in the other arrays.

创建一个具有唯一array值的数组，每个值不包含在其他给定的数组中。（注：即创建一个新数组，这个数组中的值，为第一个数字（array 参数）排除了给定数组中的值。）该方法使用 SameValueZero做相等比较。结果值的顺序是由第一个数组中的顺序确定。
 
```js
// Lodash
_.difference([1, 2, 3, 4, 5], [5, 2, 10])   // output: [1, 3, 4]

var arrays = [[1, 2, 3, 4, 5], [5, 2, 10]];

// Native
arrays.reduce(function(a, b) {
  return a.filter(function(value) {
    return !b.includes(value);
  });
});  // output: [1, 3, 4]

// ES6
arrays.reduce((a, b) => a.filter(c => !b.includes(c))) // output: [1, 3, 4]

```

### drop

Creates a slice of array with n elements dropped from the beginning

创建一个切片数组，去除array前面的n个元素 (n默认值为1)

```js
// Lodash
_.drop([1, 2, 3]);     // => [2, 3]
_.drop([1, 2, 3], 2);  // => [3]

// Native
[1, 2, 3].slice(1);    // => [2, 3]
[1, 2, 3].slice(2);    // => [3]
```

### dropRight

reates a slice of array with n elements dropped at the end.

```js
// Lodash
_.dropRight([1, 2, 3]);    // => [1, 2]
_.dropRight([1, 2, 3], 2); // => [1]

// Native
[1, 2, 3].slice(0, -1);    // => [1, 2]
[1, 2, 3].slice(0, -2);    // => [1]
```

### fill 

> fill(array, value, [start=0], [end=array.length])

Fills elements of array with value from start up to, but not including, end. Note that fill is a mutable method in both native and Lodash/Underscore.

使用 value 值来填充（替换） array，从start位置开始, 到end位置结束（但不包含end位置）

```js
// Lodash
var array = [1, 2, 3]

_.fill(array, 'a')               // output: ['a', 'a', 'a']
_.fill(Array(3), 2)              // output: [2, 2, 2]
_.fill([4, 6, 8, 10], '*', 1, 3) // output: [4, '*', '*', 10]

// Native
var array = [1, 2, 3]
array.fill('a')                   // output: ['a', 'a', 'a']
Array(3).fill(2)                  // output: [2, 2, 2]
[4, 6, 8, 10].fill('*', 1, 3)     // output: [4, '*', '*', 10]
```

### find

> _.findIndex(array, [predicate=_.identity], [fromIndex=0])

Returns the value of the first element in the array that satisfies the provided testing function. Otherwise undefined is returned.

该方法类似_.find，区别是该方法返回第一个通过 predicate 判断为真值的元素的索引值（index），而不是元素本身。

```js
// Lodash
var users = [
  { 'user': 'barney',  'age': 36, 'active': true },
  { 'user': 'fred',    'age': 40, 'active': false },
  { 'user': 'pebbles', 'age': 1,  'active': true }
]

_.find(users, function (o) { return o.age < 40; }) // output: object for 'barney'


// Native
var users = [
  { 'user': 'barney',  'age': 36, 'active': true },
  { 'user': 'fred',    'age': 40, 'active': false },
  { 'user': 'pebbles', 'age': 1,  'active': true }
]

users.find(function (o) { return o.age < 40; }) // output: object for 'barney'

```

### findIndex
Returns the index of the first element in the array that satisfies the provided testing function. Otherwise -1 is returned.

```js
// Lodash
var users = [
  { 'user': 'barney',  'age': 36, 'active': true },
  { 'user': 'fred',    'age': 40, 'active': false },
  { 'user': 'pebbles', 'age': 1,  'active': true }
]

var index = _.findIndex(users, function (o) { return o.age >= 40; }) // output: 1

// Native
var users = [
  { 'user': 'barney',  'age': 36, 'active': true },
  { 'user': 'fred',    'age': 40, 'active': false },
  { 'user': 'pebbles', 'age': 1,  'active': true }
]

var index = users.findIndex(function (o) { return o.age >= 40; }) // output: 1
```

### first
Returns the first element of an array. Passing n will return the first n elements of the array.
```js
// Lodash
_.first([1, 2, 3, 4, 5]);  // => 1

_.first([1, 2, 3, 4, 5], 2); // => [1, 2]

// Native
[1, 2, 3, 4, 5][0]; // => 1

[].concat(1, 2, 3, 4, 5).shift() // => 1

[].concat([1, 2, 3, 4, 5]).shift() // => 1

[].concat(undefined).shift() // => undefined

[1, 2, 3, 4, 5].slice(0, 2); // => [1, 2]
```

### flatten
Flattens array a single level deep.

```js
// Lodash
_.flatten([1, [2, [3, [4]], 5]]); // => [1, 2, [3, [4]], 5]

// Native
const flatten = [1, [2, [3, [4]], 5]].reduce((a, b) => a.concat(b), []) // => [1, 2, [3, [4]], 5]
const flatten = [].concat(...[1, [2, [3, [4]], 5]]) // => [1, 2, [3, [4]], 5]

// Native(ES2019)
const flatten = [1, [2, [3, [4]], 5]].flat() // => [1, 2, [3, [4]], 5]
const flatten = [1, [2, [3, [4]], 5]].flatMap(number => number) // => [1, 2, [3, [4]], 5]
```


### flattenDeep

Recursively flattens array.

```js
// Lodash
_.flattenDeep([1, [2, [3, [4]], 5]]); // => [1, 2, 3, 4, 5]

// Native
const flattenDeep = (arr) => Array.isArray(arr)
  ? arr.reduce((a, b) => a.concat(flattenDeep(b)) , [])
  : [arr]

flattenDeep([1, [[2], [3, [4]], 5]]) // => [1, 2, 3, 4, 5]

// Native(ES2019)
[1, [2, [3, [4]], 5]].flat(Infinity) // => [1, 2, 3, 4, 5]

const flattenDeep = (arr) => arr.flatMap((subArray, index) => Array.isArray(subArray) ? flattenDeep(subArray) : subArray)

flattenDeep([1, [[2], [3, [4]], 5]]) // => [1, 2, 3, 4, 5]

```

### fromPairs

Returns an object composed from key-value pairs.

```js
// Lodash
_.fromPairs([['a', 1], ['b', 2]]); // => { 'a': 1, 'b': 2 }

// Native
const fromPairs = function(arr) {
  return arr.reduce(function(accumulator, value) {
    accumulator[value[0]] = value[1];
    return accumulator;
  }, {})
}

// Compact form
const fromPairs = (arr) => arr.reduce((acc, val) => (acc[val[0]] = val[1], acc), {})

fromPairs([['a', 1], ['b', 2]]); // => { 'a': 1, 'b': 2 }

// Native(ES2019)
Object.fromEntries([['a', 1], ['b', 2]]) // => { 'a': 1, 'b': 2 }
```

### head and tail
Gets the first element or all but the first element.
head: 获取数组 array 的第一个元素。
tail: 获取除了array数组第一个元素以外的全部元素

```js
const array = [1, 2, 3]

// Lodash: _.first, _.head
_.head(array) // output: 1

// Lodash: _.tail
_.tail(array) // output: [2, 3]


// Native
const [ head, ...tail ] = array
console.log(head) // output: 1
console.log(tail) // output [2, 3]
```

### indexOf

Returns the first index at which a given element can be found in the array, or -1 if it is not present.

```js
// Underscore/Lodash
var array = [2, 9, 9]
var result = _.indexOf(array, 2)// output: 0

// Native
var array = [2, 9, 9]
var result = array.indexOf(2) // output: 0
```

### intersection
Returns an array that is the intersection of all the arrays. Each value in the result is present in each of the arrays.
创建唯一值的数组，这个数组包含所有给定数组都包含的元素，使用 SameValueZero进行相等性比较。（注：可以理解为给定数组的交集）

```js
// Lodash
_.intersection([1, 2, 3], [101, 2, 1, 10], [2, 1]) // output: [1, 2]

// Native
var arrays = [[1, 2, 3], [101, 2, 1, 10], [2, 1]];
arrays.reduce(function(a, b) {
  return a.filter(function(value) {
    return b.includes(value);
  });
}); // output: [1, 2]

// ES6
let arrays = [[1, 2, 3], [101, 2, 1, 10], [2, 1]];
arrays.reduce((a, b) => a.filter(c => b.includes(c))); // output: [1, 2]

```

### takeRight

Creates a slice of array with n elements taken from the end. ❗ Native slice does not behave entirely the same as the Lodash method. See example below to understand why.

```js
// Lodash
_.takeRight([1, 2, 3]); // => [3]
_.takeRight([1, 2, 3], 2); // => [2, 3]
_.takeRight([1, 2, 3], 5); // => [1, 2, 3]

// Native
[1, 2, 3].slice(-1); // => [3]
[1, 2, 3].slice(-2); // => [2, 3]
[1, 2, 3].slice(-5); // => [1, 2, 3]

// Difference in compatibility
// Lodash
_.takeRight([1, 2, 3], 0); // => []

// Native
[1, 2, 3].slice(0); // => [1, 2, 3]
```

### isArray
Returns true if given value is an array.
```js
// Lodash
var array = []
_.isArray(array) // output: true

// Native
var array = []
Array.isArray(array) // output: true

if (!Array.isArray) {
  Array.isArray = function (arg) {
    return Object.prototype.toString.call(arg) === '[object Array]'
  }
}

// 在JavaScript中,想要判断某个对象值属于哪种内置类型,最靠谱的做法就是通过Object.prototype.toString方法.
```

### isArrayBuffer
Checks if value is classified as an ArrayBuffer object.

```js
// Lodash
_.isArrayBuffer(new ArrayBuffer(2)) // output: true

// Native
new ArrayBuffer(2) instanceof ArrayBuffer // output: true
```





### 相关文章
 - [You-Dont-Need-Lodash-Underscore](https://github.com/you-dont-need/You-Dont-Need-Lodash-Underscore#_chunk)
 - [https://www.lodashjs.com/](https://www.lodashjs.com/)
 - [https://lodash.com/](https://lodash.com/)
 - [https://underscorejs.org/](https://underscorejs.org/)
